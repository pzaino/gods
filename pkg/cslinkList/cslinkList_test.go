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

// Package cslinkList provides a concurrent-safe linked list.
package cslinkList

import (
	"reflect"
	"sync"
	"testing"
)

const (
	errListNotEmpty = "Expected list to be empty, but it was not"
	errListIsEmpty  = "Expected list to not be empty, but it was"
)

func TestNew(t *testing.T) {
	list := NewCSLinkList[int]()
	if list == nil {
		t.Error("Expected list to be initialized, but got nil")
	}
}

func TestAppend(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list to have 1 item, but got %v", list.Size())
	}
}

func TestPrepend(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Prepend(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list to have 1 item, but got %v", list.Size())
	}
}

func TestRemove(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestRemoveEmpty(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestIsEmpty(t *testing.T) {
	list := NewCSLinkList[int]()
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestIsEmptyAfterAppend(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
}

func TestIsEmptyAfterPrepend(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Prepend(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
}

func TestIsEmptyAfterRemove(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
	if list.Size() != 0 {
		t.Errorf("Expected list to have 0 items, but got %v", list.Size())
	}
}

// Test concurrent operations on the list

func TestConcurrentAppend(t *testing.T) {
	list := NewCSLinkList[int]()
	var wg sync.WaitGroup

	// Append 1000 items concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			list.Append(val)
		}(i)
	}

	wg.Wait()

	if list.Size() != 1000 {
		t.Errorf("Expected list size 1000, got %d", list.Size())
	}
}

func TestConcurrentPrepend(t *testing.T) {
	list := NewCSLinkList[int]()
	var wg sync.WaitGroup

	// Prepend 1000 items concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			list.Prepend(val)
		}(i)
	}

	wg.Wait()

	if list.Size() != 1000 {
		t.Errorf("Expected list size 1000, got %d", list.Size())
	}
}

func TestConcurrentDeleteWithValue(t *testing.T) {
	list := NewCSLinkList[int]()

	// Append 1000 items
	for i := 0; i < 1000; i++ {
		list.Append(i)
	}

	var wg sync.WaitGroup

	// Delete all 1000 items concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			list.DeleteWithValue(val)
		}(i)
	}

	wg.Wait()

	if list.Size() != 0 {
		t.Errorf("Expected list size 0, got %d", list.Size())
	}
}

func TestInsertAt(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(3)

	err := list.InsertAt(1, 2)
	if err != nil {
		t.Error(err)
	}

	if list.Size() != 3 {
		t.Errorf("Expected list size 3, got %d", list.Size())
	}

	if list.ToSlice()[1] != 2 {
		t.Errorf("Expected list to be [1, 2, 3], but got %v", list.ToSlice())
	}
}

func TestInsertAtOutOfBounds(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)

	err := list.InsertAt(3, 3)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	if list.Size() != 2 {
		t.Errorf("Expected list size 2, got %d", list.Size())
	}

	if list.ToSlice()[0] != 1 || list.ToSlice()[1] != 2 {
		t.Errorf("Expected list to be [1, 2], but got %v", list.ToSlice())
	}
}

func TestInsertAtEmptyList(t *testing.T) {
	list := NewCSLinkList[int]()

	err := list.InsertAt(0, 1)
	if err != nil {
		t.Error(err)
	}

	if list.Size() != 1 {
		t.Errorf("Expected list size 1, got %d", list.Size())
	}

	if list.ToSlice()[0] != 1 {
		t.Errorf("Expected list to be [1], but got %v", list.ToSlice())
	}
}

func TestInsertAtComplexObjects(t *testing.T) {
	type ComplexObject struct {
		Name string
		Age  int
	}

	list := NewCSLinkList[ComplexObject]()
	list.Append(ComplexObject{"Alice", 25})
	list.Append(ComplexObject{"Bob", 30})

	err := list.InsertAt(1, ComplexObject{"Charlie", 35})
	if err != nil {
		t.Error(err)
	}

	if list.Size() != 3 {
		t.Errorf("Expected list size 3, got %d", list.Size())
	}

	if list.ToSlice()[1] != (ComplexObject{"Charlie", 35}) {
		t.Errorf("Expected list to be [Alice, Charlie, Bob], but got %v", list.ToSlice())
	}

	err = list.InsertAt(0, ComplexObject{"David", 40})
	if err != nil {
		t.Error(err)
	}

	if list.Size() != 4 {
		t.Errorf("Expected list size 4, got %d", list.Size())
	}

	if list.ToSlice()[0] != (ComplexObject{"David", 40}) {
		t.Errorf("Expected list to be [David, Alice, Charlie, Bob], but got %v", list.ToSlice())
	}
}

func TestConcurrentInsertAt(t *testing.T) {
	list := NewCSLinkList[int]()
	var wg sync.WaitGroup

	// Insert 1000 items at the beginning concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			err := list.InsertAt(0, val)
			if err != nil {
				t.Error(err)
			}
		}(i)
	}

	wg.Wait()

	if list.Size() != 1000 {
		t.Errorf("Expected list size 1000, got %d", list.Size())
	}
}

func TestConcurrentGetAt(t *testing.T) {
	list := NewCSLinkList[int]()

	// Append 1000 items
	for i := 0; i < 1000; i++ {
		list.Append(i)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	errorCount := 0

	// Concurrently get all items and check their values
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			node, err := list.GetAt(index)
			mu.Lock()
			defer mu.Unlock()
			if err != nil || node.Value != index {
				errorCount++
			}
		}(i)
	}

	wg.Wait()

	if errorCount != 0 {
		t.Errorf("Expected 0 errors, got %d", errorCount)
	}
}

func TestConcurrentClear(t *testing.T) {
	list := NewCSLinkList[int]()

	// Append 1000 items
	for i := 0; i < 1000; i++ {
		list.Append(i)
	}

	var wg sync.WaitGroup

	// Concurrently clear the list
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			list.Clear()
		}()
	}

	wg.Wait()

	if list.Size() != 0 {
		t.Errorf("Expected list size 0, got %d", list.Size())
	}
}

func TestConcurrentMerge(t *testing.T) {
	list1 := NewCSLinkList[int]()
	list2 := NewCSLinkList[int]()

	// Append 500 items to each list
	for i := 0; i < 500; i++ {
		list1.Append(i)
		list2.Append(i + 500)
	}

	var wg sync.WaitGroup

	// Concurrently merge list2 into list1
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			list1.Merge(list2)
		}()
	}

	wg.Wait()

	// Size should be at least 1000 if merged correctly multiple times concurrently
	if list1.Size() < 1000 {
		t.Errorf("Expected list size at least 1000, got %d", list1.Size())
	}
}

func TestConcurrentFind(t *testing.T) {
	list := NewCSLinkList[int]()

	// Append 1000 items
	for i := 0; i < 1000; i++ {
		list.Append(i)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	errorCount := 0

	// Concurrently find all items and check their values
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			node, err := list.Find(val)
			mu.Lock()
			defer mu.Unlock()
			if err != nil || node.Value != val {
				errorCount++
			}
		}(i)
	}

	wg.Wait()

	if errorCount != 0 {
		t.Errorf("Expected 0 errors, got %d", errorCount)
	}
}

func TestForEach(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	var sum int
	list.ForEach(func(value *int) {
		sum += *value
	})

	expectedSum := 1 + 2 + 3
	if sum != expectedSum {
		t.Errorf("Expected sum to be %d, but got %d", expectedSum, sum)
	}
}

func TestToSlice(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	slice := list.ToSlice()
	expectedSlice := []int{1, 2, 3}

	if len(slice) != len(expectedSlice) {
		t.Errorf("Expected slice length %d, got %d", len(expectedSlice), len(slice))
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != expectedSlice[i] {
			t.Errorf("Expected slice element %d to be %d, got %d", i, expectedSlice[i], slice[i])
		}
	}
}

func TestReverse(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Reverse()

	// Check if the list is reversed correctly
	if list.ToSlice()[0] != 3 || list.ToSlice()[1] != 2 || list.ToSlice()[2] != 1 {
		t.Errorf("Expected list to be reversed, but got %v", list.ToSlice())
	}
}

func TestReverseEmptyList(t *testing.T) {
	list := NewCSLinkList[int]()

	list.Reverse()

	// Check if reversing an empty list doesn't cause any errors
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestReverseConcurrent(t *testing.T) {
	list := NewCSLinkList[int]()

	const max = 100

	// Append 1000 items
	for i := 0; i < max; i++ {
		list.Append(i)
	}

	var wg sync.WaitGroup

	// Concurrently reverse the list
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			list.Reverse()
		}()
	}

	wg.Wait()

	// Check if the list is reversed correctly
	for i := 0; i < max; i++ {
		item, err := list.GetAt(i)
		if err != nil {
			t.Error(err)
		}
		if item == nil {
			t.Errorf("Expected item not to be nil")
		} else if item.Value != (max-1)-i {
			t.Errorf("Expected list to be reversed, but got %v", list.ToSlice())
		}
	}
}

func TestGetFirst(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	first := list.GetFirst()
	if first == nil {
		t.Error("Expected first node to be non-nil")
	} else if first.Value != 1 {
		t.Errorf("Expected first node value to be 1, but got %v", first.Value)
	}
}

func TestGetFirstEmptyList(t *testing.T) {
	list := NewCSLinkList[int]()

	first := list.GetFirst()
	if first != nil {
		t.Errorf("Expected first node to be nil, but got %v", first)
	}
}

func TestGetLast(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	last := list.GetLast()
	if last == nil {
		t.Error("Expected last node to be non-nil")
	} else if last.Value != 3 {
		t.Errorf("Expected last node value to be 3, but got %v", last.Value)
	}
}

func TestGetLastEmptyList(t *testing.T) {
	list := NewCSLinkList[int]()

	last := list.GetLast()
	if last != nil {
		t.Errorf("Expected last node to be nil, but got %v", last)
	}
}

func TestDeleteAt(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.DeleteAt(1)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if list.Size() != 2 {
		t.Errorf("Expected list size 2, got %d", list.Size())
	}

	if list.ToSlice()[0] != 1 || list.ToSlice()[1] != 3 {
		t.Errorf("Expected list to be [1, 3], but got %v", list.ToSlice())
	}
}

func TestDeleteAtOutOfBounds(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.DeleteAt(3)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	if list.Size() != 3 {
		t.Errorf("Expected list size 3, got %d", list.Size())
	}

	if list.ToSlice()[0] != 1 || list.ToSlice()[1] != 2 || list.ToSlice()[2] != 3 {
		t.Errorf("Expected list to be [1, 2, 3], but got %v", list.ToSlice())
	}
}

func TestDeleteAtEmptyList(t *testing.T) {
	list := NewCSLinkList[int]()

	err := list.DeleteAt(0)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	if list.Size() != 0 {
		t.Errorf("Expected list size 0, got %d", list.Size())
	}
}

func TestDeleteAtConcurrent(t *testing.T) {
	list := NewCSLinkList[int]()

	const max = 100

	// Append 1000 items
	for i := 0; i < max; i++ {
		list.Append(i)
	}

	var wg sync.WaitGroup

	// Concurrently delete all items
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := list.DeleteAt(0)
			if err != nil {
				t.Error(err)
			}
		}()
	}

	wg.Wait()

	if list.Size() != max-3 {
		t.Errorf("Expected list size %d, got %d", max-3, list.Size())
	}
}

func TestCopy(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	copyList := list.Copy()

	// Check if the copied list has the same values as the original list
	if !reflect.DeepEqual(copyList.ToSlice(), list.ToSlice()) {
		t.Errorf("Expected copied list to be %v, but got %v", list.ToSlice(), copyList.ToSlice())
	}

	// Check if modifying the copied list doesn't affect the original list
	copyList.Append(4)
	if reflect.DeepEqual(copyList.ToSlice(), list.ToSlice()) {
		t.Errorf("Expected modifying copied list to not affect the original list")
	}

	// Check if modifying the original list doesn't affect the copied list
	list.Append(5)
	if reflect.DeepEqual(copyList.ToSlice(), list.ToSlice()) {
		t.Errorf("Expected modifying original list to not affect the copied list")
	}
}

func TestCopyEmptyList(t *testing.T) {
	list := NewCSLinkList[int]()

	copyList := list.Copy()

	// Check if the copied list is empty
	if !copyList.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestCopyConcurrent(t *testing.T) {
	list := NewCSLinkList[int]()

	const max = 100

	// Append 1000 items
	for i := 0; i < max; i++ {
		list.Append(i)
	}

	var wg sync.WaitGroup

	// Concurrently copy the list
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			copyList := list.Copy()
			if !reflect.DeepEqual(copyList.ToSlice(), list.ToSlice()) {
				t.Errorf("Expected copied list to be %v, but got %v", list.ToSlice(), copyList.ToSlice())
			}
		}()
	}

	wg.Wait()
}

func TestMap(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Define the mapping function
	mapFunc := func(value int) int {
		return value * 2
	}

	// Apply the mapping function to all nodes
	list.Map(mapFunc)

	// Check if the values have been mapped correctly
	expectedSlice := []int{2, 4, 6}
	slice := list.ToSlice()

	if len(slice) != len(expectedSlice) {
		t.Errorf("Expected slice length %d, got %d", len(expectedSlice), len(slice))
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != expectedSlice[i] {
			t.Errorf("Expected slice element %d to be %d, got %d", i, expectedSlice[i], slice[i])
		}
	}
}

func TestMapEmptyList(t *testing.T) {
	list := NewCSLinkList[int]()

	// Define the mapping function
	mapFunc := func(value int) int {
		return value * 2
	}

	// Apply the mapping function to all nodes
	list.Map(mapFunc)

	// Check if mapping an empty list doesn't cause any errors
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestFilter(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Filter out even numbers
	list.Filter(func(value int) bool {
		return value%2 != 0
	})

	// Check if the list contains only odd numbers
	if list.Size() != 3 {
		t.Errorf("Expected list size 3, got %d", list.Size())
	}

	if !list.Contains(1) || !list.Contains(3) || !list.Contains(5) {
		t.Errorf("Expected list to contain only odd numbers, got %v", list.ToSlice())
	}
}

func TestFilterEmptyList(t *testing.T) {
	list := NewCSLinkList[int]()

	// Filter on an empty list
	list.Filter(func(value int) bool {
		return value > 0
	})

	// Check if the list remains empty
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestFilterAll(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Filter out all numbers
	list.Filter(func(value int) bool {
		return false
	})

	// Check if the list becomes empty
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestFilterNone(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Filter out no numbers
	list.Filter(func(value int) bool {
		return true
	})

	// Check if the list remains unchanged
	if list.Size() != 5 {
		t.Errorf("Expected list size 5, got %d", list.Size())
	}

	expectedSlice := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(list.ToSlice(), expectedSlice) {
		t.Errorf("Expected list to remain unchanged, got %v", list.ToSlice())
	}
}

func TestReduce(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	sum := list.Reduce(func(acc, val int) int {
		return acc + val
	}, 0)

	expectedSum := 1 + 2 + 3
	if sum != expectedSum {
		t.Errorf("Expected sum to be %d, but got %d", expectedSum, sum)
	}
}

func TestReduceEmptyList(t *testing.T) {
	list := NewCSLinkList[int]()

	sum := list.Reduce(func(acc, val int) int {
		return acc + val
	}, 0)

	if sum != 0 {
		t.Errorf("Expected sum to be 0, but got %d", sum)
	}
}

func TestReduceWithInitialValue(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	product := list.Reduce(func(acc, val int) int {
		return acc * val
	}, 1)

	expectedProduct := 1 * 2 * 3
	if product != expectedProduct {
		t.Errorf("Expected product to be %d, but got %d", expectedProduct, product)
	}
}

func TestReduceConcurrent(t *testing.T) {
	list := NewCSLinkList[int]()

	const max = 100

	// Append 1000 items
	for i := 0; i < max; i++ {
		list.Append(i)
	}

	var wg sync.WaitGroup

	// Concurrently reduce the list
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum := list.Reduce(func(acc, val int) int {
				return acc + val
			}, 0)
			if sum != 4950 {
				t.Errorf("Expected sum to be 4950, got %d", sum)
			}
		}()
	}

	wg.Wait()
}

func TestAny(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test with a predicate that returns true for all values
	allTrue := list.Any(func(value int) bool {
		return true
	})
	if !allTrue {
		t.Error("Expected Any to return true for all true predicate")
	}

	// Test with a predicate that returns false for all values
	allFalse := list.Any(func(value int) bool {
		return false
	})
	if allFalse {
		t.Error("Expected Any to return false for all false predicate")
	}

	// Test with a predicate that returns true for some values
	someTrue := list.Any(func(value int) bool {
		return value > 1
	})
	if !someTrue {
		t.Error("Expected Any to return true for some true predicate")
	}

	// Test with an empty list
	emptyList := NewCSLinkList[int]()
	emptyResult := emptyList.Any(func(value int) bool {
		return true
	})
	if emptyResult {
		t.Error("Expected Any to return false for empty list")
	}
}

func TestAll(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test with a predicate that returns true for all values
	allTrue := list.All(func(value int) bool {
		return true
	})
	if !allTrue {
		t.Error("Expected All to return true for all true predicate")
	}

	// Test with a predicate that returns false for all values
	allFalse := list.All(func(value int) bool {
		return false
	})
	if allFalse {
		t.Error("Expected All to return false for all false predicate")
	}

	// Test with a predicate that returns false for some values
	someFalse := list.All(func(value int) bool {
		return value > 1
	})
	if someFalse {
		t.Error("Expected All to return false for some false predicate")
	}

	// Test with an empty list
	emptyList := NewCSLinkList[int]()
	emptyResult := emptyList.All(func(value int) bool {
		return true
	})
	if emptyResult {
		t.Error("Expected All to return false for empty list")
	}
}

func TestContains(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test for existing value
	if !list.Contains(2) {
		t.Error("Expected list to contain value 2")
	}

	// Test for non-existing value
	if list.Contains(4) {
		t.Error("Expected list to not contain value 4")
	}
}

func TestContainsEmptyList(t *testing.T) {
	list := NewCSLinkList[int]()

	// Test for non-existing value in an empty list
	if list.Contains(1) {
		t.Error("Expected empty list to not contain any value")
	}
}

func TestIndexOf(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	index := list.IndexOf(2)
	if index != 1 {
		t.Errorf("Expected index 1, but got %d", index)
	}

	index = list.IndexOf(4)
	if index != -1 {
		t.Errorf("Expected index -1, but got %d", index)
	}
}

func TestIndexOfEmptyList(t *testing.T) {
	list := NewCSLinkList[int]()

	index := list.IndexOf(1)
	if index != -1 {
		t.Errorf("Expected index -1, but got %d", index)
	}
}

func TestIndexOfDuplicateValues(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(2)
	list.Append(3)

	index := list.IndexOf(2)
	if index != 1 {
		t.Errorf("Expected index 1, but got %d", index)
	}

	index = list.IndexOf(3)
	if index != 3 {
		t.Errorf("Expected index 3, but got %d", index)
	}
}

func TestLastIndexOf(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(2)

	index := list.LastIndexOf(2)
	expectedIndex := 3
	if index != expectedIndex {
		t.Errorf("Expected last index of 2 to be %d, but got %d", expectedIndex, index)
	}

	index = list.LastIndexOf(4)
	expectedIndex = -1
	if index != expectedIndex {
		t.Errorf("Expected last index of 4 to be %d, but got %d", expectedIndex, index)
	}
}

func TestLastIndexOfEmptyList(t *testing.T) {
	list := NewCSLinkList[int]()

	index := list.LastIndexOf(1)
	expectedIndex := -1
	if index != expectedIndex {
		t.Errorf("Expected last index of 1 to be %d, but got %d", expectedIndex, index)
	}
}

func TestLastIndexOfConcurrent(t *testing.T) {
	list := NewCSLinkList[int]()

	const max = 100

	// Append 1000 items
	for i := 0; i < max; i++ {
		list.Append(i % 10)
	}

	var wg sync.WaitGroup

	// Concurrently find the last index of 5
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			index := list.LastIndexOf(5)
			if index < 0 || index >= max {
				t.Errorf("Expected last index of 5 to be between 0 and %d, but got %d", max-1, index)
			}
		}()
	}

	wg.Wait()
}

func TestFindIndex(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test case 1: FindIndex returns the correct index for a matching value
	index1 := list.FindIndex(func(value int) bool {
		return value == 2
	})
	if index1 != 1 {
		t.Errorf("Expected index 1, but got %d", index1)
	}

	// Test case 2: FindIndex returns -1 for a non-matching value
	index2 := list.FindIndex(func(value int) bool {
		return value == 4
	})
	if index2 != -1 {
		t.Errorf("Expected index -1, but got %d", index2)
	}

	// Test case 3: FindIndex returns the correct index for the first matching value
	list.Append(2)
	index3 := list.FindIndex(func(value int) bool {
		return value == 2
	})
	if index3 != 1 {
		t.Errorf("Expected index 1, but got %d", index3)
	}

	// Test case 4: FindIndex returns -1 for an empty list
	emptyList := NewCSLinkList[int]()
	index4 := emptyList.FindIndex(func(value int) bool {
		return value == 1
	})
	if index4 != -1 {
		t.Errorf("Expected index -1, but got %d", index4)
	}
}

func TestFindLastIndex(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(2)
	list.Append(4)

	// Test case 1: Find the last index of value 2
	index1 := list.FindLastIndex(func(value int) bool {
		return value == 2
	})
	if index1 != 3 {
		t.Errorf("Expected last index of value 2 to be 3, but got %d", index1)
	}

	// Test case 2: Find the last index of value 5 (not found)
	index2 := list.FindLastIndex(func(value int) bool {
		return value == 5
	})
	if index2 != -1 {
		t.Errorf("Expected last index of value 5 to be -1, but got %d", index2)
	}

	// Test case 3: Find the last index of value 4
	index3 := list.FindLastIndex(func(value int) bool {
		return value == 4
	})
	if index3 != 4 {
		t.Errorf("Expected last index of value 4 to be 4, but got %d", index3)
	}
}

func TestFindAll(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test case 1: FindAll returns a new list with all nodes that match the predicate
	result1 := list.FindAll(func(value int) bool {
		return value%2 == 0
	})
	expectedResult1 := NewCSLinkList[int]()
	expectedResult1.Append(2)
	if !reflect.DeepEqual(result1.ToSlice(), expectedResult1.ToSlice()) {
		t.Errorf("Expected result1 to be %v, but got %v", expectedResult1.ToSlice(), result1.ToSlice())
	}

	// Test case 2: FindAll returns an empty list if no nodes match the predicate
	result2 := list.FindAll(func(value int) bool {
		return value > 3
	})
	expectedResult2 := NewCSLinkList[int]()
	if !reflect.DeepEqual(result2.ToSlice(), expectedResult2.ToSlice()) {
		t.Errorf("Expected result2 to be %v, but got %v", expectedResult2.ToSlice(), result2.ToSlice())
	}

	// Test case 3: FindAll returns a new list with all nodes if all nodes match the predicate
	result3 := list.FindAll(func(value int) bool {
		return value > 0
	})
	expectedResult3 := list.Copy()
	if !reflect.DeepEqual(result3.ToSlice(), expectedResult3.ToSlice()) {
		t.Errorf("Expected result3 to be %v, but got %v", expectedResult3.ToSlice(), result3.ToSlice())
	}
}

func TestFindLast(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test case 1: Find the last node with value 2
	node, err := list.FindLast(func(value int) bool {
		return value == 2
	})
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if node == nil {
		t.Error("Expected node to be non-nil")
	} else if node.Value != 2 {
		t.Errorf("Expected node value to be 2, but got %v", node.Value)
	}

	// Test case 2: Find the last node with value 4 (not found)
	node, err = list.FindLast(func(value int) bool {
		return value == 4
	})
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if node != nil {
		t.Errorf("Expected node to be nil, but got %v", node)
	}
}

func TestFindLastEmptyList(t *testing.T) {
	list := NewCSLinkList[int]()

	// Test case: Find the last node in an empty list
	node, err := list.FindLast(func(value int) bool {
		return value == 1
	})
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if node != nil {
		t.Errorf("Expected node to be nil, but got %v", node)
	}
}

func TestFindAllIndexes(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(2)
	list.Append(4)
	list.Append(2)

	indexes := list.FindAllIndexes(func(value int) bool {
		return value == 2
	})

	expectedIndexes := []int{1, 3, 5}

	if len(indexes) != len(expectedIndexes) {
		t.Errorf("Expected indexes length %d, got %d", len(expectedIndexes), len(indexes))
	}

	for i := 0; i < len(indexes); i++ {
		if indexes[i] != expectedIndexes[i] {
			t.Errorf("Expected index %d to be %d, got %d", i, expectedIndexes[i], indexes[i])
		}
	}
}

func TestFindAllIndexesEmptyList(t *testing.T) {
	list := NewCSLinkList[int]()

	indexes := list.FindAllIndexes(func(value int) bool {
		return value == 2
	})

	if len(indexes) != 0 {
		t.Errorf("Expected indexes length 0, got %d", len(indexes))
	}
}

func TestFindAllIndexesNoMatches(t *testing.T) {
	list := NewCSLinkList[int]()
	list.Append(1)
	list.Append(3)
	list.Append(5)

	indexes := list.FindAllIndexes(func(value int) bool {
		return value == 2
	})

	if len(indexes) != 0 {
		t.Errorf("Expected indexes length 0, got %d", len(indexes))
	}
}
