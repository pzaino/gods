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
	"fmt"
	"reflect"
	"strconv"
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
func TestString(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	expected := "[1 2 3]"
	result := stack.String()
	if result != expected {
		t.Errorf("Expected string representation to be %q, but got %q", expected, result)
	}
}

func TestStringEmpty(t *testing.T) {
	stack := New[int]()
	expected := "[]"
	result := stack.String()
	if result != expected {
		t.Errorf("Expected string representation to be %q, but got %q", expected, result)
	}
}

func TestPopN(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Test case 1: Pop 2 items
	items, err := stack.PopN(2)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if len(items) != 2 {
		t.Errorf("Expected 2 items, but got %v", len(items))
	}
	if items[0] != 2 {
		t.Errorf("Expected first item to be 2, but got %v", items[0])
	}
	if items[1] != 3 {
		t.Errorf("Expected second item to be 3, but got %v", items[1])
	}
	if stack.Size() != 1 {
		t.Errorf("Expected stack size to be 1, but got %v", stack.Size())
	}

	// Test case 2: Pop 1 item
	items, err = stack.PopN(1)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if len(items) != 1 {
		t.Errorf("Expected 1 item, but got %v", len(items))
	}
	if items[0] != 1 {
		t.Errorf("Expected item to be 1, but got %v", items[0])
	}
	if stack.Size() != 0 {
		t.Errorf("Expected stack size to be 0, but got %v", stack.Size())
	}

	// Test case 3: Pop from empty stack
	items, err = stack.PopN(1)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if items != nil {
		t.Errorf("Expected items to be nil, but got %v", items)
	}
}

func TestPushN(t *testing.T) {
	stack := New[int]()
	stack.PushN(1, 2, 3)
	slice := stack.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected stack to have 3 items, but got %v", len(slice))
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

func TestPopAll(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	items := stack.PopAll()
	if len(items) != 3 {
		t.Errorf("Expected %d items, but got %d", 3, len(items))
	}
	if stack.Size() != 0 {
		t.Errorf("Expected stack to be empty, but it has %d items", stack.Size())
	}
	if items[0] != 3 {
		t.Errorf("Expected first item to be 3, but got %v", items[0])
	}
	if items[1] != 2 {
		t.Errorf("Expected second item to be 2, but got %v", items[1])
	}
	if items[2] != 1 {
		t.Errorf("Expected third item to be 1, but got %v", items[2])
	}
}

func TestPushAll(t *testing.T) {
	stack := New[int]()
	items := []int{1, 2, 3}
	stack.PushAll(items)

	if stack.Size() != len(items) {
		t.Errorf("Expected stack size to be %d, but got %d", len(items), stack.Size())
	}

	slice := stack.ToSlice()
	for i, item := range items {
		if slice[i] != item {
			t.Errorf("Expected item at index %d to be %d, but got %d", i, item, slice[i])
		}
	}
}

func TestFilter(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	// Test filtering even numbers
	stack.Filter(func(item int) bool {
		return item%2 == 0
	})

	// Check if the stack contains only even numbers
	if !stack.Contains(2) {
		t.Error("Expected stack to contain 2, but it did not")
	}
	if !stack.Contains(4) {
		t.Error("Expected stack to contain 4, but it did not")
	}
	if stack.Contains(1) {
		t.Error("Expected stack to not contain 1, but it did")
	}
	if stack.Contains(3) {
		t.Error("Expected stack to not contain 3, but it did")
	}
}

func TestMap(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Test mapping to double the values
	doubledStack := stack.Map(func(item int) int {
		return item * 2
	})

	// Check if the original stack is unchanged
	if stack.Size() != 3 {
		t.Errorf("Expected original stack to have 3 items, but got %v", stack.Size())
	}

	// Check if the new stack has the mapped values
	doubledSlice := doubledStack.ToSlice()
	if len(doubledSlice) != 3 {
		t.Errorf("Expected doubled stack to have 3 items, but got %v", len(doubledSlice))
	}
	if doubledSlice[0] != 2 {
		t.Errorf("Expected first item to be 2, but got %v", doubledSlice[0])
	}
	if doubledSlice[1] != 4 {
		t.Errorf("Expected second item to be 4, but got %v", doubledSlice[1])
	}
	if doubledSlice[2] != 6 {
		t.Errorf("Expected third item to be 6, but got %v", doubledSlice[2])
	}
}

func TestReduce(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Test case 1: Sum of all items
	sum, err := stack.Reduce(func(a, b int) int {
		return a + b
	})
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if sum != 6 {
		t.Errorf("Expected sum to be 6, but got %v", sum)
	}

	// Test case 2: Product of all items
	product, err := stack.Reduce(func(a, b int) int {
		return a * b
	})
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if product != 6 {
		t.Errorf("Expected product to be 6, but got %v", product)
	}

	// Test case 3: Concatenation of all items as strings
	concatenation, err := stack.Reduce(func(a, b int) int {
		strA := strconv.Itoa(a)
		strB := strconv.Itoa(b)
		concatenated, _ := strconv.Atoi(strA + strB)
		return concatenated
	})
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if concatenation != 123 {
		t.Errorf("Expected concatenation to be 123, but got %v", concatenation)
	}

	// Test case 4: Error when stack is empty
	emptyStack := New[int]()
	_, err = emptyStack.Reduce(func(a, b int) int {
		return a + b
	})
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}

func TestForEach(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Define a function to be applied to each item
	fn := func(item int) {
		// Perform some action on the item
		fmt.Println(item)
	}

	// Apply the function to each item in the stack
	stack.ForEach(fn)

	// Add assertions here if needed
}

func TestAny(t *testing.T) {
	stack := New[int]()
	stack.Push(2)
	stack.Push(4)
	stack.Push(6)

	// Test with a predicate that returns true for even numbers
	isEven := func(n int) bool {
		return n%2 == 0
	}

	if !stack.Any(isEven) {
		t.Error("Expected stack to have at least one even number, but it did not")
	}

	// Test with a predicate that returns true for odd numbers
	isOdd := func(n int) bool {
		return n%2 != 0
	}

	if stack.Any(isOdd) {
		t.Error("Expected stack to not have any odd numbers, but it did")
	}

	// Test with an empty stack
	emptyStack := New[int]()
	if emptyStack.Any(isEven) {
		t.Error("Expected empty stack to not have any even numbers, but it did")
	}
}

func TestAll(t *testing.T) {
	stack := New[int]()
	stack.Push(2)
	stack.Push(4)
	stack.Push(6)

	// Test case 1: All items are even
	isEven := func(n int) bool {
		return n%2 == 0
	}
	if !stack.All(isEven) {
		t.Error("Expected all items to be even, but they were not")
	}

	// Test case 2: Not all items are odd
	isOdd := func(n int) bool {
		return n%2 != 0
	}
	if stack.All(isOdd) {
		t.Error("Expected not all items to be odd, but they were")
	}

	// Test case 3: Stack is empty
	emptyStack := New[int]()
	if !emptyStack.All(isEven) {
		t.Error("Expected all items to be even in an empty stack, but they were not")
	}
}

func TestFind(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Test case 1: Item exists in the stack
	item, err := stack.Find(func(i int) bool {
		return i == 2
	})
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if item == nil {
		t.Error("Expected item to not be nil, but it was")
	} else if *item != 2 {
		t.Errorf("Expected item to be 2, but got %v", *item)
	}

	// Test case 2: Item does not exist in the stack
	item, err = stack.Find(func(i int) bool {
		return i == 4
	})
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if item != nil {
		t.Errorf("Expected item to be nil, but got %v", *item)
	}
}

func TestFindIndex(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	index, err := stack.FindIndex(func(item int) bool {
		return item == 2
	})
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if index != 1 {
		t.Errorf("Expected index to be 1, but got %v", index)
	}

	index, err = stack.FindIndex(func(item int) bool {
		return item == 4
	})
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if index != -1 {
		t.Errorf("Expected index to be -1, but got %v", index)
	}
}

func TestFindLast(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Test case 1: Item exists in the stack
	item, err := stack.FindLast(func(i int) bool {
		return i == 2
	})
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if item == nil {
		t.Error("Expected item to not be nil, but it was")
	} else if *item != 2 {
		t.Errorf("Expected item to be 2, but got %v", *item)
	}

	// Test case 2: Item does not exist in the stack
	item, err = stack.FindLast(func(i int) bool {
		return i == 4
	})
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if item != nil {
		t.Errorf("Expected item to be nil, but got %v", *item)
	}
}

func TestFindLastIndex(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	index, err := stack.FindLastIndex(func(item int) bool {
		return item == 2
	})
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if index != 1 {
		t.Errorf("Expected index to be 1, but got %v", index)
	}

	index, err = stack.FindLastIndex(func(item int) bool {
		return item == 4
	})
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if index != -1 {
		t.Errorf("Expected index to be -1, but got %v", index)
	}
}

func TestFindAll(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	// Test case 1: Find all even numbers
	evenPredicate := func(item int) bool {
		return item%2 == 0
	}
	evenItems := stack.FindAll(evenPredicate)
	expectedEvenItems := []int{2, 4}
	if !reflect.DeepEqual(evenItems, expectedEvenItems) {
		t.Errorf("Expected even items to be %v, but got %v", expectedEvenItems, evenItems)
	}

	// Test case 2: Find all odd numbers
	oddPredicate := func(item int) bool {
		return item%2 != 0
	}
	oddItems := stack.FindAll(oddPredicate)
	expectedOddItems := []int{1, 3, 5}
	if !reflect.DeepEqual(oddItems, expectedOddItems) {
		t.Errorf("Expected odd items to be %v, but got %v", expectedOddItems, oddItems)
	}

	// Test case 3: Find all numbers greater than 3
	greaterThanThreePredicate := func(item int) bool {
		return item > 3
	}
	greaterThanThreeItems := stack.FindAll(greaterThanThreePredicate)
	expectedGreaterThanThreeItems := []int{4, 5}
	if !reflect.DeepEqual(greaterThanThreeItems, expectedGreaterThanThreeItems) {
		t.Errorf("Expected items greater than 3 to be %v, but got %v", expectedGreaterThanThreeItems, greaterThanThreeItems)
	}
}

func TestFindIndices(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(2)
	stack.Push(4)
	stack.Push(2)

	indices := stack.FindIndices(func(item int) bool {
		return item == 2
	})

	expectedIndices := []int{1, 3, 5}
	if !reflect.DeepEqual(indices, expectedIndices) {
		t.Errorf("Expected indices to be %v, but got %v", expectedIndices, indices)
	}

	indices = stack.FindIndices(func(item int) bool {
		return item == 5
	})

	if indices != nil {
		t.Errorf("Expected indices to be %v, but got %v", expectedIndices, indices)
	}
}
