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

// Package pqueue provides a non-concurrent-safe , max-heap, priority queue.
package pqueue_test

import (
	"fmt"
	"testing"

	"github.com/pzaino/gods/pkg/pqueue"
)

func TestNew(t *testing.T) {
	pq := pqueue.New[int]()
	if pq == nil {
		t.Fatal("Expected new priority queue to be non-nil")
	}
	if !pq.IsEmpty() {
		t.Fatal("Expected new priority queue to be empty")
	}
}

func TestEnqueueAndDequeue(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	if pq.Size() != 2 {
		t.Fatal("Expected priority queue size to be 2")
	}

	val, err := pq.Dequeue()
	if err != nil {
		t.Fatal(err)
	}
	if val != 20 {
		t.Fatalf("Expected dequeued value to be 20, got %d", val)
	}

	val, err = pq.Dequeue()
	if err != nil {
		t.Fatal(err)
	}
	if val != 10 {
		t.Fatalf("Expected dequeued value to be 10, got %d", val)
	}

	if !pq.IsEmpty() {
		t.Fatal("Expected priority queue to be empty after dequeueing all elements")
	}
}

func TestPeek(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	val, err := pq.Peek()
	if err != nil {
		t.Fatal(err)
	}
	if val != 10 {
		t.Fatalf("Expected peek value to be 10, got %d", val)
	}
}

func TestClear(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Clear()
	if !pq.IsEmpty() {
		t.Fatal("Expected priority queue to be empty after clear")
	}
}

func TestContains(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	if !pq.Contains(10) {
		t.Fatal("Expected priority queue to contain value 10")
	}
	if pq.Contains(20) {
		t.Fatal("Expected priority queue to not contain value 20")
	}
}

func TestEquals(t *testing.T) {
	pq1 := pqueue.New[int]()
	pq2 := pqueue.New[int]()
	pq1.Enqueue(10, 1)
	pq2.Enqueue(10, 1)
	if !pq1.Equals(pq2) {
		t.Fatal("Expected priority queues to be equal")
	}
	pq2.Enqueue(20, 2)
	if pq1.Equals(pq2) {
		t.Fatal("Expected priority queues to be not equal")
	}
}

func TestCopy(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	copy := pq.Copy()
	if !copy.Equals(pq) {
		t.Fatal("Expected copy to be equal to original")
	}
	copy.Enqueue(20, 2)
	if copy.Equals(pq) {
		t.Fatal("Expected copy to not be equal to original after modification")
	}
}

func TestString(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	str := pq.String(func(val int) string {
		return fmt.Sprintf("%d", val)
	})
	expected := "[10]"
	if str != expected {
		t.Fatalf("Expected string representation to be %s, got %s", expected, str)
	}
}

func TestMap(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	mapped := pq.Map(func(val int) int {
		return val * 2
	})
	expectedValues := []int{40, 20}
	for i, val := range mapped.Values() {
		if val != expectedValues[i] {
			t.Fatalf("Expected mapped value %d, got %d", expectedValues[i], val)
		}
	}
}

func TestFilter(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.Filter(func(val int) bool {
		return val > 10
	})
	if pq.Size() != 1 || pq.Contains(10) {
		t.Fatal("Expected priority queue to contain only values greater than 10")
	}
}

func TestReduce(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	result := pq.Reduce(func(a, b int) int {
		return a + b
	}, 0)
	if result != 30 {
		t.Fatalf("Expected reduced result to be 30, got %d", result)
	}
}

func TestForEach(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	err := pq.ForEach(func(val *int) error {
		*val += 5
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	expectedValues := []int{25, 15}
	for i, val := range pq.Values() {
		if val != expectedValues[i] {
			t.Fatalf("Expected for each modified value %d, got %d", expectedValues[i], val)
		}
	}
}

func TestAny(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	if !pq.Any(func(val int) bool { return val > 15 }) {
		t.Fatal("Expected any to return true for values greater than 15")
	}
	if pq.Any(func(val int) bool { return val > 25 }) {
		t.Fatal("Expected any to return false for values greater than 25")
	}
}

func TestAll(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	if !pq.All(func(val int) bool { return val > 5 }) {
		t.Fatal("Expected all to return true for values greater than 5")
	}
	if pq.All(func(val int) bool { return val > 15 }) {
		t.Fatal("Expected all to return false for values greater than 15")
	}
}

func TestIndexOf(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	index, err := pq.IndexOf(10)
	if err != nil {
		t.Fatal(err)
	}
	if index != 1 {
		t.Fatalf("Expected index of 10 to be 0, got %d", index)
	}

	_, err = pq.IndexOf(30)
	if err == nil {
		t.Fatal("Expected index of non-existing value to return error")
	}
}

func TestLastIndexOf(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.Enqueue(10, 3)
	index, err := pq.LastIndexOf(10)
	if err != nil {
		t.Fatal(err)
	}
	if index != 1 {
		t.Fatalf("Expected last index of 10 to be 0, got %d", index)
	}
}

func TestFindIndex(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	index, err := pq.FindIndex(func(val int) bool { return val > 15 })
	if err != nil {
		t.Fatal(err)
	}
	if index != 0 {
		t.Fatalf("Expected find index for value greater than 15 to be 0, got %d", index)
	}

	_, err = pq.FindIndex(func(val int) bool { return val > 25 })
	if err == nil {
		t.Fatal("Expected find index for non-existing value to return error")
	}
}

func TestFindLastIndex(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.Enqueue(10, 3)
	index, err := pq.FindLastIndex(func(val int) bool { return val == 10 })
	if err != nil {
		t.Fatal(err)
	}
	if index != 1 {
		t.Fatalf("Expected find last index for value 10 to be 0, got %d", index)
	}
}

func TestFindAll(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.Enqueue(10, 3)
	result := pq.FindAll(func(val int) bool { return val == 10 })
	if result.Size() != 2 {
		t.Fatalf("Expected find all to return 2 elements, got %d", result.Size())
	}
}

func TestFindLast(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	val, err := pq.FindLast(func(val int) bool { return val == 10 })
	if err != nil {
		t.Fatal(err)
	}
	if val != 10 {
		t.Fatalf("Expected find last value to be 10, got %d", val)
	}
	_, err = pq.FindLast(func(val int) bool { return val == 30 })
	if err == nil {
		t.Fatal("Expected find last to return error for non-existing value")
	}
}

func TestFindAllIndexes(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.Enqueue(10, 3)
	indexes := pq.FindAllIndexes(func(val int) bool { return val == 10 })
	if len(indexes) != 2 {
		t.Fatalf("Expected find all indexes to return 2 elements, got %d", len(indexes))
	}
}

func TestValues(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	values := pq.Values()
	expectedValues := []int{20, 10}
	for i, val := range values {
		if val != expectedValues[i] {
			t.Fatalf("Expected value %d, got %d", expectedValues[i], val)
		}
	}
}

func TestDequeueAll(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.Enqueue(30, 3)

	values, err := pq.DequeueAll()
	if err != nil {
		t.Fatal(err)
	}
	expectedValues := []int{30, 20, 10}
	if len(values) != len(expectedValues) {
		t.Fatalf("Expected dequeued values length to be %d, got %d", len(expectedValues), len(values))
	}
	for i, val := range values {
		if val != expectedValues[i] {
			t.Fatalf("Expected dequeued value at index %d to be %d, got %d", i, expectedValues[i], val)
		}
	}

	if !pq.IsEmpty() {
		t.Fatal("Expected priority queue to be empty after dequeuing all elements")
	}
}

func TestDequeueN(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.Enqueue(30, 3)

	values, err := pq.DequeueN(2)
	if err != nil {
		t.Fatal(err)
	}
	expectedValues := []int{30, 20}
	if len(values) != len(expectedValues) {
		t.Fatalf("Expected dequeued values length to be %d, got %d", len(expectedValues), len(values))
	}
	for i, val := range values {
		if val != expectedValues[i] {
			t.Fatalf("Expected dequeued value at index %d to be %d, got %d", i, expectedValues[i], val)
		}
	}

	if pq.Size() != 1 || !pq.Contains(10) {
		t.Fatal("Expected priority queue to contain only value 10 after dequeuing 2 elements")
	}

	_, err = pq.DequeueN(2)
	if err == nil {
		t.Fatal("Expected error when dequeuing more elements than available")
	}
}

func TestUpdatePriority(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.Enqueue(30, 3)

	err := pq.UpdatePriority(20, 4)
	if err != nil {
		t.Fatal(err)
	}

	if pq.Size() != 3 {
		t.Fatal("Expected priority queue size to be 3")
	}

	val, err := pq.Dequeue()
	if err != nil {
		t.Fatal(err)
	}
	if val != 20 {
		t.Fatalf("Expected dequeued value to be 30, got %d", val)
	}

	val, err = pq.Dequeue()
	if err != nil {
		t.Fatal(err)
	}
	if val != 30 {
		t.Fatalf("Expected dequeued value to be 20, got %d", val)
	}

	val, err = pq.Dequeue()
	if err != nil {
		t.Fatal(err)
	}
	if val != 10 {
		t.Fatalf("Expected dequeued value to be 10, got %d", val)
	}
}

func TestUpdateValue(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	err := pq.UpdateValue(10, 30)
	if err != nil {
		t.Fatal(err)
	}
	if !pq.Contains(30) {
		t.Fatal("Expected priority queue to contain value 30 after updating value")
	}
	if pq.Contains(10) {
		t.Fatal("Expected priority queue to not contain value 10 after updating value")
	}
	err = pq.UpdateValue(40, 50)
	if err == nil {
		t.Fatal("Expected error when updating non-existing value")
	}
}

func TestMerge(t *testing.T) {
	pq1 := pqueue.New[int]()
	pq1.Enqueue(10, 1)
	pq1.Enqueue(20, 2)

	pq2 := pqueue.New[int]()
	pq2.Enqueue(30, 3)
	pq2.Enqueue(40, 4)

	pq1.Merge(pq2)

	expectedValues := []int{40, 30, 20, 10}
	for i, val := range pq1.Values() {
		if val != expectedValues[i] {
			t.Fatalf("Expected merged value %d, got %d", expectedValues[i], val)
		}
	}

	if pq2.Size() != 0 {
		t.Fatal("Expected merged priority queue to be empty")
	}
}

func TestCheckSize(t *testing.T) {
	pq := pqueue.New[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.Enqueue(30, 3)

	pq.CheckSize()

	if pq.Size() != 3 {
		t.Fatal("Expected priority queue size to be 3 after calling CheckSize")
	}
}
