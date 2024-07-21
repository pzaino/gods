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
	errListNotEmpty = "Expected list to be empty, but it was not"
	errListIsEmpty  = "Expected list to not be empty, but it was"
	errWrongSize    = "Expected list to have size %v, but got %v"
	errNoError      = "Expected no error, but got %v"
	errYesError     = "Expected an error, but got nil"
	errWrongValue   = "Expected value to be %v, but got %v"
)

func TestNew(t *testing.T) {
	list := New[int]()
	if list == nil {
		t.Error("Expected list to be initialized, but got nil")
	}
}

func TestAppend(t *testing.T) {
	list := New[int]()
	list.Append(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list to have 1 item, but got %v", list.Size())
	}
}

func TestPrepend(t *testing.T) {
	list := New[int]()
	list.Prepend(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list to have 1 item, but got %v", list.Size())
	}
}

func TestRemove(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestRemoveEmpty(t *testing.T) {
	list := New[int]()
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestIsEmpty(t *testing.T) {
	list := New[int]()
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestIsEmptyAfterAppend(t *testing.T) {
	list := New[int]()
	list.Append(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
}

func TestSize(t *testing.T) {
	list := New[int]()
	if list.Size() != 0 {
		t.Errorf("Expected list to have size 0, but got %v", list.Size())
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
	list := New[int]()
	list.Prepend(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
}

func TestIsEmptyAfterRemove(t *testing.T) {
	list := New[int]()
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
	list := New[int]()
	list.Append(1)
	if list.GetFirst().Value != 1 {
		t.Errorf("Expected first element to be 1, but got %v", list.GetFirst().Value)
	}
}

func TestGetLast(t *testing.T) {
	list := New[int]()
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
	list := New[int]()
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
	list := New[int]()
	_, err := list.GetAt(0)
	if err == nil {
		t.Error(errYesError)
	}
}

func TestGetAtOutOfBound(t *testing.T) {
	list := New[int]()
	list.Append(1)
	_, err := list.GetAt(1)
	if err == nil {
		t.Error(errYesError)
	}
}

func TestInsertAt(t *testing.T) {
	list := New[int]()
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
		t.Errorf("Expected value at index 1 to be 2, but got %v", item.Value)
	}
}

func TestInsertAtEmpty(t *testing.T) {
	list := New[int]()
	err := list.InsertAt(1, 1)
	if err == nil {
		t.Error(errYesError)
	}
}

func TestInsertAtOutOfBound(t *testing.T) {
	list := New[int]()
	list.Append(1)
	err := list.InsertAt(2, 2)
	if err == nil {
		t.Error(errYesError)
	}
}

func TestInsertAtHead(t *testing.T) {
	list := New[int]()
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
		t.Errorf("Expected value at index 0 to be 2, but got %v", item.Value)
	}
}

func TestInsertAtTail(t *testing.T) {
	list := New[int]()
	list.Append(1)
	err := list.InsertAt(1, 2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if list.Size() != 2 {
		t.Errorf("Expected list to have size 2, but got %v", list.Size())
	}
	item, err := list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf("Expected value at index 1 to be 2, but got %v", item.Value)
	}
}

func TestInsertAtMiddle(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(3)
	err := list.InsertAt(1, 2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if list.Size() != 3 {
		t.Errorf("Expected list to have size 3, but got %v", list.Size())
	}
	item, err := list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf("Expected value at index 1 to be 2, but got %v", item.Value)
	}
}

func TestRemoveAt(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	err := list.RemoveAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if list.Size() != 2 {
		t.Errorf("Expected list to have size 2, but got %v", list.Size())
	}
	item, err := list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 3 {
		t.Errorf("Expected value at index 1 to be 3, but got %v", item.Value)
	}
}

func TestRemoveAtEmpty(t *testing.T) {
	list := New[int]()
	err := list.RemoveAt(0)
	if err == nil {
		t.Error(errYesError)
	}
}

func TestRemoveAtOutOfBound(t *testing.T) {
	list := New[int]()
	list.Append(1)
	err := list.RemoveAt(1)
	if err == nil {
		t.Error(errYesError)
	}
}

func TestRemoveAtHead(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	err := list.RemoveAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list to have size 1, but got %v", list.Size())
	}
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf("Expected value at index 0 to be 2, but got %v", item.Value)
	}
}

func TestRemoveAtTail(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	err := list.RemoveAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list to have size 1, but got %v", list.Size())
	}
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 0 to be 1, but got %v", item.Value)
	}
}

func TestRemoveAtMiddle(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	err := list.RemoveAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if list.Size() != 2 {
		t.Errorf("Expected list to have size 2, but got %v", list.Size())
	}
	item, err := list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 3 {
		t.Errorf("Expected value at index 1 to be 3, but got %v", item.Value)
	}
}

func TestReverse(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Reverse()
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 3 {
		t.Errorf("Expected value at index 0 to be 3, but got %v", item.Value)
	}
	item, err = list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf("Expected value at index 1 to be 2, but got %v", item.Value)
	}
	item, err = list.GetAt(2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 2 to be 1, but got %v", item.Value)
	}
}

func TestReverseEmpty(t *testing.T) {
	list := New[int]()
	list.Reverse()
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestReverseSingle(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Reverse()
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 0 to be 1, but got %v", item.Value)
	}
}

func TestReverseDouble(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Reverse()
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf("Expected value at index 0 to be 2, but got %v", item.Value)
	}
	item, err = list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 1 to be 1, but got %v", item.Value)
	}
}

func TestReverseTriple(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Reverse()
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 3 {
		t.Errorf("Expected value at index 0 to be 3, but got %v", item.Value)
	}
	item, err = list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf("Expected value at index 1 to be 2, but got %v", item.Value)
	}
	item, err = list.GetAt(2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 2 to be 1, but got %v", item.Value)
	}
}

func TestCopy(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	newList := list.Copy()
	if list.Size() != newList.Size() {
		t.Errorf("Expected new list to have size %v, but got %v", list.Size(), newList.Size())
	}
	for i := 0; i < list.Size(); i++ {
		item, err := list.GetAt(i)
		if err != nil {
			t.Errorf(errNoError, err)
		}
		newItem, err := newList.GetAt(i)
		if err != nil {
			t.Errorf(errNoError, err)
		}
		if item.Value != newItem.Value {
			t.Errorf("Expected value at index %v to be %v, but got %v", i, item.Value, newItem.Value)
		}
	}
}

func TestMerge(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	newList := New[int]()
	newList.Append(4)
	newList.Append(5)
	newList.Append(6)
	list.Merge(newList)
	if list.Size() != 6 {
		t.Errorf("Expected list to have size 6, but got %v", list.Size())
	}
	for i := 0; i < list.Size(); i++ {
		item, err := list.GetAt(i)
		if err != nil {
			t.Errorf(errNoError, err)
		}
		if item.Value != i+1 {
			t.Errorf("Expected value at index %v to be %v, but got %v", i, i+1, item.Value)
		}
	}
}

func TestMergeEmpty(t *testing.T) {
	list := New[int]()
	newList := New[int]()
	list.Merge(newList)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestMergeEmptyList(t *testing.T) {
	list := New[int]()
	list.Append(1)
	newList := New[int]()
	list.Merge(newList)
	if list.Size() != 1 {
		t.Errorf("Expected list to have size 1, but got %v", list.Size())
	}
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 0 to be 1, but got %v", item.Value)
	}
}

func TestMergeEmptyList2(t *testing.T) {
	list := New[int]()
	newList := New[int]()
	newList.Append(1)
	list.Merge(newList)
	if list.Size() != 1 {
		t.Errorf("Expected list to have size 1, but got %v", list.Size())
	}
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 0 to be 1, but got %v", item.Value)
	}
}

func TestMergeEmptyList3(t *testing.T) {
	list := New[int]()
	newList := New[int]()
	list.Append(1)
	list.Merge(newList)
	if list.Size() != 1 {
		t.Errorf("Expected list to have size 1, but got %v", list.Size())
	}
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 0 to be 1, but got %v", item.Value)
	}
}

func TestMergeEmptyList4(t *testing.T) {
	list := New[int]()
	newList := New[int]()
	newList.Append(1)
	list.Merge(newList)
	if list.Size() != 1 {
		t.Errorf("Expected list to have size 1, but got %v", list.Size())
	}
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 0 to be 1, but got %v", item.Value)
	}
}

func TestMergeEmptyList5(t *testing.T) {
	list := New[int]()
	newList := New[int]()
	list.Append(1)
	newList.Merge(list)
	if newList.Size() != 1 {
		t.Errorf("Expected list to have size 1, but got %v", newList.Size())
	}
	item, err := newList.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 0 to be 1, but got %v", item.Value)
	}
}

func TestMergeEmptyList6(t *testing.T) {
	list := New[int]()
	newList := New[int]()
	newList.Append(1)
	newList.Merge(list)
	if newList.Size() != 1 {
		t.Errorf("Expected list to have size 1, but got %v", newList.Size())
	}
	item, err := newList.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 0 to be 1, but got %v", item.Value)
	}
}

func TestMergeEmptyList7(t *testing.T) {
	list := New[int]()
	newList := New[int]()
	list.Merge(newList)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestMergeEmptyList8(t *testing.T) {
	list := New[int]()
	list.Append(1)
	newList := New[int]()
	newList.Merge(list)
	if newList.IsEmpty() {
		t.Error(errListIsEmpty)
	}
}

func TestMergeEmptyList9(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Remove(1)
	newList := New[int]()
	list.Merge(newList)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestMergeEmptyList10(t *testing.T) {
	list := New[int]()
	list.Append(1)
	newList := New[int]()
	newList.Merge(list)
	newList.Remove(1)
	if !newList.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestMergeEmptyList11(t *testing.T) {
	list := New[int]()
	newList := New[int]()
	list.Append(1)
	list.Merge(newList)
	if list.Size() != 1 {
		t.Errorf("Expected list to have size 1, but got %v", list.Size())
	}
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 0 to be 1, but got %v", item.Value)
	}
}

func TestMergeEmptyList12(t *testing.T) {
	list := New[int]()
	newList := New[int]()
	newList.Append(1)
	list.Merge(newList)
	if list.Size() != 1 {
		t.Errorf("Expected list to have size 1, but got %v", list.Size())
	}
	item, err := list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 0 to be 1, but got %v", item.Value)
	}
}

func TestMergeEmptyList13(t *testing.T) {
	list := New[int]()
	newList := New[int]()
	list.Append(1)
	newList.Merge(list)
	if newList.Size() != 1 {
		t.Errorf("Expected list to have size 1, but got %v", newList.Size())
	}
	item, err := newList.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 0 to be 1, but got %v", item.Value)
	}
}

func TestMergeEmptyList14(t *testing.T) {
	list := New[int]()
	newList := New[int]()
	newList.Append(1)
	newList.Merge(list)
	if newList.Size() != 1 {
		t.Errorf("Expected list to have size 1, but got %v", newList.Size())
	}
	item, err := newList.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 0 to be 1, but got %v", item.Value)
	}
}

func TestMergeEmptyList15(t *testing.T) {
	list := New[int]()
	list.Append(1)
	newList := New[int]()
	list.Merge(newList)
	list.Clear()
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestInsert(t *testing.T) {
	list := New[int]()
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
	list := New[int]()
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
	list := New[int]()
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
	list := New[int]()
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
	list := New[int]()
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
	list := New[int]()

	slice := list.ToSlice()

	if len(slice) != 0 {
		t.Errorf("Expected empty slice, but got length %d", len(slice))
	}
}

func TestFind(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test case 1: Value exists in the list
	node, err := list.Find(2)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if node.Value != 2 {
		t.Errorf("Expected value to be 2, but got %v", node.Value)
	}

	// Test case 2: Value does not exist in the list
	_, err = list.Find(4)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}

func TestDelete(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test deleting a value that exists in the list
	list.Delete(2)
	if list.Size() != 2 {
		t.Errorf("Expected list to have size 2, but got %v", list.Size())
	}
	if list.Contains(2) {
		t.Error("Expected list to not contain value 2")
	}

	// Test deleting a value that doesn't exist in the list
	list.Delete(4)
	if list.Size() != 2 {
		t.Errorf("Expected list to have size 2, but got %v", list.Size())
	}

	// Test deleting the first value in the list
	list.Delete(1)
	if list.Size() != 1 {
		t.Errorf("Expected list to have size 1, but got %v", list.Size())
	}
	if list.Contains(1) {
		t.Error("Expected list to not contain value 1")
	}

	// Test deleting the last value in the list
	list.Delete(3)
	if list.Size() != 0 {
		t.Errorf("Expected list to have size 0, but got %v", list.Size())
	}
	if list.Contains(3) {
		t.Error("Expected list to not contain value 3")
	}
}

func TestDeleteEmpty(t *testing.T) {
	list := New[int]()
	list.Delete(1)
	if list.Size() != 0 {
		t.Errorf("Expected list to have size 0, but got %v", list.Size())
	}
}

func TestInsertEmpty(t *testing.T) {
	list := New[int]()

	err := list.Insert(1)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	expected := []int{1}
	actual := list.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestInsertAfter(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.InsertAfter(2, 4)

	expected := []int{1, 2, 4, 3}
	actual := list.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestInsertAfterNotFound(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.InsertAfter(4, 5)

	expected := []int{1, 2, 3}
	actual := list.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestInsertBefore(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.InsertBefore(2, 4)

	expected := []int{1, 4, 2, 3}
	actual := list.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestInsertBeforeNotFound(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.InsertBefore(4, 5)

	expected := []int{1, 2, 3}
	actual := list.ToSlice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestDeleteLast(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.DeleteLast()
	if list.Size() != 2 {
		t.Errorf("Expected list to have size 2, but got %v", list.Size())
	}

	item, err := list.GetAt(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 2 {
		t.Errorf("Expected value at index 1 to be 2, but got %v", item.Value)
	}

	list.DeleteLast()
	if list.Size() != 1 {
		t.Errorf("Expected list to have size 1, but got %v", list.Size())
	}

	item, err = list.GetAt(0)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item.Value != 1 {
		t.Errorf("Expected value at index 0 to be 1, but got %v", item.Value)
	}

	list.DeleteLast()
	if list.Size() != 0 {
		t.Errorf("Expected list to have size 0, but got %v", list.Size())
	}

	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestDeleteLastEmpty(t *testing.T) {
	list := New[int]()
	list.DeleteLast()
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestDeleteFirst(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.DeleteFirst()

	if list.Size() != 2 {
		t.Errorf("Expected list to have size 2, but got %v", list.Size())
	}

	if list.GetFirst().Value != 2 {
		t.Errorf("Expected first element to be 2, but got %v", list.GetFirst().Value)
	}
}

func TestToSliceReverse(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	expected := []int{3, 2, 1}
	result := list.ToSliceReverse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestToSliceReverseEmpty(t *testing.T) {
	list := New[int]()

	expected := []int{}
	result := list.ToSliceReverse()

	if result != nil {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestToSliceReverseSingle(t *testing.T) {
	list := New[int]()
	list.Append(1)

	expected := []int{1}
	result := list.ToSliceReverse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestToSliceReverseDouble(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)

	expected := []int{2, 1}
	result := list.ToSliceReverse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestToSliceReverseTriple(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	expected := []int{4, 3, 2, 1}
	result := list.ToSliceReverse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestToSliceFromIndex(t *testing.T) {
	list := New[int]()
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
		result := list.ToSliceFromIndex(test.index)
		if result == nil && (test.index != 5 && test.index != 6) {
			t.Errorf("Expected ToSliceFromIndex(%d) to return %v, but got %v", test.index, test.expected, result)
		}
	}
}

func TestToSliceReverseFromIndex(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Test from index 0
	expected := []int{5, 4, 3, 2, 1}
	result := list.ToSliceReverseFromIndex(0)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test from index 1
	expected = []int{4, 3, 2, 1}
	result = list.ToSliceReverseFromIndex(1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test from index 2
	expected = []int{3, 2, 1}
	result = list.ToSliceReverseFromIndex(2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test from index 3
	expected = []int{2, 1}
	result = list.ToSliceReverseFromIndex(3)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test from index 4
	expected = []int{1}
	result = list.ToSliceReverseFromIndex(4)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test from index 5 (out of bounds)
	expected = []int{}
	result = list.ToSliceReverseFromIndex(5)
	if result != nil {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestForEach(t *testing.T) {
	list := New[int]()
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
	list := New[int]()

	list.ForEach(func(value *int) {
		t.Error("ForEach should not be called on an empty list")
	})
}

func TestAny(t *testing.T) {
	list := New[int]()
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
	list := New[int]()
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
	list := New[int]()
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
