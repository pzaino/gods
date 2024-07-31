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

// Package circularLinkList provides a non-concurrent-safe circular linked list.
package circularLinkList

import (
	"errors"
)

const (
	errIndexOutOfBound = "index out of bounds"
	errListIsEmpty     = "list is empty"
)

// Node represents a node in the circular linked list
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// CircularLinkList represents a circular linked list
type CircularLinkList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
	size uint64
}

// New creates a new CircularLinkList
func New[T comparable]() *CircularLinkList[T] {
	return &CircularLinkList[T]{}
}

// NewFromSlice creates a new CircularLinkList from a slice
func NewFromSlice[T comparable](items []T) *CircularLinkList[T] {
	l := New[T]()
	for i := 0; i < len(items); i++ {
		l.Append(items[i])
	}
	return l
}

// Append adds a new node to the end of the list
func (l *CircularLinkList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		newNode.Next = newNode
		l.size++
		return
	}

	l.Tail.Next = newNode
	newNode.Next = l.Head
	l.Tail = newNode
	l.size++
}

// Prepend adds a new node to the beginning of the list
func (l *CircularLinkList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		newNode.Next = newNode
		l.size++
		return
	}

	newNode.Next = l.Head
	l.Head = newNode
	l.Tail.Next = newNode
	l.size++
}

// DeleteWithValue deletes the first node with the given value
func (l *CircularLinkList[T]) DeleteWithValue(value T) {
	if l.Head == nil {
		return
	}

	// Special case: head node needs to be deleted
	if l.Head.Value == value {
		if l.Head == l.Tail {
			l.Head = nil
			l.Tail = nil
			l.size = 0
			return
		}
		l.Head = l.Head.Next
		l.Tail.Next = l.Head
		l.size--
		return
	}

	current := l.Head
	for current.Next != l.Head {
		if current.Next.Value == value {
			if current.Next == l.Tail {
				l.Tail = current
			}
			current.Next = current.Next.Next
			l.size--
			return
		}
		current = current.Next
	}
}

// ToSlice returns the list as a slice
func (l *CircularLinkList[T]) ToSlice() []T {
	var result []T

	if l.Head == nil {
		return result
	}

	current := l.Head
	for {
		result = append(result, current.Value)
		current = current.Next
		if current == l.Head {
			break
		}
	}

	return result
}

// IsEmpty checks if the list is empty
func (l *CircularLinkList[T]) IsEmpty() bool {
	return l.Head == nil
}

// Find returns the first node with the given value
func (l *CircularLinkList[T]) Find(value T) (*Node[T], error) {
	if l.Head == nil {
		return nil, errors.New("value not found")
	}

	current := l.Head
	for {
		if current.Value == value {
			return current, nil
		}
		current = current.Next
		if current == l.Head {
			break
		}
	}

	return nil, errors.New("value not found")
}

// Reverse reverses the list
func (l *CircularLinkList[T]) Reverse() {
	if l.Head == nil {
		return
	}

	var prev, next *Node[T]
	current := l.Head
	l.Tail = l.Head

	for {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
		if current == l.Head {
			break
		}
	}

	l.Head.Next = prev
	l.Head = prev
}

// Size returns the number of nodes in the list
func (l *CircularLinkList[T]) Size() uint64 {
	return l.size
}

// CheckSize recalculate the size of the list
func (l *CircularLinkList[T]) CheckSize() {
	size := uint64(0)

	if l.Head == nil {
		l.size = 0
		return
	}

	current := l.Head
	for {
		size++
		current = current.Next
		if current == l.Head {
			break
		}
	}

	l.size = size
}

// GetFirst returns the first node in the list
func (l *CircularLinkList[T]) GetFirst() *Node[T] {
	return l.Head
}

// GetLast returns the last node in the list
func (l *CircularLinkList[T]) GetLast() *Node[T] {
	return l.Tail
}

// GetAt returns the node at the given index
func (l *CircularLinkList[T]) GetAt(index uint64) (*Node[T], error) {
	if l.Head == nil {
		return nil, errors.New(errIndexOutOfBound)
	}

	if index > l.size {
		// This is a circular list, so when the index is bigger than the size
		// we need to calculate the real index
		index = index % l.size
	}

	current := l.Head
	for i := uint64(0); i < index; i++ {
		current = current.Next
		if current == l.Head {
			return nil, errors.New(errIndexOutOfBound)
		}
	}

	return current, nil
}

// InsertAt inserts a new node at the given index
func (l *CircularLinkList[T]) InsertAt(index uint64, value T) error {
	if index > l.size {
		// This is a circular list, so when the index is bigger than the size
		// we need to calculate the real index
		index = index % l.size
	}

	if index == 0 {
		l.Prepend(value)
		return nil
	}

	current := l.Head
	for i := uint64(0); i < index-1; i++ {
		current = current.Next
		if current == l.Head {
			return errors.New(errIndexOutOfBound)
		}
	}

	newNode := &Node[T]{Value: value}
	newNode.Next = current.Next
	current.Next = newNode

	if current == l.Tail {
		l.Tail = newNode
	}

	return nil
}

// DeleteAt deletes the node at the given index
func (l *CircularLinkList[T]) DeleteAt(index uint64) error {
	if index > l.size {
		// This is a circular list, so when the index is bigger than the size
		// we need to calculate the real index
		index = index % l.size
	}

	if index == 0 {
		if l.Head == nil {
			return errors.New(errIndexOutOfBound)
		}
		if l.Head == l.Tail {
			l.Head = nil
			l.Tail = nil
			return nil
		}
		l.Head = l.Head.Next
		l.Tail.Next = l.Head
		return nil
	}

	current := l.Head
	for i := uint64(0); i < index-1; i++ {
		current = current.Next
		if current == l.Head {
			return errors.New(errIndexOutOfBound)
		}
	}

	if current.Next == l.Head {
		return errors.New(errIndexOutOfBound)
	}

	if current.Next == l.Tail {
		l.Tail = current
	}

	current.Next = current.Next.Next

	return nil
}

// Clear removes all nodes from the list
func (l *CircularLinkList[T]) Clear() {
	l.Head = nil
	l.Tail = nil
	l.size = 0
}

// Copy returns a copy of the list
func (l *CircularLinkList[T]) Copy() *CircularLinkList[T] {
	newList := New[T]()

	if l.Head == nil {
		return newList
	}

	current := l.Head
	for {
		newList.Append(current.Value)
		current = current.Next
		if current == l.Head {
			break
		}
	}

	return newList
}

// Merge appends all the nodes from another list to the current list
func (l *CircularLinkList[T]) Merge(list *CircularLinkList[T]) {
	if list.Head == nil {
		return
	}

	current := list.Head
	for {
		l.Append(current.Value)
		current = current.Next
		if current == list.Head {
			break
		}
	}
	list.Clear()
}

// Map generates a new list by applying the function to all the nodes in the list
func (l *CircularLinkList[T]) Map(f func(T) T) *CircularLinkList[T] {
	newList := New[T]()

	if l.Head == nil {
		return newList
	}

	current := l.Head
	for {
		newList.Append(f(current.Value))
		current = current.Next
		if current == l.Head {
			break
		}
	}

	return newList
}

// MapFrom generates a new list by applying the function to all the nodes in the list starting from the specified index
func (l *CircularLinkList[T]) MapFrom(start uint64, f func(T) T) (*CircularLinkList[T], error) {
	if l.Head == nil {
		return nil, errors.New(errIndexOutOfBound)
	}

	if start > l.size {
		// This is a circular list, so when the index is bigger than the size
		// we need to calculate the real index
		start = start % l.size
	}

	newList := New[T]()

	current := l.Head
	for i := uint64(0); i < start; i++ {
		current = current.Next
		if current == l.Head {
			return nil, errors.New(errIndexOutOfBound)
		}
	}

	for {
		newList.Append(f(current.Value))
		current = current.Next
		if current == l.Head {
			break
		}
	}

	return newList, nil
}

// MapRange generates a new list by applying the function to all the nodes in the list in the range [start, end)
func (l *CircularLinkList[T]) MapRange(start, end uint64, f func(T) T) (*CircularLinkList[T], error) {
	if l.Head == nil {
		return nil, errors.New(errIndexOutOfBound)
	}

	if start > l.size {
		// This is a circular list, so when the index is bigger than the size
		// we need to calculate the real index
		start = start % l.size
	}

	if end > l.size {
		// This is a circular list, so when the index is bigger than the size
		// we need to calculate the real index
		end = end % l.size
	}

	if start > end {
		return nil, errors.New(errIndexOutOfBound)
	}

	newList := New[T]()

	current := l.Head
	for i := uint64(0); i < start; i++ {
		current = current.Next
		if current == l.Head {
			return nil, errors.New(errIndexOutOfBound)
		}
	}

	for i := start; i < end; i++ {
		newList.Append(f(current.Value))
		current = current.Next
		if current == l.Head {
			break
		}
	}

	return newList, nil
}

// ForEach applies the function to each node in the list
func (l *CircularLinkList[T]) ForEach(f func(*T)) {
	if l.Head == nil {
		return
	}

	current := l.Head
	for {
		f(&current.Value)
		current = current.Next
		if current == l.Head {
			break
		}
	}
}

// ForRange applies the function to each node in the list in the range [start, end]
func (l *CircularLinkList[T]) ForRange(start, end uint64, f func(*T)) error {
	if l.Head == nil {
		return errors.New(errIndexOutOfBound)
	}

	if start > l.size {
		// This is a circular list, so when the index is bigger than the size
		// we need to calculate the real index
		start = start % l.size
	}

	if end > l.size {
		// This is a circular list, so when the index is bigger than the size
		// we need to calculate the real index
		end = end % l.size
	}

	if start > end {
		return errors.New(errIndexOutOfBound)
	}

	current := l.Head
	for i := uint64(0); i < start; i++ {
		current = current.Next
		if current == l.Head {
			return errors.New(errIndexOutOfBound)
		}
	}

	for i := start; i <= end; i++ {
		f(&current.Value)
		current = current.Next
		if current == l.Head {
			break
		}
	}

	return nil
}

// ForFrom applies the function to each node in the list starting from the index
func (l *CircularLinkList[T]) ForFrom(start uint64, f func(*T)) error {
	if l.Head == nil {
		return errors.New(errIndexOutOfBound)
	}

	if start > l.size {
		// This is a circular list, so when the index is bigger than the size
		// we need to calculate the real index
		start = start % l.size
	}

	current := l.Head
	for i := uint64(0); i < start; i++ {
		current = current.Next
		if current == l.Head {
			return errors.New(errIndexOutOfBound)
		}
	}

	for {
		f(&current.Value)
		current = current.Next
		if current == l.Head {
			break
		}
	}

	return nil
}

// Filter removes nodes from the list that don't match the predicate
func (l *CircularLinkList[T]) Filter(f func(T) bool) {
	if l.Head == nil {
		return
	}

	// Handle the head separately
	for l.Head != nil && !f(l.Head.Value) {
		// Remove the head node
		if l.Head == l.Tail {
			// List has only one element
			l.Head = nil
			l.Tail = nil
			l.size = 0
			return
		} else {
			l.Head = l.Head.Next
			l.Tail.Next = l.Head
			l.size--
		}
	}

	if l.Head == nil {
		return
	}

	current := l.Head

	for current.Next != l.Head {
		if !f(current.Next.Value) {
			// Remove the node
			if current.Next == l.Tail {
				l.Tail = current
				l.Tail.Next = l.Head
			} else {
				current.Next = current.Next.Next
			}
			l.size--
		} else {
			current = current.Next
		}
	}
}

// Reduce reduces the list to a single value
func (l *CircularLinkList[T]) Reduce(f func(T, T) T) (T, error) {
	if l.Head == nil {
		var rVal T
		return rVal, errors.New(errListIsEmpty)
	}

	result := l.Head.Value
	current := l.Head.Next
	for {
		result = f(result, current.Value)
		current = current.Next
		if current == l.Head {
			break
		}
	}

	return result, nil
}

// ReduceFrom reduces the list to a single value starting from the index
func (l *CircularLinkList[T]) ReduceFrom(start uint64, f func(T, T) T) (T, error) {
	if l.Head == nil || l.size == 0 {
		var rVal T
		return rVal, errors.New(errListIsEmpty)
	}

	if start > l.size {
		// This is a circular list, so when the index is bigger than the size
		// we need to calculate the real index
		start = start % l.size
	}

	current := l.Head
	for i := uint64(0); i < start; i++ {
		current = current.Next
		if current == l.Head {
			break
		}
	}

	result := current.Value
	current = current.Next
	for {
		result = f(result, current.Value)
		current = current.Next
		if current == l.Head {
			break
		}
	}

	return result, nil
}

// ReduceRange reduces the list to a single value in the range [start, end)
func (l *CircularLinkList[T]) ReduceRange(start, end uint64, f func(T, T) T) (T, error) {
	if l.Head == nil {
		var rVal T
		return rVal, errors.New(errListIsEmpty)
	}

	if start > l.size {
		// This is a circular list, so when the index is bigger than the size
		// we need to calculate the real index
		start = start % l.size
	}

	if end > l.size {
		// This is a circular list, so when the index is bigger than the size
		// we need to calculate the real index
		end = end % l.size
	}

	if start > end {
		var rVal T
		return rVal, errors.New(errIndexOutOfBound)
	}

	current := l.Head
	for i := uint64(0); i < start; i++ {
		current = current.Next
		if current == l.Head {
			break
		}
	}

	result := current.Value
	current = current.Next
	for i := start; i < end; i++ {
		result = f(result, current.Value)
		current = current.Next
		if current == l.Head {
			break
		}
	}

	return result, nil
}
