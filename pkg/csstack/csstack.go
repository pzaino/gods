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

// Package csstack provides a concurrency-safe stack (LIFO) using stack package.
package csstack

import (
	"errors"
	"sync"

	stack "github.com/pzaino/gods/pkg/stack"
)

// CSStack is a concurrency-safe stack.
type CSStack[T comparable] struct {
	mu sync.RWMutex
	s  *stack.Stack[T]
}

// New creates a new concurrency-safe stack.
func New[T comparable]() *CSStack[T] {
	return &CSStack[T]{s: stack.New[T]()}
}

// NewFromSlice creates a new concurrency-safe stack from a slice.
func NewFromSlice[T comparable](items []T) *CSStack[T] {
	cs := New[T]()
	cs.s.PushAll(items)
	return cs
}

// Push adds an item to the stack.
func (cs *CSStack[T]) Push(item T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.s.Push(item)
}

// IsEmpty checks if the stack is empty.
func (cs *CSStack[T]) IsEmpty() bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.s.IsEmpty()
}

// Pop removes and returns the top item from the stack.
func (cs *CSStack[T]) Pop() (*T, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.s.Pop()
}

// ToSlice returns the stack as a slice.
func (cs *CSStack[T]) ToSlice() []T {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s.ToSlice()
}

// ToStack returns the stack as a stack (non-concurrent-safe).
func (cs *CSStack[T]) ToStack() *stack.Stack[T] {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s
}

// Reverse reverses the stack.
func (cs *CSStack[T]) Reverse() {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.s.Reverse()
}

// Swap swaps the top two items on the stack.
func (cs *CSStack[T]) Swap() error {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.s.Swap()
}

// Top returns the top item from the stack without removing it.
func (cs *CSStack[T]) Top() (*T, error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s.Top()
}

// Peek is a wrapper around Top (for those more used to using Peek).
func (cs *CSStack[T]) Peek() (*T, error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s.Peek()
}

// Size returns the number of items in the stack.
func (cs *CSStack[T]) Size() uint64 {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s.Size()
}

// Clear removes all items from the stack.
func (cs *CSStack[T]) Clear() {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.s.Clear()
}

// Contains checks if the stack contains an item.
func (cs *CSStack[T]) Contains(item T) bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s.Contains(item)
}

// Copy returns a new CSStack with the same items.
func (cs *CSStack[T]) Copy() *CSStack[T] {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return &CSStack[T]{s: cs.s.Copy()}
}

// Equal checks if two stacks are equal.
func (cs *CSStack[T]) Equal(other *CSStack[T]) bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	other.mu.Lock()
	defer other.mu.Unlock()
	return cs.s.Equal(other.s)
}

// String returns a string representation of the stack.
func (cs *CSStack[T]) String() string {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s.String()
}

func (cs *CSStack[T]) PopN(n uint64) ([]T, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	if cs.s.Size() < n {
		return nil, errors.New("Stack has less than n items")
	}
	return cs.s.PopN(n)
}

// PushN adds multiple items to the stack.
func (cs *CSStack[T]) PushN(items ...T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.s.PushN(items...)
}

// PopAll removes and returns all items from the stack.
func (cs *CSStack[T]) PopAll() []T {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.s.PopAll()
}

// PushAll adds multiple items to the stack.
func (cs *CSStack[T]) PushAll(items []T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.s.PushAll(items)
}

// Filter removes items from the stack that don't match the predicate.
func (cs *CSStack[T]) Filter(predicate func(T) bool) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.s.Filter(predicate)
}

// Map creates a new stack with the results of applying the function to each item.
func (cs *CSStack[T]) Map(fn func(T) T) (*CSStack[T], error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	csStack := &CSStack[T]{}
	var err error
	csStack.s, err = cs.s.Map(fn)
	return csStack, err
}

// Reduce reduces the stack to a single value.
func (cs *CSStack[T]) Reduce(fn func(T, T) T) (T, error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s.Reduce(fn)
}

// ForEach applies the function to each item in the stack.
func (cs *CSStack[T]) ForEach(fn func(*T) error) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.s.ForEach(fn)
}

// ForRange applies the function to each item in the stack in the range [start, end).
func (cs *CSStack[T]) ForRange(start, end uint64, fn func(*T) error) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.s.ForRange(start, end, fn)
}

// ForFrom applies the function to each item in the stack starting from the index.
func (cs *CSStack[T]) ForFrom(start uint64, fn func(*T) error) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.s.ForFrom(start, fn)
}

// Any checks if any item in the stack matches the predicate.
func (cs *CSStack[T]) Any(predicate func(T) bool) bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s.Any(predicate)
}

// All checks if all items in the stack match the predicate.
func (cs *CSStack[T]) All(predicate func(T) bool) bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s.All(predicate)
}

// Find returns the first item that matches the predicate.
func (cs *CSStack[T]) Find(predicate func(T) bool) (*T, error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s.Find(predicate)
}

// FindIndex returns the index of the first item that matches the predicate.
func (cs *CSStack[T]) FindIndex(predicate func(T) bool) (uint64, error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s.FindIndex(predicate)
}

// FindLast returns the last item that matches the predicate.
func (cs *CSStack[T]) FindLast(predicate func(T) bool) (*T, error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s.FindLast(predicate)
}

// FindLastIndex returns the index of the last item that matches the predicate.
func (cs *CSStack[T]) FindLastIndex(predicate func(T) bool) (uint64, error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s.FindLastIndex(predicate)
}

// FindAll returns all items that match the predicate.
func (cs *CSStack[T]) FindAll(predicate func(T) bool) []T {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s.FindAll(predicate)
}

// FindIndices returns the indices of all items that match the predicate.
func (cs *CSStack[T]) FindIndices(predicate func(T) bool) []uint64 {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.s.FindIndices(predicate)
}
