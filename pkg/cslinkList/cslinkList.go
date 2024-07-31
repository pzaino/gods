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

// Package cslinkList provides a concurrency-safe linked list using linkList package.
package cslinkList

import (
	"sync"

	linkList "github.com/pzaino/gods/pkg/linkList"
)

// CSLinkList is a concurrency-safe linked list.
type CSLinkList[T comparable] struct {
	mu sync.Mutex
	l  *linkList.LinkList[T]
}

// New creates a new concurrency-safe linked list.
func New[T comparable]() *CSLinkList[T] {
	return &CSLinkList[T]{l: linkList.New[T]()}
}

// NewFromSlice creates a new concurrency-safe linked list from a slice.
func NewFromSlice[T comparable](items []T) *CSLinkList[T] {
	cs := New[T]()
	cs.l = linkList.NewFromSlice(items)
	return cs
}

// Append adds a new node to the end of the list.
func (cs *CSLinkList[T]) Append(value T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.Append(value)
}

// Prepend adds a new node to the beginning of the list.
func (cs *CSLinkList[T]) Prepend(value T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.Prepend(value)
}

// DeleteWithValue deletes the first node with the given value.
func (cs *CSLinkList[T]) DeleteWithValue(value T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.DeleteWithValue(value)
}

// ToSlice returns the list as a slice.
func (cs *CSLinkList[T]) ToSlice() []T {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.ToSlice()
}

// IsEmpty checks if the list is empty.
func (cs *CSLinkList[T]) IsEmpty() bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.IsEmpty()
}

// Find returns the first node with the given value.
func (cs *CSLinkList[T]) Find(value T) (*linkList.Node[T], error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.Find(value)
}

// Reverse reverses the list.
func (cs *CSLinkList[T]) Reverse() {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.Reverse()
}

// Size returns the number of nodes in the list.
func (cs *CSLinkList[T]) Size() uint64 {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.Size()
}

// GetFirst returns the first node in the list.
func (cs *CSLinkList[T]) GetFirst() *linkList.Node[T] {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.GetFirst()
}

// GetLast returns the last node in the list.
func (cs *CSLinkList[T]) GetLast() *linkList.Node[T] {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.GetLast()
}

// GetAt returns the node at the given index.
func (cs *CSLinkList[T]) GetAt(index uint64) (*linkList.Node[T], error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.GetAt(index)
}

// InsertAt inserts a new node at the given index.
func (cs *CSLinkList[T]) InsertAt(index uint64, value T) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.InsertAt(index, value)
}

// DeleteAt deletes the node at the given index.
func (cs *CSLinkList[T]) DeleteAt(index uint64) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.DeleteAt(index)
}

// Remove is just an alias for DeleteWithValue.
func (cs *CSLinkList[T]) Remove(value T) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.Remove(value)
}

// Clear removes all nodes from the list.
func (cs *CSLinkList[T]) Clear() {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.Clear()
}

// Copy returns a copy of the list.
func (cs *CSLinkList[T]) Copy() *CSLinkList[T] {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return &CSLinkList[T]{l: cs.l.Copy()}
}

// Merge appends all the nodes from another list to the current list.
func (cs *CSLinkList[T]) Merge(list *CSLinkList[T]) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	list.mu.Lock()
	defer list.mu.Unlock()
	cs.l.Merge(list.l)
	list.l.Clear()
}

// Map generates a new list by applying the function to all the nodes in the list.
func (cs *CSLinkList[T]) Map(f func(T) T) *CSLinkList[T] {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	newList := cs.l.Map(f)
	newCSList := New[T]()
	newCSList.l = newList
	return newCSList
}

// MapFrom generates a new list by applying the function to all the nodes in the list starting from the specified index.
func (cs *CSLinkList[T]) MapFrom(start uint64, f func(T) T) (*CSLinkList[T], error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	newList, err := cs.l.MapFrom(start, f)
	if err != nil {
		return nil, err
	}

	newCSList := New[T]()
	newCSList.l = newList
	return newCSList, nil
}

// MapRange generates a new list by applying the function to all the nodes in the list in the range [start, end).
func (cs *CSLinkList[T]) MapRange(start, end uint64, f func(T) T) (*CSLinkList[T], error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	newList, err := cs.l.MapRange(start, end, f)
	if err != nil {
		return nil, err
	}

	newCSList := New[T]()
	newCSList.l = newList
	return newCSList, nil
}

// Filter removes nodes from the list that don't match the predicate.
func (cs *CSLinkList[T]) Filter(f func(T) bool) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.Filter(f)
}

// Reduce reduces the list to a single value.
func (cs *CSLinkList[T]) Reduce(f func(T, T) T, initial T) T {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.Reduce(f, initial)
}

// ForEach applies the function to all the nodes in the list.
func (cs *CSLinkList[T]) ForEach(f func(*T)) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.l.ForEach(f)
}

// ForRange applies the function to all the nodes in the list in the range [start, end).
func (cs *CSLinkList[T]) ForRange(start, end uint64, f func(*T)) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.ForRange(start, end, f)
}

// ForFrom applies the function to all the nodes in the list starting from the index.
func (cs *CSLinkList[T]) ForFrom(start uint64, f func(*T)) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.ForFrom(start, f)
}

// Any checks if any node in the list matches the predicate.
func (cs *CSLinkList[T]) Any(f func(T) bool) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.Any(f)
}

// All checks if all nodes in the list match the predicate.
func (cs *CSLinkList[T]) All(f func(T) bool) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.All(f)
}

// Contains checks if the list contains the given value.
func (cs *CSLinkList[T]) Contains(value T) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.Contains(value)
}

// IndexOf returns the index of the first node with the given value.
func (cs *CSLinkList[T]) IndexOf(value T) (uint64, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.IndexOf(value)
}

// LastIndexOf returns the index of the last node with the given value.
func (cs *CSLinkList[T]) LastIndexOf(value T) (uint64, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.LastIndexOf(value)
}

// FindIndex returns the index of the first node that matches the predicate.
func (cs *CSLinkList[T]) FindIndex(f func(T) bool) (uint64, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.FindIndex(f)
}

// FindLastIndex returns the index of the last node that matches the predicate.
func (cs *CSLinkList[T]) FindLastIndex(f func(T) bool) (uint64, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.FindLastIndex(f)
}

// FindAll returns all nodes that match the predicate.
func (cs *CSLinkList[T]) FindAll(f func(T) bool) *CSLinkList[T] {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return &CSLinkList[T]{l: cs.l.FindAll(f)}
}

// FindLast returns the last node that matches the predicate.
func (cs *CSLinkList[T]) FindLast(f func(T) bool) (*linkList.Node[T], error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.FindLast(f)
}

// FindAllIndexes returns the indexes of all nodes that match the predicate.
func (cs *CSLinkList[T]) FindAllIndexes(f func(T) bool) []uint64 {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.l.FindAllIndexes(f)
}
