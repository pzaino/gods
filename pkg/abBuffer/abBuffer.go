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

// Package abBuffer provides a non-concurrent-safe A/B buffer.
package abBuffer

import (
	"errors"
	"runtime"
	"sync"
)

const (
	errBufferOverflow = "buffer overflow"
	errInvalidBuffer  = "invalid buffer"
	errBufferEmpty    = "buffer is empty"
)

// Buffer represents a double-buffered structure
type Buffer[T comparable] struct {
	A        []T
	B        []T
	active   *[]T
	size     uint64
	capacity uint64
}

// NewBuffer creates a new Buffer with a given capacity
func NewBuffer[T comparable](capacity uint64) *Buffer[T] {
	a := make([]T, 0, capacity)
	b := make([]T, 0, capacity)
	return &Buffer[T]{
		A:        a,
		B:        b,
		active:   &a,
		capacity: capacity,
	}
}

// Append adds a new element to the active buffer
func (b *Buffer[T]) Append(value T) error {
	if uint64(len(*b.active)) >= b.capacity {
		return errors.New(errBufferOverflow)
	}

	*b.active = append(*b.active, value)
	b.size++
	return nil
}

// Clear clears the active buffer
func (b *Buffer[T]) Clear() {
	*b.active = (*b.active)[:0]
	b.size = 0
}

// Swap swaps the active buffer with the inactive one
func (b *Buffer[T]) Swap() {
	if b.active == &b.A {
		b.active = &b.B
	} else {
		b.active = &b.A
	}
}

// GetActive returns the active buffer
func (b *Buffer[T]) GetActive() []T {
	return *b.active
}

// GetInactive returns the inactive buffer
func (b *Buffer[T]) GetInactive() []T {
	if b.active == &b.A {
		return b.B
	}
	return b.A
}

// Size returns the number of elements in the active buffer
func (b *Buffer[T]) Size() uint64 {
	return b.size
}

// Capacity returns the capacity of the buffer
func (b *Buffer[T]) Capacity() uint64 {
	return b.capacity
}

// IsEmpty checks if the active buffer is empty
func (b *Buffer[T]) IsEmpty() bool {
	return b.size == 0
}

// ToSlice returns the active buffer as a slice
func (b *Buffer[T]) ToSlice() []T {
	return append([]T(nil), (*b.active)...)
}

// Find returns the first index of the given value in the active buffer
func (b *Buffer[T]) Find(value T) (int, error) {
	for i, v := range *b.active {
		if v == value {
			return i, nil
		}
	}
	return -1, errors.New(errBufferEmpty)
}

// Remove removes the element at the given index in the active buffer
func (b *Buffer[T]) Remove(index int) error {
	if index < 0 || index >= len(*b.active) {
		return errors.New(errInvalidBuffer)
	}

	*b.active = append((*b.active)[:index], (*b.active)[index+1:]...)
	b.size--
	return nil
}

// InsertAt inserts a new element at the given index in the active buffer
func (b *Buffer[T]) InsertAt(index int, value T) error {
	if index < 0 || index > len(*b.active) {
		return errors.New(errInvalidBuffer)
	}

	if uint64(len(*b.active)) >= b.capacity {
		return errors.New(errBufferOverflow)
	}

	*b.active = append((*b.active)[:index], append([]T{value}, (*b.active)[index:]...)...)
	b.size++
	return nil
}

// ForEach applies the function to all elements in the active buffer
func (b *Buffer[T]) ForEach(f func(T)) {
	for _, v := range *b.active {
		f(v)
	}
}

// ForFrom applies the function to all elements in the active buffer starting from the given index
func (b *Buffer[T]) ForFrom(index int, f func(T)) error {
	if index < 0 || index >= len(*b.active) {
		return errors.New(errInvalidBuffer)
	}

	for i := index; i < len(*b.active); i++ {
		f((*b.active)[i])
	}
	return nil
}

// ForRange applies the function to all elements in the active buffer in the range [start, end)
func (b *Buffer[T]) ForRange(start, end int, f func(T)) error {
	if start < 0 || start >= len(*b.active) || end < 0 || end > len(*b.active) {
		return errors.New(errInvalidBuffer)
	}

	for i := start; i < end; i++ {
		f((*b.active)[i])
	}
	return nil
}

// Map generates a new buffer by applying the function to all elements in the active buffer
func (b *Buffer[T]) Map(f func(T) T) (*Buffer[T], error) {
	newBuffer := NewBuffer[T](b.capacity)
	for _, v := range *b.active {
		err := newBuffer.Append(f(v))
		if err != nil {
			return nil, err
		}
	}
	return newBuffer, nil
}

// MapFrom generates a new buffer by applying the function to all elements in the active buffer starting from the given index
func (b *Buffer[T]) MapFrom(index int, f func(T) T) (*Buffer[T], error) {
	if index < 0 || index >= len(*b.active) {
		return nil, errors.New(errInvalidBuffer)
	}

	newBuffer := NewBuffer[T](b.capacity)
	for i := index; i < len(*b.active); i++ {
		err := newBuffer.Append(f((*b.active)[i]))
		if err != nil {
			return nil, err
		}
	}
	return newBuffer, nil
}

// MapRange generates a new buffer by applying the function to all elements in the active buffer in the range [start, end]
func (b *Buffer[T]) MapRange(start, end int, f func(T) T) (*Buffer[T], error) {
	if start < 0 || start >= len(*b.active) || end < 0 || end > len(*b.active) {
		return nil, errors.New(errInvalidBuffer)
	}

	newBuffer := NewBuffer[T](b.capacity)
	for i := start; i < end; i++ {
		err := newBuffer.Append(f((*b.active)[i]))
		if err != nil {
			return nil, err
		}
	}
	return newBuffer, nil
}

// Filter filter the active buffer by removing elements that don't match the predicate
func (b *Buffer[T]) Filter(f func(T) bool) {
	newBuffer := make([]T, 0, b.capacity)
	var size uint64
	for _, v := range *b.active {
		if f(v) {
			newBuffer = append(newBuffer, v)
			size++
		}
	}
	*b.active = newBuffer
	b.size = size
}

// Reduce reduces the buffer to a single value using the given function and initial value
func (b *Buffer[T]) Reduce(f func(T, T) T, initial T) T {
	result := initial
	for _, v := range *b.active {
		result = f(result, v)
	}
	return result
}

// ReduceFrom reduces the buffer to a single value starting from the given index using the given function and initial value
func (b *Buffer[T]) ReduceFrom(index int, f func(T, T) T, initial T) (T, error) {
	if index < 0 || index >= len(*b.active) {
		return initial, errors.New(errInvalidBuffer)
	}

	result := initial
	for i := index; i < len(*b.active); i++ {
		result = f(result, (*b.active)[i])
	}
	return result, nil
}

// ReduceRange reduces the buffer to a single value in the range [start, end) using the given function and initial value
func (b *Buffer[T]) ReduceRange(start, end int, f func(T, T) T, initial T) (T, error) {
	if start < 0 || start >= len(*b.active) || end < 0 || end > len(*b.active) {
		return initial, errors.New(errInvalidBuffer)
	}

	result := initial
	for i := start; i < end; i++ {
		result = f(result, (*b.active)[i])
	}
	return result, nil
}

// Contains checks if the active buffer contains the given value
func (b *Buffer[T]) Contains(value T) bool {
	_, err := b.Find(value)
	return err == nil
}

// Any checks if any element in the active buffer matches the predicate
func (b *Buffer[T]) Any(f func(T) bool) bool {
	for _, v := range *b.active {
		if f(v) {
			return true
		}
	}
	return false
}

// All checks if all elements in the active buffer match the predicate
func (b *Buffer[T]) All(f func(T) bool) bool {
	for _, v := range *b.active {
		if !f(v) {
			return false
		}
	}
	return true
}

// IndexOf returns the index of the first element with the given value in the active buffer
func (b *Buffer[T]) IndexOf(value T) (int, error) {
	return b.Find(value)
}

// LastIndexOf returns the index of the last element with the given value in the active buffer
func (b *Buffer[T]) LastIndexOf(value T) (int, error) {
	for i := len(*b.active) - 1; i >= 0; i-- {
		if (*b.active)[i] == value {
			return i, nil
		}
	}
	return -1, errors.New(errBufferEmpty)
}

// Copy creates a new buffer with the same elements as the active buffer
func (b *Buffer[T]) Copy() *Buffer[T] {
	newBuffer := NewBuffer[T](b.capacity)
	newBuffer.A = append([]T(nil), b.A...)
	newBuffer.B = append([]T(nil), b.B...)
	newBuffer.size = b.size
	return newBuffer
}

// CopyInactive creates a new buffer with the same elements as the inactive buffer
func (b *Buffer[T]) CopyInactive() *Buffer[T] {
	newBuffer := NewBuffer[T](b.capacity)
	if b.active == &b.A {
		newBuffer.A = append([]T(nil), b.B...)
		newBuffer.B = append([]T(nil), b.A...)
	} else {
		newBuffer.A = append([]T(nil), b.A...)
		newBuffer.B = append([]T(nil), b.B...)
	}
	newBuffer.size = b.size
	return newBuffer
}

// Merge merges the active buffer with another buffer
func (b *Buffer[T]) Merge(other *Buffer[T]) error {
	if b.capacity < other.size+b.size {
		return errors.New(errBufferOverflow)
	}

	*b.active = append(*b.active, *other.active...)
	b.size += other.size
	return nil
}

// MergeInactive merges the inactive buffer with another buffer
func (b *Buffer[T]) MergeInactive(other *Buffer[T]) error {
	if b.capacity < other.size+b.size {
		return errors.New(errBufferOverflow)
	}

	if b.active == &b.A {
		b.B = append(b.B, *other.active...)
	} else {
		b.A = append(b.A, *other.active...)
	}
	b.size += other.size
	return nil
}

// Blitter overwrite the values of the active buffer with the values of the other buffer using the "blitting" function
func (b *Buffer[T]) Blitter(other *Buffer[T], f func(T, T) T) error {
	if b.capacity < other.size {
		return errors.New(errBufferOverflow)
	}

	// Parallelize the blitting process for large buffers
	const minParallelSize = 1024 // Minimum size to consider parallel execution
	if b.size >= minParallelSize {
		numCPU := runtime.NumCPU()
		var wg sync.WaitGroup
		chunkSize := (int(b.size) + numCPU - 1) / numCPU // Determine chunk size

		wg.Add(numCPU)
		for i := 0; i < numCPU; i++ {
			start := i * chunkSize
			end := start + chunkSize
			if end > int(b.size) {
				end = int(b.size)
			}

			go func(start, end int) {
				defer wg.Done()
				for j := start; j < end; j++ {
					(*b.active)[j] = f((*b.active)[j], (*other.active)[j])
				}
			}(start, end)
		}
		wg.Wait()
	} else {
		// Single-threaded blitting for small buffers
		for i := uint64(0); i < b.size; i++ {
			(*b.active)[i] = f((*b.active)[i], (*other.active)[i])
		}
	}

	return nil
}
