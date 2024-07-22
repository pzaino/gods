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

// Package dlinkList provides a non-concurrent-safe doubly linked list.
package dlinkList

import "errors"

const errIndexOutOfBound = "index out of bounds"

// Node is a representation of a node in a doubly linked list
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

// DLinkList is a representation of a doubly linked list
type DLinkList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

// NewDLinkList creates a new doubly linked list
func NewDLinkList[T comparable]() *DLinkList[T] {
	return &DLinkList[T]{}
}

// Append adds a new node to the end of the doubly linked list
func (l *DLinkList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	newNode.Prev = l.Tail
	l.Tail.Next = newNode
	l.Tail = newNode
}

// Prepend adds a new node to the beginning of the doubly linked list
func (l *DLinkList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	newNode.Next = l.Head
	l.Head.Prev = newNode
	l.Head = newNode
}

// Insert inserts a new node with the given value at first available index
// note: this is just an alias for Append
func (l *DLinkList[T]) Insert(value T) error {
	ln := l.Size()
	l.Append(value)
	if ln == l.Size() {
		return errors.New("failed to insert")
	}
	return nil
}

// InsertAfter inserts a new node with the given value after the node with the given value
func (l *DLinkList[T]) InsertAfter(value, newValue T) {
	node, err := l.Find(value)
	if err != nil {
		return
	}

	newNode := &Node[T]{Value: newValue}
	newNode.Next = node.Next
	newNode.Prev = node
	node.Next = newNode
	if newNode.Next != nil {
		newNode.Next.Prev = newNode
	}
}

// InsertBefore inserts a new node with the given value before the node with the given value
func (l *DLinkList[T]) InsertBefore(value, newValue T) {
	node, err := l.Find(value)
	if err != nil {
		return
	}

	newNode := &Node[T]{Value: newValue}
	newNode.Next = node
	newNode.Prev = node.Prev
	node.Prev = newNode
	if newNode.Prev != nil {
		newNode.Prev.Next = newNode
	}
}

// InsertAt inserts a new node with the given value at the given index
func (l *DLinkList[T]) InsertAt(index int, value T) error {
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
	newNode.Prev = current
	current.Next = newNode
	if newNode.Next != nil {
		newNode.Next.Prev = newNode
	}

	return nil
}

// DeleteWithValue deletes the first occurrence of a node with the given value
func (l *DLinkList[T]) DeleteWithValue(value T) {
	if l.Head == nil {
		return
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		if l.Head != nil {
			l.Head.Prev = nil
		}
		return
	}

	current := l.Head
	for current.Next != nil {
		if current == nil {
			return
		}
		if current.Next.Value == value {
			current.Next = current.Next.Next
			if current.Next != nil {
				current.Next.Prev = current
			}
			return
		}
		current = current.Next
	}
}

func (l *DLinkList[T]) Remove(value T) {
	l.DeleteWithValue(value)
}

func (l *DLinkList[T]) RemoveAt(index int) error {
	return l.DeleteAt(index)
}

// Delete deletes the first node with the given value
func (l *DLinkList[T]) Delete(value T) {
	node, err := l.Find(value)
	if err != nil {
		return
	}

	if node.Prev == nil {
		l.Head = node.Next
		if l.Head != nil {
			l.Head.Prev = nil
		}
		return
	}

	if node.Next == nil {
		l.Tail = node.Prev
		l.Tail.Next = nil
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// DeleteLast deletes the last node in the doubly linked list
func (l *DLinkList[T]) DeleteLast() {
	if l.Tail == nil {
		return
	}

	if l.Tail.Prev == nil {
		l.Head = nil
		l.Tail = nil
		return
	}

	l.Tail = l.Tail.Prev
	l.Tail.Next = nil
}

// DeleteFirst deletes the first node in the doubly linked list
func (l *DLinkList[T]) DeleteFirst() {
	if l.Head == nil {
		return
	}

	if l.Head.Next == nil {
		l.Head = nil
		l.Tail = nil
		return
	}

	l.Head = l.Head.Next
	l.Head.Prev = nil
}

// DeleteAt deletes the node at the given index
func (l *DLinkList[T]) DeleteAt(index int) error {
	if index < 0 {
		return errors.New(errIndexOutOfBound)
	}

	if index == 0 {
		if l.Head == nil {
			return errors.New(errIndexOutOfBound)
		}
		l.Head = l.Head.Next
		l.Head.Prev = nil
		return nil
	}

	current := l.Head
	for i := 0; i < index; i++ {
		if current == nil {
			return errors.New(errIndexOutOfBound)
		}
		current = current.Next
	}

	if current == nil {
		return errors.New(errIndexOutOfBound)
	}

	if current.Next == nil {
		current.Prev.Next = nil
		return nil
	}

	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev

	return nil
}

// ToSlice converts the doubly linked list to a slice
func (l *DLinkList[T]) ToSlice() []T {
	var result []T

	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result
}

// ToSliceReverse converts the doubly linked list to a slice in reverse order
func (l *DLinkList[T]) ToSliceReverse() []T {
	var result []T

	current := l.Tail
	for current != nil {
		result = append(result, current.Value)
		current = current.Prev
	}

	return result
}

// ToSliceFromIndex converts the doubly linked list to a slice starting from the given index
func (l *DLinkList[T]) ToSliceFromIndex(index int) []T {
	var result []T

	current, err := l.GetAt(index)
	if err != nil {
		return result
	}

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result
}

// ToSliceReverseFromIndex converts the doubly linked list to a slice in reverse order starting from the given index
func (l *DLinkList[T]) ToSliceReverseFromIndex(index int) []T {
	var result []T

	current, err := l.GetAt((l.Size() - 1) - index)
	if err != nil {
		return result
	}

	for current != nil {
		result = append(result, current.Value)
		current = current.Prev
	}

	return result
}

// Reverse reverses the doubly linked list
func (l *DLinkList[T]) Reverse() {
	current := l.Head
	var prev *Node[T]

	for current != nil {
		next := current.Next
		current.Next = prev
		current.Prev = next
		prev = current
		current = next
	}

	l.Head, l.Tail = l.Tail, l.Head
}

// Find returns the first node with the given value
func (l *DLinkList[T]) Find(value T) (*Node[T], error) {
	current := l.Head
	for current != nil {
		if current.Value == value {
			return current, nil
		}
		current = current.Next
	}

	return nil, errors.New("value not found")
}

// IsEmpty returns true if the doubly linked list is empty
func (l *DLinkList[T]) IsEmpty() bool {
	return l.Head == nil
}

// GetAt returns the node at the given index
func (l *DLinkList[T]) GetAt(index int) (*Node[T], error) {
	if index < 0 {
		return nil, errors.New(errIndexOutOfBound)
	}

	current := l.Head
	if current == nil {
		return nil, errors.New(errIndexOutOfBound)
	}
	if index == 0 {
		return current, nil
	}

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

// GetLast returns the last node in the doubly linked list
func (l *DLinkList[T]) GetLast() *Node[T] {
	return l.Tail
}

// GetFirst returns the first node in the doubly linked list
func (l *DLinkList[T]) GetFirst() *Node[T] {
	return l.Head
}

// Size returns the number of nodes in the doubly linked list
func (l *DLinkList[T]) Size() int {
	size := 0
	current := l.Head
	for current != nil {
		size++
		current = current.Next
	}

	return size
}

// Clear removes all nodes from the doubly linked list
func (l *DLinkList[T]) Clear() {
	l.Head = nil
	l.Tail = nil
}

// Contains returns true if the doubly linked list contains the given value
func (l *DLinkList[T]) Contains(value T) bool {
	current := l.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}

	return false
}

// ForEach traverses the doubly linked list and applies the given function to each node
func (l *DLinkList[T]) ForEach(f func(*T)) {
	current := l.Head
	for current != nil {
		f(&current.Value)
		current = current.Next
	}
}

// Any returns true if the given function returns true for any node in the doubly linked list
func (l *DLinkList[T]) Any(f func(T) bool) bool {
	current := l.Head
	for current != nil {
		if f(current.Value) {
			return true
		}
		current = current.Next
	}

	return false
}

// All returns true if the given function returns true for all nodes in the doubly linked list
func (l *DLinkList[T]) All(f func(T) bool) bool {
	current := l.Head
	for current != nil {
		if !f(current.Value) {
			return false
		}
		current = current.Next
	}

	return true
}

// IndexOf returns the index of the first occurrence of the given value in the doubly linked list
func (l *DLinkList[T]) IndexOf(value T) int {
	current := l.Head
	index := 0
	for current != nil {
		if current.Value == value {
			return index
		}
		index++
		current = current.Next
	}

	return -1
}

// LastIndexOf returns the index of the last occurrence of the given value in the doubly linked list
func (l *DLinkList[T]) LastIndexOf(value T) int {
	current := l.Tail
	index := l.Size() - 1
	for current != nil {
		if current.Value == value {
			return index
		}
		index--
		current = current.Prev
	}

	return -1
}

// Filter returns a new doubly linked list containing only the nodes that satisfy the given function
func (l *DLinkList[T]) Filter(f func(T) bool) *DLinkList[T] {
	result := NewDLinkList[T]()

	current := l.Head
	for current != nil {
		if f(current.Value) {
			result.Append(current.Value)
		}
		current = current.Next
	}

	return result
}

// Map returns a new doubly linked list containing the result of applying the given function to each node
func (l *DLinkList[T]) Map(f func(T) T) *DLinkList[T] {
	result := NewDLinkList[T]()

	current := l.Head
	for current != nil {
		result.Append(f(current.Value))
		current = current.Next
	}

	return result
}

// Reduce reduces the doubly linked list to a single value using the given function
func (l *DLinkList[T]) Reduce(f func(T, T) T) T {
	if l.IsEmpty() {
		var rVal T
		return rVal
	}

	result := l.Head.Value
	current := l.Head.Next
	for current != nil {
		result = f(result, current.Value)
		current = current.Next
	}

	return result
}

// Copy returns a new doubly linked list with the same nodes as the original doubly linked list
func (l *DLinkList[T]) Copy() *DLinkList[T] {
	newList := NewDLinkList[T]()

	current := l.Head
	for current != nil {
		newList.Append(current.Value)
		current = current.Next
	}

	return newList
}

// Merge appends the nodes of the given doubly linked list to the original doubly linked list
func (l *DLinkList[T]) Merge(list *DLinkList[T]) {
	if list.IsEmpty() {
		return
	}
	current := list.Head
	for current != nil {
		l.Append(current.Value)
		current = current.Next
	}
}

// ReverseCopy returns a new doubly linked list with the nodes of the original doubly linked list in reverse order
func (l *DLinkList[T]) ReverseCopy() *DLinkList[T] {
	newList := NewDLinkList[T]()

	current := l.Tail
	for current != nil {
		newList.Append(current.Value)
		current = current.Prev
	}

	return newList
}

// ReverseMerge appends the nodes of the given doubly linked list to the original doubly linked list in reverse order
func (l *DLinkList[T]) ReverseMerge(list *DLinkList[T]) {
	current := list.Tail
	for current != nil {
		l.Append(current.Value)
		current = current.Prev
	}
}

// Equal returns true if the given doubly linked list is equal to the original doubly linked list
func (l *DLinkList[T]) Equal(list *DLinkList[T]) bool {
	current1 := l.Head
	current2 := list.Head

	for current1 != nil && current2 != nil {
		if current1.Value != current2.Value {
			return false
		}
		current1 = current1.Next
		current2 = current2.Next
	}

	return current1 == nil && current2 == nil
}

// Swap swaps the nodes at the given indices
func (l *DLinkList[T]) Swap(i, j int) error {
	node1, err := l.GetAt(i)
	if err != nil {
		return err
	}

	node2, err := l.GetAt(j)
	if err != nil {
		return err
	}

	node1.Value, node2.Value = node2.Value, node1.Value

	return nil
}

// Sort sorts the doubly linked list according to the given function
// for example, to sort a list of integers in ascending order, use:
// list.Sort(func(a, b int) bool { return a < b })
func (l *DLinkList[T]) Sort(f func(T, T) bool) {
	if l.IsEmpty() {
		return
	}

	if l.Size() < 2 {
		return
	}

	nodes := make([]*Node[T], 0, l.Size())
	current := l.Head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	quickSort(nodes, f, 0, len(nodes)-1)

	l.Head = nodes[0]
	l.Tail = nodes[len(nodes)-1]

	var i int
	for i = 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
		nodes[i+1].Prev = nodes[i]
	}
	nodes[i].Next = nil
}

func quickSort[T comparable](nodes []*Node[T], f func(T, T) bool, low, high int) {
	if low < high {
		p := partition(nodes, f, low, high)
		quickSort(nodes, f, low, p-1)
		quickSort(nodes, f, p+1, high)
	}
}

func partition[T comparable](nodes []*Node[T], f func(T, T) bool, low, high int) int {
	pivot := nodes[high]
	i := low

	for j := low; j < high; j++ {
		if f(nodes[j].Value, pivot.Value) {
			nodes[i], nodes[j] = nodes[j], nodes[i]
			i++
		}
	}

	nodes[i], nodes[high] = nodes[high], nodes[i]

	return i
}

// FindAll returns a new doubly linked list containing all nodes that satisfy the given function
func (l *DLinkList[T]) FindAll(f func(T) bool) *DLinkList[T] {
	newList := NewDLinkList[T]()

	current := l.Head
	for current != nil {
		if f(current.Value) {
			newList.Append(current.Value)
		}
		current = current.Next
	}

	return newList
}

// FindLast returns the last node that satisfies the given function
func (l *DLinkList[T]) FindLast(f func(T) bool) (*Node[T], error) {
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

// FindLastIndex returns the index of the last node that satisfies the given function
func (l *DLinkList[T]) FindLastIndex(f func(T) bool) int {
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

// FindIndex returns the index of the first node that satisfies the given function
func (l *DLinkList[T]) FindIndex(f func(T) bool) int {
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
