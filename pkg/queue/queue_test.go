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
package queue_test

import (
	"strconv"
	"testing"

	queue "github.com/pzaino/gods/pkg/queue"
)

const (
	errExpectedQueueEmpty = "expected queue to be empty"
	errExpectedNoError    = "expected no error, got %v"
	errPeekNoError        = "Peek should not return an error"
	errDeqShouldReturn    = "Dequeue should return %d"
)

func TestQueue(t *testing.T) {
	q := queue.New[int]()
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
		t.Errorf(errPeekNoError)
	}
	if item != 1 {
		t.Errorf("Peek should return 1")
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if item != 1 {
		t.Errorf(errDeqShouldReturn, 1)
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
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item, err := q.Dequeue()
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if item != 1 {
		t.Errorf(errDeqShouldReturn, 1)
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if item != 2 {
		t.Errorf(errDeqShouldReturn, 2)
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if item != 3 {
		t.Errorf(errDeqShouldReturn, 3)
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
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item, err := q.Peek()
	if err != nil {
		t.Errorf(errPeekNoError)
	}
	if item != 1 {
		t.Errorf("Peek should return 1")
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if item != 1 {
		t.Errorf(errDeqShouldReturn, 1)
	}

	item, err = q.Peek()
	if err != nil {
		t.Errorf(errPeekNoError)
	}
	if item != 2 {
		t.Errorf("Peek should return 2")
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if item != 2 {
		t.Errorf(errDeqShouldReturn, 2)
	}

	item, err = q.Peek()
	if err != nil {
		t.Errorf(errPeekNoError)
	}
	if item != 3 {
		t.Errorf("Peek should return 3")
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if item != 3 {
		t.Errorf(errDeqShouldReturn, 3)
	}

	_, err = q.Peek()
	if err == nil {
		t.Errorf("Peek should return an error when the queue is empty")
	}
}

func TestValues(t *testing.T) {
	q := queue.New[int]()
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
	q := queue.New[int]()
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
	q1 := queue.New[int]()
	q1.Enqueue(1)
	q1.Enqueue(2)
	q1.Enqueue(3)

	q2 := queue.New[int]()
	q2.Enqueue(1)
	q2.Enqueue(2)
	q2.Enqueue(3)

	if !q1.Equals(q2) {
		t.Errorf("Equals should return true for equal queues")
	}

	q3 := queue.New[int]()
	q3.Enqueue(1)
	q3.Enqueue(2)

	if q1.Equals(q3) {
		t.Errorf("Equals should return false for queues with different sizes")
	}

	q4 := queue.New[int]()
	q4.Enqueue(1)
	q4.Enqueue(2)
	q4.Enqueue(4)

	if q1.Equals(q4) {
		t.Errorf("Equals should return false for queues with different elements")
	}
}

func TestCopy(t *testing.T) {
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	copy := q.Copy()

	if !q.Equals(copy) {
		t.Errorf("Copy should create an equal queue")
	}

	item, err := q.Dequeue()
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if item != 1 {
		t.Errorf(errDeqShouldReturn, 1)
	}

	if q.Equals(copy) {
		t.Errorf("Copy should create a separate copy of the queue")
	}
}

func TestString(t *testing.T) {
	q := queue.New[int]()
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
	q := queue.New[int]()
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
	q := queue.New[int]()
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
		t.Errorf(errExpectedNoError, err)
	}
	if item != 2 {
		t.Errorf(errDeqShouldReturn, 2)
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
	q := queue.New[int]()
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
	emptyQueue := queue.New[int]()
	emptyResult := emptyQueue.Reduce(f, 0)

	// Check the result
	if emptyResult != 0 {
		t.Errorf("Reduce on empty queue returned incorrect result, got: %d, want: %d", emptyResult, 0)
	}
}

func TestForEach(t *testing.T) {
	q := queue.New[int]()
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

func TestAny(t *testing.T) {
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Define the predicate function
	predicate := func(elem int) bool {
		return elem%2 == 0
	}

	// Check if any element in the queue matches the predicate
	if !q.Any(predicate) {
		t.Errorf("Any should return true for at least one element that matches the predicate")
	}

	// Define a predicate that doesn't match any element
	nonMatchingPredicate := func(elem int) bool {
		return elem > 10
	}

	// Check if any element in the queue matches the non-matching predicate
	if q.Any(nonMatchingPredicate) {
		t.Errorf("Any should return false for all elements that don't match the predicate")
	}

	// Check if any element in an empty queue matches the predicate
	emptyQueue := queue.New[int]()
	if emptyQueue.Any(predicate) {
		t.Errorf("Any should return false for an empty queue")
	}
}

func TestAll(t *testing.T) {
	q := queue.New[int]()
	q.Enqueue(2)
	q.Enqueue(4)
	q.Enqueue(6)

	// Define the predicate function
	predicate := func(elem int) bool {
		return elem%2 == 0
	}

	// Check if all elements in the queue match the predicate
	if !q.All(predicate) {
		t.Errorf("All should return true for all elements that match the predicate")
	}

	// Define a predicate that doesn't match all elements
	nonMatchingPredicate := func(elem int) bool {
		return elem > 4
	}

	// Check if all elements in the queue match the non-matching predicate
	if q.All(nonMatchingPredicate) {
		t.Errorf("All should return false for at least one element that doesn't match the predicate")
	}

	// Check if all elements in an empty queue match the predicate
	emptyQueue := queue.New[int]()
	if emptyQueue.All(predicate) {
		t.Errorf("All should return true for an empty queue")
	}
}

func TestIndexOf(t *testing.T) {
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	index, err := q.IndexOf(2)
	if err != nil {
		t.Errorf("IndexOf returned an unexpected error: %v", err)
	}
	if index != 1 {
		t.Errorf("IndexOf returned incorrect index, got: %d, want: %d", index, 1)
	}

	_, err = q.IndexOf(4)
	if err == nil {
		t.Errorf("IndexOf should return an error for a value not found in the queue")
	}

	emptyQueue := queue.New[int]()
	_, err = emptyQueue.IndexOf(1)
	if err == nil {
		t.Errorf("IndexOf should return an error for an empty queue")
	}
}

func TestLastIndexOf(t *testing.T) {
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(2)

	// Test for an existing value
	index, err := q.LastIndexOf(2)
	if err != nil {
		t.Errorf("LastIndexOf returned an unexpected error: %v", err)
	}
	if index != 3 {
		t.Errorf("LastIndexOf returned incorrect index, got: %d, want: %d", index, 3)
	}

	// Test for a non-existing value
	_, err = q.LastIndexOf(4)
	if err == nil {
		t.Errorf("LastIndexOf should return an error for a non-existing value")
	}

	// Test for an empty queue
	emptyQueue := queue.New[int]()
	_, err = emptyQueue.LastIndexOf(1)
	if err == nil {
		t.Errorf("LastIndexOf should return an error for an empty queue")
	}
}

func TestFindIndex(t *testing.T) {
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Define the predicate function
	predicate := func(elem int) bool {
		return elem%2 == 0
	}

	// Find the index of the first even element
	index, err := q.FindIndex(predicate)
	if err != nil {
		t.Errorf("FindIndex returned an error: %v", err)
	}
	if index != 1 {
		t.Errorf("FindIndex returned incorrect index, got: %d, want: %d", index, 1)
	}

	// Define a predicate that doesn't match any element
	nonMatchingPredicate := func(elem int) bool {
		return elem > 10
	}

	// Find the index of the first element that doesn't match the non-matching predicate
	_, err = q.FindIndex(nonMatchingPredicate)
	if err == nil {
		t.Errorf("FindIndex should return an error when no element matches the predicate")
	}

	// Find the index of the first element in an empty queue
	emptyQueue := queue.New[int]()
	_, err = emptyQueue.FindIndex(predicate)
	if err == nil {
		t.Errorf("FindIndex should return an error when the queue is empty")
	}
}

func TestFindLastIndex(t *testing.T) {
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Define the predicate function
	predicate := func(elem int) bool {
		return elem%2 == 0
	}

	// Find the last index of an even element
	index, err := q.FindLastIndex(predicate)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if index != 1 {
		t.Errorf("FindLastIndex should return index 1")
	}

	// Define a predicate that doesn't match any element
	nonMatchingPredicate := func(elem int) bool {
		return elem > 10
	}

	// Find the last index of an element that doesn't exist
	_, err = q.FindLastIndex(nonMatchingPredicate)
	if err == nil {
		t.Errorf("FindLastIndex should return an error when the element is not found")
	}

	// Find the last index of an element in an empty queue
	emptyQueue := queue.New[int]()
	_, err = emptyQueue.FindLastIndex(predicate)
	if err == nil {
		t.Errorf("FindLastIndex should return an error when the queue is empty")
	}
}

func TestFindAll(t *testing.T) {
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	// Define the predicate function
	predicate := func(elem int) bool {
		return elem%2 == 0
	}

	// Find all even elements in the queue
	result := q.FindAll(predicate)

	// Check the size of the result queue
	if result.Size() != 2 {
		t.Errorf("FindAll should return a queue with 2 elements")
	}

	// Check the values of the result queue
	values := result.Values()
	if values[0] != 2 {
		t.Errorf("FindAll should return a queue with value 2 at index 0")
	}
	if values[1] != 4 {
		t.Errorf("FindAll should return a queue with value 4 at index 1")
	}

	// Find all elements that are greater than 5
	greaterThanFive := func(elem int) bool {
		return elem > 5
	}

	// Find all elements that are greater than 5 in an empty queue
	emptyQueue := queue.New[int]()
	emptyResult := emptyQueue.FindAll(greaterThanFive)

	// Check the size of the result queue
	if emptyResult.Size() != 0 {
		t.Errorf("FindAll on empty queue should return a queue with 0 elements")
	}
}

func TestFindLast(t *testing.T) {
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Define the predicate function
	predicate := func(elem int) bool {
		return elem%2 == 0
	}

	// Find the last element that matches the predicate
	result, err := q.FindLast(predicate)
	if err != nil {
		t.Errorf("FindLast returned an error: %v", err)
	}
	if result != 2 {
		t.Errorf("FindLast returned incorrect result, got: %d, want: %d", result, 2)
	}

	// Define a predicate that doesn't match any element
	nonMatchingPredicate := func(elem int) bool {
		return elem > 10
	}

	// Find the last element that matches the non-matching predicate
	_, err = q.FindLast(nonMatchingPredicate)
	if err == nil {
		t.Errorf("FindLast should return an error when no element matches the predicate")
	}

	// Find the last element in an empty queue
	emptyQueue := queue.New[int]()
	_, err = emptyQueue.FindLast(predicate)
	if err == nil {
		t.Errorf("FindLast should return an error when the queue is empty")
	}
}

func TestFindAllIndexes(t *testing.T) {
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Define the predicate function
	predicate := func(elem int) bool {
		return elem%2 == 0
	}

	// Find all indexes of elements that match the predicate
	indexes := q.FindAllIndexes(predicate)

	// Check the number of indexes found
	if len(indexes) != 1 {
		t.Errorf("FindAllIndexes should return 1 index")
	}

	// Check the value of the found index
	if indexes[0] != 1 {
		t.Errorf("FindAllIndexes should return index 1")
	}

	// Find all indexes of elements that don't match the predicate
	nonMatchingPredicate := func(elem int) bool {
		return elem > 10
	}
	nonMatchingIndexes := q.FindAllIndexes(nonMatchingPredicate)

	// Check that no indexes are found
	if len(nonMatchingIndexes) != 0 {
		t.Errorf("FindAllIndexes should not return any indexes")
	}

	// Find all indexes of elements in an empty queue
	emptyQueue := queue.New[int]()
	emptyIndexes := emptyQueue.FindAllIndexes(predicate)

	// Check that no indexes are found in an empty queue
	if len(emptyIndexes) != 0 {
		t.Errorf("FindAllIndexes should not return any indexes in an empty queue")
	}
}
