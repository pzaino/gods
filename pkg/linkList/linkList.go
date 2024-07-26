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

// Package linkList provides a non-concurrent-safe linked list.
package linkList

import "errors"

const (
	errIndexOutOfBound = "index out of bounds"
)

// Node represents a node in the linked list
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// LinkList represents a linked list
type LinkList[T comparable] struct {
	Head *Node[T]
}

// NewLinkList creates a new LinkList
func NewLinkList[T comparable]() *LinkList[T] {
	return &LinkList[T]{}
}

// NewLinkListFromSlice creates a new LinkList from a slice
func NewLinkListFromSlice[T comparable](items []T) *LinkList[T] {
	l := NewLinkList[T]()
	for i := 0; i < len(items); i++ {
		l.Append(items[i])
	}
	return l
}

// Append adds a new node to the end of the list
func (l *LinkList[T]) Append(value T) {
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
func (l *LinkList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value}

	newNode.Next = l.Head
	l.Head = newNode
}

// DeleteWithValue deletes the first node with the given value
func (l *LinkList[T]) DeleteWithValue(value T) {
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
func (l *LinkList[T]) ToSlice() []T {
	var result []T

	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result
}

// IsEmpty checks if the list is empty
func (l *LinkList[T]) IsEmpty() bool {
	return l.Head == nil
}

// Find returns the first node with the given value
func (l *LinkList[T]) Find(value T) (*Node[T], error) {
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
func (l *LinkList[T]) Reverse() {
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
func (l *LinkList[T]) Size() int {
	size := 0
	current := l.Head
	for current != nil {
		size++
		current = current.Next
	}

	return size
}

// Values returns all the values in the list
func (l *LinkList[T]) GetFirst() *Node[T] {
	if l == nil {
		return nil
	}

	return l.Head
}

// GetLast returns the last node in the list
func (l *LinkList[T]) GetLast() *Node[T] {
	if l == nil {
		return nil
	}

	if l.Head == nil {
		return nil
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	return current
}

// GetAt returns the node at the given index
func (l *LinkList[T]) GetAt(index int) (*Node[T], error) {
	if index < 0 {
		return nil, errors.New(errIndexOutOfBound)
	}

	if l == nil {
		return nil, nil
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
func (l *LinkList[T]) InsertAt(index int, value T) error {
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
func (l *LinkList[T]) DeleteAt(index int) error {
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

// Remove is just an alias for DeleteWithValue
func (l *LinkList[T]) Remove(value T) {
	l.DeleteWithValue(value)
}

// Clear removes all nodes from the list
func (l *LinkList[T]) Clear() {
	l.Head = nil
}

// Copy returns a copy of the list
func (l *LinkList[T]) Copy() *LinkList[T] {
	newList := NewLinkList[T]()

	current := l.Head
	for current != nil {
		newList.Append(current.Value)
		current = current.Next
	}

	return newList
}

// Merge appends all the nodes from another list to the current list
func (l *LinkList[T]) Merge(list *LinkList[T]) {
	current := list.Head
	for current != nil {
		l.Append(current.Value)
		current = current.Next
	}
	// Clear the list
	list.Clear()
}

// Map generates a new list by applying the function to all the nodes in the list
func (l *LinkList[T]) Map(f func(T) T) *LinkList[T] {
	newList := NewLinkList[T]()
	current := l.Head
	for current != nil {
		newList.Append(f(current.Value))
		current = current.Next
	}
	return newList
}

// MapFrom generates a new list by applying the function to all the nodes in the list starting from the specified index
func (l *LinkList[T]) MapFrom(start int, f func(T) T) (*LinkList[T], error) {
	if start < 0 {
		return nil, errors.New(errIndexOutOfBound)
	}

	newList := NewLinkList[T]()
	current, err := l.GetAt(start)
	if err != nil {
		return nil, err
	}

	for current != nil {
		newList.Append(f(current.Value))
		current = current.Next
	}

	return newList, nil
}

// MapRange generates a new list by applying the function to all the nodes in the list within the specified range
func (l *LinkList[T]) MapRange(start, end int, f func(T) T) (*LinkList[T], error) {
	if start < 0 || end < 0 {
		return nil, errors.New(errIndexOutOfBound)
	}

	if start > end {
		return nil, errors.New("start index cannot be greater than end index")
	}

	if end >= l.Size() {
		return nil, errors.New(errIndexOutOfBound)
	}

	newList := NewLinkList[T]()
	current, err := l.GetAt(start)
	if err != nil {
		return nil, err
	}

	for i := start; i <= end; i++ {
		newList.Append(f(current.Value))
		current = current.Next
	}

	return newList, nil
}

// Filter removes nodes from the list that don't match the predicate
func (l *LinkList[T]) Filter(f func(T) bool) {
	if l.Head == nil {
		return
	}

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
func (l *LinkList[T]) Reduce(f func(T, T) T, initial T) T {
	result := initial

	current := l.Head
	for current != nil {
		result = f(result, current.Value)
		current = current.Next
	}

	return result
}

// ForEach applies the function to all the nodes in the list
func (l *LinkList[T]) ForEach(f func(*T)) {
	current := l.Head
	for current != nil {
		f(&current.Value)
		current = current.Next
	}
}

// ForRange applies the function to all the nodes in the list within the specified range
func (l *LinkList[T]) ForRange(start, end int, f func(*T)) error {
	if start < 0 || end < 0 {
		return errors.New(errIndexOutOfBound)
	}

	if start > end {
		return errors.New("start index cannot be greater than end index")
	}

	if end >= l.Size() {
		return errors.New(errIndexOutOfBound)
	}

	current, err := l.GetAt(start)
	if err != nil {
		return err
	}

	for i := start; i <= end; i++ {
		f(&current.Value)
		current = current.Next
		if current == nil {
			break
		}
	}

	return nil
}

// ForFrom applies the function to all the nodes in the list starting from the specified index
func (l *LinkList[T]) ForFrom(start int, f func(*T)) error {
	if start < 0 {
		return errors.New(errIndexOutOfBound)
	}

	current, err := l.GetAt(start)
	if err != nil {
		return err
	}

	for current != nil {
		f(&current.Value)
		current = current.Next
		if current == nil {
			break
		}
	}

	return nil
}

// Any checks if any node in the list matches the predicate
func (l *LinkList[T]) Any(f func(T) bool) bool {
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
func (l *LinkList[T]) All(f func(T) bool) bool {
	if l == nil {
		return false
	}
	if l.Head == nil {
		return false
	}

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
func (l *LinkList[T]) Contains(value T) bool {
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
func (l *LinkList[T]) IndexOf(value T) int {
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
func (l *LinkList[T]) LastIndexOf(value T) int {
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
func (l *LinkList[T]) FindIndex(f func(T) bool) int {
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
func (l *LinkList[T]) FindLastIndex(f func(T) bool) int {
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
func (l *LinkList[T]) FindAll(f func(T) bool) *LinkList[T] {
	newList := NewLinkList[T]()

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
func (l *LinkList[T]) FindLast(f func(T) bool) (*Node[T], error) {
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
func (l *LinkList[T]) FindAllIndexes(f func(T) bool) []int {
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
