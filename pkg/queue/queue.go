// Copyright 2023 Paolo Fabio Zaino
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
)

// Queue is a FIFO data structure
type Queue[T comparable] struct {
	data []T
}

// New creates a new Queue
func New[T comparable]() *Queue[T] {
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
	copy := New[T]()
	for _, e := range q.data {
		copy.Enqueue(e)
	}
	return copy
}

// String returns a string representation of the queue
func (q *Queue[T]) String(f func(T) string) string {
	return q.dataString(f)
}

func (q *Queue[T]) dataString(f func(T) string) string {
	str := "["
	for i, e := range q.data {
		str += f(e)
		if i < len(q.data)-1 {
			str += ", "
		}
	}
	str += "]"
	return str
}

// StringWithFormat returns a string representation of the queue using a custom format
func (q *Queue[T]) StringWithFormat(format func(T) string) string {
	return q.dataString(format)
}

// StringWithFormatAndSeparator returns a string representation of the queue using a custom format and separator
func (q *Queue[T]) StringWithFormatAndSeparator(format func(T) string, separator string) string {
	return q.dataString(func(e T) string {
		return format(e) + separator
	})
}
