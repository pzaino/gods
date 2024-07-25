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

// Package linkList provides a non-concurrent-safe linked list.
package linkList

import (
	"fmt"
	"testing"
)

const (
	errListNotEmpty        = "Expected list to be empty, but it was not"
	errListIsEmpty         = "Expected list to not be empty, but it was"
	errExpectedIndex       = "Expected index %d, but got %d"
	errExpectedSliceLength = "Expected slice length %d, but got %d"
	errExpectedNoError     = "Expected no error, but got %v"
	errExpectedErr         = "Expected an error, but got nil"
	errExpectedSliceElem   = "Expected slice element %d to be %d, but got %d"
	errExpectedNodeValue   = "Expected node value to be %v, but got %v"
)

func TestNew(t *testing.T) {
	list := NewLinkList[int]()
	if list == nil {
		t.Error("Expected list to be initialized, but got nil")
	}
}

func TestAppend(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list to have 1 item, but got %v", list.Size())
	}
}

func TestPrepend(t *testing.T) {
	list := NewLinkList[int]()
	list.Prepend(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list to have 1 item, but got %v", list.Size())
	}
}

func TestRemove(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestRemoveEmpty(t *testing.T) {
	list := NewLinkList[int]()
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestIsEmpty(t *testing.T) {
	list := NewLinkList[int]()
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestIsEmptyAfterAppend(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
}

func TestIsEmptyAfterPrepend(t *testing.T) {
	list := NewLinkList[int]()
	list.Prepend(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
}

func TestIsEmptyAfterRemove(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
	if list.Size() != 0 {
		t.Errorf("Expected list to have 0 items, but got %v", list.Size())
	}
}

func TestToSlice(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	slice := list.ToSlice()
	expected := []int{1, 2, 3}

	if len(slice) != len(expected) {
		t.Errorf(errExpectedSliceLength, len(expected), len(slice))
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != expected[i] {
			t.Errorf(errExpectedSliceElem, i, expected[i], slice[i])
		}
	}
}

func TestReverse(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Reverse()

	slice := list.ToSlice()
	expected := []int{3, 2, 1}

	if len(slice) != len(expected) {
		t.Errorf(errExpectedSliceLength, len(expected), len(slice))
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != expected[i] {
			t.Errorf(errExpectedSliceElem, i, expected[i], slice[i])
		}
	}
}

func TestReverseEmptyList(t *testing.T) {
	list := NewLinkList[int]()

	list.Reverse()

	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestReverseSingleElementList(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)

	list.Reverse()

	slice := list.ToSlice()
	expected := []int{1}

	if len(slice) != len(expected) {
		t.Errorf(errExpectedSliceLength, len(expected), len(slice))
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != expected[i] {
			t.Errorf(errExpectedSliceElem, i, expected[i], slice[i])
		}
	}
}

func TestGetFirst(t *testing.T) {
	list := NewLinkList[int]()
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

func TestGetLast(t *testing.T) {
	list := NewLinkList[int]()
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
	list := NewLinkList[int]()

	last := list.GetLast()
	if last != nil {
		t.Error("Expected last node to be nil")
	}
}

func TestGetAt(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test getting a valid index
	node, err := list.GetAt(1)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if node.Value != 2 {
		t.Errorf(errExpectedNodeValue, 2, node.Value)
	}

	// Test getting an index out of bounds
	_, err = list.GetAt(3)
	if err == nil {
		t.Error(errExpectedErr)
	}
}

func TestInsertAt(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.InsertAt(1, 4)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}

	slice := list.ToSlice()
	expected := []int{1, 4, 2, 3}

	if len(slice) != len(expected) {
		t.Errorf(errExpectedSliceLength, len(expected), len(slice))
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != expected[i] {
			t.Errorf(errExpectedSliceElem, i, expected[i], slice[i])
		}
	}
}

func TestInsertAtNegativeIndex(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.InsertAt(-1, 4)
	if err == nil {
		t.Error(errExpectedErr)
	}
}

func TestInsertAtOutOfBoundsIndex(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.InsertAt(4, 4)
	if err == nil {
		t.Error(errExpectedErr)
	}
}

func TestInsertAtZeroIndex(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.InsertAt(0, 4)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}

	slice := list.ToSlice()
	expected := []int{4, 1, 2, 3}

	if len(slice) != len(expected) {
		t.Errorf(errExpectedSliceLength, len(expected), len(slice))
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != expected[i] {
			t.Errorf(errExpectedSliceElem, i, expected[i], slice[i])
		}
	}
}

func TestInsertAtEmptyList(t *testing.T) {
	list := NewLinkList[int]()

	err := list.InsertAt(0, 1)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}

	slice := list.ToSlice()
	expected := []int{1}

	if len(slice) != len(expected) {
		t.Errorf(errExpectedSliceLength, len(expected), len(slice))
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != expected[i] {
			t.Errorf(errExpectedSliceElem, i, expected[i], slice[i])
		}
	}
}

func TestDeleteAt(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.DeleteAt(1)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}

	slice := list.ToSlice()
	expected := []int{1, 3}

	if len(slice) != len(expected) {
		t.Errorf(errExpectedSliceLength, len(expected), len(slice))
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != expected[i] {
			t.Errorf(errExpectedSliceElem, i, expected[i], slice[i])
		}
	}
}

func TestDeleteAtNegativeIndex(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.DeleteAt(-1)
	if err == nil {
		t.Error(errExpectedErr)
	}
}

func TestDeleteAtOutOfBoundsIndex(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.DeleteAt(3)
	if err == nil {
		t.Error(errExpectedErr)
	}
}

func TestDeleteAtZeroIndex(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.DeleteAt(0)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}

	slice := list.ToSlice()
	expected := []int{2, 3}

	if len(slice) != len(expected) {
		t.Errorf(errExpectedSliceLength, len(expected), len(slice))
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != expected[i] {
			t.Errorf(errExpectedSliceElem, i, expected[i], slice[i])
		}
	}
}

func TestDeleteAtEmptyList(t *testing.T) {
	list := NewLinkList[int]()

	err := list.DeleteAt(0)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestClear(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Clear()

	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
	if list.Size() != 0 {
		t.Errorf("Expected list to have 0 items, but got %v", list.Size())
	}
}

func TestCopy(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	copyList := list.Copy()

	// Check if the copy is a different instance
	if list == copyList {
		t.Error("Expected the copy to be a different instance")
	}

	// Check if the copy has the same values
	listSlice := list.ToSlice()
	copySlice := copyList.ToSlice()

	if len(listSlice) != len(copySlice) {
		t.Errorf("Expected the copy to have %d items, but got %d", len(listSlice), len(copySlice))
	}

	for i := 0; i < len(listSlice); i++ {
		if listSlice[i] != copySlice[i] {
			t.Errorf("Expected the copy element at index %d to be %d, but got %d", i, listSlice[i], copySlice[i])
		}
	}

	// Check if modifying the copy doesn't affect the original list
	copyList.Append(4)

	if list.Size() != 3 {
		t.Errorf("Expected the original list to have 3 items, but got %d", list.Size())
	}
}

func TestMerge(t *testing.T) {
	list1 := NewLinkList[int]()
	list1.Append(1)
	list1.Append(2)
	list1.Append(3)

	list2 := NewLinkList[int]()
	list2.Append(4)
	list2.Append(5)
	list2.Append(6)

	list1.Merge(list2)

	expected := []int{1, 2, 3, 4, 5, 6}
	slice := list1.ToSlice()

	if len(slice) != len(expected) {
		t.Errorf(errExpectedSliceLength, len(expected), len(slice))
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != expected[i] {
			t.Errorf(errExpectedSliceElem, i, expected[i], slice[i])
		}
	}

	// Check if the original list2 is empty
	if !list2.IsEmpty() {
		t.Error("Expected list2 to be empty after merge")
	}
}

func TestMap(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test mapping the values to their squares
	newList := list.Map(func(value int) int {
		return value * value
	})

	slice := newList.ToSlice()
	expected := []int{1, 4, 9}

	if len(slice) != len(expected) {
		t.Errorf(errExpectedSliceLength, len(expected), len(slice))
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != expected[i] {
			t.Errorf(errExpectedSliceElem, i, expected[i], slice[i])
		}
	}
}

func TestMapEmptyList(t *testing.T) {
	list := NewLinkList[int]()

	// Test mapping an empty list
	list.Map(func(value int) int {
		return value * value
	})

	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestFilter(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Filter out even numbers
	list.Filter(func(value int) bool {
		return value%2 != 0
	})

	slice := list.ToSlice()
	expected := []int{1, 3, 5}

	if len(slice) != len(expected) {
		t.Errorf(errExpectedSliceLength, len(expected), len(slice))
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != expected[i] {
			t.Errorf(errExpectedSliceElem, i, expected[i], slice[i])
		}
	}
}

func TestFilterCleanList(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(3)
	list.Append(5)

	// Filter out the full list (basically return an empty list)
	list.Filter(func(value int) bool {
		return false
	})

	slice := list.ToSlice()

	if len(slice) != 0 {
		for i := 0; i < len(slice); i++ {
			fmt.Printf(" element: %v\n", slice[i])
		}
		t.Errorf(errExpectedSliceLength, 0, len(slice))
	}
}

func TestFilterEmptyList(t *testing.T) {
	list := NewLinkList[int]()

	// Filter on an empty list
	list.Filter(func(value int) bool {
		return value%2 != 0
	})

	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestReduce(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test reducing the list with addition
	sum := list.Reduce(func(a, b int) int {
		return a + b
	}, 0)
	if sum != 6 {
		t.Errorf("Expected sum to be 6, but got %d", sum)
	}

	// Test reducing the list with multiplication
	product := list.Reduce(func(a, b int) int {
		return a * b
	}, 1)
	if product != 6 {
		t.Errorf("Expected product to be 6, but got %d", product)
	}

	// Test reducing an empty list
	emptyList := NewLinkList[int]()
	result := emptyList.Reduce(func(a, b int) int {
		return a + b
	}, 0)
	if result != 0 {
		t.Errorf("Expected result to be 0, but got %d", result)
	}
}

func TestForEach(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test function that prints the values
	var result []int
	list.ForEach(func(value *int) {
		result = append(result, *value)
	})

	expected := []int{1, 2, 3}
	if len(result) != len(expected) {
		t.Errorf("Expected result length %d, but got %d", len(expected), len(result))
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected result element %d to be %d, but got %d", i, expected[i], result[i])
		}
	}

	// Test function that doubles the values
	list.ForEach(func(value *int) {
		*value *= 2
	})

	result = list.ToSlice()
	expected = []int{2, 4, 6}
	if len(result) != len(expected) {
		t.Errorf("Expected list length %d, but got %d", len(expected), len(result))
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected list element %d to be %d, but got %d", i, expected[i], result[i])
		}
	}
}

func TestAny(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test when the predicate returns true for at least one node
	any := list.Any(func(value int) bool {
		return value > 2
	})
	if !any {
		t.Error("Expected Any to return true, but got false")
	}

	// Test when the predicate returns false for all nodes
	any = list.Any(func(value int) bool {
		return value > 3
	})
	if any {
		t.Error("Expected Any to return false, but got true")
	}

	// Test with an empty list
	emptyList := NewLinkList[int]()
	any = emptyList.Any(func(value int) bool {
		return value > 0
	})
	if any {
		t.Error("Expected Any to return false for an empty list, but got true")
	}
}

func TestAll(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(-2)
	list.Append(4)
	list.Append(6)

	// Test when all elements satisfy the predicate
	allEven := list.All(func(value int) bool {
		return value%2 == 0
	})
	if !allEven {
		t.Error("Expected all elements to be even")
	}

	// Test when not all elements satisfy the predicate
	allPositive := list.All(func(value int) bool {
		return value > 0
	})
	if allPositive {
		t.Error("Expected not all elements to be positive")
	}

	// Test with an empty list
	emptyList := NewLinkList[int]()
	allEmpty := emptyList.All(func(value int) bool {
		return value == 0
	})
	if allEmpty {
		t.Error("Expected no elements to be checked in an empty list")
	}
}

func TestContains(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test when the list contains the value
	if !list.Contains(2) {
		t.Error("Expected list to contain the value 2")
	}

	// Test when the list does not contain the value
	if list.Contains(4) {
		t.Error("Expected list to not contain the value 4")
	}

	// Test when the list is empty
	emptyList := NewLinkList[int]()
	if emptyList.Contains(1) {
		t.Error("Expected empty list to not contain any value")
	}
}

func TestIndexOf(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(2)

	// Test finding an existing value
	index := list.IndexOf(2)
	if index != 1 {
		t.Errorf("Expected index to be 1, but got %d", index)
	}

	// Test finding a non-existing value
	index = list.IndexOf(4)
	if index != -1 {
		t.Errorf("Expected index to be -1, but got %d", index)
	}
}

func TestIndexOfEmptyList(t *testing.T) {
	list := NewLinkList[int]()

	// Test finding a value in an empty list
	index := list.IndexOf(1)
	if index != -1 {
		t.Errorf("Expected index to be -1, but got %d", index)
	}
}

func TestLastIndexOf(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(2)
	list.Append(4)

	index := list.LastIndexOf(2)
	expected := 3

	if index != expected {
		t.Errorf(errExpectedIndex, expected, index)
	}

	index = list.LastIndexOf(5)
	expected = -1

	if index != expected {
		t.Errorf(errExpectedIndex, expected, index)
	}
}

func TestLastIndexOfEmptyList(t *testing.T) {
	list := NewLinkList[int]()

	index := list.LastIndexOf(1)
	expected := -1

	if index != expected {
		t.Errorf(errExpectedIndex, expected, index)
	}
}

func TestFindIndex(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Test finding an existing value
	index := list.FindIndex(func(value int) bool {
		return value == 3
	})
	if index != 2 {
		t.Errorf("Expected index 2, but got %d", index)
	}

	// Test finding a non-existing value
	index = list.FindIndex(func(value int) bool {
		return value == 6
	})
	if index != -1 {
		t.Errorf("Expected index -1, but got %d", index)
	}
}

func TestFindLastIndex(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(2)
	list.Append(4)

	// Test finding the last index of a value that exists in the list
	index := list.FindLastIndex(func(value int) bool {
		return value == 2
	})
	if index != 3 {
		t.Errorf("Expected last index of value 2 to be 3, but got %d", index)
	}

	// Test finding the last index of a value that doesn't exist in the list
	index = list.FindLastIndex(func(value int) bool {
		return value == 5
	})
	if index != -1 {
		t.Errorf("Expected last index of value 5 to be -1, but got %d", index)
	}
}

func TestFindAll(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Test finding all even numbers
	evenList := list.FindAll(func(value int) bool {
		return value%2 == 0
	})

	expectedEven := []int{2, 4}
	evenSlice := evenList.ToSlice()

	if len(evenSlice) != len(expectedEven) {
		t.Errorf("Expected even slice length %d, but got %d", len(expectedEven), len(evenSlice))
	}

	for i := 0; i < len(evenSlice); i++ {
		if evenSlice[i] != expectedEven[i] {
			t.Errorf("Expected even slice element %d to be %d, but got %d", i, expectedEven[i], evenSlice[i])
		}
	}

	// Test finding all odd numbers
	oddList := list.FindAll(func(value int) bool {
		return value%2 != 0
	})

	expectedOdd := []int{1, 3, 5}
	oddSlice := oddList.ToSlice()

	if len(oddSlice) != len(expectedOdd) {
		t.Errorf("Expected odd slice length %d, but got %d", len(expectedOdd), len(oddSlice))
	}

	for i := 0; i < len(oddSlice); i++ {
		if oddSlice[i] != expectedOdd[i] {
			t.Errorf("Expected odd slice element %d to be %d, but got %d", i, expectedOdd[i], oddSlice[i])
		}
	}
}

func TestFindLast(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test finding a value that exists
	node, err := list.FindLast(func(value int) bool {
		return value == 2
	})
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if node == nil {
		t.Error("Expected node to be non-nil")
	} else if node.Value != 2 {
		t.Errorf(errExpectedNodeValue, 2, node.Value)
	}

	// Test finding a value that doesn't exist
	_, err = list.FindLast(func(value int) bool {
		return value == 4
	})
	if err == nil {
		t.Error(errExpectedErr)
	}
}

func TestFindAllIndexes(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(2)
	list.Append(4)
	list.Append(2)

	// Test finding indexes of all occurrences of 2
	indexes := list.FindAllIndexes(func(value int) bool {
		return value == 2
	})

	expected := []int{1, 3, 5}

	if len(indexes) != len(expected) {
		t.Errorf("Expected indexes length %d, but got %d", len(expected), len(indexes))
	}

	for i := 0; i < len(indexes); i++ {
		if indexes[i] != expected[i] {
			t.Errorf("Expected index %d to be %d, but got %d", i, expected[i], indexes[i])
		}
	}

	// Test finding indexes of all occurrences of 5 (not present in the list)
	indexes = list.FindAllIndexes(func(value int) bool {
		return value == 5
	})

	expected = []int{}

	if len(indexes) != len(expected) {
		t.Errorf("Expected indexes length %d, but got %d", len(expected), len(indexes))
	}

	for i := 0; i < len(indexes); i++ {
		if indexes[i] != expected[i] {
			t.Errorf("Expected index %d to be %d, but got %d", i, expected[i], indexes[i])
		}
	}
}

func TestFind(t *testing.T) {
	list := NewLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test finding a value that exists in the list
	node, err := list.Find(2)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if node.Value != 2 {
		t.Errorf(errExpectedNodeValue, 2, node.Value)
	}

	// Test finding a value that doesn't exist in the list
	_, err = list.Find(4)
	if err == nil {
		t.Error(errExpectedErr)
	}
}

func TestFindEmptyList(t *testing.T) {
	list := NewLinkList[int]()

	// Test finding a value in an empty list
	_, err := list.Find(1)
	if err == nil {
		t.Error(errExpectedErr)
	}
}
