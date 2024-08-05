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

// Package csdlinkList provides a concurrency-safe doubly linked list using dlinkList package.
package csdlinkList

import (
	"sync"

	dlinkList "github.com/pzaino/gods/pkg/dlinkList"
)

// CSDLinkList is a concurrency-safe doubly linked list.
type CSDLinkList[T comparable] struct {
	mu sync.RWMutex
	l  *dlinkList.DLinkList[T]
}

// New creates a new concurrency-safe doubly linked list.
func New[T comparable]() *CSDLinkList[T] {
	return &CSDLinkList[T]{l: dlinkList.New[T]()}
}

// Append adds a new node to the end of the doubly linked list.
func (cs *CSDLinkList[T]) Append(value T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.Append(value)
}

// Prepend adds a new node to the beginning of the doubly linked list.
func (cs *CSDLinkList[T]) Prepend(value T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.Prepend(value)
}

// Insert inserts a new node with the given value at the first available index.
func (cs *CSDLinkList[T]) Insert(value T) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.Insert(value)
}

// InsertAfter inserts a new node with the given value after the node with the given value.
func (cs *CSDLinkList[T]) InsertAfter(value, newValue T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.InsertAfter(value, newValue)
}

// InsertBefore inserts a new node with the given value before the node with the given value.
func (cs *CSDLinkList[T]) InsertBefore(value, newValue T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.InsertBefore(value, newValue)
}

// InsertAt inserts a new node with the given value at the given index.
func (cs *CSDLinkList[T]) InsertAt(index uint64, value T) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.InsertAt(index, value)
}

// DeleteWithValue deletes the first occurrence of a node with the given value.
func (cs *CSDLinkList[T]) DeleteWithValue(value T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.DeleteWithValue(value)
}

// Remove is an alias for DeleteWithValue.
func (cs *CSDLinkList[T]) Remove(value T) {
	cs.DeleteWithValue(value)
}

// RemoveAt deletes the node at the given index.
func (cs *CSDLinkList[T]) RemoveAt(index uint64) error {
	return cs.DeleteAt(index)
}

// Delete deletes the first node with the given value.
func (cs *CSDLinkList[T]) Delete(value T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.Delete(value)
}

// DeleteLast deletes the last node in the doubly linked list.
func (cs *CSDLinkList[T]) DeleteLast() {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.DeleteLast()
}

// DeleteFirst deletes the first node in the doubly linked list.
func (cs *CSDLinkList[T]) DeleteFirst() {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.DeleteFirst()
}

// DeleteAt deletes the node at the given index.
func (cs *CSDLinkList[T]) DeleteAt(index uint64) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.DeleteAt(index)
}

// ToSlice converts the doubly linked list to a slice.
func (cs *CSDLinkList[T]) ToSlice() []T {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.ToSlice()
}

// ToSliceReverse converts the doubly linked list to a slice in reverse order.
func (cs *CSDLinkList[T]) ToSliceReverse() []T {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.ToSliceReverse()
}

// ToSliceFromIndex converts the doubly linked list to a slice starting from the given index.
func (cs *CSDLinkList[T]) ToSliceFromIndex(index uint64) []T {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.ToSliceFromIndex(index)
}

// ToSliceReverseFromIndex converts the doubly linked list to a slice in reverse order starting from the given index.
func (cs *CSDLinkList[T]) ToSliceReverseFromIndex(index uint64) []T {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.ToSliceReverseFromIndex(index)
}

// Reverse reverses the doubly linked list.
func (cs *CSDLinkList[T]) Reverse() {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.Reverse()
}

// Find returns the first node with the given value.
func (cs *CSDLinkList[T]) Find(value T) (*dlinkList.Node[T], error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.Find(value)
}

// IsEmpty returns true if the doubly linked list is empty.
func (cs *CSDLinkList[T]) IsEmpty() bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.IsEmpty()
}

// GetAt returns the node at the given index.
func (cs *CSDLinkList[T]) GetAt(index uint64) (*dlinkList.Node[T], error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.GetAt(index)
}

// GetLast returns the last node in the doubly linked list.
func (cs *CSDLinkList[T]) GetLast() *dlinkList.Node[T] {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.GetLast()
}

// GetFirst returns the first node in the doubly linked list.
func (cs *CSDLinkList[T]) GetFirst() *dlinkList.Node[T] {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.GetFirst()
}

// Size returns the number of nodes in the doubly linked list.
func (cs *CSDLinkList[T]) Size() uint64 {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.Size()
}

// Clear removes all nodes from the doubly linked list.
func (cs *CSDLinkList[T]) Clear() {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.Clear()
}

// Contains returns true if the doubly linked list contains the given value.
func (cs *CSDLinkList[T]) Contains(value T) bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.Contains(value)
}

// ForEach traverses the doubly linked list and applies the given function to each node.
func (cs *CSDLinkList[T]) ForEach(f func(*T)) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.ForEach(f)
}

// ForFrom traverses the doubly linked list starting from the given index and applies the given function to each node.
func (cs *CSDLinkList[T]) ForFrom(index uint64, f func(*T)) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.ForFrom(index, f)
}

// ForReverseFrom traverses the doubly linked list in reverse order starting from the given index and applies the given function to each node.
func (cs *CSDLinkList[T]) ForReverseFrom(index uint64, f func(*T)) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.ForReverseFrom(index, f)
}

// ForEachReverse traverses the doubly linked list in reverse order and applies the given function to each node.
func (cs *CSDLinkList[T]) ForEachReverse(f func(*T)) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.ForEachReverse(f)
}

// ForRange traverses the doubly linked list in the given range and applies the given function to each node.
func (cs *CSDLinkList[T]) ForRange(start, end uint64, f func(*T)) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.ForRange(start, end, f)
}

// ForReverseRange traverses the doubly linked list in reverse order in the given range and applies the given function to each node.
func (cs *CSDLinkList[T]) ForReverseRange(start, end uint64, f func(*T)) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.ForReverseRange(start, end, f)
}

// Any returns true if the given function returns true for any node in the doubly linked list.
func (cs *CSDLinkList[T]) Any(f func(T) bool) bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.Any(f)
}

// All returns true if the given function returns true for all nodes in the doubly linked list.
func (cs *CSDLinkList[T]) All(f func(T) bool) bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.All(f)
}

// IndexOf returns the index of the first occurrence of the given value in the doubly linked list.
func (cs *CSDLinkList[T]) IndexOf(value T) int {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.IndexOf(value)
}

// LastIndexOf returns the index of the last occurrence of the given value in the doubly linked list.
func (cs *CSDLinkList[T]) LastIndexOf(value T) (uint64, error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.LastIndexOf(value)
}

// Filter returns a new doubly linked list containing only the nodes that satisfy the given function.
func (cs *CSDLinkList[T]) Filter(f func(T) bool) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.Filter(f)
}

// Map returns a new doubly linked list containing the result of applying the given function to each node.
func (cs *CSDLinkList[T]) Map(f func(T) T) *CSDLinkList[T] {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return &CSDLinkList[T]{l: cs.l.Map(f)}
}

// MapFrom returns a new doubly linked list containing the result of applying the given function to each node starting from the given index.
func (cs *CSDLinkList[T]) MapFrom(index uint64, f func(T) T) *CSDLinkList[T] {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return &CSDLinkList[T]{l: cs.l.MapFrom(index, f)}
}

// MapRange returns a new doubly linked list containing the result of applying the given function to each node in the given range.
func (cs *CSDLinkList[T]) MapRange(start, end uint64, f func(T) T) *CSDLinkList[T] {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return &CSDLinkList[T]{l: cs.l.MapRange(start, end, f)}
}

// Reduce reduces the doubly linked list to a single value using the given function.
func (cs *CSDLinkList[T]) Reduce(f func(T, T) T) T {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.Reduce(f)
}

// Copy returns a new doubly linked list with the same nodes as the original doubly linked list.
func (cs *CSDLinkList[T]) Copy() *CSDLinkList[T] {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return &CSDLinkList[T]{l: cs.l.Copy()}
}

// Merge appends the nodes of the given doubly linked list to the original doubly linked list.
func (cs *CSDLinkList[T]) Merge(list *CSDLinkList[T]) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	list.mu.Lock()
	defer list.mu.Unlock()
	cs.l.Merge(list.l)
}

// ReverseCopy returns a new doubly linked list with the nodes of the original doubly linked list in reverse order.
func (cs *CSDLinkList[T]) ReverseCopy() *CSDLinkList[T] {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return &CSDLinkList[T]{l: cs.l.ReverseCopy()}
}

// ReverseMerge appends the nodes of the given doubly linked list to the original doubly linked list in reverse order.
func (cs *CSDLinkList[T]) ReverseMerge(list *CSDLinkList[T]) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	list.mu.Lock()
	defer list.mu.Unlock()
	cs.l.ReverseMerge(list.l)
}

// Equal returns true if the given doubly linked list is equal to the original doubly linked list.
func (cs *CSDLinkList[T]) Equal(list *CSDLinkList[T]) bool {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	list.mu.RLock()
	defer list.mu.RUnlock()
	return cs.l.Equal(list.l)
}

// Swap swaps the nodes at the given indices.
func (cs *CSDLinkList[T]) Swap(i, j uint64) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.Swap(i, j)
}

// Sort sorts the doubly linked list according to the given function.
func (cs *CSDLinkList[T]) Sort(f func(T, T) bool) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.Sort(f)
}

// FindAll returns a new doubly linked list containing all nodes that satisfy the given function.
func (cs *CSDLinkList[T]) FindAll(f func(T) bool) *CSDLinkList[T] {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return &CSDLinkList[T]{l: cs.l.FindAll(f)}
}

// FindLast returns the last node that satisfies the given function.
func (cs *CSDLinkList[T]) FindLast(f func(T) bool) (*dlinkList.Node[T], error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.FindLast(f)
}

// FindLastIndex returns the index of the last node that satisfies the given function.
func (cs *CSDLinkList[T]) FindLastIndex(f func(T) bool) int {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.FindLastIndex(f)
}

// FindIndex returns the index of the first node that satisfies the given function.
func (cs *CSDLinkList[T]) FindIndex(f func(T) bool) int {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.l.FindIndex(f)
}
