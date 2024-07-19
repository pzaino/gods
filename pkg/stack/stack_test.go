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

// Package stack provides a non-concurrent-safe stack (LIFO).
package stack

import (
	"testing"
)

func TestNew(t *testing.T) {
	stack := New[int]()
	if stack == nil {
		t.Error("Expected stack to be initialized, but got nil")
	}
}

func TestPush(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	if stack.IsEmpty() {
		t.Error("Expected stack to not be empty, but it was")
	}
	if len(stack.items) != 1 {
		t.Errorf("Expected stack to have 1 item, but got %v", len(stack.items))
	}
}

func TestPop(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	item, err := stack.Pop()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if item == nil {
		t.Error("Expected item to not be nil, but it was")
	} else if *item != 1 {
		t.Errorf("Expected item to be 1, but got %v", *item)
	}
}

func TestPopEmpty(t *testing.T) {
	stack := New[int]()
	_, err := stack.Pop()
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}

func TestIsEmpty(t *testing.T) {
	stack := New[int]()
	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty, but it was not")
	}
}

func TestIsEmptyAfterPush(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	if stack.IsEmpty() {
		t.Error("Expected stack to not be empty, but it was")
	}
}

func TestIsEmptyAfterPop(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	item, err := stack.Pop()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if item == nil {
		t.Error("Expected item to not be nil, but it was")
	}
	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty, but it was not")
	}
}

func TestToSlice(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	slice := stack.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected slice to have 3 items, but got %v", len(slice))
	}
	if slice[0] != 1 {
		t.Errorf("Expected first item to be 1, but got %v", slice[0])
	}
	if slice[1] != 2 {
		t.Errorf("Expected second item to be 2, but got %v", slice[1])
	}
	if slice[2] != 3 {
		t.Errorf("Expected third item to be 3, but got %v", slice[2])
	}
}

func TestReverse(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Reverse()
	slice := stack.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected slice to have 3 items, but got %v", len(slice))
	}
	if slice[0] != 3 {
		t.Errorf("Expected first item to be 3, but got %v", slice[0])
	}
	if slice[1] != 2 {
		t.Errorf("Expected second item to be 2, but got %v", slice[1])
	}
	if slice[2] != 1 {
		t.Errorf("Expected third item to be 1, but got %v", slice[2])
	}
}

func TestSwap(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	err := stack.Swap()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	slice := stack.ToSlice()
	if len(slice) != 2 {
		t.Errorf("Expected slice to have 2 items, but got %v", len(slice))
	}
	if slice[0] != 2 {
		t.Errorf("Expected first item to be 2, but got %v", slice[0])
	}
	if slice[1] != 1 {
		t.Errorf("Expected second item to be 1, but got %v", slice[1])
	}
}

func TestSwapEmpty(t *testing.T) {
	stack := New[int]()
	err := stack.Swap()
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}

func TestTop(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	item, err := stack.Top()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if item == nil {
		t.Error("Expected item to not be nil, but it was")
	} else if *item != 1 {
		t.Errorf("Expected item to be 1, but got %v", *item)
	}
}

func TestTopEmpty(t *testing.T) {
	stack := New[int]()
	_, err := stack.Top()
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}

func TestTopAfterPop(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	item, err := stack.Pop()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if item == nil {
		t.Error("Expected item to not be nil, but it was")
	}
	top, err := stack.Top()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if top == nil {
		t.Error("Expected top to not be nil, but it was")
	} else if *top != 1 {
		t.Errorf("Expected top to be 1, but got %v", *top)
	}
}

func TestTopAfterSwap(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	err := stack.Swap()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	top, err := stack.Top()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if top == nil {
		t.Error("Expected top to not be nil, but it was")
	} else if *top != 1 {
		t.Errorf("Expected top to be 1, but got %v", *top)
	}
}

func TestPeek(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	item, err := stack.Peek()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if item == nil {
		t.Error("Expected item to not be nil, but it was")
	} else if *item != 1 {
		t.Errorf("Expected item to be 1, but got %v", *item)
	}
}

func TestPeekEmpty(t *testing.T) {
	stack := New[int]()
	_, err := stack.Peek()
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}

func TestPeekAfterPop(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	item, err := stack.Pop()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if item == nil {
		t.Error("Expected item to not be nil, but it was")
	}
	top, err := stack.Peek()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if top == nil {
		t.Error("Expected top to not be nil, but it was")
	} else if *top != 1 {
		t.Errorf("Expected top to be 1, but got %v", *top)
	}
}

func TestPeekAfterSwap(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	err := stack.Swap()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	top, err := stack.Peek()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if top == nil {
		t.Error("Expected top to not be nil, but it was")
	} else if *top != 1 {
		t.Errorf("Expected top to be 1, but got %v", *top)
	}
}

func TestSize(t *testing.T) {
	stack := New[int]()
	if stack.Size() != 0 {
		t.Errorf("Expected stack to have 0 items, but got %v", stack.Size())
	}
	stack.Push(1)
	if stack.Size() != 1 {
		t.Errorf("Expected stack to have 1 item, but got %v", stack.Size())
	}
	stack.Push(2)
	if stack.Size() != 2 {
		t.Errorf("Expected stack to have 2 items, but got %v", stack.Size())
	}
	stack.Push(3)
	if stack.Size() != 3 {
		t.Errorf("Expected stack to have 3 items, but got %v", stack.Size())
	}
}

func TestClear(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Clear()
	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty, but it was not")
	}
}

func TestContains(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if !stack.Contains(1) {
		t.Error("Expected stack to contain 1, but it did not")
	}
	if !stack.Contains(2) {
		t.Error("Expected stack to contain 2, but it did not")
	}
	if !stack.Contains(3) {
		t.Error("Expected stack to contain 3, but it did not")
	}
	if stack.Contains(4) {
		t.Error("Expected stack to not contain 4, but it did")
	}
}

func TestCopy(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	copy := stack.Copy()
	if copy == nil {
		t.Error("Expected copy to not be nil, but it was")
	}
	if copy == stack {
		t.Error("Expected copy to be a different instance, but it was the same")
	}
	if copy.Size() != 3 {
		t.Errorf("Expected copy to have 3 items, but got %v", copy.Size())
	}
}

func TestEqual(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	other := New[int]()
	other.Push(1)
	other.Push(2)
	other.Push(3)
	if !stack.Equal(other) {
		t.Error("Expected stacks to be equal, but they were not")
	}
}

func TestEqualDifferentSize(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	other := New[int]()
	other.Push(1)
	other.Push(2)
	if stack.Equal(other) {
		t.Error("Expected stacks to not be equal, but they were")
	}
}

func TestEqualDifferentItems(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	other := New[int]()
	other.Push(1)
	other.Push(2)
	other.Push(4)
	if stack.Equal(other) {
		t.Error("Expected stacks to not be equal, but they were")
	}
}

func TestEqualDifferentOrder(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	other := New[int]()
	other.Push(3)
	other.Push(2)
	other.Push(1)
	if stack.Equal(other) {
		t.Error("Expected stacks to not be equal, but they were")
	}
}

func TestEqualEmpty(t *testing.T) {
	stack := New[int]()
	other := New[int]()
	if !stack.Equal(other) {
		t.Error("Expected stacks to be equal, but they were not")
	}
}

func TestEqualEmptyOther(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	other := New[int]()
	if stack.Equal(other) {
		t.Error("Expected stacks to not be equal, but they were")
	}
}

func TestEqualEmptyBoth(t *testing.T) {
	stack := New[int]()
	other := New[int]()
	if !stack.Equal(other) {
		t.Error("Expected stacks to be equal, but they were not")
	}
}

func TestEqualNil(t *testing.T) {
	stack := New[int]()
	if stack.Equal(nil) {
		t.Error("Expected stack to not be equal to nil, but it was")
	}
}
