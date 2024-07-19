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
