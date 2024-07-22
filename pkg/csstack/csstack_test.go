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

// Package csstack is a concurrent-safe stack library (LIFO).
package csstack

import (
	"reflect"
	"strings"
	"testing"
)

// TestCSStack tests the CSStack type.
func TestCSStack(t *testing.T) {
	s := NewCSStack[int]()

	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.IsEmpty() {
		t.Errorf("IsEmpty() = true; want false")
	}

	item, err := s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 3 {
		t.Errorf("Pop() = %v; want 3", *item)
	}

	if s.IsEmpty() {
		t.Errorf("IsEmpty() = true; want false")
	}

	item, err = s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 2 {
		t.Errorf("Pop() = %v; want 2", *item)
	}

	item, err = s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 1 {
		t.Errorf("Pop() = %v; want 1", *item)
	}

	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}

	item, err = s.Pop()
	if err == nil {
		t.Errorf("Pop() error = nil; want error")
	}
	if item != nil {
		t.Errorf("Pop() = %v; want nil", item)
	}
}

// TestCSStackTop tests the CSStack Top method.
func TestCSStackTop(t *testing.T) {
	s := NewCSStack[int]()

	_, err := s.Top()
	if err == nil {
		t.Errorf("Top() error = nil; want error")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	item, err := s.Top()
	if err != nil {
		t.Errorf("Top() error = %v; want nil", err)
	}
	if *item != 3 {
		t.Errorf("Top() = %v; want 3", *item)
	}
}

// TestCSStackPeek tests the CSStack Peek method.
func TestCSStackPeek(t *testing.T) {
	s := NewCSStack[int]()

	_, err := s.Peek()
	if err == nil {
		t.Errorf("Peek() error = nil; want error")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	item, err := s.Peek()
	if err != nil {
		t.Errorf("Peek() error = %v; want nil", err)
	}
	if *item != 3 {
		t.Errorf("Peek() = %v; want 3", *item)
	}
}

// TestCSStackSize tests the CSStack Size method.
func TestCSStackSize(t *testing.T) {
	s := NewCSStack[int]()

	if size := s.Size(); size != 0 {
		t.Errorf("Size() = %v; want 0", size)
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if size := s.Size(); size != 3 {
		t.Errorf("Size() = %v; want 3", size)
	}
}

// TestCSStackClear tests the CSStack Clear method.
func TestCSStackClear(t *testing.T) {
	s := NewCSStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	s.Clear()

	if size := s.Size(); size != 0 {
		t.Errorf("Size() = %v; want 0", size)
	}
}

// TestCSStackContains tests the CSStack Contains method.
func TestCSStackContains(t *testing.T) {
	s := NewCSStack[int]()

	if contains := s.Contains(1); contains {
		t.Errorf("Contains() = true; want false")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if contains := s.Contains(2); !contains {
		t.Errorf("Contains() = false; want true")
	}
}

// TestCSStackCopy tests the CSStack Copy method.
func TestCSStackCopy(t *testing.T) {
	s := NewCSStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	stack := s.Copy()

	if size := stack.Size(); size != 3 {
		t.Errorf("Size() = %v; want 3", size)
	}

	item, err := stack.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 3 {
		t.Errorf("Pop() = %v; want 3", *item)
	}

	item, err = stack.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 2 {
		t.Errorf("Pop() = %v; want 2", *item)
	}

	item, err = stack.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 1 {
		t.Errorf("Pop() = %v; want 1", *item)
	}

	if !stack.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}
}

// TestCSStackEqual tests the CSStack Equal method.
func TestCSStackEqual(t *testing.T) {
	s1 := NewCSStack[int]()
	s2 := NewCSStack[int]()

	if equal := s1.Equal(s2); !equal {
		t.Errorf("Equal() = false; want true")
	}

	s1.Push(1)
	s1.Push(2)
	s1.Push(3)

	if equal := s1.Equal(s2); equal {
		t.Errorf("Equal() = true; want false")
	}

	s2.Push(1)
	s2.Push(2)
	s2.Push(3)

	if equal := s1.Equal(s2); !equal {
		t.Errorf("Equal() = false; want true")
	}
}

// TestCSStackString tests the CSStack String method.
func TestCSStackString(t *testing.T) {
	s := NewCSStack[int]()

	if str := s.String(); str != "[]" {
		t.Errorf("String() = %v; want []", str)
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if str := s.String(); str != "[1 2 3]" {
		t.Errorf("String() = %v; want [1 2 3]", str)
	}
}

// TestCSStackConcurrent tests the CSStack type in a concurrent environment.
func TestCSStackConcurrent(t *testing.T) {
	s := NewCSStack[int]()

	const n = 100
	const m = 10

	done := make(chan struct{})
	defer close(done)

	for i := 0; i < m; i++ {
		go func() {
			for j := 0; j < n; j++ {
				s.Push(j)
				item, err := s.Pop()
				if err != nil {
					if strings.ToLower(strings.TrimSpace(err.Error())) != "stack is empty" {
						t.Errorf("Pop() error = %v; want nil", err)
					}
				}
				if err == nil && item == nil {
					t.Errorf("Expected non-nil item")
				}
				item, err = s.Top()
				if err != nil {
					if strings.ToLower(strings.TrimSpace(err.Error())) != "stack is empty" {
						t.Errorf("Top() error = %v; want nil", err)
					}
				}
				if err == nil && item == nil {
					t.Errorf("Top() = nil; want item")
				}
				item, err = s.Peek()
				if err != nil {
					if err.Error() != "stack is empty" {
						t.Errorf("Peek() error = %v; want nil", err)
					}
				}
				if err == nil && item == nil {
					t.Errorf("Peek() = nil; want item")
				}
				s.Size()
				s.Clear()
				s.Contains(j)
				s.Copy()
			}
			done <- struct{}{}
		}()
	}

	for i := 0; i < m; i++ {
		<-done
	}
}

// TestCSStackPopN tests the CSStack PopN method.
func TestCSStackPopN(t *testing.T) {
	s := NewCSStack[int]()

	_, err := s.PopN(1)
	if err == nil {
		t.Errorf("PopN() error = nil; want error")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	items, err := s.PopN(2)
	if err != nil {
		t.Errorf("PopN() error = %v; want nil", err)
	}
	if len(items) != 2 {
		t.Errorf("PopN() returned %d items; want 2", len(items))
	}
	if items[0] != 3 {
		t.Errorf("PopN() item[0] = %v; want 3", items[0])
	}
	if items[1] != 2 {
		t.Errorf("PopN() item[1] = %v; want 2", items[1])
	}

	if size := s.Size(); size != 1 {
		t.Errorf("Size() = %v; want 1", size)
	}

	item, err := s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 1 {
		t.Errorf("Pop() = %v; want 1", *item)
	}

	if size := s.Size(); size != 0 {
		t.Errorf("Size() = %v; want 0", size)
	}

	_, err = s.PopN(1)
	if err == nil {
		t.Errorf("PopN() error = nil; want error")
	}
}

// TestCSStackPushN tests the CSStack PushN method.
func TestCSStackPushN(t *testing.T) {
	s := NewCSStack[int]()

	s.PushN(1, 2, 3)

	if size := s.Size(); size != 3 {
		t.Errorf("Size() = %v; want 3", size)
	}

	item, err := s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 3 {
		t.Errorf("Pop() = %v; want 3", *item)
	}

	item, err = s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 2 {
		t.Errorf("Pop() = %v; want 2", *item)
	}

	item, err = s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 1 {
		t.Errorf("Pop() = %v; want 1", *item)
	}

	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}
}

// TestCSStackPopAll tests the CSStack PopAll method.
func TestCSStackPopAll(t *testing.T) {
	s := NewCSStack[int]()

	items := s.PopAll()
	if len(items) != 0 {
		t.Errorf("PopAll() returned %d items; want 0", len(items))
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	items = s.PopAll()
	if len(items) != 3 {
		t.Errorf("PopAll() returned %d items; want 3", len(items))
	}
	if items[0] != 3 {
		t.Errorf("PopAll() item[0] = %v; want 3", items[0])
	}
	if items[1] != 2 {
		t.Errorf("PopAll() item[1] = %v; want 2", items[1])
	}
	if items[2] != 1 {
		t.Errorf("PopAll() item[2] = %v; want 1", items[2])
	}

	if size := s.Size(); size != 0 {
		t.Errorf("Size() = %v; want 0", size)
	}
}

// TestCSStackPushAll tests the CSStack PushAll method.
func TestCSStackPushAll(t *testing.T) {
	s := NewCSStack[int]()

	items := []int{1, 2, 3}
	s.PushAll(items)

	if size := s.Size(); size != 3 {
		t.Errorf("Size() = %v; want 3", size)
	}

	item, err := s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 3 {
		t.Errorf("Pop() = %v; want 3", *item)
	}

	item, err = s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 2 {
		t.Errorf("Pop() = %v; want 2", *item)
	}

	item, err = s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 1 {
		t.Errorf("Pop() = %v; want 1", *item)
	}

	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}
}

// TestCSStackFilter tests the CSStack Filter method.
func TestCSStackFilter(t *testing.T) {
	s := NewCSStack[int]()

	// Test filtering an empty stack
	s.Filter(func(item int) bool {
		return item > 0
	})
	if !s.IsEmpty() {
		t.Errorf("Filter() on empty stack modified the stack")
	}

	// Test filtering a non-empty stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	s.Filter(func(item int) bool {
		return item%2 == 0
	})

	// Verify that only even numbers are left in the stack
	expected := []int{2}
	actual := s.ToSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Filter() failed, expected %v, got %v", expected, actual)
	}
}

// TestCSStackMap tests the CSStack Map method.
func TestCSStackMap(t *testing.T) {
	s := NewCSStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Define the mapping function
	fn := func(item int) int {
		return item * 2
	}

	// Apply the mapping function to the stack
	mappedStack := s.Map(fn)

	// Verify the size of the mapped stack
	if size := mappedStack.Size(); size != 3 {
		t.Errorf("Size() = %v; want 3", size)
	}

	// Verify the items in the mapped stack
	expectedItems := []int{6, 4, 2}
	for _, expectedItem := range expectedItems {
		item, err := mappedStack.Pop()
		if err != nil {
			t.Errorf("Pop() error = %v; want nil", err)
		}
		if *item != expectedItem {
			t.Errorf("Pop() = %v; want %v", *item, expectedItem)
		}
	}

	// Verify that the original stack is unchanged
	if size := s.Size(); size != 3 {
		t.Errorf("Size() = %v; want 3", size)
	}
}

// TestCSStackReduce tests the CSStack Reduce method.
func TestCSStackReduce(t *testing.T) {
	s := NewCSStack[int]()

	// Test reducing an empty stack
	_, err := s.Reduce(func(a, b int) int { return a + b })
	if err == nil {
		t.Errorf("Reduce() error = nil; want error")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test reducing the stack with the addition function
	result, err := s.Reduce(func(a, b int) int { return a + b })
	if err != nil {
		t.Errorf("Reduce() error = %v; want nil", err)
	}
	if result != 6 {
		t.Errorf("Reduce() result = %v; want 6", result)
	}

	// Test reducing the stack with the multiplication function
	result, err = s.Reduce(func(a, b int) int { return a * b })
	if err != nil {
		t.Errorf("Reduce() error = %v; want nil", err)
	}
	if result != 6 {
		t.Errorf("Reduce() result = %v; want 6", result)
	}

	// Test reducing the stack with the subtraction function
	result, err = s.Reduce(func(a, b int) int { return a - b })
	if err != nil {
		t.Errorf("Reduce() error = %v; want nil", err)
	}
	if result != -4 {
		t.Errorf("Reduce() result = %v; want -4", result)
	}
}

// TestCSStackForEach tests the CSStack ForEach method.
func TestCSStackForEach(t *testing.T) {
	s := NewCSStack[int]()

	// Test ForEach on an empty stack
	s.ForEach(func(item *int) {
		t.Errorf("ForEach() called on an empty stack")
	})

	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test ForEach on a non-empty stack
	var sum int
	s.ForEach(func(item *int) {
		sum += *item
	})

	if sum != 6 {
		t.Errorf("ForEach() sum = %v; want 6", sum)
	}
}

// TestCSStackAny tests the CSStack Any method.
func TestCSStackAny(t *testing.T) {
	s := NewCSStack[int]()

	// Test with an empty stack
	any := s.Any(func(item int) bool {
		return item > 0
	})
	if any {
		t.Errorf("Any() = true; want false")
	}

	// Test with a non-empty stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	any = s.Any(func(item int) bool {
		return item > 2
	})
	if !any {
		t.Errorf("Any() = false; want true")
	}

	any = s.Any(func(item int) bool {
		return item < 0
	})
	if any {
		t.Errorf("Any() = true; want false")
	}
}

// TestCSStackAll tests the CSStack All method.
func TestCSStackAll(t *testing.T) {
	s := NewCSStack[int]()

	// Test with an empty stack
	if all := s.All(func(item int) bool {
		return item > 0
	}); all {
		t.Errorf("All() = true; want false")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test with a predicate that returns true for all items
	if all := s.All(func(item int) bool {
		return item > 0
	}); !all {
		t.Errorf("All() = false; want true")
	}

	// Test with a predicate that returns false for all items
	if all := s.All(func(item int) bool {
		return item > 3
	}); all {
		t.Errorf("All() = true; want false")
	}

	// Test with a predicate that returns true for some items
	if all := s.All(func(item int) bool {
		return item > 1
	}); all {
		t.Errorf("All() = true; want false")
	}
}

// TestCSStackFind tests the CSStack Find method.
func TestCSStackFind(t *testing.T) {
	s := NewCSStack[int]()

	// Test finding an item in an empty stack
	item, err := s.Find(func(i int) bool {
		return i == 1
	})
	if err == nil {
		t.Errorf("Find() error = nil; want error")
	}
	if item != nil {
		t.Errorf("Find() = %v; want nil", item)
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test finding an item that exists in the stack
	item, err = s.Find(func(i int) bool {
		return i == 2
	})
	if err != nil {
		t.Errorf("Find() error = %v; want nil", err)
	}
	if *item != 2 {
		t.Errorf("Find() = %v; want 2", *item)
	}

	// Test finding an item that doesn't exist in the stack
	item, err = s.Find(func(i int) bool {
		return i == 4
	})
	if err == nil {
		t.Errorf("Find() error = nil; want error")
	}
	if item != nil {
		t.Errorf("Find() = %v; want nil", item)
	}
}

// TestCSStackFindIndex tests the CSStack FindIndex method.
func TestCSStackFindIndex(t *testing.T) {
	s := NewCSStack[int]()

	// Test finding an item in an empty stack
	index, err := s.FindIndex(func(item int) bool {
		return item == 1
	})
	if err == nil {
		t.Errorf("FindIndex() error = nil; want error")
	}
	if index != -1 {
		t.Errorf("FindIndex() index = %v; want -1", index)
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test finding an item that exists in the stack
	index, err = s.FindIndex(func(item int) bool {
		return item == 2
	})
	if err != nil {
		t.Errorf("FindIndex() error = %v; want nil", err)
	}
	if index != 1 {
		t.Errorf("FindIndex() index = %v; want 1", index)
	}

	// Test finding an item that doesn't exist in the stack
	index, err = s.FindIndex(func(item int) bool {
		return item == 4
	})
	if err == nil {
		t.Errorf("FindIndex() error = nil; want error")
	}
	if index != -1 {
		t.Errorf("FindIndex() index = %v; want -1", index)
	}
}

// TestCSStackFindLast tests the CSStack FindLast method.
func TestCSStackFindLast(t *testing.T) {
	s := NewCSStack[int]()

	// Test finding in an empty stack
	item, err := s.FindLast(func(i int) bool {
		return i == 1
	})
	if err == nil {
		t.Errorf("FindLast() error = nil; want error")
	}
	if item != nil {
		t.Errorf("FindLast() = %v; want nil", item)
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test finding the last item that matches the predicate
	item, err = s.FindLast(func(i int) bool {
		return i%2 == 0
	})
	if err != nil {
		t.Errorf("FindLast() error = %v; want nil", err)
	}
	if *item != 2 {
		t.Errorf("FindLast() = %v; want 2", *item)
	}

	// Test finding the last item that doesn't match the predicate
	item, err = s.FindLast(func(i int) bool {
		return i == 4
	})
	if err == nil {
		t.Errorf("FindLast() error = nil; want error")
	}
	if item != nil {
		t.Errorf("FindLast() = %v; want nil", item)
	}
}

// TestCSStackFindLastIndex tests the CSStack FindLastIndex method.
func TestCSStackFindLastIndex(t *testing.T) {
	s := NewCSStack[int]()

	// Test finding the last index of an item in an empty stack
	index, err := s.FindLastIndex(func(item int) bool {
		return item == 1
	})
	if err == nil {
		t.Errorf("FindLastIndex() error = nil; want error")
	}
	if index != -1 {
		t.Errorf("FindLastIndex() = %v; want -1", index)
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test finding the last index of an item that exists in the stack
	index, err = s.FindLastIndex(func(item int) bool {
		return item == 2
	})
	if err != nil {
		t.Errorf("FindLastIndex() error = %v; want nil", err)
	}
	if index != 1 {
		t.Errorf("FindLastIndex() = %v; want 1", index)
	}

	// Test finding the last index of an item that doesn't exist in the stack
	index, err = s.FindLastIndex(func(item int) bool {
		return item == 4
	})
	if err == nil {
		t.Errorf("FindLastIndex() error = nil; want item not found")
	}
	if index != -1 {
		t.Errorf("FindLastIndex() = %v; want -1", index)
	}
}

// TestCSStackFindAll tests the CSStack FindAll method.
func TestCSStackFindAll(t *testing.T) {
	s := NewCSStack[int]()

	// Test finding all items in an empty stack
	items := s.FindAll(func(item int) bool {
		return item > 0
	})
	if len(items) != 0 {
		t.Errorf("FindAll() returned %d items; want 0", len(items))
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test finding all items in a non-empty stack
	items = s.FindAll(func(item int) bool {
		return item > 1
	})
	if len(items) != 2 {
		t.Errorf("FindAll() returned %d items; want 2", len(items))
	}
	if items[0] != 2 {
		t.Errorf("FindAll() item[0] = %v; want 2", items[0])
	}
	if items[1] != 3 {
		t.Errorf("FindAll() item[1] = %v; want 3", items[1])
	}

	// Test finding all items that don't exist in the stack
	items = s.FindAll(func(item int) bool {
		return item > 10
	})
	if len(items) != 0 {
		t.Errorf("FindAll() returned %d items; want 0", len(items))
	}
}

// TestCSStackFindIndices tests the CSStack FindIndices method.
func TestCSStackFindIndices(t *testing.T) {
	s := NewCSStack[int]()

	// Test finding indices in an empty stack
	indices := s.FindIndices(func(item int) bool {
		return item > 0
	})
	if len(indices) != 0 {
		t.Errorf("FindIndices() returned %d indices; want 0", len(indices))
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test finding indices in a non-empty stack
	indices = s.FindIndices(func(item int) bool {
		return item > 1
	})
	expectedIndices := []int{1, 2}
	if !reflect.DeepEqual(indices, expectedIndices) {
		t.Errorf("FindIndices() returned %v indices; want %v", indices, expectedIndices)
	}

	// Test finding indices with no matching items
	indices = s.FindIndices(func(item int) bool {
		return item > 10
	})
	if len(indices) != 0 {
		t.Errorf("FindIndices() returned %d indices; want 0", len(indices))
	}
}

func TestCSStackToStack(t *testing.T) {
	s := NewCSStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	stack := s.ToStack()

	if size := stack.Size(); size != 3 {
		t.Errorf("Size() = %v; want 3", size)
	}

	item, err := stack.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 3 {
		t.Errorf("Pop() = %v; want 3", *item)
	}

	item, err = stack.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 2 {
		t.Errorf("Pop() = %v; want 2", *item)
	}

	item, err = stack.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 1 {
		t.Errorf("Pop() = %v; want 1", *item)
	}

	if !stack.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}
}

// TestCSStackReverse tests the CSStack Reverse method.
func TestCSStackReverse(t *testing.T) {
	s := NewCSStack[int]()

	// Test reversing an empty stack
	s.Reverse()
	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}

	// Test reversing a stack with one item
	s.Push(1)
	s.Reverse()
	if size := s.Size(); size != 1 {
		t.Errorf("Size() = %v; want 1", size)
	}
	item, err := s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 1 {
		t.Errorf("Pop() = %v; want 1", *item)
	}

	// Test reversing a stack with multiple items
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Reverse()
	if size := s.Size(); size != 3 {
		t.Errorf("Size() = %v; want 3", size)
	}
	item, err = s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 1 {
		t.Errorf("Pop() = %v; want 1", *item)
	}
	item, err = s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 2 {
		t.Errorf("Pop() = %v; want 2", *item)
	}
	item, err = s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 3 {
		t.Errorf("Pop() = %v; want 3", *item)
	}
}

// TestCSStackSwap tests the CSStack Swap method.
func TestCSStackSwap(t *testing.T) {
	s := NewCSStack[int]()

	// Test swapping an empty stack
	err := s.Swap()
	if err == nil {
		t.Errorf("Swap() error = nil; want error")
	}

	s.Push(1)

	// Test swapping a stack with one item
	err = s.Swap()
	if err == nil {
		t.Errorf("Swap() error = nil; want error")
	}

	s.Push(2)
	s.Push(3)

	// Test swapping a stack with multiple items
	err = s.Swap()
	if err != nil {
		t.Errorf("Swap() error = %v; want nil", err)
	}

	item, err := s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 2 {
		t.Errorf("Pop() = %v; want 2", *item)
	}

	item, err = s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 3 {
		t.Errorf("Pop() = %v; want 1", *item)
	}
}
