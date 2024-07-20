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
	"errors"
	"sync"
)

const (
	errIndexOutOfBound = "index out of bounds"
)

// Node represents a node in the linked list
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// CSLinkList represents a concurrent-safe linked list
type CSLinkList[T comparable] struct {
	mu   sync.Mutex
	Head *Node[T]
}

// CSLinkListNew creates a new CSLinkList
func CSLinkListNew[T comparable]() *CSLinkList[T] {
	return &CSLinkList[T]{}
}

// Append adds a new node to the end of the list
func (l *CSLinkList[T]) Append(value T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	newNode := &Node[T]{Value: value}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

// Prepend adds a new node to the beginning of the list
func (l *CSLinkList[T]) Prepend(value T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	newNode := &Node[T]{Value: value}
	newNode.Next = l.Head
	l.Head = newNode
}

// DeleteWithValue deletes the first node with the given value
func (l *CSLinkList[T]) DeleteWithValue(value T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.Head == nil {
		return
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

// ToSlice returns the list as a slice
func (l *CSLinkList[T]) ToSlice() []T {
	l.mu.Lock()
	defer l.mu.Unlock()

	var result []T
	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

// IsEmpty checks if the list is empty
func (l *CSLinkList[T]) IsEmpty() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.Head == nil
}

// Find returns the first node with the given value
func (l *CSLinkList[T]) Find(value T) (*Node[T], error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.Head
	for current != nil {
		if current.Value == value {
			return current, nil
		}
		current = current.Next
	}
	return nil, errors.New("value not found")
}

// Reverse reverses the list
func (l *CSLinkList[T]) Reverse() {
	l.mu.Lock()
	defer l.mu.Unlock()

	var prev *Node[T]
	current := l.Head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}

// Size returns the number of nodes in the list
func (l *CSLinkList[T]) Size() int {
	l.mu.Lock()
	defer l.mu.Unlock()

	size := 0
	current := l.Head
	for current != nil {
		size++
		current = current.Next
	}
	return size
}

// GetFirst returns the first node in the list
func (l *CSLinkList[T]) GetFirst() *Node[T] {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.Head
}

// GetLast returns the last node in the list
func (l *CSLinkList[T]) GetLast() *Node[T] {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	return current
}

// GetAt returns the node at the given index
func (l *CSLinkList[T]) GetAt(index int) (*Node[T], error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if index < 0 {
		return nil, errors.New(errIndexOutOfBound)
	}

	current := l.Head
	for i := 0; i < index; i++ {
		if current == nil {
			return nil, errors.New(errIndexOutOfBound)
		}
		current = current.Next
	}
	if current == nil {
		return nil, errors.New(errIndexOutOfBound)
	}
	return current, nil
}

// InsertAt inserts a new node at the given index
func (l *CSLinkList[T]) InsertAt(index int, value T) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if index < 0 {
		return errors.New(errIndexOutOfBound)
	}

	if index == 0 {
		l.Prepend(value)
		return nil
	}

	current := l.Head
	for i := 0; i < index-1; i++ {
		if current == nil {
			return errors.New(errIndexOutOfBound)
		}
		current = current.Next
	}

	if current == nil {
		return errors.New(errIndexOutOfBound)
	}

	newNode := &Node[T]{Value: value}
	newNode.Next = current.Next
	current.Next = newNode

	return nil
}

// DeleteAt deletes the node at the given index
func (l *CSLinkList[T]) DeleteAt(index int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if index < 0 {
		return errors.New(errIndexOutOfBound)
	}

	if index == 0 {
		if l.Head == nil {
			return errors.New(errIndexOutOfBound)
		}
		l.Head = l.Head.Next
		return nil
	}

	current := l.Head
	for i := 0; i < index-1; i++ {
		if current == nil {
			return errors.New(errIndexOutOfBound)
		}
		current = current.Next
	}

	if current == nil || current.Next == nil {
		return errors.New(errIndexOutOfBound)
	}

	current.Next = current.Next.Next
	return nil
}

// Remove is an alias for DeleteWithValue
func (l *CSLinkList[T]) Remove(value T) {
	l.DeleteWithValue(value)
}

// Clear removes all nodes from the list
func (l *CSLinkList[T]) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.Head = nil
}

// Copy returns a copy of the list
func (l *CSLinkList[T]) Copy() *CSLinkList[T] {
	l.mu.Lock()
	defer l.mu.Unlock()

	newList := CSLinkListNew[T]()
	current := l.Head
	for current != nil {
		newList.Append(current.Value)
		current = current.Next
	}
	return newList
}

// Merge appends all the nodes from another list to the current list
func (l *CSLinkList[T]) Merge(list *CSLinkList[T]) {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := list.Head
	for current != nil {
		l.Append(current.Value)
		current = current.Next
	}
}

// Map applies the function to all the nodes in the list
func (l *CSLinkList[T]) Map(f func(T) T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.Head
	for current != nil {
		current.Value = f(current.Value)
		current = current.Next
	}
}

// Filter removes nodes from the list that don't match the predicate
func (l *CSLinkList[T]) Filter(f func(T) bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for l.Head != nil && !f(l.Head.Value) {
		l.Head = l.Head.Next
	}

	current := l.Head
	for current != nil && current.Next != nil {
		if !f(current.Next.Value) {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
}

// Reduce reduces the list to a single value
func (l *CSLinkList[T]) Reduce(f func(T, T) T, initial T) T {
	l.mu.Lock()
	defer l.mu.Unlock()

	result := initial
	current := l.Head
	for current != nil {
		result = f(result, current.Value)
		current = current.Next
	}
	return result
}

// ForEach applies the function to all the nodes in the list
func (l *CSLinkList[T]) ForEach(f func(T)) {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.Head
	for current != nil {
		f(current.Value)
		current = current.Next
	}
}

// Any checks if any node in the list matches the predicate
func (l *CSLinkList[T]) Any(f func(T) bool) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.Head
	for current != nil {
		if f(current.Value) {
			return true
		}
		current = current.Next
	}
	return false
}

// All checks if all nodes in the list match the predicate
func (l *CSLinkList[T]) All(f func(T) bool) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.Head
	for current != nil {
		if !f(current.Value) {
			return false
		}
		current = current.Next
	}
	return true
}

// Contains checks if the list contains the given value
func (l *CSLinkList[T]) Contains(value T) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

// IndexOf returns the index of the first node with the given value
func (l *CSLinkList[T]) IndexOf(value T) int {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.Head
	index := 0
	for current != nil {
		if current.Value == value {
			return index
		}
		current = current.Next
		index++
	}
	return -1
}

// LastIndexOf returns the index of the last node with the given value
func (l *CSLinkList[T]) LastIndexOf(value T) int {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.Head
	index := -1
	i := 0
	for current != nil {
		if current.Value == value {
			index = i
		}
		current = current.Next
		i++
	}
	return index
}

// FindIndex returns the index of the first node that matches the predicate
func (l *CSLinkList[T]) FindIndex(f func(T) bool) int {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.Head
	index := 0
	for current != nil {
		if f(current.Value) {
			return index
		}
		current = current.Next
		index++
	}
	return -1
}

// FindLastIndex returns the index of the last node that matches the predicate
func (l *CSLinkList[T]) FindLastIndex(f func(T) bool) int {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.Head
	index := -1
	i := 0
	for current != nil {
		if f(current.Value) {
			index = i
		}
		current = current.Next
		i++
	}
	return index
}

// FindAll returns all nodes that match the predicate
func (l *CSLinkList[T]) FindAll(f func(T) bool) *CSLinkList[T] {
	l.mu.Lock()
	defer l.mu.Unlock()

	newList := CSLinkListNew[T]()
	current := l.Head
	for current != nil {
		if f(current.Value) {
			newList.Append(current.Value)
		}
		current = current.Next
	}
	return newList
}

// FindLast returns the last node that matches the predicate
func (l *CSLinkList[T]) FindLast(f func(T) bool) (*Node[T], error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var result *Node[T]
	current := l.Head
	for current != nil {
		if f(current.Value) {
			result = current
		}
		current = current.Next
	}
	if result == nil {
		return nil, errors.New("value not found")
	}
	return result, nil
}

// FindAllIndexes returns the indexes of all nodes that match the predicate
func (l *CSLinkList[T]) FindAllIndexes(f func(T) bool) []int {
	l.mu.Lock()
	defer l.mu.Unlock()

	var result []int
	current := l.Head
	index := 0
	for current != nil {
		if f(current.Value) {
			result = append(result, index)
		}
		current = current.Next
		index++
	}
	return result
}
