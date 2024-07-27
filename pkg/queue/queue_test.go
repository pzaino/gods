// Copyright 2024 Paolo Fabio Zaino
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package queue provides a non-concurrent-safe queue (FIFO).
package queue

import (
	"strconv"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	if q.Size() != 0 {
		t.Errorf("Queue should be empty")
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Errorf("Queue should have 3 elements")
	}

	item, err := q.Peek()
	if err != nil {
		t.Errorf("Peek should not return an error")
	}
	if item != 1 {
		t.Errorf("Peek should return 1")
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue should not return an error")
	}
	if item != 1 {
		t.Errorf("Dequeue should return 1")
	}

	if q.Size() != 2 {
		t.Errorf("Queue should have 2 elements")
	}

	if q.IsEmpty() {
		t.Errorf("Queue should not be empty")
	}

	q.Clear()
	if !q.IsEmpty() {
		t.Errorf("Queue should be empty")
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item, err := q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue should not return an error")
	}
	if item != 1 {
		t.Errorf("Dequeue should return 1")
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue should not return an error")
	}
	if item != 2 {
		t.Errorf("Dequeue should return 2")
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue should not return an error")
	}
	if item != 3 {
		t.Errorf("Dequeue should return 3")
	}

	if !q.IsEmpty() {
		t.Errorf("Queue should be empty after dequeueing all elements")
	}

	_, err = q.Dequeue()
	if err == nil {
		t.Errorf("Dequeue should return an error when the queue is empty")
	}
}

func TestPeek(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item, err := q.Peek()
	if err != nil {
		t.Errorf("Peek should not return an error")
	}
	if item != 1 {
		t.Errorf("Peek should return 1")
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue should not return an error")
	}
	if item != 1 {
		t.Errorf("Dequeue should return 1")
	}

	item, err = q.Peek()
	if err != nil {
		t.Errorf("Peek should not return an error")
	}
	if item != 2 {
		t.Errorf("Peek should return 2")
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue should not return an error")
	}
	if item != 2 {
		t.Errorf("Dequeue should return 2")
	}

	item, err = q.Peek()
	if err != nil {
		t.Errorf("Peek should not return an error")
	}
	if item != 3 {
		t.Errorf("Peek should return 3")
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue should not return an error")
	}
	if item != 3 {
		t.Errorf("Dequeue should return 3")
	}

	_, err = q.Peek()
	if err == nil {
		t.Errorf("Peek should return an error when the queue is empty")
	}
}

func TestValues(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	values := q.Values()
	if len(values) != 3 {
		t.Errorf("Values should return a slice with 3 elements")
	}
	if values[0] != 1 {
		t.Errorf("Values[0] should be 1")
	}
	if values[1] != 2 {
		t.Errorf("Values[1] should be 2")
	}
	if values[2] != 3 {
		t.Errorf("Values[2] should be 3")
	}

	q.Clear()
	values = q.Values()
	if len(values) != 0 {
		t.Errorf("Values should return an empty slice when the queue is empty")
	}
}

func TestContains(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if !q.Contains(1) {
		t.Errorf("Contains should return true for existing element")
	}

	if q.Contains(4) {
		t.Errorf("Contains should return false for non-existing element")
	}
}

func TestEquals(t *testing.T) {
	q1 := NewQueue[int]()
	q1.Enqueue(1)
	q1.Enqueue(2)
	q1.Enqueue(3)

	q2 := NewQueue[int]()
	q2.Enqueue(1)
	q2.Enqueue(2)
	q2.Enqueue(3)

	if !q1.Equals(q2) {
		t.Errorf("Equals should return true for equal queues")
	}

	q3 := NewQueue[int]()
	q3.Enqueue(1)
	q3.Enqueue(2)

	if q1.Equals(q3) {
		t.Errorf("Equals should return false for queues with different sizes")
	}

	q4 := NewQueue[int]()
	q4.Enqueue(1)
	q4.Enqueue(2)
	q4.Enqueue(4)

	if q1.Equals(q4) {
		t.Errorf("Equals should return false for queues with different elements")
	}
}

func TestCopy(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	copy := q.Copy()

	if !q.Equals(copy) {
		t.Errorf("Copy should create an equal queue")
	}

	item, err := q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue should not return an error")
	}
	if item != 1 {
		t.Errorf("Dequeue should return 1")
	}

	if q.Equals(copy) {
		t.Errorf("Copy should create a separate copy of the queue")
	}
}

func TestString(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	expected := "[1, 2, 3]"
	result := q.String(func(elem int) string {
		return strconv.Itoa(elem)
	})

	if result != expected {
		t.Errorf("String returned incorrect result, got: %s, want: %s", result, expected)
	}

	q.Clear()
	expected = "[]"
	result = q.String(func(elem int) string {
		return strconv.Itoa(elem)
	})

	if result != expected {
		t.Errorf("String returned incorrect result, got: %s, want: %s", result, expected)
	}
}

func TestMap(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Define the mapping function
	f := func(elem int) int {
		return elem * 2
	}

	// Apply the mapping function to the queue
	mappedQueue := q.Map(f)

	// Check the size of the mapped queue
	if mappedQueue.Size() != 3 {
		t.Errorf("Mapped queue should have 3 elements")
	}

	// Check the values of the mapped queue
	values := mappedQueue.Values()
	if values[0] != 2 {
		t.Errorf("Mapped queue should have value 2 at index 0")
	}
	if values[1] != 4 {
		t.Errorf("Mapped queue should have value 4 at index 1")
	}
	if values[2] != 6 {
		t.Errorf("Mapped queue should have value 6 at index 2")
	}
}

func TestFilter(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Define the predicate function
	predicate := func(elem int) bool {
		return elem%2 == 0
	}

	// Apply the filter function to the queue
	q.Filter(predicate)

	// Check the size of the filtered queue
	if q.Size() != 1 {
		t.Errorf("Filtered queue should have 1 element")
	}

	// Check the values of the filtered queue
	item, err := q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue should not return an error")
	}
	if item != 2 {
		t.Errorf("Dequeue should return 2")
	}

	if !q.IsEmpty() {
		t.Errorf("Filtered queue should be empty after dequeuing all elements")
	}

	_, err = q.Dequeue()
	if err == nil {
		t.Errorf("Dequeue should return an error when the queue is empty")
	}
}

func TestReduce(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Define the reduce function
	f := func(a, b int) int {
		return a + b
	}

	// Reduce the queue to a single value
	result := q.Reduce(f, 0)

	// Check the result
	if result != 6 {
		t.Errorf("Reduce returned incorrect result, got: %d, want: %d", result, 6)
	}

	// Reduce an empty queue
	emptyQueue := NewQueue[int]()
	emptyResult := emptyQueue.Reduce(f, 0)

	// Check the result
	if emptyResult != 0 {
		t.Errorf("Reduce on empty queue returned incorrect result, got: %d, want: %d", emptyResult, 0)
	}
}

func TestForEach(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Define the function to be applied to each element
	f := func(elem *int) {
		*elem = *elem * 2
	}

	// Apply the function to each element in the queue
	q.ForEach(f)

	// Check the values of the queue after applying the function
	values := q.Values()
	if values[0] != 2 {
		t.Errorf("ForEach did not apply the function correctly to element at index 0")
	}
	if values[1] != 4 {
		t.Errorf("ForEach did not apply the function correctly to element at index 1")
	}
	if values[2] != 6 {
		t.Errorf("ForEach did not apply the function correctly to element at index 2")
	}
}
