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

// Package dlinkList provides a non-concurrent-safe doubly linked list.
package dlinkList

import (
	"reflect"
	"testing"
)

const (
	errListNotEmpty         = "Expected list to be empty, but it was not"
	errListIsEmpty          = "Expected list to not be empty, but it was"
	errWrongSize            = "Expected list to have size %v, but got %v"
	errNoError              = "Expected no error, but got %v"
	errYesError             = "Expected an error, but got nil"
	errWrongValue           = "Expected value to be %v, but got %v"
	errExpectedX            = "Expected %v, but got %v"
	errExpectedValToBe      = "Expected value at index %d to be %v, but got %v"
	errExpectedFilteredList = "Expected filtered list to be %v, but got %v"
	errExpectedIndex        = "Expected index to be %v, but got %v"
)

func TestNew(t *testing.T) {
	list := NewDLinkList[int]()
	if list == nil {
		t.Error("Expected list to be initialized, but got nil")
	}
}

func TestAppend(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list to have 1 item, but got %v", list.Size())
	}
}

func TestPrepend(t *testing.T) {
	list := NewDLinkList[int]()
	list.Prepend(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list to have 1 item, but got %v", list.Size())
	}
}

func TestRemove(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestRemoveEmpty(t *testing.T) {
	list := NewDLinkList[int]()
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestIsEmpty(t *testing.T) {
	list := NewDLinkList[int]()
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestIsEmptyAfterAppend(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
}

func TestSize(t *testing.T) {
	list := NewDLinkList[int]()
	if list.Size() != 0 {
		t.Errorf(errWrongSize, 0, list.Size())
	}
	list.Append(1)
	if list.Size() != 1 {
		t.Errorf(errWrongSize, 1, list.Size())
	}
	list.Append(2)
	if list.Size() != 2 {
		t.Errorf(errWrongSize, 2, list.Size())
	}
	list.Remove(1)
	if list.Size() != 1 {
		t.Errorf(errWrongSize, 1, list.Size())
	}
	list.Remove(2)
	if list.Size() != 0 {
		t.Errorf(errWrongSize, 0, list.Size())
	}
}

func TestIsEmptyAfterPrepend(t *testing.T) {
	list := NewDLinkList[int]()
	list.Prepend(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
}

func TestIsEmptyAfterRemove(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
	if list.Size() != 0 {
		t.Errorf(errWrongSize, 0, list.Size())
	}
}

func TestGetFirst(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	if list.GetFirst().Value != 1 {
		t.Errorf("Expected first element to be 1, but got %v", list.GetFirst().Value)
	}
}

func TestGetLast(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	if list.GetLast().Value != 1 {
		t.Errorf("Expected last element to be 1, but got %v", list.GetLast().Value)
	}
	list.Append(2)
	if list.GetLast().Value != 2 {
		t.Errorf("Expected last element to be 2, but got %v", list.GetLast().Value)
	}
}

func TestGetAt(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	node, err := list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if node.Value != 2 {
		t.Errorf(errWrongValue, 2, node.Value)
	}
}

func TestGetAtEmpty(t *testing.T) {
	list := NewDLinkList[int]()
	_, err := list.GetAt(0)
	if err == nil {
		t.Error(errYesError)
	}
}

func TestGetAtOutOfBound(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	_, err := list.GetAt(1)
	if err == nil {
		t.Error(errYesError)
	}
}

func TestInsertAt(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	err := list.InsertAt(1, 2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if list.Size() != 2 {
		t.Errorf(errWrongSize, 2, list.Size())
	}
	item, err := list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf(errExpectedValToBe, 1, 2, item.Value)
	}
}

func TestInsertAtEmpty(t *testing.T) {
	list := NewDLinkList[int]()
	err := list.InsertAt(1, 1)
	if err == nil {
		t.Error(errYesError)
	}
}

func TestInsertAtOutOfBound(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	err := list.InsertAt(2, 2)
	if err == nil {
		t.Error(errYesError)
	}
}

func TestInsertAtHead(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	err := list.InsertAt(0, 2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if list.Size() != 2 {
		t.Errorf(errWrongSize, 2, list.Size())
	}
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf(errExpectedValToBe, 0, 2, item.Value)
	}
}

func TestInsertAtTail(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	err := list.InsertAt(2, 2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if list.Size() != 3 {
		t.Errorf(errWrongSize, 3, list.Size())
	}
	item, err := list.GetAt(2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf(errExpectedValToBe, 1, 2, item.Value)
	}
}

func TestInsertAtMiddle(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(3)
	err := list.InsertAt(1, 2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if list.Size() != 3 {
		t.Errorf(errWrongSize, 3, list.Size())
	}
	item, err := list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf(errExpectedValToBe, 1, 2, item.Value)
	}
}

func TestRemoveAt(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	err := list.RemoveAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if list.Size() != 1 {
		t.Errorf(errWrongSize, 2, list.Size())
	}
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf(errExpectedValToBe, 1, 3, item.Value)
	}
}

func TestRemoveAtEmpty(t *testing.T) {
	list := NewDLinkList[int]()
	err := list.RemoveAt(0)
	if err == nil {
		t.Error(errYesError)
	}
}

func TestRemoveAtOutOfBound(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	err := list.RemoveAt(1)
	if err == nil {
		t.Error(errYesError)
	}
}

func TestRemoveAtHead(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	err := list.RemoveAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if list.Size() != 1 {
		t.Errorf(errWrongSize, 1, list.Size())
	}
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf(errExpectedValToBe, 0, 2, item.Value)
	}
}

func TestRemoveAtTail(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	err := list.RemoveAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if list.Size() != 1 {
		t.Errorf(errWrongSize, 1, list.Size())
	}
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf(errExpectedValToBe, 0, 1, item.Value)
	}
}

func TestRemoveAtMiddle(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	err := list.RemoveAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if list.Size() != 2 {
		t.Errorf(errWrongSize, 2, list.Size())
	}
	item, err := list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 3 {
		t.Errorf(errExpectedValToBe, 1, 3, item.Value)
	}
}

func TestReverse(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Reverse()
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 3 {
		t.Errorf(errExpectedValToBe, 0, 3, item.Value)
	}
	item, err = list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf(errExpectedValToBe, 1, 2, item.Value)
	}
	item, err = list.GetAt(2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf(errExpectedValToBe, 2, 1, item.Value)
	}
}

func TestReverseEmpty(t *testing.T) {
	list := NewDLinkList[int]()
	list.Reverse()
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestReverseSingle(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Reverse()
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf(errExpectedValToBe, 0, 1, item.Value)
	}
}

func TestReverseDouble(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Reverse()
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf(errExpectedValToBe, 0, 2, item.Value)
	}
	item, err = list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf(errExpectedValToBe, 1, 1, item.Value)
	}
}

func TestReverseTriple(t *testing.T) {
	type test struct {
		index int
		value int
	}

	list := NewDLinkList[test]()
	list.Append(test{1, 1})
	list.Append(test{2, 1})
	list.Append(test{3, 1})
	list.Reverse()
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value.index != 3 {
		t.Errorf(errExpectedValToBe, 0, 3, item.Value)
	}
	item, err = list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value.index != 2 {
		t.Errorf(errExpectedValToBe, 1, 2, item.Value)
	}
	item, err = list.GetAt(2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value.index != 1 {
		t.Errorf(errExpectedValToBe, 2, 1, item.Value)
	}
}

func TestCopy(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	newList := list.Copy()
	if list.Size() != newList.Size() {
		t.Errorf("Expected new list to have size %v, but got %v", list.Size(), newList.Size())
	}
	for i := uint64(0); i < list.size; i++ {
		item, err := list.GetAt(i)
		if err != nil {
			t.Errorf(errNoError, err)
		}
		newItem, err := newList.GetAt(i)
		if err != nil {
			t.Errorf(errNoError, err)
		}
		if item.Value != newItem.Value {
			t.Errorf(errExpectedValToBe, i, item.Value, newItem.Value)
		}
	}
}

func TestMerge(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	newList := NewDLinkList[int]()
	newList.Append(4)
	newList.Append(5)
	newList.Append(6)
	list.Merge(newList)
	if list.Size() != 6 {
		t.Errorf(errWrongSize, 6, list.Size())
	}
	for i := uint64(0); i < list.size; i++ {
		item, err := list.GetAt(i)
		if err != nil {
			t.Errorf(errNoError, err)
		}
		if item.Value != int(i)+1 {
			t.Errorf(errExpectedValToBe, i, i+1, item.Value)
		}
	}
}

func TestMergeEmpty(t *testing.T) {
	list := NewDLinkList[int]()
	newList := NewDLinkList[int]()
	list.Merge(newList)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestMergeEmptyList(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	newList := NewDLinkList[int]()
	list.Merge(newList)
	if list.Size() != 1 {
		t.Errorf(errWrongSize, 1, list.Size())
	}
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf(errExpectedValToBe, 0, 1, item.Value)
	}
}

func TestMergeEmptyList2(t *testing.T) {
	list := NewDLinkList[int]()
	newList := NewDLinkList[int]()
	newList.Append(1)
	list.Merge(newList)
	if list.Size() != 1 {
		t.Errorf(errWrongSize, 1, list.Size())
	}
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf(errExpectedValToBe, 0, 1, item.Value)
	}
}

func TestMergeEmptyList3(t *testing.T) {
	list := NewDLinkList[int]()
	newList := NewDLinkList[int]()
	list.Append(1)
	list.Merge(newList)
	if list.Size() != 1 {
		t.Errorf(errWrongSize, 1, list.Size())
	}
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf(errExpectedValToBe, 0, 1, item.Value)
	}
}

func TestMergeEmptyList5(t *testing.T) {
	list := NewDLinkList[int]()
	newList := NewDLinkList[int]()
	list.Append(1)
	newList.Merge(list)
	if newList.Size() != 1 {
		t.Errorf(errWrongSize, 1, newList.Size())
	}
	item, err := newList.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf(errExpectedValToBe, 0, 1, item.Value)
	}
}

func TestMergeEmptyList6(t *testing.T) {
	list := NewDLinkList[int]()
	newList := NewDLinkList[int]()
	newList.Append(1)
	newList.Merge(list)
	if newList.Size() != 1 {
		t.Errorf(errWrongSize, 1, newList.Size())
	}
	item, err := newList.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf(errExpectedValToBe, 0, 1, item.Value)
	}
}

func TestMergeEmptyList8(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	newList := NewDLinkList[int]()
	newList.Merge(list)
	if newList.IsEmpty() {
		t.Error(errListIsEmpty)
	}
}

func TestMergeEmptyList9(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Remove(1)
	newList := NewDLinkList[int]()
	list.Merge(newList)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestMergeEmptyList10(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	newList := NewDLinkList[int]()
	newList.Merge(list)
	newList.Remove(1)
	if !newList.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestMergeEmptyList15(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	newList := NewDLinkList[int]()
	list.Merge(newList)
	list.Clear()
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestInsert(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.InsertAt(1, 4)
	if err != nil {
		t.Errorf(errNoError, err)
	}

	if list.Size() != 4 {
		t.Errorf(errWrongSize, 4, list.Size())
	}

	item, err := list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 4 {
		t.Errorf(errWrongValue, 4, item.Value)
	}

	item, err = list.GetAt(2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf(errWrongValue, 2, item.Value)
	}
}

func TestInsertAtStart(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.InsertAt(0, 4)
	if err != nil {
		t.Errorf(errNoError, err)
	}

	if list.Size() != 4 {
		t.Errorf(errWrongSize, 4, list.Size())
	}

	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 4 {
		t.Errorf(errWrongValue, 4, item.Value)
	}

	item, err = list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf(errWrongValue, 1, item.Value)
	}
}

func TestInsertAtEnd(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.InsertAt(3, 4)
	if err != nil {
		t.Errorf(errNoError, err)
	}

	if list.Size() != 4 {
		t.Errorf(errWrongSize, 4, list.Size())
	}

	item, err := list.GetAt(3)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 4 {
		t.Errorf(errWrongValue, 4, item.Value)
	}

	item, err = list.GetAt(2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 3 {
		t.Errorf(errWrongValue, 3, item.Value)
	}
}

func TestInsertOutOfBounds(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.InsertAt(5, 4)
	if err == nil {
		t.Error(errYesError)
	}

	if list.Size() != 3 {
		t.Errorf(errWrongSize, 3, list.Size())
	}
}

func TestToSlice(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	slice := list.ToSlice()
	expected := []int{1, 2, 3}

	if len(slice) != len(expected) {
		t.Errorf("Expected slice length %d, but got %d", len(expected), len(slice))
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != expected[i] {
			t.Errorf("Expected value at index %d to be %d, but got %d", i, expected[i], slice[i])
		}
	}
}

func TestToSliceEmpty(t *testing.T) {
	list := NewDLinkList[int]()

	slice := list.ToSlice()

	if len(slice) != 0 {
		t.Errorf("Expected empty slice, but got length %d", len(slice))
	}
}

func TestFind(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test case 1: Value exists in the list
	node, err := list.Find(2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if node.Value != 2 {
		t.Errorf("Expected value to be 2, but got %v", node.Value)
	}

	// Test case 2: Value does not exist in the list
	_, err = list.Find(4)
	if err == nil {
		t.Error(errYesError)
	}
}

func TestDelete(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test deleting a value that exists in the list
	list.Delete(2)
	if list.Size() != 2 {
		t.Errorf(errWrongSize, 2, list.Size())
	}
	if list.Contains(2) {
		t.Error("Expected list to not contain value 2")
	}

	// Test deleting a value that doesn't exist in the list
	list.Delete(4)
	if list.Size() != 2 {
		t.Errorf(errWrongSize, 2, list.Size())
	}

	// Test deleting the first value in the list
	list.Delete(1)
	if list.Size() != 1 {
		t.Errorf(errWrongSize, 1, list.Size())
	}
	if list.Contains(1) {
		t.Error("Expected list to not contain value 1")
	}

	// Test deleting the last value in the list
	list.Delete(3)
	if list.Size() != 0 {
		t.Errorf(errWrongSize, 0, list.Size())
	}
	if list.Contains(3) {
		t.Error("Expected list to not contain value 3")
	}
}

func TestDeleteEmpty(t *testing.T) {
	list := NewDLinkList[int]()
	list.Delete(1)
	if list.Size() != 0 {
		t.Errorf(errWrongSize, 0, list.Size())
	}
}

func TestInsertEmpty(t *testing.T) {
	list := NewDLinkList[int]()

	err := list.Insert(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}

	expected := []int{1}
	actual := list.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedX, expected, actual)
	}
}

func TestInsertAfter(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.InsertAfter(2, 4)

	expected := []int{1, 2, 4, 3}
	actual := list.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedX, expected, actual)
	}
}

func TestInsertAfterNotFound(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.InsertAfter(4, 5)

	expected := []int{1, 2, 3}
	actual := list.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedX, expected, actual)
	}
}

func TestInsertBefore(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.InsertBefore(2, 4)

	expected := []int{1, 4, 2, 3}
	actual := list.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedX, expected, actual)
	}
}

func TestInsertBeforeNotFound(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.InsertBefore(4, 5)

	expected := []int{1, 2, 3}
	actual := list.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedX, expected, actual)
	}
}

func TestDeleteLast(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.DeleteLast()
	if list.Size() != 2 {
		t.Errorf(errWrongSize, 2, list.Size())
	}

	item, err := list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf(errExpectedValToBe, 1, 2, item.Value)
	}

	list.DeleteLast()
	if list.Size() != 1 {
		t.Errorf(errWrongSize, 1, list.Size())
	}

	item, err = list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf(errExpectedValToBe, 0, 1, item.Value)
	}

	list.DeleteLast()
	if list.Size() != 0 {
		t.Errorf(errWrongSize, 0, list.Size())
	}

	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestDeleteLastEmpty(t *testing.T) {
	list := NewDLinkList[int]()
	list.DeleteLast()
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestDeleteFirst(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.DeleteFirst()

	if list.Size() != 2 {
		t.Errorf(errWrongSize, 2, list.Size())
	}

	if list.GetFirst().Value != 2 {
		t.Errorf("Expected first element to be 2, but got %v", list.GetFirst().Value)
	}
}

func TestToSliceReverse(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	expected := []int{3, 2, 1}
	result := list.ToSliceReverse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}
}

func TestToSliceReverseEmpty(t *testing.T) {
	list := NewDLinkList[int]()

	expected := []int{}
	result := list.ToSliceReverse()

	if result != nil {
		t.Errorf(errExpectedX, expected, result)
	}
}

func TestToSliceReverseSingle(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)

	expected := []int{1}
	result := list.ToSliceReverse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}
}

func TestToSliceReverseDouble(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)

	expected := []int{2, 1}
	result := list.ToSliceReverse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}
}

func TestToSliceReverseTriple(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	expected := []int{4, 3, 2, 1}
	result := list.ToSliceReverse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}
}

func TestToSliceFromIndex(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	tests := []struct {
		index    int
		expected []int
	}{
		{0, []int{1, 2, 3, 4, 5}},
		{1, []int{2, 3, 4, 5}},
		{2, []int{3, 4, 5}},
		{3, []int{4, 5}},
		{4, []int{5}},
		{5, []int{}},
		{6, []int{}},
	}

	for _, test := range tests {
		result := list.ToSliceFromIndex(uint64(test.index))
		if result == nil && (test.index != 5 && test.index != 6) {
			t.Errorf("Expected ToSliceFromIndex(%d) to return %v, but got %v", test.index, test.expected, result)
		}
	}
}

func TestToSliceReverseFromIndex(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Test from index 0
	expected := []int{5, 4, 3, 2, 1}
	result := list.ToSliceReverseFromIndex(0)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}

	// Test from index 1
	expected = []int{4, 3, 2, 1}
	result = list.ToSliceReverseFromIndex(1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}

	// Test from index 2
	expected = []int{3, 2, 1}
	result = list.ToSliceReverseFromIndex(2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}

	// Test from index 3
	expected = []int{2, 1}
	result = list.ToSliceReverseFromIndex(3)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}

	// Test from index 4
	expected = []int{1}
	result = list.ToSliceReverseFromIndex(4)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}

	// Test from index 5 (out of bounds)
	expected = []int{}
	result = list.ToSliceReverseFromIndex(5)
	if result != nil {
		t.Errorf(errExpectedX, expected, result)
	}
}

func TestForEach(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	sum := 0
	list.ForEach(func(value *int) {
		sum += *value
	})

	expectedSum := 1 + 2 + 3
	if sum != expectedSum {
		t.Errorf("Expected sum to be %d, but got %d", expectedSum, sum)
	}
}

func TestForEachEmpty(t *testing.T) {
	list := NewDLinkList[int]()

	list.ForEach(func(value *int) {
		t.Error("ForEach should not be called on an empty list")
	})
}

func TestAny(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test case 1: Function returns true for at least one node
	result1 := list.Any(func(value int) bool {
		return value > 2
	})
	if !result1 {
		t.Error("Expected Any to return true, but got false")
	}

	// Test case 2: Function returns false for all nodes
	result2 := list.Any(func(value int) bool {
		return value > 5
	})
	if result2 {
		t.Error("Expected Any to return false, but got true")
	}
}

func TestAll(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	// Test case 1: All elements are even
	allEven := list.All(func(n int) bool {
		return n%2 == 0
	})
	if allEven {
		t.Error("Expected all elements not to be even")
	}

	// Test case 2: All elements are odd
	allOdd := list.All(func(n int) bool {
		return n%2 != 0
	})
	if allOdd {
		t.Error("Expected not all elements to be odd")
	}

	// Test case 3: All elements are greater than 0
	allGreaterThanZero := list.All(func(n int) bool {
		return n > 0
	})
	if !allGreaterThanZero {
		t.Error("Expected all elements to be greater than 0")
	}

	// Test case 4: All elements are less than 10
	allLessThanTen := list.All(func(n int) bool {
		return n < 10
	})
	if !allLessThanTen {
		t.Error("Expected all elements to be less than 10")
	}

	// Test case 5: All elements are equal to 1
	allEqualToOne := list.All(func(n int) bool {
		return n == 1
	})
	if allEqualToOne {
		t.Error("Expected all elements not to be equal to 1")
	}
}

func TestIndexOf(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	index := list.IndexOf(2)
	if index != 1 {
		t.Errorf("Expected index of 2 to be 1, but got %v", index)
	}

	index = list.IndexOf(4)
	if index != -1 {
		t.Errorf("Expected index of 4 to be -1, but got %v", index)
	}
}

func TestLastIndexOf(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(2)
	list.Append(4)

	index, err := list.LastIndexOf(2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if index != 3 {
		t.Errorf("Expected last index of 2 to be 3, but got %v", index)
	}

	_, err = list.LastIndexOf(5)
	if err == nil {
		t.Errorf(errYesError)
	}

	index, err = list.LastIndexOf(4)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if index != 4 {
		t.Errorf("Expected last index of 4 to be 4, but got %v", index)
	}
}

func TestFilter(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Test filtering even numbers
	filtered := list.Filter(func(value int) bool {
		return value%2 == 0
	})

	expected := []int{2, 4}
	actual := filtered.ToSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedFilteredList, expected, actual)
	}

	// Test filtering odd numbers
	filtered = list.Filter(func(value int) bool {
		return value%2 != 0
	})

	expected = []int{1, 3, 5}
	actual = filtered.ToSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedFilteredList, expected, actual)
	}

	// Test filtering with an empty list
	emptyList := NewDLinkList[int]()
	filtered = emptyList.Filter(func(value int) bool {
		return value > 0
	})

	expected = []int{}
	actual = filtered.ToSlice()
	if actual != nil {
		t.Errorf(errExpectedFilteredList, expected, actual)
	}
}

func TestMap(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	mappedList := list.Map(func(value int) int {
		return value * 2
	})

	expectedList := NewDLinkList[int]()
	expectedList.Append(2)
	expectedList.Append(4)
	expectedList.Append(6)

	if !mappedList.Equal(expectedList) {
		t.Errorf("Expected mapped list to be equal to %v, but got %v", expectedList.ToSlice(), mappedList.ToSlice())
	}
}

func TestMapEmptyList(t *testing.T) {
	list := NewDLinkList[int]()

	mappedList := list.Map(func(value int) int {
		return value * 2
	})

	if !mappedList.IsEmpty() {
		t.Error("Expected mapped list to be empty")
	}
}

func TestReduce(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	sum := list.Reduce(func(a, b int) int {
		return a + b
	})

	if sum != 6 {
		t.Errorf("Expected sum to be 6, but got %v", sum)
	}

	product := list.Reduce(func(a, b int) int {
		return a * b
	})

	if product != 6 {
		t.Errorf("Expected product to be 6, but got %v", product)
	}

	max := list.Reduce(func(a, b int) int {
		if a > b {
			return a
		}
		return b
	})

	if max != 3 {
		t.Errorf("Expected max to be 3, but got %v", max)
	}

	min := list.Reduce(func(a, b int) int {
		if a < b {
			return a
		}
		return b
	})

	if min != 1 {
		t.Errorf("Expected min to be 1, but got %v", min)
	}
}

func TestReduceEmptyList(t *testing.T) {
	list := NewDLinkList[int]()

	sum := list.Reduce(func(a, b int) int {
		return a + b
	})

	if sum != 0 {
		t.Errorf("Expected sum to be 0, but got %v", sum)
	}
}

func TestReverseCopy(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	reverseCopy := list.ReverseCopy()

	// Check if the original list is not modified
	if list.Size() != 3 {
		t.Errorf("Expected original list to have size 3, but got %v", list.Size())
	}

	// Check if the reverse copy is correct
	if reverseCopy.Size() != 3 {
		t.Errorf("Expected reverse copy to have size 3, but got %v", reverseCopy.Size())
	}

	item, err := reverseCopy.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 3 {
		t.Errorf(errExpectedValToBe, 0, 3, item.Value)
	}

	item, err = reverseCopy.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf(errExpectedValToBe, 1, 2, item.Value)
	}

	item, err = reverseCopy.GetAt(2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf(errExpectedValToBe, 2, 1, item.Value)
	}
}

func TestReverseMerge(t *testing.T) {
	list1 := NewDLinkList[int]()
	list1.Append(1)
	list1.Append(2)
	list1.Append(3)

	list2 := NewDLinkList[int]()
	list2.Append(4)
	list2.Append(5)
	list2.Append(6)

	list1.ReverseMerge(list2)

	expected := []int{1, 2, 3, 6, 5, 4}
	actual := list1.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedX, expected, actual)
	}
}

func TestReverseMergeEmptyList(t *testing.T) {
	list1 := NewDLinkList[int]()

	list2 := NewDLinkList[int]()
	list2.Append(1)
	list2.Append(2)
	list2.Append(3)

	list1.ReverseMerge(list2)

	expected := []int{3, 2, 1}
	actual := list1.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedX, expected, actual)
	}
}

func TestReverseMergeEmptyLists(t *testing.T) {
	list1 := NewDLinkList[int]()

	list2 := NewDLinkList[int]()

	list1.ReverseMerge(list2)

	expected := []int{}
	actual := list1.ToSlice()

	if actual != nil {
		t.Errorf(errExpectedX, expected, actual)
	}
}

func TestEqual(t *testing.T) {
	list1 := NewDLinkList[int]()
	list2 := NewDLinkList[int]()

	// Test when both lists are empty
	if !list1.Equal(list2) {
		t.Error("Expected two empty lists to be equal")
	}

	// Test when one list is empty and the other is not
	list1.Append(1)
	if list1.Equal(list2) {
		t.Error("Expected list with one element to not be equal to an empty list")
	}

	// Test when both lists have the same elements in the same order
	list2.Append(1)
	if !list1.Equal(list2) {
		t.Error("Expected two lists with the same elements to be equal")
	}

	// Test when both lists have the same elements in a different order
	list1.Append(2)
	list2.Prepend(2)
	if list1.Equal(list2) {
		t.Error("Expected two lists with the same elements in a different order to not be equal")
	}

	// Test when both lists have different elements
	list2.DeleteFirst()
	list2.Append(3)
	if list1.Equal(list2) {
		t.Error("Expected two lists with different elements to not be equal")
	}
}

func TestSwap(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Swap nodes at index 0 and 2
	err := list.Swap(0, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify the values after swapping
	expected := []int{3, 2, 1}
	actual := list.ToSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedX, expected, actual)
	}

	// Swap nodes at index 1 and 1 (same index)
	err = list.Swap(1, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify the values after swapping (no change)
	expected = []int{3, 2, 1}
	actual = list.ToSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedX, expected, actual)
	}

	// Swap nodes at out-of-bound indices
	err = list.Swap(20, 3)
	if err == nil {
		t.Error(errYesError)
	}

	// Verify the values after swapping (no change)
	expected = []int{3, 2, 1}
	actual = list.ToSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedX, expected, actual)
	}
}

func TestSort(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(3)
	list.Append(1)
	list.Append(2)
	list.Append(5)
	list.Append(4)

	list.Sort(func(a, b int) bool {
		return a < b
	})

	expected := []int{1, 2, 3, 4, 5}
	actual := list.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedX, expected, actual)
	}
}

func TestSortEmpty(t *testing.T) {
	list := NewDLinkList[int]()

	list.Sort(func(a, b int) bool {
		return a < b
	})

	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestSortSingle(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)

	list.Sort(func(a, b int) bool {
		return a < b
	})

	expected := []int{1}
	actual := list.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedX, expected, actual)
	}
}

func TestSortAlreadySorted(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Sort(func(a, b int) bool {
		return a < b
	})

	expected := []int{1, 2, 3}
	actual := list.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedX, expected, actual)
	}
}

func TestSortDescendingOrder(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(3)
	list.Append(2)
	list.Append(1)

	list.Sort(func(a, b int) bool {
		return a < b
	})

	expected := []int{1, 2, 3}
	actual := list.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedX, expected, actual)
	}
}

func TestSortCustomType(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	list := NewDLinkList[Person]()
	list.Append(Person{Name: "Alice", Age: 25})
	list.Append(Person{Name: "Bob", Age: 20})
	list.Append(Person{Name: "Charlie", Age: 30})

	list.Sort(func(a, b Person) bool {
		return a.Age < b.Age
	})

	expected := []Person{
		{Name: "Bob", Age: 20},
		{Name: "Alice", Age: 25},
		{Name: "Charlie", Age: 30},
	}
	actual := list.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedX, expected, actual)
	}
}

func TestFindAll(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test case 1: FindAll returns a new doubly linked list containing all nodes that satisfy the given function
	result := list.FindAll(func(value int) bool {
		return value%2 == 0
	})
	expectedResult := NewDLinkList[int]()
	expectedResult.Append(2)
	if !expectedResult.Equal(result) {
		t.Errorf("Expected result to be %v, but got %v", expectedResult.ToSlice(), result.ToSlice())
	}

	// Test case 2: FindAll returns an empty doubly linked list if no nodes satisfy the given function
	result = list.FindAll(func(value int) bool {
		return value > 10
	})
	expectedResult = NewDLinkList[int]()
	if !expectedResult.Equal(result) {
		t.Errorf("Expected result to be %v, but got %v", expectedResult.ToSlice(), result.ToSlice())
	}
}

func TestFindLast(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test case 1: Value exists in the list
	node, err := list.FindLast(func(value int) bool {
		return value == 2
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if node.Value != 2 {
		t.Errorf("Expected value to be 2, but got %v", node.Value)
	}

	// Test case 2: Value does not exist in the list
	_, err = list.FindLast(func(value int) bool {
		return value == 4
	})
	if err == nil {
		t.Error(errYesError)
	}
}

func TestFindLastIndex(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(2)
	list.Append(4)

	index := list.FindLastIndex(func(value int) bool {
		return value == 2
	})

	if index != 3 {
		t.Errorf(errExpectedIndex, 3, index)
	}

	index = list.FindLastIndex(func(value int) bool {
		return value == 5
	})

	if index != -1 {
		t.Errorf(errExpectedIndex, -1, index)
	}
}

func TestFindLastIndexEmpty(t *testing.T) {
	list := NewDLinkList[int]()
	index := list.FindLastIndex(func(value int) bool {
		return value == 1
	})

	if index != -1 {
		t.Errorf(errExpectedIndex, -1, index)
	}
}

func TestFindIndex(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	index := list.FindIndex(func(value int) bool {
		return value == 2
	})

	if index != 1 {
		t.Errorf(errExpectedIndex, 1, index)
	}

	index = list.FindIndex(func(value int) bool {
		return value == 4
	})

	if index != -1 {
		t.Errorf(errExpectedIndex, -1, index)
	}
}

func TestFindIndexEmpty(t *testing.T) {
	list := NewDLinkList[int]()
	index := list.FindIndex(func(value int) bool {
		return value == 1
	})

	if index != -1 {
		t.Errorf(errExpectedIndex, -1, index)
	}
}
func TestForFrom(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	var result []int
	list.ForFrom(1, func(value *int) {
		result = append(result, *value)
	})

	expected := []int{2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}
}

func TestForFromEmpty(t *testing.T) {
	list := NewDLinkList[int]()

	var result []int
	list.ForFrom(0, func(value *int) {
		result = append(result, *value)
	})

	if len(result) != 0 {
		t.Errorf("Expected an empty result, but got %v", result)
	}
}

func TestForFromOutOfBound(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	var result []int
	list.ForFrom(3, func(value *int) {
		result = append(result, *value)
	})

	if len(result) != 0 {
		t.Errorf("Expected an empty result, but got %v", result)
	}
}

func TestForRange(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Test case 1: start = 0, end = 2
	var result []int
	list.ForRange(0, 2, func(value *int) {
		result = append(result, *value)
	})

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}

	// Test case 2: start = 1, end = 3
	result = nil
	list.ForRange(1, 3, func(value *int) {
		result = append(result, *value)
	})

	expected = []int{2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}

	// Test case 3: start = 2, end = 4
	result = nil
	list.ForRange(2, 4, func(value *int) {
		result = append(result, *value)
	})

	expected = []int{3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}

	// Test case 4: start = 0, end = 0
	result = nil
	list.ForRange(0, 0, func(value *int) {
		result = append(result, *value)
	})

	expected = []int{1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}

	// Test case 5: start = 4, end = 4
	result = nil
	list.ForRange(4, 4, func(value *int) {
		result = append(result, *value)
	})

	expected = []int{5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}

	// Test case 6: start = 0, end = 5 (out of bounds)
	result = nil
	list.ForRange(0, 5, func(value *int) {
		result = append(result, *value)
	})

	expected = []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}

	// Test case 7: start = 5, end = 0 (invalid range)
	result = nil
	list.ForRange(5, 0, func(value *int) {
		result = append(result, *value)
	})

	expected = []int{}
	if result != nil {
		t.Errorf(errExpectedX, expected, result)
	}
}

func TestForReverseRange(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Test case 1: Reverse range from index 0 to 2
	var result1 []int
	list.ForReverseRange(0, 2, func(value *int) {
		result1 = append(result1, *value)
	})
	expected1 := []int{5, 4, 3}
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf(errExpectedFilteredList, expected1, result1)
	}

	// Test case 2: Reverse range from index 1 to 3
	var result2 []int
	list.ForReverseRange(1, 3, func(value *int) {
		result2 = append(result2, *value)
	})
	expected2 := []int{4, 3, 2}
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf(errExpectedFilteredList, expected2, result2)
	}

	// Test case 3: Reverse range from index 2 to 4
	var result3 []int
	list.ForReverseRange(2, 4, func(value *int) {
		result3 = append(result3, *value)
	})
	expected3 := []int{3, 2, 1}
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf(errExpectedFilteredList, expected3, result3)
	}

	// Test case 4: Reverse range from index 0 to 4
	var result4 []int
	list.ForReverseRange(0, 4, func(value *int) {
		result4 = append(result4, *value)
	})
	expected4 := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf(errExpectedFilteredList, expected4, result4)
	}

	// Test case 5: Reverse range from index 1 to 1
	var result5 []int
	list.ForReverseRange(1, 1, func(value *int) {
		result5 = append(result5, *value)
	})
	expected5 := []int{4}
	if !reflect.DeepEqual(result5, expected5) {
		t.Errorf(errExpectedFilteredList, expected5, result5)
	}

	// Test case 6: Reverse range from index 4 to 2
	var result6 []int
	list.ForReverseRange(4, 2, func(value *int) {
		result6 = append(result6, *value)
	})
	expected6 := []int{1}
	if result6 != nil {
		t.Errorf(errExpectedFilteredList, expected6, result6)
	}
}

func TestForReverseFrom(t *testing.T) {
	list := NewDLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	var result []int
	list.ForReverseFrom(1, func(value *int) {
		result = append(result, *value)
	})

	expected := []int{2, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(errExpectedX, expected, result)
	}

	// Test with empty list
	emptyList := NewDLinkList[int]()
	emptyList.ForReverseFrom(0, func(value *int) {
		t.Error("Should not execute the callback function for an empty list")
	})
}
