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

// Package queue provides a non-concurrent-safe queue (FIFO).
package queue

import (
	"errors"
	"strings"
)

const (
	errQueueIsEmpty  = "queue is empty"
	errValueNotFound = "value not found"
)

// Queue is a FIFO data structure
type Queue[T comparable] struct {
	data []T
	size uint64
}

// NewQueue creates a new Queue
func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{}
}

// IsEmpty returns true if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// Enqueue adds an element to the end of the queue
func (q *Queue[T]) Enqueue(elem T) {
	q.data = append(q.data, elem)
	q.size++
}

// Dequeue removes and returns the first element in the queue
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var rVal T
		return rVal, errors.New(errQueueIsEmpty)
	}
	elem := q.data[0]
	q.data = q.data[1:]
	q.size--
	return elem, nil
}

// Peek returns the first element in the queue without removing it
func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var rVal T
		return rVal, errors.New(errQueueIsEmpty)
	}
	return q.data[0], nil
}

// Size returns the number of elements in the queue
func (q *Queue[T]) Size() uint64 {
	return q.size
}

// Clear removes all elements from the queue
func (q *Queue[T]) Clear() {
	q.data = []T{}
	q.size = 0
}

// Values returns all elements in the queue
func (q *Queue[T]) Values() []T {
	return q.data
}

// Contains returns true if the queue contains the given element
func (q *Queue[T]) Contains(elem T) bool {
	if q.size == 0 {
		return false
	}

	for i := uint64(0); i < q.size; i++ {
		if q.data[i] == elem {
			return true
		}
	}
	return false
}

// Equals returns true if the queue is equal to another queue
func (q *Queue[T]) Equals(other *Queue[T]) bool {
	if q.Size() != other.Size() {
		return false
	}

	for i := uint64(0); i < q.size; i++ {
		if q.data[i] != other.data[i] {
			return false
		}
	}
	return true
}

// Copy returns a copy of the queue
func (q *Queue[T]) Copy() *Queue[T] {
	copy := NewQueue[T]()
	copy.data = append(copy.data, q.data...)
	copy.size = q.size
	return copy
}

// String returns a string representation of the queue
func (q *Queue[T]) String(f func(T) string) string {
	return q.dataString(f)
}

func (q *Queue[T]) dataString(f func(T) string) string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, e := range q.data {
		sb.WriteString(f(e))
		if i < len(q.data)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

// Map creates a new queue with the results of applying the function to all elements in the queue
func (q *Queue[T]) Map(f func(T) T) *Queue[T] {
	newQueue := NewQueue[T]()

	if q.size == 0 {
		return newQueue
	}

	for i := uint64(0); i < q.size; i++ {
		newQueue.Enqueue(f(q.data[i]))
	}
	return newQueue
}

// Filter removes elements from the queue that don't match the predicate
func (q *Queue[T]) Filter(f func(T) bool) {
	var newData []T
	var size uint64
	for i := uint64(0); i < q.size; i++ {
		if f(q.data[i]) {
			newData = append(newData, q.data[i])
			size++
		}
	}
	q.data = newData
	q.size = size
}

// Reduce reduces the queue to a single value
func (q *Queue[T]) Reduce(f func(T, T) T, initial T) T {
	result := initial
	for i := uint64(0); i < q.size; i++ {
		result = f(result, q.data[i])
	}
	return result
}

// ForEach applies the function to all the elements in the queue
func (q *Queue[T]) ForEach(f func(*T)) {
	if q.size == 0 {
		return
	}
	for i := uint64(0); i < q.size; i++ {
		f(&q.data[i])
	}
}

// Any checks if any element in the queue matches the predicate
func (q *Queue[T]) Any(f func(T) bool) bool {
	if q.size == 0 {
		return false
	}
	for i := uint64(0); i < q.size; i++ {
		if f(q.data[i]) {
			return true
		}
	}
	return false
}

// All checks if all elements in the queue match the predicate
func (q *Queue[T]) All(f func(T) bool) bool {
	if q.size == 0 {
		return false
	}
	for i := uint64(0); i < q.size; i++ {
		if !f(q.data[i]) {
			return false
		}
	}
	return true
}

// IndexOf returns the index of the first element with the given value
func (q *Queue[T]) IndexOf(value T) (uint64, error) {
	if q.size == 0 {
		return 0, errors.New(errQueueIsEmpty)
	}

	for i := uint64(0); i < q.size; i++ {
		if q.data[i] == value {
			return i, nil
		}
	}
	return 0, errors.New(errValueNotFound)
}

// LastIndexOf returns the index of the last element with the given value
func (q *Queue[T]) LastIndexOf(value T) (uint64, error) {
	if q.size == 0 {
		return 0, errors.New(errQueueIsEmpty)
	}

	index := uint64(0)
	found := false
	for i := uint64(0); i < q.size; i++ {
		if q.data[i] == value {
			index = i
			found = true
		}
	}
	if !found {
		return 0, errors.New(errValueNotFound)
	}
	return index, nil
}

// FindIndex returns the index of the first element that matches the predicate
func (q *Queue[T]) FindIndex(f func(T) bool) (uint64, error) {
	if q.size == 0 {
		return 0, errors.New(errQueueIsEmpty)
	}

	for i := uint64(0); i < q.size; i++ {
		if f(q.data[i]) {
			return i, nil
		}
	}
	return 0, errors.New(errValueNotFound)
}

// FindLastIndex returns the index of the last element that matches the predicate
func (q *Queue[T]) FindLastIndex(f func(T) bool) (uint64, error) {
	if q.size == 0 {
		return 0, errors.New(errQueueIsEmpty)
	}

	index := uint64(0)
	found := false
	for i := uint64(0); i < q.size; i++ {
		if f(q.data[i]) {
			index = i
			found = true
		}
	}
	if !found {
		return 0, errors.New(errValueNotFound)
	}
	return index, nil
}

// FindAll returns all elements that match the predicate
func (q *Queue[T]) FindAll(f func(T) bool) *Queue[T] {
	newQueue := NewQueue[T]()
	for i := uint64(0); i < q.size; i++ {
		if f(q.data[i]) {
			newQueue.Enqueue(q.data[i])
		}
	}
	return newQueue
}

// FindLast returns the last element that matches the predicate
func (q *Queue[T]) FindLast(f func(T) bool) (T, error) {
	var result T
	if q.size == 0 {
		return result, errors.New(errQueueIsEmpty)
	}
	found := false
	for i := uint64(0); i < q.size; i++ {
		if f(q.data[i]) {
			result = q.data[i]
			found = true
		}
	}
	if !found {
		return result, errors.New(errValueNotFound)
	}
	return result, nil
}

// FindAllIndexes returns the indexes of all elements that match the predicate
func (q *Queue[T]) FindAllIndexes(f func(T) bool) []uint64 {
	var result []uint64
	for i := uint64(0); i < q.size; i++ {
		if f(q.data[i]) {
			result = append(result, i)
		}
	}
	return result
}
