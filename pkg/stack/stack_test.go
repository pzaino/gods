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
package stack_test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	stack "github.com/pzaino/gods/pkg/stack"
)

const (
	errNoError         = "Expected no error, but got %v"
	errYesError        = "Expected an error, but got nil"
	errItemNotNil      = "Expected item to not be nil, but it was"
	errExpectedItemX   = "Expected item to be %v, but got %v"
	errExpectedXItemY  = "Expected %v item to be %v, but got %v"
	errStackNotEmpty   = "Expected stack to be empty, but it was not"
	errExpectedStack   = "Expected stack to be %v, but got %v"
	errExpectedResult  = "Expected result to be %v, but got %v"
	errExpected2Stacks = "Expected both stacks to be equal, but they were not"
)

func TestNew(t *testing.T) {
	stack := stack.NewStack[int]()
	if stack == nil {
		t.Error("Expected stack to be initialized, but got nil")
	}
}

func TestPush(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	if s.IsEmpty() {
		t.Error("Expected stack to not be empty, but it was")
	}
	if s.Size() != 1 {
		t.Errorf("Expected stack to have 1 item, but got %v", s.Size())
	}
}

func TestPop(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	item, err := s.Pop()
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item == nil {
		t.Error(errItemNotNil)
	} else if *item != 1 {
		t.Errorf(errExpectedItemX, 1, *item)
	}
}

func TestPopEmpty(t *testing.T) {
	s := stack.NewStack[int]()
	_, err := s.Pop()
	if err == nil {
		t.Error(errYesError)
	}
}

func TestIsEmpty(t *testing.T) {
	s := stack.NewStack[int]()
	if !s.IsEmpty() {
		t.Error(errStackNotEmpty)
	}
}

func TestIsEmptyAfterPush(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	if s.IsEmpty() {
		t.Error("Expected stack to not be empty, but it was")
	}
}

func TestIsEmptyAfterPop(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	item, err := s.Pop()
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item == nil {
		t.Error(errItemNotNil)
	}
	if !s.IsEmpty() {
		t.Error(errStackNotEmpty)
	}
}

func TestToSlice(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected slice to have 3 items, but got %v", len(slice))
	}
	if slice[0] != 1 {
		t.Errorf(errExpectedXItemY, "first", 1, slice[0])
	}
	if slice[1] != 2 {
		t.Errorf(errExpectedXItemY, "second", 2, slice[1])
	}
	if slice[2] != 3 {
		t.Errorf(errExpectedXItemY, "third", 3, slice[2])
	}
}

func TestReverse(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Reverse()
	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected slice to have 3 items, but got %v", len(slice))
	}
	if slice[0] != 3 {
		t.Errorf(errExpectedXItemY, "first", 3, slice[0])
	}
	if slice[1] != 2 {
		t.Errorf(errExpectedXItemY, "second", 2, slice[1])
	}
	if slice[2] != 1 {
		t.Errorf(errExpectedXItemY, "third", 1, slice[2])
	}
}

func TestSwap(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	err := s.Swap()
	if err != nil {
		t.Errorf(errNoError, err)
	}
	slice := s.ToSlice()
	if len(slice) != 2 {
		t.Errorf("Expected slice to have 2 items, but got %v", len(slice))
	}
	if slice[0] != 2 {
		t.Errorf(errExpectedXItemY, "first", 2, slice[0])
	}
	if slice[1] != 1 {
		t.Errorf(errExpectedXItemY, "second", 1, slice[1])
	}
}

func TestSwapEmpty(t *testing.T) {
	s := stack.NewStack[int]()
	err := s.Swap()
	if err == nil {
		t.Error(errYesError)
	}
}

func TestTop(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	item, err := s.Top()
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item == nil {
		t.Error(errItemNotNil)
	} else if *item != 1 {
		t.Errorf(errExpectedItemX, 1, *item)
	}
}

func TestTopEmpty(t *testing.T) {
	s := stack.NewStack[int]()
	_, err := s.Top()
	if err == nil {
		t.Error(errYesError)
	}
}

func TestTopAfterPop(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	item, err := s.Pop()
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item == nil {
		t.Error(errItemNotNil)
	}
	top, err := s.Top()
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if top == nil {
		t.Error("Expected the top of the stack to not be nil, but it was")
	} else if *top != 1 {
		t.Errorf("Expected the top to be 1, but got %v", *top)
	}
}

func TestTopAfterSwap(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	err := s.Swap()
	if err != nil {
		t.Errorf(errNoError, err)
	}
	top, err := s.Top()
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if top == nil {
		t.Error("Expected stack's top to not be nil, but it was")
	} else if *top != 1 {
		t.Errorf("Expected stack's top to be 1, but got %v", *top)
	}
}

func TestPeek(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	item, err := s.Peek()
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item == nil {
		t.Error(errItemNotNil)
	} else if *item != 1 {
		t.Errorf(errExpectedItemX, 1, *item)
	}
}

func TestPeekEmpty(t *testing.T) {
	s := stack.NewStack[int]()
	_, err := s.Peek()
	if err == nil {
		t.Error(errYesError)
	}
}

func TestPeekAfterPop(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	item, err := s.Pop()
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item == nil {
		t.Error(errItemNotNil)
	}
	top, err := s.Peek()
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if top == nil {
		t.Error("Expected top to not be nil, but it was")
	} else if *top != 1 {
		t.Errorf("Expected top to be 1, but got %v", *top)
	}
}

func TestPeekAfterSwap(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	err := s.Swap()
	if err != nil {
		t.Errorf(errNoError, err)
	}
	top, err := s.Peek()
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if top == nil {
		t.Error("Expected top to not be nil, but it was")
	} else if *top != 1 {
		t.Errorf("Expected top to be 1, but got %v", *top)
	}
}

func TestSize(t *testing.T) {
	s := stack.NewStack[int]()
	if s.Size() != 0 {
		t.Errorf("Expected stack to have 0 items, but got %v", s.Size())
	}
	s.Push(1)
	if s.Size() != 1 {
		t.Errorf("Expected stack to have 1 item, but got %v", s.Size())
	}
	s.Push(2)
	if s.Size() != 2 {
		t.Errorf("Expected stack to have 2 items, but got %v", s.Size())
	}
	s.Push(3)
	if s.Size() != 3 {
		t.Errorf("Expected stack to have 3 items, but got %v", s.Size())
	}
}

func TestClear(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Clear()
	if !s.IsEmpty() {
		t.Error(errStackNotEmpty)
	}
}

func TestContains(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if !s.Contains(1) {
		t.Error("Expected stack to contain 1, but it did not")
	}
	if !s.Contains(2) {
		t.Error("Expected stack to contain 2, but it did not")
	}
	if !s.Contains(3) {
		t.Error("Expected stack to contain 3, but it did not")
	}
	if s.Contains(4) {
		t.Error("Expected stack to not contain 4, but it did")
	}
}

func TestCopy(t *testing.T) {
	stack := stack.NewStack[int]()
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
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	other := stack.NewStack[int]()
	other.Push(1)
	other.Push(2)
	other.Push(3)
	if !s.Equal(other) {
		t.Error("Expected stacks to be equal, but they were not")
	}
}

func TestEqualDifferentSize(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	other := stack.NewStack[int]()
	other.Push(1)
	other.Push(2)
	if s.Equal(other) {
		t.Error("Expected stacks to not be identical, but they were")
	}
}

func TestEqualDifferentItems(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	other := stack.NewStack[int]()
	other.Push(1)
	other.Push(2)
	other.Push(4)
	if s.Equal(other) {
		t.Error("Expected stacks to not be the same, but they were")
	}
}

func TestEqualDifferentOrder(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	other := stack.NewStack[int]()
	other.Push(3)
	other.Push(2)
	other.Push(1)
	if s.Equal(other) {
		t.Error("Expected stacks to not be equal, but they were")
	}
}

func TestEqualEmpty(t *testing.T) {
	s := stack.NewStack[int]()
	other := stack.NewStack[int]()
	if !s.Equal(other) {
		t.Error("Expected stacks to be equal, but they were not")
	}
}

func TestEqualEmptyOther(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	other := stack.NewStack[int]()
	if s.Equal(other) {
		t.Error("Expected stacks to not be equal, but they were")
	}
}

func TestEqualNil(t *testing.T) {
	s := stack.NewStack[int]()
	if s.Equal(nil) {
		t.Error("Expected stack to not be equal to nil, but it was")
	}
}
func TestString(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	expected := "[1 2 3]"
	result := s.String()
	if result != expected {
		t.Errorf("Expected string representation to be %q, but got %q", expected, result)
	}
}

func TestStringEmpty(t *testing.T) {
	s := stack.NewStack[int]()
	expected := "[]"
	result := s.String()
	if result != expected {
		t.Errorf("Expected string representation to be %q, but got %q", expected, result)
	}
}

func TestPopN(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test case 1: Pop 2 items
	items, err := s.PopN(2)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if len(items) != 2 {
		t.Errorf("Expected 2 items, but got %v", len(items))
	}
	if items[0] != 3 {
		t.Errorf(errExpectedXItemY, "first", 2, items[0])
	}
	if items[1] != 2 {
		t.Errorf(errExpectedXItemY, "second", 3, items[1])
	}
	if s.Size() != 1 {
		t.Errorf("Expected stack size to be 1, but got %v", s.Size())
	}

	// Test case 2: Pop 1 item
	items, err = s.PopN(1)
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if len(items) != 1 {
		t.Errorf("Expected 1 item, but got %v", len(items))
	}
	if items[0] != 1 {
		t.Errorf(errExpectedItemX, 1, items[0])
	}
	if s.Size() != 0 {
		t.Errorf("Expected stack size to be 0, but got %v", s.Size())
	}

	// Test case 3: Pop from empty stack
	items, err = s.PopN(1)
	if err == nil {
		t.Error(errYesError)
	}
	if items != nil {
		t.Errorf("Expected items to be nil, but got %v", items)
	}
}

func TestPushN(t *testing.T) {
	s := stack.NewStack[int]()
	s.PushN(1, 2, 3)
	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected stack to have 3 items, but got %v", len(slice))
	}
	if slice[0] != 1 {
		t.Errorf(errExpectedXItemY, "first", 1, slice[0])
	}
	if slice[1] != 2 {
		t.Errorf(errExpectedXItemY, "second", 3, slice[1])
	}
	if slice[2] != 3 {
		t.Errorf(errExpectedXItemY, "third", 3, slice[2])
	}
}

func TestPopAll(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	items := s.PopAll()
	if len(items) != 3 {
		t.Errorf("Expected %d items, but got %d", 3, len(items))
	}
	if s.Size() != 0 {
		t.Errorf("Expected stack to be empty, but it has %d items", s.Size())
	}
	if items[0] != 3 {
		t.Errorf(errExpectedXItemY, "first", 3, items[0])
	}
	if items[1] != 2 {
		t.Errorf(errExpectedXItemY, "second", 2, items[1])
	}
	if items[2] != 1 {
		t.Errorf(errExpectedXItemY, "third", 1, items[2])
	}
}

func TestPushAll(t *testing.T) {
	s := stack.NewStack[int]()
	items := []int{1, 2, 3}
	s.PushAll(items)

	if s.Size() != uint64(len(items)) {
		t.Errorf("Expected stack size to be %d, but got %d", len(items), s.Size())
	}

	slice := s.ToSlice()
	for i, item := range items {
		if slice[i] != item {
			t.Errorf("Expected item at index %d to be %d, but got %d", i, item, slice[i])
		}
	}
}

func TestFilter(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	// Test filtering even numbers
	s.Filter(func(item int) bool {
		return item%2 == 0
	})

	// Check if the stack contains only even numbers
	if !s.Contains(2) {
		t.Error("Expected stack to contain 2, but it did not")
	}
	if !s.Contains(4) {
		t.Error("Expected stack to contain 4, but it did not")
	}
	if s.Contains(1) {
		t.Error("Expected stack to not contain 1, but it did")
	}
	if s.Contains(3) {
		t.Error("Expected stack to not contain 3, but it did")
	}
}

func TestMap(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test mapping to double the values
	doubledStack := s.Map(func(item int) int {
		return item * 2
	})

	// Check if the original stack is unchanged
	if s.Size() != 3 {
		t.Errorf("Expected original stack to have 3 items, but got %v", s.Size())
	}

	// Check if the new stack has the mapped values
	doubledSlice := doubledStack.ToSlice()
	if len(doubledSlice) != 3 {
		t.Errorf("Expected doubled stack to have 3 items, but got %v", len(doubledSlice))
	}
	if doubledSlice[0] != 2 {
		t.Errorf(errExpectedXItemY, "first", 2, doubledSlice[0])
	}
	if doubledSlice[1] != 4 {
		t.Errorf(errExpectedXItemY, "second", 4, doubledSlice[1])
	}
	if doubledSlice[2] != 6 {
		t.Errorf(errExpectedXItemY, "third", 6, doubledSlice[2])
	}
}

func TestReduce(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test case 1: Sum of all items
	sum, err := s.Reduce(func(a, b int) int {
		return a + b
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if sum != 6 {
		t.Errorf("Expected sum to be 6, but got %v", sum)
	}

	// Test case 2: Product of all items
	product, err := s.Reduce(func(a, b int) int {
		return a * b
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if product != 6 {
		t.Errorf("Expected product to be 6, but got %v", product)
	}

	// Test case 3: Concatenation of all items as strings
	concatenation, err := s.Reduce(func(a, b int) int {
		strA := strconv.Itoa(a)
		strB := strconv.Itoa(b)
		concatenated, _ := strconv.Atoi(strA + strB)
		return concatenated
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if concatenation != 123 {
		t.Errorf("Expected concatenation to be 123, but got %v", concatenation)
	}

	// Test case 4: Error when stack is empty
	emptyStack := stack.NewStack[int]()
	_, err = emptyStack.Reduce(func(a, b int) int {
		return a + b
	})
	if err == nil {
		t.Error(errYesError)
	}
}

func TestForEach(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Define a function to be applied to each item
	fn := func(item *int) {
		// Perform some action on the item
		fmt.Println(*item)
	}

	// Apply the function to each item in the stack
	s.ForEach(fn)
}

func TestAny(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(2)
	s.Push(4)
	s.Push(6)

	// Test with a predicate that returns true for even numbers
	isEven := func(n int) bool {
		return n%2 == 0
	}

	if !s.Any(isEven) {
		t.Error("Expected stack to have at least one even number, but it did not")
	}

	// Test with a predicate that returns true for odd numbers
	isOdd := func(n int) bool {
		return n%2 != 0
	}

	if s.Any(isOdd) {
		t.Error("Expected stack to not have any odd numbers, but it did")
	}

	// Test with an empty stack
	emptyStack := stack.NewStack[int]()
	if emptyStack.Any(isEven) {
		t.Error("Expected empty stack to not have any even numbers, but it did")
	}
}

func TestAll(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(2)
	s.Push(4)
	s.Push(6)

	// Test case 1: All items are even
	isEven := func(n int) bool {
		return n%2 == 0
	}
	if !s.All(isEven) {
		t.Error("Expected all items to be even, but they were not")
	}

	// Test case 2: Not all items are odd
	isOdd := func(n int) bool {
		return n%2 != 0
	}
	if s.All(isOdd) {
		t.Error("Expected not all items to be odd, but they were")
	}

	// Test case 3: Stack is empty
	emptyStack := stack.NewStack[int]()
	if emptyStack.All(isEven) {
		t.Error("Expected an empty stack to match no items, but it did")
	}
}

func TestFind(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test case 1: Item exists in the stack
	item, err := s.Find(func(i int) bool {
		return i == 2
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item == nil {
		t.Error(errItemNotNil)
	} else if *item != 2 {
		t.Errorf(errExpectedItemX, 2, *item)
	}

	// Test case 2: Item does not exist in the stack
	item, err = s.Find(func(i int) bool {
		return i == 4
	})
	if err == nil {
		t.Error(errYesError)
	}
	if item != nil {
		t.Errorf(errExpectedItemX, nil, *item)
	}
}

func TestFindIndex(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	index, err := s.FindIndex(func(item int) bool {
		return item == 2
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if index != 1 {
		t.Errorf("Expected index to be 1, but got %v", index)
	}

	_, err = s.FindIndex(func(item int) bool {
		return item == 4
	})
	if err == nil {
		t.Error(errYesError)
	}
}

func TestFindLast(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test case 1: Item exists in the stack
	item, err := s.FindLast(func(i int) bool {
		return i == 2
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if item == nil {
		t.Error(errItemNotNil)
	} else if *item != 2 {
		t.Errorf(errExpectedItemX, 2, *item)
	}

	// Test case 2: Item does not exist in the stack
	item, err = s.FindLast(func(i int) bool {
		return i == 4
	})
	if err == nil {
		t.Error(errYesError)
	}
	if item != nil {
		t.Errorf(errExpectedItemX, nil, *item)
	}
}

func TestFindLastIndex(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	index, err := s.FindLastIndex(func(item int) bool {
		return item == 2
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	if index != 1 {
		t.Errorf("Expected index to be 1, but got %v", index)
	}

	_, err = s.FindLastIndex(func(item int) bool {
		return item == 4
	})
	if err == nil {
		t.Error(errYesError)
	}
}

func TestFindAll(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	// Test case 1: Find all even numbers
	evenPredicate := func(item int) bool {
		return item%2 == 0
	}
	evenItems := s.FindAll(evenPredicate)
	expectedEvenItems := []int{2, 4}
	if !reflect.DeepEqual(evenItems, expectedEvenItems) {
		t.Errorf("Expected even items to be %v, but got %v", expectedEvenItems, evenItems)
	}

	// Test case 2: Find all odd numbers
	oddPredicate := func(item int) bool {
		return item%2 != 0
	}
	oddItems := s.FindAll(oddPredicate)
	expectedOddItems := []int{1, 3, 5}
	if !reflect.DeepEqual(oddItems, expectedOddItems) {
		t.Errorf("Expected odd items to be %v, but got %v", expectedOddItems, oddItems)
	}

	// Test case 3: Find all numbers greater than 3
	greaterThanThreePredicate := func(item int) bool {
		return item > 3
	}
	greaterThanThreeItems := s.FindAll(greaterThanThreePredicate)
	expectedGreaterThanThreeItems := []int{4, 5}
	if !reflect.DeepEqual(greaterThanThreeItems, expectedGreaterThanThreeItems) {
		t.Errorf("Expected items greater than 3 to be %v, but got %v", expectedGreaterThanThreeItems, greaterThanThreeItems)
	}
}

func TestFindIndices(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(2)
	s.Push(4)
	s.Push(2)

	indices := s.FindIndices(func(item int) bool {
		return item == 2
	})

	expectedIndices := []uint64{1, 3, 5}
	if !reflect.DeepEqual(indices, expectedIndices) {
		t.Errorf("Expected indices to be %v, but got %v", expectedIndices, indices)
	}

	indices = s.FindIndices(func(item int) bool {
		return item == 5
	})

	if indices != nil {
		t.Errorf("Expected indices to be %v, but got %v", expectedIndices, indices)
	}
}

func TestForRange(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test case 1: Apply function to each item within the range [0, 1]
	err := s.ForRange(0, 1, func(item *int) {
		*item *= 2
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	expected := []int{2, 4, 3}
	actual := s.ToSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedStack, expected, actual)
	}

	// Test case 2: Apply function to each item within the range [1, 2]
	err = s.ForRange(1, 2, func(item *int) {
		*item *= 3
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	expected = []int{2, 12, 9}
	actual = s.ToSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedStack, expected, actual)
	}

	// Test case 3: Apply function to each item within the range [2, 2]
	err = s.ForRange(2, 2, func(item *int) {
		*item *= 4
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	expected = []int{2, 12, 36}
	actual = s.ToSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(errExpectedStack, expected, actual)
	}

	// Test case 4: Start index out of range
	err = s.ForRange(3, 4, func(item *int) {
		*item *= 5
	})
	if err == nil {
		t.Error(errYesError)
	}

	// Test case 5: End index out of range
	err = s.ForRange(1, 3, func(item *int) {
		*item *= 6
	})
	if err == nil {
		t.Error(errYesError)
	}

	// Test case 6: Start index is greater than end index
	err = s.ForRange(2, 1, func(item *int) {
		*item *= 7
	})
	if err == nil {
		t.Error(errYesError)
	}
}

func TestForFrom(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test case 1: Apply function to each item starting from index 0
	var result []int
	err := s.ForFrom(0, func(item *int) {
		result = append(result, *item)
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	expectedResult := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf(errExpectedResult, expectedResult, result)
	}

	// Test case 2: Apply function to each item starting from index 1
	result = nil
	err = s.ForFrom(1, func(item *int) {
		result = append(result, *item)
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	expectedResult = []int{2, 3}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf(errExpectedResult, expectedResult, result)
	}

	// Test case 3: Apply function to each item starting from index 2
	result = nil
	err = s.ForFrom(2, func(item *int) {
		result = append(result, *item)
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	expectedResult = []int{3}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf(errExpectedResult, expectedResult, result)
	}

	// Test case 4: Apply function to each item starting from index out of range
	err = s.ForFrom(3, func(item *int) {
		// Function should not be called
		t.Error("Function should not be called")
	})
	if err == nil {
		t.Error(errYesError)
	} else if err.Error() != "start index out of range" {
		t.Errorf("Expected error message to be 'start index out of range', but got '%v'", err.Error())
	}
}

func TestMapFrom(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Stack:
	// 3
	// 2
	// 1

	// Test case 1: MapFrom starting from index 0
	result1, err1 := s.MapFrom(0, func(item int) int {
		return item * 2
	})
	if err1 != nil {
		t.Errorf(errNoError, err1)
	}
	expected1 := stack.NewStack[int]()
	expected1.Push(2)
	expected1.Push(4)
	expected1.Push(6)

	if !result1.Equal(expected1) {
		t.Error(errExpected2Stacks)
	}

	test := s.ToSlice()
	for i, item := range test {
		fmt.Println(i, item)
	}

	// Test case 2: MapFrom starting from index 1
	result2, err2 := s.MapFrom(1, func(item int) int {
		return item * 3
	})
	if err2 != nil {
		t.Errorf(errNoError, err2)
	}
	expected2 := stack.NewStack[int]()
	expected2.Push(3)
	expected2.Push(6)
	if !result2.Equal(expected2) {
		t.Error(errExpected2Stacks)
	}

	// Test case 3: MapFrom starting from index 2
	result3, err3 := s.MapFrom(2, func(item int) int {
		return item * 4
	})
	if err3 != nil {
		t.Errorf(errNoError, err3)
	}
	expected3 := stack.NewStack[int]()
	expected3.Push(4)
	if !result3.Equal(expected3) {
		t.Error(errExpected2Stacks)
	}

	// Test case 4: MapFrom with start index out of range
	_, err4 := s.MapFrom(3, func(item int) int {
		return item * 5
	})
	if err4 == nil {
		t.Error(errYesError)
	}
}

func TestMapRange(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	// Test case 1: MapRange from index 0 to 2
	result, err := s.MapRange(0, 2, func(item int) int {
		return item * 2
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	expected := []int{6, 8, 10}
	if !result.Equal(stack.NewStackFromSlice(expected)) {
		t.Errorf(errExpectedStack, expected, result.ToSlice())
	}

	// Test case 2: MapRange from index 1 to 3
	result, err = s.MapRange(1, 3, func(item int) int {
		return item + 1
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	expected = []int{3, 4, 5}
	if !result.Equal(stack.NewStackFromSlice(expected)) {
		t.Errorf(errExpectedStack, expected, result.ToSlice())
	}

	// Test case 3: MapRange from index 2 to 4
	result, err = s.MapRange(2, 3, func(item int) int {
		return item - 1
	})
	if err != nil {
		t.Errorf(errNoError, err)
	}
	expected = []int{1, 2}
	if !result.Equal(stack.NewStackFromSlice(expected)) {
		t.Errorf(errExpectedStack, expected, result.ToSlice())
	}

	// Test case 5: MapRange with invalid end index
	_, err = s.MapRange(0, 5, func(item int) int {
		return item * 2
	})
	if err == nil {
		t.Error(errYesError)
	}

	// Test case 6: MapRange with start index greater than end index
	_, err = s.MapRange(3, 2, func(item int) int {
		return item * 2
	})
	if err == nil {
		t.Error(errYesError)
	}
}

func TestCheckSize(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.CheckSize()
	if s.Size() != 3 {
		t.Errorf("Expected stack to have 3 items after CheckSize, but got %v", s.Size())
	}

	s.Clear()
	s.CheckSize()
	if s.Size() != 0 {
		t.Errorf("Expected stack to have 0 items after CheckSize, but got %v", s.Size())
	}
}
