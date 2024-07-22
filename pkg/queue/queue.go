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

// Queue is a FIFO data structure
type Queue[T comparable] struct {
	data []T
}

// New creates a new Queue
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
}

// Dequeue removes and returns the first element in the queue
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var rVal T
		return rVal, errors.New("queue is empty")
	}
	elem := q.data[0]
	q.data = q.data[1:]
	return elem, nil
}

// Peek returns the first element in the queue without removing it
func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var rVal T
		return rVal, errors.New("queue is empty")
	}
	return q.data[0], nil
}

// Size returns the number of elements in the queue
func (q *Queue[T]) Size() int {
	return len(q.data)
}

// Clear removes all elements from the queue
func (q *Queue[T]) Clear() {
	q.data = []T{}
}

// Values returns all elements in the queue
func (q *Queue[T]) Values() []T {
	return q.data
}

// Contains returns true if the queue contains the given element
func (q *Queue[T]) Contains(elem T) bool {
	for _, e := range q.data {
		if e == elem {
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
	for i, e := range q.data {
		if e != other.data[i] {
			return false
		}
	}
	return true
}

// Copy returns a copy of the queue
func (q *Queue[T]) Copy() *Queue[T] {
	copy := NewQueue[T]()
	copy.data = append(copy.data, q.data...)
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

// Map applies the function to all the elements in the queue
func (q *Queue[T]) Map(f func(T) T) {
	for i, e := range q.data {
		q.data[i] = f(e)
	}
}

// Filter removes elements from the queue that don't match the predicate
func (q *Queue[T]) Filter(f func(T) bool) {
	var newData []T
	for _, e := range q.data {
		if f(e) {
			newData = append(newData, e)
		}
	}
	q.data = newData
}

// Reduce reduces the queue to a single value
func (q *Queue[T]) Reduce(f func(T, T) T, initial T) T {
	result := initial
	for _, e := range q.data {
		result = f(result, e)
	}
	return result
}

// ForEach applies the function to all the elements in the queue
func (q *Queue[T]) ForEach(f func(*T)) {
	for i := range q.data {
		f(&q.data[i])
	}
}

// Any checks if any element in the queue matches the predicate
func (q *Queue[T]) Any(f func(T) bool) bool {
	for _, e := range q.data {
		if f(e) {
			return true
		}
	}
	return false
}

// All checks if all elements in the queue match the predicate
func (q *Queue[T]) All(f func(T) bool) bool {
	for _, e := range q.data {
		if !f(e) {
			return false
		}
	}
	return true
}

// IndexOf returns the index of the first element with the given value
func (q *Queue[T]) IndexOf(value T) int {
	for i, e := range q.data {
		if e == value {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last element with the given value
func (q *Queue[T]) LastIndexOf(value T) int {
	index := -1
	for i, e := range q.data {
		if e == value {
			index = i
		}
	}
	return index
}

// FindIndex returns the index of the first element that matches the predicate
func (q *Queue[T]) FindIndex(f func(T) bool) int {
	for i, e := range q.data {
		if f(e) {
			return i
		}
	}
	return -1
}

// FindLastIndex returns the index of the last element that matches the predicate
func (q *Queue[T]) FindLastIndex(f func(T) bool) int {
	index := -1
	for i, e := range q.data {
		if f(e) {
			index = i
		}
	}
	return index
}

// FindAll returns all elements that match the predicate
func (q *Queue[T]) FindAll(f func(T) bool) *Queue[T] {
	newQueue := NewQueue[T]()
	for _, e := range q.data {
		if f(e) {
			newQueue.Enqueue(e)
		}
	}
	return newQueue
}

// FindLast returns the last element that matches the predicate
func (q *Queue[T]) FindLast(f func(T) bool) (T, error) {
	var result T
	found := false
	for _, e := range q.data {
		if f(e) {
			result = e
			found = true
		}
	}
	if !found {
		return result, errors.New("value not found")
	}
	return result, nil
}

// FindAllIndexes returns the indexes of all elements that match the predicate
func (q *Queue[T]) FindAllIndexes(f func(T) bool) []int {
	var result []int
	for i, e := range q.data {
		if f(e) {
			result = append(result, i)
		}
	}
	return result
}
