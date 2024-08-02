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

// Package ringBuffer provides a high-performance circular buffer.
package ringBuffer

import (
	"errors"
)

const (
	ErrCircularBufferEmpty = "ring buffer is empty"
)

// CircularBuffer represents a circular buffer data structure.
type CircularBuffer[T comparable] struct {
	data     []T
	capacity uint64
	head     uint64
	tail     uint64
	size     uint64
}

// New creates a new CircularBuffer with a given capacity.
func New[T comparable](capacity uint64) *CircularBuffer[T] {
	// Capacity should be a power of two for optimal performance with bitwise operations
	return &CircularBuffer[T]{
		data:     make([]T, capacity),
		capacity: capacity,
		head:     0,
		tail:     0,
		size:     0,
	}
}

// Append adds a new element to the buffer, overwriting the oldest data if the buffer is full.
func (cb *CircularBuffer[T]) Append(value T) {
	cb.data[cb.tail] = value
	cb.tail = (cb.tail + 1) % cb.capacity
	if cb.size < cb.capacity {
		cb.size++
	} else {
		cb.head = (cb.head + 1) % cb.capacity // Advance head when full
	}
}

// Remove removes the oldest element from the buffer.
func (cb *CircularBuffer[T]) Remove() (T, error) {
	if cb.IsEmpty() {
		var zero T
		return zero, errors.New(ErrCircularBufferEmpty)
	}

	value := cb.data[cb.head]
	cb.head = (cb.head + 1) % cb.capacity
	cb.size--

	return value, nil
}

// Get returns the element at a given index in the buffer (0 being the oldest).
func (cb *CircularBuffer[T]) Get(index uint64) (T, error) {
	if index >= cb.size {
		var zero T
		return zero, errors.New(ErrCircularBufferEmpty)
	}
	pos := (cb.head + index) & (cb.capacity - 1)
	return cb.data[pos], nil
}

// Size returns the current number of elements in the buffer.
func (cb *CircularBuffer[T]) Size() uint64 {
	return cb.size
}

// Capacity returns the capacity of the buffer.
func (cb *CircularBuffer[T]) Capacity() uint64 {
	return cb.capacity
}

// IsEmpty checks if the buffer is empty.
func (cb *CircularBuffer[T]) IsEmpty() bool {
	if cb == nil {
		return true
	}
	return cb.size == 0
}

// IsFull checks if the buffer is full.
func (cb *CircularBuffer[T]) IsFull() bool {
	return cb.size == cb.capacity
}

// Clear resets the buffer, making it empty.
func (cb *CircularBuffer[T]) Clear() {
	cb.head = 0
	cb.tail = 0
	cb.size = 0
}

// ToSlice returns the buffer content as a slice (oldest to newest).
func (cb *CircularBuffer[T]) ToSlice() []T {
	result := make([]T, cb.size)
	for i := uint64(0); i < cb.size; i++ {
		result[i] = cb.data[(cb.head+i)%cb.capacity]
	}
	return result
}

// ForEach applies a function to all elements in the buffer from oldest to newest.
func (cb *CircularBuffer[T]) ForEach(f func(T)) {
	for i := uint64(0); i < cb.size; i++ {
		f(cb.data[(cb.head+i)%cb.capacity])
	}
}

// Contains checks if the buffer contains a given value.
func (cb *CircularBuffer[T]) Contains(value T) bool {
	for i := uint64(0); i < cb.size; i++ {
		if cb.data[(cb.head+i)%cb.capacity] == value {
			return true
		}
	}
	return false
}
