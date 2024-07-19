// Copyright 2023 Paolo Fabio Zaino
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
	"errors"
	"fmt"
)

// Error messages
const (
	errItemNotFound = "item not found"
	errStackIsEmpty = "stack is empty"
)

// Stack is a non-concurrent-safe stack.
type Stack[T comparable] struct {
	items []T
}

// New creates a new Stack.
func New[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an item to the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Pop removes and returns the top item from the stack.
func (s *Stack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New(errStackIsEmpty)
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return &item, nil
}

// ToSlice returns the stack as a slice.
func (s *Stack[T]) ToSlice() []T {
	items := make([]T, len(s.items))
	copy(items, s.items)
	return items
}

// Reverse reverses the stack.
func (s *Stack[T]) Reverse() {
	for i := 0; i < len(s.items)/2; i++ {
		j := len(s.items) - i - 1
		s.items[i], s.items[j] = s.items[j], s.items[i]
	}
}

// Swap swaps the top two items on the stack.
func (s *Stack[T]) Swap() error {
	if len(s.items) < 2 {
		return errors.New("Stack has less than 2 items")
	}

	s.items[len(s.items)-1], s.items[len(s.items)-2] = s.items[len(s.items)-2], s.items[len(s.items)-1]
	return nil
}

// Top returns the top item from the stack without removing it.
func (s *Stack[T]) Top() (*T, error) {
	if len(s.items) == 0 {
		return nil, errors.New(errStackIsEmpty)
	}

	item := s.items[len(s.items)-1]
	return &item, nil
}

// Peek is a wrapper around Top (for who's more used to use Peek).
func (s *Stack[T]) Peek() (*T, error) {
	return s.Top()
}

// Size returns the number of items in the stack.
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Clear removes all items from the stack.
func (s *Stack[T]) Clear() {
	s.items = s.items[:0]
}

// Contains checks if the stack contains an item.
func (s *Stack[T]) Contains(item T) bool {
	for _, v := range s.items {
		if v == item {
			return true
		}
	}
	return false
}

// Copy returns a new Stack with the same items.
func (s *Stack[T]) Copy() *Stack[T] {
	stack := New[T]()
	for _, item := range s.items {
		stack.Push(item)
	}
	return stack
}

// Equal checks if two stacks are equal.
func (s *Stack[T]) Equal(other *Stack[T]) bool {
	if s == nil && other == nil {
		return true
	}

	if (s != nil && other == nil) || (s == nil && other != nil) {
		return false
	}

	if s.Size() != other.Size() {
		return false
	}

	for i, v := range s.items {
		if v != other.items[i] {
			return false
		}
	}
	return true
}

// String returns a string representation of the stack.
func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", s.items)
}

// PopN removes and returns the top n items from the stack.
func (s *Stack[T]) PopN(n int) ([]T, error) {
	if len(s.items) < n {
		return nil, errors.New("Stack has less than n items")
	}

	items := s.items[len(s.items)-n:]
	s.items = s.items[:len(s.items)-n]
	return items, nil
}

// PushN adds multiple items to the stack.
func (s *Stack[T]) PushN(items ...T) {
	s.items = append(s.items, items...)
}

// PopAll removes and returns all items from the stack.
func (s *Stack[T]) PopAll() []T {
	items := make([]T, len(s.items))
	copy(items, s.items)
	s.items = s.items[:0]
	return items
}

// PushAll adds multiple items to the stack.
func (s *Stack[T]) PushAll(items []T) {
	s.items = append(s.items, items...)
}

// Filter removes items from the stack that don't match the predicate.
func (s *Stack[T]) Filter(predicate func(T) bool) {
	var items []T
	for _, item := range s.items {
		if predicate(item) {
			items = append(items, item)
		}
	}
	s.items = items
}

// Map creates a new stack with the results of applying the function to each item.
func (s *Stack[T]) Map(fn func(T) T) *Stack[T] {
	stack := New[T]()
	for _, item := range s.items {
		stack.Push(fn(item))
	}
	return stack
}

// Reduce reduces the stack to a single value.
func (s *Stack[T]) Reduce(fn func(T, T) T) (T, error) {
	if len(s.items) == 0 {
		var rVal T
		return rVal, errors.New(errStackIsEmpty)
	}

	result := s.items[0]
	for i := 1; i < len(s.items); i++ {
		result = fn(result, s.items[i])
	}
	return result, nil
}

// ForEach applies the function to each item in the stack.
func (s *Stack[T]) ForEach(fn func(T)) {
	for _, item := range s.items {
		fn(item)
	}
}

// Any checks if any item in the stack matches the predicate.
func (s *Stack[T]) Any(predicate func(T) bool) bool {
	for _, item := range s.items {
		if predicate(item) {
			return true
		}
	}
	return false
}

// All checks if all items in the stack match the predicate.
func (s *Stack[T]) All(predicate func(T) bool) bool {
	for _, item := range s.items {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Find returns the first item that matches the predicate.
func (s *Stack[T]) Find(predicate func(T) bool) (*T, error) {
	for _, item := range s.items {
		if predicate(item) {
			return &item, nil
		}
	}
	return nil, errors.New(errItemNotFound)
}

// FindIndex returns the index of the first item that matches the predicate.
func (s *Stack[T]) FindIndex(predicate func(T) bool) (int, error) {
	for i, item := range s.items {
		if predicate(item) {
			return i, nil
		}
	}
	return -1, errors.New(errItemNotFound)
}

// FindLast returns the last item that matches the predicate.
func (s *Stack[T]) FindLast(predicate func(T) bool) (*T, error) {
	for i := len(s.items) - 1; i >= 0; i-- {
		if predicate(s.items[i]) {
			return &s.items[i], nil
		}
	}
	return nil, errors.New(errItemNotFound)
}

// FindLastIndex returns the index of the last item that matches the predicate.
func (s *Stack[T]) FindLastIndex(predicate func(T) bool) (int, error) {
	for i := len(s.items) - 1; i >= 0; i-- {
		if predicate(s.items[i]) {
			return i, nil
		}
	}
	return -1, errors.New(errItemNotFound)
}

// FindAll returns all items that match the predicate.
func (s *Stack[T]) FindAll(predicate func(T) bool) []T {
	var items []T
	for _, item := range s.items {
		if predicate(item) {
			items = append(items, item)
		}
	}
	return items
}

// FindIndices returns the indices of all items that match the predicate.
func (s *Stack[T]) FindIndices(predicate func(T) bool) []int {
	var indices []int
	for i, item := range s.items {
		if predicate(item) {
			indices = append(indices, i)
		}
	}
	return indices
}
