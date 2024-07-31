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

// Package pqueue provides a non-concurrent-safe , max-heap, priority queue.
package pqueue

import (
	"errors"
	"strings"
)

const (
	ErrQueueIsEmpty    = "queue is empty"
	ErrIndexOutOfBound = "index out of bound"
	ErrValueNotFound   = "value not found"
)

// Element represents an element in the priority queue with a value and a priority.
type Element[T comparable] struct {
	Value    T
	Priority int
}

// PriorityQueue is a priority queue data structure
type PriorityQueue[T comparable] struct {
	data []Element[T]
	size uint64
}

// Helper functions for heap operations

// upHeap moves the element at the given index up the heap to restore the heap property
func (pq *PriorityQueue[T]) upHeap(index uint64) {
	for index > 0 {
		parent := (index - 1) / 2
		if pq.data[index].Priority <= pq.data[parent].Priority {
			break
		}
		pq.data[index], pq.data[parent] = pq.data[parent], pq.data[index]
		index = parent
	}
}

// downHeap moves the element at the given index down the heap to restore the heap property
func (pq *PriorityQueue[T]) downHeap(index uint64) {
	element := pq.data[index]
	lastIndex := pq.size - 1
	for {
		left := 2*index + 1
		if left > lastIndex {
			break
		}
		right := left + 1
		child := left
		if right <= lastIndex && pq.data[right].Priority > pq.data[left].Priority {
			child = right
		}
		if element.Priority >= pq.data[child].Priority {
			break
		}
		pq.data[index] = pq.data[child]
		index = child
	}
	pq.data[index] = element
}

// New creates a new PriorityQueue
func New[T comparable]() *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

// IsEmpty returns true if the priority queue is empty
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.size == 0
}

// Enqueue adds an element to the priority queue
func (pq *PriorityQueue[T]) Enqueue(value T, priority int) {
	element := Element[T]{Value: value, Priority: priority}
	pq.data = append(pq.data, element)
	pq.size++
	pq.upHeap(pq.size - 1)
}

// Dequeue removes and returns the highest priority element in the queue
func (pq *PriorityQueue[T]) Dequeue() (T, error) {
	if pq.IsEmpty() {
		var rVal T
		return rVal, errors.New(ErrQueueIsEmpty)
	}

	element := pq.data[0]
	lastIndex := pq.size - 1

	// Move the last element to the root and reduce the size
	pq.data[0] = pq.data[lastIndex]
	pq.size--
	pq.data = pq.data[:pq.size]

	if pq.size > 0 {
		pq.downHeap(0)
	}

	return element.Value, nil
}

// DequeueAll removes and returns all elements in the priority queue
// The returned list should be ordered by priority
func (pq *PriorityQueue[T]) DequeueAll() ([]T, error) {
	values := make([]T, pq.size)
	for i := uint64(0); !pq.IsEmpty(); i++ {
		item, err := pq.Dequeue()
		if err != nil {
			return values, err
		}
		values[i] = item
	}
	pq.Clear()
	return values, nil
}

// DequeueN removes and returns the first n elements in the priority queue
// The returned list should be ordered by priority
func (pq *PriorityQueue[T]) DequeueN(n uint64) ([]T, error) {
	if pq.IsEmpty() {
		return nil, errors.New(ErrQueueIsEmpty)
	}
	if n > pq.size {
		return nil, errors.New(ErrIndexOutOfBound)
	}
	values := make([]T, n)
	for i := uint64(0); i < n; i++ {
		item, err := pq.Dequeue()
		if err != nil {
			return nil, err
		}
		values[i] = item
	}
	return values, nil
}

// UpdatePriority updates the priority of an element in the priority queue
func (pq *PriorityQueue[T]) UpdatePriority(value T, newPriority int) error {
	if pq.IsEmpty() {
		return errors.New(ErrQueueIsEmpty)
	}

	for i, e := range pq.data {
		if e.Value == value {
			pq.data[i].Priority = newPriority
			pq.upHeap(uint64(i))
			pq.downHeap(uint64(i))
			return nil
		}
	}
	return errors.New(ErrValueNotFound)
}

// UpdatePriorityAt updates the priority of an element at the given index
func (pq *PriorityQueue[T]) UpdatePriorityAt(index uint64, newPriority int) error {
	if pq.IsEmpty() {
		return errors.New(ErrQueueIsEmpty)
	}
	if index >= pq.size {
		return errors.New(ErrValueNotFound)
	}

	oldPriority := pq.data[index].Priority
	pq.data[index].Priority = newPriority

	if newPriority > oldPriority {
		pq.upHeap(index)
	} else {
		pq.downHeap(index)
	}

	return nil
}

// UpdateValue updates the value of an element in the priority queue
func (pq *PriorityQueue[T]) UpdateValue(value T, newValue T) error {
	if pq.IsEmpty() {
		return errors.New(ErrQueueIsEmpty)
	}

	for i, e := range pq.data {
		if e.Value == value {
			pq.data[i].Value = newValue
			return nil
		}
	}
	return errors.New(ErrValueNotFound)
}

// Peek returns the highest priority element in the queue without removing it
func (pq *PriorityQueue[T]) Peek() (T, error) {
	if pq.IsEmpty() {
		var rVal T
		return rVal, errors.New(ErrQueueIsEmpty)
	}
	return pq.data[0].Value, nil
}

// Size returns the number of elements in the priority queue
func (pq *PriorityQueue[T]) Size() uint64 {
	return pq.size
}

// CheckSize recalculate the size of the priority queue
func (pq *PriorityQueue[T]) CheckSize() {
	pq.size = uint64(len(pq.data))
}

// Clear removes all elements from the priority queue
func (pq *PriorityQueue[T]) Clear() {
	pq.data = []Element[T]{}
	pq.size = 0
}

// Values returns all elements in the priority queue (it does not remove them!)
func (pq *PriorityQueue[T]) Values() []T {
	values := make([]T, len(pq.data))
	for i, element := range pq.data {
		values[i] = element.Value
	}
	return values
}

// Contains returns true if the priority queue contains the given element
func (pq *PriorityQueue[T]) Contains(value T) bool {
	if pq.size == 0 {
		return false
	}

	for _, e := range pq.data {
		if e.Value == value {
			return true
		}
	}
	return false
}

// Equals returns true if the priority queue is equal to another priority queue
func (pq *PriorityQueue[T]) Equals(other *PriorityQueue[T]) bool {
	if pq.Size() != other.Size() {
		return false
	}
	for i, e := range pq.data {
		if e.Value != other.data[i].Value || e.Priority != other.data[i].Priority {
			return false
		}
	}
	return true
}

// Copy returns a copy of the priority queue
func (pq *PriorityQueue[T]) Copy() *PriorityQueue[T] {
	copy := New[T]()
	copy.data = append(copy.data, pq.data...)
	copy.size = pq.size
	return copy
}

// Merge merges two priority queues (it considers the priority)
func (pq *PriorityQueue[T]) Merge(other *PriorityQueue[T]) {
	// Merge the two slices considering the priority
	for _, e := range other.data {
		pq.Enqueue(e.Value, e.Priority)
	}
	// Clear the other queue
	other.Clear()
}

// String returns a string representation of the priority queue
func (pq *PriorityQueue[T]) String(f func(T) string) string {
	return pq.dataString(f)
}

func (pq *PriorityQueue[T]) dataString(f func(T) string) string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, e := range pq.data {
		sb.WriteString(f(e.Value))
		if i < len(pq.data)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

// Map creates a new priority queue with the results of applying the function to each element
func (pq *PriorityQueue[T]) Map(f func(T) T) *PriorityQueue[T] {
	newQueue := New[T]()
	for i := 0; i < len(pq.data); i++ {
		newQueue.Enqueue(f(pq.data[i].Value), pq.data[i].Priority)
	}
	return newQueue
}

// Filter removes elements from the priority queue that don't match the predicate
func (pq *PriorityQueue[T]) Filter(f func(T) bool) {
	var newData []Element[T]
	var size uint64
	for i := uint64(0); i < pq.size; i++ {
		if f(pq.data[i].Value) {
			newData = append(newData, pq.data[i])
			size++
		}
	}
	pq.data = newData
	pq.size = size
}

// Reduce reduces the priority queue to a single value
func (pq *PriorityQueue[T]) Reduce(f func(T, T) T, initial T) T {
	result := initial
	for i := uint64(0); i < pq.size; i++ {
		result = f(result, pq.data[i].Value)
	}
	return result
}

// ForEach applies the function to all the elements in the priority queue
func (pq *PriorityQueue[T]) ForEach(f func(*T)) {
	for i := uint64(0); i < pq.size; i++ {
		f(&pq.data[i].Value)
	}
}

// Any checks if any element in the priority queue matches the predicate
func (pq *PriorityQueue[T]) Any(f func(T) bool) bool {
	for i := uint64(0); i < pq.size; i++ {
		if f(pq.data[i].Value) {
			return true
		}
	}
	return false
}

// All checks if all elements in the priority queue match the predicate
func (pq *PriorityQueue[T]) All(f func(T) bool) bool {
	for i := uint64(0); i < pq.size; i++ {
		if !f(pq.data[i].Value) {
			return false
		}
	}
	return true
}

// IndexOf returns the index of the first element with the given value
func (pq *PriorityQueue[T]) IndexOf(value T) (uint64, error) {
	for i := uint64(0); i < pq.size; i++ {
		if pq.data[i].Value == value {
			return i, nil
		}
	}
	return 0, errors.New(ErrValueNotFound)
}

// LastIndexOf returns the index of the last element with the given value
func (pq *PriorityQueue[T]) LastIndexOf(value T) (uint64, error) {
	index := uint64(0)
	found := false
	for i := uint64(0); i < pq.size; i++ {
		if pq.data[i].Value == value {
			index = i
			found = true
		}
	}
	if !found {
		return 0, errors.New(ErrValueNotFound)
	}
	return index, nil
}

// FindIndex returns the index of the first element that matches the predicate
func (pq *PriorityQueue[T]) FindIndex(f func(T) bool) (uint64, error) {
	for i := uint64(0); i < pq.size; i++ {
		if f(pq.data[i].Value) {
			return i, nil
		}
	}
	return 0, errors.New(ErrValueNotFound)
}

// FindLastIndex returns the index of the last element that matches the predicate
func (pq *PriorityQueue[T]) FindLastIndex(f func(T) bool) (uint64, error) {
	index := uint64(0)
	found := false
	for i := uint64(0); i < pq.size; i++ {
		if f(pq.data[i].Value) {
			index = i
			found = true
		}
	}
	if !found {
		return 0, errors.New(ErrValueNotFound)
	}
	return index, nil
}

// FindAll returns all elements that match the predicate
func (pq *PriorityQueue[T]) FindAll(f func(T) bool) *PriorityQueue[T] {
	newQueue := New[T]()
	for i := uint64(0); i < pq.size; i++ {
		if f(pq.data[i].Value) {
			newQueue.Enqueue(pq.data[i].Value, pq.data[i].Priority)
		}
	}
	return newQueue
}

// FindLast returns the last element that matches the predicate
func (pq *PriorityQueue[T]) FindLast(f func(T) bool) (T, error) {
	var result T
	found := false
	for i := uint64(0); i < pq.size; i++ {
		if f(pq.data[i].Value) {
			result = pq.data[i].Value
			found = true
		}
	}
	if !found {
		return result, errors.New(ErrValueNotFound)
	}
	return result, nil
}

// FindAllIndexes returns the indexes of all elements that match the predicate
func (pq *PriorityQueue[T]) FindAllIndexes(f func(T) bool) []uint64 {
	var result []uint64
	for i := uint64(0); i < pq.size; i++ {
		if f(pq.data[i].Value) {
			result = append(result, i)
		}
	}
	return result
}
