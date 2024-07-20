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
	"sync"
	"testing"
)

const (
	errListNotEmpty = "Expected list to be empty, but it was not"
	errListIsEmpty  = "Expected list to not be empty, but it was"
)

func TestNew(t *testing.T) {
	list := CSLinkListNew[int]()
	if list == nil {
		t.Error("Expected list to be initialized, but got nil")
	}
}

func TestAppend(t *testing.T) {
	list := CSLinkListNew[int]()
	list.Append(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list to have 1 item, but got %v", list.Size())
	}
}

func TestPrepend(t *testing.T) {
	list := CSLinkListNew[int]()
	list.Prepend(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list to have 1 item, but got %v", list.Size())
	}
}

func TestRemove(t *testing.T) {
	list := CSLinkListNew[int]()
	list.Append(1)
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestRemoveEmpty(t *testing.T) {
	list := CSLinkListNew[int]()
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestIsEmpty(t *testing.T) {
	list := CSLinkListNew[int]()
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestIsEmptyAfterAppend(t *testing.T) {
	list := CSLinkListNew[int]()
	list.Append(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
}

func TestIsEmptyAfterPrepend(t *testing.T) {
	list := CSLinkListNew[int]()
	list.Prepend(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
}

func TestIsEmptyAfterRemove(t *testing.T) {
	list := CSLinkListNew[int]()
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
	list := CSLinkListNew[int]()
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
	list := CSLinkListNew[int]()
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
	list := CSLinkListNew[int]()

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

func TestConcurrentInsertAt(t *testing.T) {
	list := CSLinkListNew[int]()
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
	list := CSLinkListNew[int]()

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
	list := CSLinkListNew[int]()

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
	list1 := CSLinkListNew[int]()
	list2 := CSLinkListNew[int]()

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
	list := CSLinkListNew[int]()

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
