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
	"errors"
	"fmt"
)

// Error messages
const (
	errItemNotFound  = "item not found"
	errStackIsEmpty  = "stack is empty"
	errStartIndexOOR = "start index out of range"
	errEndIndexOOR   = "end index out of range"
)

// Stack is a non-concurrent-safe stack.
type Stack[T comparable] struct {
	items []T
	size  uint64
}

// New creates a new Stack.
func New[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

// NewFromSlice creates a new Stack from a slice.
func NewFromSlice[T comparable](items []T) *Stack[T] {
	stack := New[T]()
	stack.PushAll(items)
	return stack
}

// Push adds an item to the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
	s.size++
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}

// Pop removes and returns the top item from the stack.
func (s *Stack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New(errStackIsEmpty)
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.size--
	return &item, nil
}

// ToSlice returns the stack as a slice.
func (s *Stack[T]) ToSlice() []T {
	items := make([]T, s.size)
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
func (s *Stack[T]) Size() uint64 {
	return s.size
}

// CheckSize recalculate the size of the stack.
func (s *Stack[T]) CheckSize() {
	s.size = uint64(len(s.items))
}

// Clear removes all items from the stack.
func (s *Stack[T]) Clear() {
	s.items = s.items[:0]
	s.size = 0
}

// Contains checks if the stack contains an item.
func (s *Stack[T]) Contains(item T) bool {
	if s.size == 0 {
		return false
	}

	if s.items[0] == item {
		return true
	}
	fmt.Printf("s.size: %d\n", s.size)
	for i := s.size - 1; i > 0; i-- {
		if s.items[i] == item {
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

	if s.size != other.size {
		return false
	}

	if s.size == 0 && other.size == 0 {
		return true
	}
	if s.items[0] != other.items[0] {
		return false
	}

	for i := s.size - 1; i > 0; i-- {
		if s.items[i] != other.items[i] {
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
func (s *Stack[T]) PopN(n uint64) ([]T, error) {
	if s.size < n {
		return nil, errors.New("Stack has less than n items")
	}

	items := make([]T, n)
	for i := uint64(0); i < n; i++ {
		item, err := s.Pop()
		if err != nil {
			return nil, err
		}
		items[i] = *item
	}
	return items, nil
}

// PushN adds multiple items to the stack.
func (s *Stack[T]) PushN(items ...T) {
	s.items = append(s.items, items...)
	s.size += uint64(len(items))
}

// PopAll removes and returns all items from the stack.
func (s *Stack[T]) PopAll() []T {
	items := make([]T, len(s.items))
	for i := len(s.items) - 1; i >= 0; i-- {
		items[len(s.items)-i-1] = s.items[i]
	}
	s.items = s.items[:0]
	s.size = 0
	return items
}

// PushAll adds multiple items to the stack.
func (s *Stack[T]) PushAll(items []T) {
	s.items = append(s.items, items...)
	s.size += uint64(len(items))
}

// Filter removes items from the stack that don't match the predicate.
func (s *Stack[T]) Filter(predicate func(T) bool) {
	var items []T
	var size uint64
	for _, item := range s.items {
		if predicate(item) {
			items = append(items, item)
			size++
		}
	}
	s.items = items
	s.size = size
}

// Map creates a new stack with the results of applying the function to each item.
func (s *Stack[T]) Map(fn func(T) T) *Stack[T] {
	stack := New[T]()
	for i := uint64(0); i < s.size; i++ {
		stack.Push(fn(s.items[i]))
	}
	return stack
}

// MapFrom creates a new stack with the results of applying the function to each item starting from the specified index.
// Please note: the start index is the top of the stack.
func (s *Stack[T]) MapFrom(start uint64, fn func(T) T) (*Stack[T], error) {
	if s.size == 0 {
		return nil, errors.New(errStackIsEmpty)
	}

	if start >= s.size {
		return nil, errors.New(errStartIndexOOR)
	}

	// calculate stack start index
	stackStart := s.size - start - 1

	stack := New[T]()
	stack.items = make([]T, s.size-(start))
	for i := stackStart; i > 0; i-- {
		stack.items[i] = fn(s.items[i])
		stack.size++
	}
	stack.items[0] = fn(s.items[0])
	stack.size++
	return stack, nil
}

// MapRange creates a new stack with the results of applying the function to each item within the specified range.
// Please note: start and end are inclusive and on a stack this means that the start index is the top of the stack.
func (s *Stack[T]) MapRange(start, end uint64, fn func(T) T) (*Stack[T], error) {
	if start >= s.size {
		return nil, errors.New(errStartIndexOOR)
	}

	if end >= s.size {
		return nil, errors.New(errEndIndexOOR)
	}

	if start > end {
		return nil, errors.New("start index is greater than end index")
	}

	// Convert the start and end index to the stack indexes
	stackStart := (s.size - start) - 1
	stackEnd := (s.size - end) - 1

	stack := New[T]()
	for i := stackEnd; i <= stackStart; i++ {
		stack.Push(fn(s.items[i]))
	}
	return stack, nil
}

// Reduce reduces the stack to a single value.
func (s *Stack[T]) Reduce(fn func(T, T) T) (T, error) {
	if s.size == 0 {
		var rVal T
		return rVal, errors.New(errStackIsEmpty)
	}

	result := s.items[0]
	for i := uint64(1); i < s.size; i++ {
		result = fn(result, s.items[i])
	}
	return result, nil
}

// ForEach applies the function to each item in the stack.
func (s *Stack[T]) ForEach(fn func(*T)) {
	for i := uint64(0); i < s.size; i++ {
		fn(&s.items[i])
	}
}

// ForRange applies the function to each item in the stack within the specified range.
func (s *Stack[T]) ForRange(start, end uint64, fn func(*T)) error {
	if start >= s.size {
		return errors.New(errStartIndexOOR)
	}

	if end >= s.size {
		return errors.New(errEndIndexOOR)
	}

	if start > end {
		return errors.New("start index is greater than end index")
	}

	for i := start; i <= end; i++ {
		fn(&s.items[i])
	}
	return nil
}

// ForFrom applies the function to each item in the stack starting from the specified index.
func (s *Stack[T]) ForFrom(start uint64, fn func(*T)) error {
	if start >= s.size {
		return errors.New(errStartIndexOOR)
	}

	for i := start; i < s.size; i++ {
		fn(&s.items[i])
	}
	return nil
}

// Any checks if any item in the stack matches the predicate.
func (s *Stack[T]) Any(predicate func(T) bool) bool {
	if s == nil {
		return false
	}
	if s.size == 0 {
		return false
	}

	for i := uint64(0); i < s.size; i++ {
		if predicate(s.items[i]) {
			return true
		}
	}
	return false
}

// All checks if all items in the stack match the predicate.
func (s *Stack[T]) All(predicate func(T) bool) bool {
	if s == nil {
		return false
	}
	if s.size == 0 {
		return false
	}

	for i := uint64(0); i < s.size; i++ {
		if !predicate(s.items[i]) {
			return false
		}
	}
	return true
}

// Find returns the first item that matches the predicate.
func (s *Stack[T]) Find(predicate func(T) bool) (*T, error) {
	if s == nil {
		return nil, errors.New(errItemNotFound)
	}
	if len(s.items) == 0 {
		return nil, errors.New(errItemNotFound)
	}

	for i := uint64(0); i < s.size; i++ {
		if predicate(s.items[i]) {
			return &s.items[i], nil
		}
	}
	return nil, errors.New(errItemNotFound)
}

// FindIndex returns the index of the first item that matches the predicate.
func (s *Stack[T]) FindIndex(predicate func(T) bool) (uint64, error) {
	for i := uint64(0); i < s.size; i++ {
		if predicate(s.items[i]) {
			return i, nil
		}
	}
	return 0, errors.New(errItemNotFound)
}

// FindLast returns the last item that matches the predicate.
func (s *Stack[T]) FindLast(predicate func(T) bool) (*T, error) {
	if s.size == 0 {
		return nil, errors.New(errItemNotFound)
	}

	for i := s.size - 1; i > 0; i-- {
		if predicate(s.items[i]) {
			return &s.items[i], nil
		}
	}
	if predicate(s.items[0]) {
		return &s.items[0], nil
	}

	return nil, errors.New(errItemNotFound)
}

// FindLastIndex returns the index of the last item that matches the predicate.
func (s *Stack[T]) FindLastIndex(predicate func(T) bool) (uint64, error) {
	if s.size == 0 {
		return 0, errors.New(errItemNotFound)
	}

	for i := s.size - 1; i > 0; i-- {
		if predicate(s.items[i]) {
			return i, nil
		}
	}
	if predicate(s.items[0]) {
		return 0, nil
	}

	return 0, errors.New(errItemNotFound)
}

// FindAll returns all items that match the predicate.
func (s *Stack[T]) FindAll(predicate func(T) bool) []T {
	var items []T
	for i := uint64(0); i < s.size; i++ {
		if predicate(s.items[i]) {
			items = append(items, s.items[i])
		}
	}
	return items
}

// FindIndices returns the indices of all items that match the predicate.
func (s *Stack[T]) FindIndices(predicate func(T) bool) []uint64 {
	var indices []uint64
	for i := uint64(0); i < s.size; i++ {
		if predicate(s.items[i]) {
			indices = append(indices, i)
		}
	}
	return indices
}
