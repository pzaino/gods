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
	size uint64
}

// New creates a new doubly linked list
func New[T comparable]() *DLinkList[T] {
	return &DLinkList[T]{}
}

// Append adds a new node to the end of the doubly linked list
func (l *DLinkList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		l.size++
		return
	}

	newNode.Prev = l.Tail
	l.Tail.Next = newNode
	l.Tail = newNode
	l.size++
}

// Prepend adds a new node to the beginning of the doubly linked list
func (l *DLinkList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		l.size++
		return
	}

	newNode.Next = l.Head
	l.Head.Prev = newNode
	l.Head = newNode
	l.size++
}

// Insert inserts a new node with the given value at first available index
// note: this is just an alias for Append
func (l *DLinkList[T]) Insert(value T) error {
	ln := l.size
	l.Append(value)
	if ln == l.size {
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
func (l *DLinkList[T]) InsertAt(index uint64, value T) error {
	if index > l.size {
		return errors.New(errIndexOutOfBound)
	}

	if index == 0 {
		l.Prepend(value)
		return nil
	}

	current := l.Head
	for i := uint64(0); i < index-1; i++ {
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
	l.size++

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
		l.size--
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
			l.size--
			return
		}
		current = current.Next
	}
}

func (l *DLinkList[T]) Remove(value T) {
	l.DeleteWithValue(value)
}

func (l *DLinkList[T]) RemoveAt(index uint64) error {
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
		l.size--
		return
	}

	if node.Next == nil {
		l.Tail = node.Prev
		l.Tail.Next = nil
		l.size--
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	l.size--
}

// DeleteLast deletes the last node in the doubly linked list
func (l *DLinkList[T]) DeleteLast() {
	if l.Tail == nil {
		return
	}

	if l.Tail.Prev == nil {
		l.Head = nil
		l.Tail = nil
		l.size--
		return
	}

	l.Tail = l.Tail.Prev
	l.Tail.Next = nil
	l.size--
}

// DeleteFirst deletes the first node in the doubly linked list
func (l *DLinkList[T]) DeleteFirst() {
	if l.Head == nil {
		return
	}

	if l.Head.Next == nil {
		l.Head = nil
		l.Tail = nil
		l.size--
		return
	}

	l.Head = l.Head.Next
	l.Head.Prev = nil
	l.size--
}

// DeleteAt deletes the node at the given index
func (l *DLinkList[T]) DeleteAt(index uint64) error {
	if index > l.size {
		return errors.New(errIndexOutOfBound)
	}

	// delete the first node
	if index == 0 {
		if l.Head == nil {
			return errors.New(errIndexOutOfBound)
		}
		l.Head = l.Head.Next
		l.Head.Prev = nil
		l.size--
		return nil
	}

	// find the node at the given index
	current := l.Head
	for i := uint64(0); i < index; i++ {
		if current == nil {
			return errors.New(errIndexOutOfBound)
		}
		current = current.Next
	}

	// Check if the node is valid
	if current == nil {
		return errors.New(errIndexOutOfBound)
	}

	// this is the last node
	if current.Next == nil {
		current.Prev.Next = nil
		l.size--
		return nil
	}

	// regular node in the middle
	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev
	l.size--

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
func (l *DLinkList[T]) ToSliceFromIndex(index uint64) []T {
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
func (l *DLinkList[T]) ToSliceReverseFromIndex(index uint64) []T {
	var result []T

	if index > l.size {
		return result
	}

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

	return nil, errors.New("value  not found")
}

// IsEmpty returns true if the doubly linked list is empty
func (l *DLinkList[T]) IsEmpty() bool {
	return l.Head == nil
}

// GetAt returns the node at the given index
func (l *DLinkList[T]) GetAt(index uint64) (*Node[T], error) {
	if index > l.size {
		return nil, errors.New(errIndexOutOfBound)
	}

	current := l.Head
	if current == nil {
		return nil, errors.New(errIndexOutOfBound)
	}
	if index == 0 {
		return current, nil
	}

	for i := uint64(0); i < index; i++ {
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
func (l *DLinkList[T]) Size() uint64 {
	return l.size
}

// CheckSize recalculates the size of the doubly linked list
func (l *DLinkList[T]) CheckSize() {
	size := uint64(0)
	current := l.Head
	for current != nil {
		size++
		current = current.Next
	}

	l.size = size
}

// Clear removes all nodes from the doubly linked list
func (l *DLinkList[T]) Clear() {
	l.Head = nil
	l.Tail = nil
	l.size = 0
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
	if l.IsEmpty() {
		return
	}

	current := l.Head
	for current != nil {
		f(&current.Value)
		current = current.Next
	}
}

// ForFrom traverses the doubly linked list starting from the given index and applies the given function to each node
func (l *DLinkList[T]) ForFrom(index uint64, f func(*T)) {
	if index > l.size {
		return
	}

	if l.IsEmpty() {
		return
	}

	current, err := l.GetAt(index)
	if err != nil {
		return
	}

	for current != nil {
		f(&current.Value)
		current = current.Next
		if current == nil {
			break
		}
	}
}

// ForEachReverse traverses the doubly linked list in reverse order and applies the given function to each node
func (l *DLinkList[T]) ForEachReverse(f func(*T)) {
	if l.IsEmpty() {
		return
	}

	current := l.Tail
	for current != nil {
		f(&current.Value)
		current = current.Prev
	}
}

// ForReverseFrom traverses the doubly linked list in reverse order starting from the given index and applies the given function to each node
func (l *DLinkList[T]) ForReverseFrom(index uint64, f func(*T)) {
	if index > l.size {
		return
	}

	if l.IsEmpty() {
		return
	}

	current, err := l.GetAt((l.Size() - 1) - index)
	if err != nil {
		return
	}

	for current != nil {
		f(&current.Value)
		current = current.Prev
		if current == nil {
			break
		}
	}
}

// ForRange traverses the doubly linked list from the start index to the end index and applies the given function to each node
func (l *DLinkList[T]) ForRange(start, end uint64, f func(*T)) {
	if start > end || start > l.size || end > l.size {
		return
	}

	if l.IsEmpty() {
		return
	}

	current, err := l.GetAt(start)
	if err != nil {
		return
	}

	for i := start; i <= end; i++ {
		f(&current.Value)
		current = current.Next
		if current == nil {
			break
		}
	}
}

// ForReverseRange traverses the doubly linked list in reverse order from the start index to the end index and applies the given function to each node
func (l *DLinkList[T]) ForReverseRange(start, end uint64, f func(*T)) {
	if start > end {
		return
	}

	if l.IsEmpty() {
		return
	}

	if l.Size() < start {
		return
	}

	if l.Size() < end {
		end = l.Size() - 1
	}

	if end < start {
		return
	}

	if start == 0 && end == 0 {
		f(&l.Head.Value)
		return
	}

	current, err := l.GetAt((l.Size() - 1) - start)
	if err != nil {
		return
	}

	for i := start; i <= end; i++ {
		f(&current.Value)
		current = current.Prev
		if current == nil {
			break
		}
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
func (l *DLinkList[T]) LastIndexOf(value T) (uint64, error) {
	current := l.Tail
	index := l.Size() - 1
	for current != nil {
		if current.Value == value {
			return index, nil
		}
		index--
		current = current.Prev
	}

	return 0, errors.New("value not found")
}

// removeNode removes a node from the doubly linked list
// note: this is a private method and should not be used outside of this package
func (l *DLinkList[T]) removeNode(node *Node[T]) {
	if node.Prev == nil {
		l.Head = node.Next
		if l.Head != nil {
			l.Head.Prev = nil
		}
	} else {
		node.Prev.Next = node.Next
	}

	if node.Next == nil {
		l.Tail = node.Prev
		if l.Tail != nil {
			l.Tail.Next = nil
		}
	} else {
		node.Next.Prev = node.Prev
	}

	l.size--
}

// Filter returns a new doubly linked list containing only the nodes that satisfy the given function
func (l *DLinkList[T]) Filter(f func(T) bool) {
	if l.size == 0 || l.Head == nil {
		return
	}

	current := l.Head
	for current != nil {
		next := current.Next // Store the next node
		if !f(current.Value) {
			l.removeNode(current)
		}
		current = next // Move to the next node
	}

	// If the list is now empty after filtering, reset the Tail pointer
	if l.size == 0 {
		l.Tail = nil
	}
}

// Map returns a new doubly linked list containing the result of applying the given function to each node
func (l *DLinkList[T]) Map(f func(T) T) *DLinkList[T] {
	result := New[T]()

	current := l.Head
	for current != nil {
		result.Append(f(current.Value))
		current = current.Next
	}

	return result
}

// MapFrom returns a new doubly linked list containing the result of applying the given function to each node starting from the given index
func (l *DLinkList[T]) MapFrom(index uint64, f func(T) T) *DLinkList[T] {
	result := New[T]()

	if index > l.size {
		return result
	}

	if l.IsEmpty() {
		return result
	}

	current, err := l.GetAt(index)
	if err != nil {
		return result
	}

	for current != nil {
		result.Append(f(current.Value))
		current = current.Next
		if current == nil {
			break
		}
	}

	return result
}

// MapRange returns a new doubly linked list containing the result of applying the given function to each node in the range [start, end)
func (l *DLinkList[T]) MapRange(start, end uint64, f func(T) T) *DLinkList[T] {
	result := New[T]()

	if start > end || start > l.size || end > l.size {
		return result
	}

	if l.IsEmpty() {
		return result
	}

	current, err := l.GetAt(start)
	if err != nil {
		return result
	}

	for i := start; i <= end; i++ {
		result.Append(f(current.Value))
		current = current.Next
		if current == nil {
			break
		}
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
	newList := New[T]()

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
	// clear the list
	list.Clear()
}

// ReverseCopy returns a new doubly linked list with the nodes of the original doubly linked list in reverse order
func (l *DLinkList[T]) ReverseCopy() *DLinkList[T] {
	newList := New[T]()

	current := l.Tail
	for current != nil {
		newList.Append(current.Value)
		current = current.Prev
	}

	return newList
}

// ReverseMerge appends the nodes of the given doubly linked list to the original doubly linked list in reverse order
func (l *DLinkList[T]) ReverseMerge(list *DLinkList[T]) {
	if list.IsEmpty() {
		return
	}

	current := list.Tail
	for current != nil {
		l.Append(current.Value)
		current = current.Prev
	}
	// clear the list
	list.Clear()
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
func (l *DLinkList[T]) Swap(i, j uint64) error {
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
	newList := New[T]()

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
