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

// Package circularLinkList provides a non-concurrent-safe circular linked list.
package circularLinkList_test

import (
	"fmt"
	"testing"

	"github.com/pzaino/gods/pkg/circularLinkList" // Adjust the import path as necessary
)

const (
	errListIsEmpty    = "list is empty"
	errValueNotFound  = "value not found"
	errExpectedLength = "expected length %d, got %d"
	errExpectedValue  = "expected %d, got %d"
	errExpectedNoErr  = "unexpected error: %v"
	errExpectedResult = "expected result %d, got %d"
	errExpectedError  = "expected error %q, got %v"
	errExpectedError2 = "expected error, got nil"
)

func TestAppend(t *testing.T) {
	list := circularLinkList.NewCircularLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	expected := []int{1, 2, 3}
	actual := list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf("expected list length %d, got %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf(errExpectedValue, v, actual[i])
		}
	}
}

func TestPrepend(t *testing.T) {
	list := circularLinkList.NewCircularLinkList[int]()
	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)

	expected := []int{3, 2, 1}
	actual := list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf("expected length %d , got %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf("expected  %d, got %d", v, actual[i])
		}
	}
}

func TestDeleteWithValue(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})
	list.DeleteWithValue(3)

	expected := []int{1, 2, 4}
	actual := list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf("expected length  %d, got %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf("expected %d , got %d", v, actual[i])
		}
	}
}

func TestFind(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})
	node, err := list.Find(3)

	if err != nil {
		t.Fatalf("unexpected  error: %v", err)
	}

	if node == nil || node.Value != 3 {
		t.Fatalf("expected to find node with value 3")
	}
}

func TestReverse(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})
	list.Reverse()

	expected := []int{4, 3, 2, 1}
	actual := list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf("expected length  %d , got %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf("expected value %d , got %d", v, actual[i])
		}
	}
}

func TestSize(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})

	expected := uint64(4)
	actual := list.Size()

	if expected != actual {
		t.Fatalf("expected size %d, got %d", expected, actual)
	}
}

func TestGetFirst(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})
	node := list.GetFirst()

	if node == nil || node.Value != 1 {
		t.Fatalf("expected to get first node with value 1")
	}
}

func TestGetLast(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})
	node := list.GetLast()

	if node == nil || node.Value != 4 {
		t.Fatalf("expected to get last node with value 4")
	}
}

func TestGetAt(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})
	node, err := list.GetAt(2)

	if err != nil {
		t.Fatalf("unexpected error:  %v", err)
	}

	if node == nil || node.Value != 3 {
		t.Fatalf("expected to get node with value 3 at index 2")
	}
}

func TestInsertAt(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 4, 5})
	err := list.InsertAt(2, 3)

	if err != nil {
		t.Fatalf(errExpectedNoErr, err)
	}

	expected := []int{1, 2, 3, 4, 5}
	actual := list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf(errExpectedLength, len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf(errExpectedValue, v, actual[i])
		}
	}
}

func TestDeleteAt(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4, 5})
	err := list.DeleteAt(2)

	if err != nil {
		t.Fatalf(errExpectedNoErr, err)
	}

	expected := []int{1, 2, 4, 5}
	actual := list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf(errExpectedLength, len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf(errExpectedValue, v, actual[i])
		}
	}
}

func TestClear(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4, 5})
	list.Clear()

	expected := uint64(0)
	actual := list.Size()

	if expected != actual {
		t.Fatalf("expected size %d, got %d", expected, actual)
	}
}

func TestIsEmpty(t *testing.T) {
	list := circularLinkList.NewCircularLinkList[int]()

	// Test when the list is empty
	if !list.IsEmpty() {
		t.Fatalf("expected list to be empty")
	}

	// Test when the list is not empty
	list.Append(1)
	if list.IsEmpty() {
		t.Fatalf("expected list not to be empty")
	}
}

func TestCopy(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})
	copyList := list.Copy()

	// Check if the copied list is not the same instance as the original list
	if copyList == list {
		t.Fatalf("expected copied list to be a different instance")
	}

	// Check if the copied list has the same values as the original list
	expected := []int{1, 2, 3, 4}
	actual := copyList.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf(errExpectedLength, len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf(errExpectedValue, v, actual[i])
		}
	}

	// Check if modifying the copied list does not affect the original list
	copyList.Append(5)

	if list.Size() != 4 {
		t.Fatalf("expected original list size to be 4, got %d", list.Size())
	}
}

func TestMerge(t *testing.T) {
	list1 := circularLinkList.NewCircularLinkList[int]()
	list1.Append(1)
	list1.Append(2)
	list1.Append(3)

	list2 := circularLinkList.NewCircularLinkList[int]()
	list2.Append(4)
	list2.Append(5)
	list2.Append(6)

	list1.Merge(list2)

	expected := []int{1, 2, 3, 4, 5, 6}
	actual := list1.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf(errExpectedLength, len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf(errExpectedValue, v, actual[i])
		}
	}

	if list2.Size() != 0 {
		t.Fatalf("expected list2 to be empty after merge")
	}
}

func TestMap(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})
	newList := list.Map(func(value int) int {
		return value * 2
	})

	expected := []int{2, 4, 6, 8}
	actual := newList.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf(errExpectedLength, len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf(errExpectedValue, v, actual[i])
		}
	}
}

func TestMapFrom(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})
	newList, err := list.MapFrom(2, func(value int) int {
		return value * 2
	})

	if err != nil {
		t.Fatalf(errExpectedNoErr, err)
	}

	expected := []int{6, 8}
	actual := newList.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf(errExpectedLength, len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf(errExpectedValue, v, actual[i])
		}
	}
}

func TestMapRange(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4, 5})
	newList, err := list.MapRange(1, 4, func(value int) int {
		return value * 2
	})

	if err != nil {
		t.Fatalf(errExpectedNoErr, err)
	}

	expected := []int{4, 6, 8}
	actual := newList.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf(errExpectedLength, len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf(errExpectedValue, v, actual[i])
		}
	}
}

func TestForEach(t *testing.T) {
	list := circularLinkList.NewCircularLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	sum := 0
	list.ForEach(func(value *int) {
		sum += *value
	})

	expectedSum := 6
	if sum != expectedSum {
		t.Fatalf("expected sum %d, got %d", expectedSum, sum)
	}
}

func TestForRange(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4, 5})

	// Test when the range is within the list size
	err := list.ForRange(1, 4, func(value *int) {
		*value *= 2
	})

	if err != nil {
		t.Fatalf(errExpectedNoErr, err)
	}

	expected := []int{1, 4, 6, 8, 10}
	actual := list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf(errExpectedLength, len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf(errExpectedValue, v, actual[i])
		}
	}

	// Test when the range exceeds the list size
	// The range will be [3, 3] since the list size is 5
	// so, in out case it will affect only the 4th element
	err = list.ForRange(3, 8, func(value *int) {
		*value *= 2
	})

	if err != nil {
		t.Fatalf(errExpectedNoErr, err)
	}

	expected = []int{1, 4, 6, 16, 10}
	actual = list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf(errExpectedLength, len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf(errExpectedValue, v, actual[i])
		}
	}

	// Test when the start index is greater than the end index
	err = list.ForRange(4, 2, func(value *int) {
		*value *= 2
	})

	if err == nil {
		t.Fatalf(errExpectedError2)
	}
}

func TestForFrom(t *testing.T) {
	list := circularLinkList.NewCircularLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Test when the start index is within the list size
	err := list.ForFrom(2, func(value *int) {
		*value *= 2
	})

	if err != nil {
		t.Fatalf(errExpectedNoErr, err)
	}

	expected := []int{1, 2, 6, 8, 10}
	actual := list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf(errExpectedLength, len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf(errExpectedValue, v, actual[i])
		}
	}

	// Test when the start index exceeds the list size
	// The start index will be 7 since the list size is 5
	// so, in our case, it will affect only the 2nd element
	err = list.ForFrom(7, func(value *int) {
		*value *= 2
	})

	if err != nil {
		t.Fatalf(errExpectedNoErr, err)
	}

	expected = []int{1, 2, 12, 16, 20}
	actual = list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf(errExpectedLength, len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf(errExpectedValue, v, actual[i])
		}
	}
}

func TestFilter(t *testing.T) {
	list := circularLinkList.NewCircularLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Test filtering even numbers
	list.Filter(func(value int) bool {
		return value%2 == 0
	})

	expected := []int{2, 4}
	actual := list.ToSlice()

	fmt.Printf("Actual: %v\n", actual)
	if len(expected) != len(actual) {
		t.Fatalf(errExpectedLength, len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf(errExpectedValue, v, actual[i])
		}
	}

	// Test filtering odd numbers
	list.Filter(func(value int) bool {
		return value%2 != 0
	})

	expected = []int{}
	actual = list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf(errExpectedLength, len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf(errExpectedValue, v, actual[i])
		}
	}
}

func TestReduce(t *testing.T) {
	list := circularLinkList.NewCircularLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Test when the list is not empty
	result, err := list.Reduce(func(a, b int) int {
		return a + b
	})

	if err != nil {
		t.Fatalf(errExpectedNoErr, err)
	}

	expected := 15
	if result != expected {
		t.Fatalf(errExpectedResult, expected, result)
	}

	// Test when the list is empty
	emptyList := circularLinkList.NewCircularLinkList[int]()
	_, err = emptyList.Reduce(func(a, b int) int {
		return a + b
	})

	if err == nil {
		t.Fatalf(errExpectedError2)
	}
}

func TestReduceFrom(t *testing.T) {
	list := circularLinkList.NewCircularLinkList[int]()

	// Test when the list is empty
	_, err := list.ReduceFrom(0, func(acc, value int) int {
		return acc + value
	})

	if err == nil {
		t.Fatalf(errExpectedError, errListIsEmpty, err)
	}

	// Load values from a slice into the list
	list = circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4, 5})

	// Test when the start index is greater than the list size
	_, err = list.ReduceFrom(10, func(acc, value int) int {
		return acc + value
	})

	if err != nil {
		t.Fatalf(errExpectedNoErr, err)
	}

	// Test when the start index is within the list size
	result, err := list.ReduceFrom(2, func(acc, value int) int {
		return acc + value
	})

	if err != nil {
		t.Fatalf(errExpectedNoErr, err)
	}

	expected := 12
	if result != expected {
		t.Fatalf(errExpectedResult, expected, result)
	}
}

func TestReduceRange(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4, 5})

	// Test when the range is within the list size
	result, err := list.ReduceRange(1, 4, func(acc, value int) int {
		return acc + value
	})

	if err != nil {
		t.Fatalf(errExpectedNoErr, err)
	}

	expected := 14
	if result != expected {
		t.Fatalf(errExpectedResult, expected, result)
	}

	// Test when the range exceeds the list size
	// The range will be [3, 3] since the list size is 5
	// so, in our case, it will affect only the 4th element
	result, err = list.ReduceRange(3, 8, func(acc, value int) int {
		return acc + value
	})

	if err != nil {
		t.Fatalf(errExpectedNoErr, err)
	}

	expected = 4
	if result != expected {
		t.Fatalf(errExpectedResult, expected, result)
	}

	// Test when the start index is greater than the end index
	_, err = list.ReduceRange(4, 2, func(acc, value int) int {
		return acc + value
	})

	if err == nil {
		t.Fatalf(errExpectedError2)
	}
}
