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

// Package buffer provides a generic buffer data structure.
package buffer

import (
	"errors"
	"runtime"
	"sync"
)

const (
	ErrBufferOverflow = "buffer overflow"
	ErrInvalidBuffer  = "invalid buffer"
	ErrBufferEmpty    = "buffer is empty"
	ErrValueNotFound  = "value not found"
)

// Buffer represent the Buffer structure used in an ABBuffer
type Buffer[T comparable] struct {
	data     []T
	size     uint64
	capacity uint64
}

// NewBuffer creates a new Buffer
func NewBuffer[T comparable]() *Buffer[T] {
	return &Buffer[T]{}
}

// IsEmpty returns true if the buffer is empty
func (b *Buffer[T]) IsEmpty() bool {
	if b == nil {
		return true
	}
	return b.size == 0
}

// IsFull returns true if the buffer is full
func (b *Buffer[T]) IsFull() bool {
	if b.capacity == 0 {
		return false
	}
	return b.size == b.capacity
}

// Append adds an element to the end of the buffer
func (b *Buffer[T]) Append(elem T) error {
	if b.IsFull() {
		return errors.New(ErrBufferOverflow)
	}
	b.data = append(b.data, elem)
	b.size++
	return nil
}

// InsertAt adds an element at the given index
func (b *Buffer[T]) InsertAt(index uint64, elem T) error {
	if index > b.size {
		return errors.New(ErrBufferOverflow)
	}
	if b.IsFull() {
		return errors.New(ErrBufferOverflow)
	}

	// Insert the element at the given index
	b.data = append(b.data[:index], append([]T{elem}, b.data[index:]...)...)
	b.size++

	return nil
}

// Put replaces the element at the given index
func (b *Buffer[T]) Put(index uint64, elem T) error {
	if index >= b.size {
		return errors.New(ErrValueNotFound)
	}
	b.data[index] = elem
	return nil
}

// Get returns the element at the given index
func (b *Buffer[T]) Get(index uint64) (T, error) {
	if index >= b.size {
		var rVal T
		return rVal, errors.New(ErrValueNotFound)
	}
	return b.data[index], nil
}

// Set sets the element at the given index
func (b *Buffer[T]) Set(index uint64, elem T) error {
	if index >= b.size {
		return errors.New(ErrValueNotFound)
	}
	b.data[index] = elem
	return nil
}

// Remove removes the element at the given index
func (b *Buffer[T]) Remove(index uint64) error {
	if index >= b.size {
		return errors.New(ErrValueNotFound)
	}
	b.data = append(b.data[:index], b.data[index+1:]...)
	b.size--
	return nil
}

// Clear removes all elements from the buffer
func (b *Buffer[T]) Clear() {
	b.data = []T{}
	b.size = 0
}

// Values returns all elements in the buffer
func (b *Buffer[T]) Values() []T {
	if b == nil {
		return nil
	}
	return b.data
}

// Size returns the number of elements in the buffer
func (b *Buffer[T]) Size() uint64 {
	if b == nil {
		return 0
	}
	return b.size
}

// Capacity returns the capacity of the buffer
func (b *Buffer[T]) Capacity() uint64 {
	return b.capacity
}

// SetCapacity sets the capacity of the buffer
func (b *Buffer[T]) SetCapacity(capacity uint64) {
	b.capacity = capacity
}

// Equals returns true if the buffer is equal to another buffer
func (b *Buffer[T]) Equals(other *Buffer[T]) bool {
	if b.Size() != other.Size() {
		return false
	}
	for i := uint64(0); i < b.Size(); i++ {
		if b.data[i] != other.data[i] {
			return false
		}
	}
	return true
}

// ToSlice returns a slice of the buffer
func (b *Buffer[T]) ToSlice() []T {
	return b.data
}

// Reverse reverses the buffer
func (b *Buffer[T]) Reverse() {
	for i := uint64(0); i < b.size/2; i++ {
		j := b.size - i - 1
		b.data[i], b.data[j] = b.data[j], b.data[i]
	}
}

// Find returns the index of the first element with the given value
func (b *Buffer[T]) Find(value T) (uint64, error) {
	for i := uint64(0); i < b.size; i++ {
		if b.data[i] == value {
			return i, nil
		}
	}
	return 0, errors.New(ErrValueNotFound)
}

// Contains returns true if the buffer contains the given element
func (b *Buffer[T]) Contains(value T) bool {
	for i := uint64(0); i < b.size; i++ {
		if b.data[i] == value {
			return true
		}
	}
	return false
}

// Copy returns a new buffer with copied elements
func (b *Buffer[T]) Copy() *Buffer[T] {
	newBuffer := NewBuffer[T]()
	newBuffer.data = make([]T, b.size)
	copy(newBuffer.data, b.data)
	newBuffer.size = b.size
	newBuffer.capacity = b.capacity
	return newBuffer
}

// NewReference returns a new buffer with the same elements (aka elements are not copied)
func (b *Buffer[T]) NewReference() *Buffer[T] {
	newBuffer := NewBuffer[T]()
	newBuffer.data = append(newBuffer.data, b.data...)
	newBuffer.size = b.size
	newBuffer.capacity = b.capacity
	return newBuffer
}

// Merge appends all elements from another buffer
func (b *Buffer[T]) Merge(other *Buffer[T]) {
	b.data = append(b.data, other.data...)
	b.size += other.size

	// Clear the other buffer
	other.Clear()

}

// PopN removes and returns the last n elements
func (b *Buffer[T]) PopN(n uint64) ([]T, error) {
	if b.size < n {
		return nil, errors.New(ErrBufferEmpty)
	}
	start := b.size - n
	end := b.size
	values := b.data[start:end]
	b.data = b.data[:start]
	b.size -= n
	return values, nil
}

// PushN adds multiple elements to the end of the buffer
func (b *Buffer[T]) PushN(items ...T) error {
	if b.size+uint64(len(items)) > b.capacity {
		return errors.New(ErrBufferOverflow)
	}
	b.data = append(b.data, items...)
	b.size += uint64(len(items))
	return nil
}

// ShiftLeft shifts all elements to the left by n positions
func (b *Buffer[T]) ShiftLeft(n uint64) {
	if b.size == 0 || n == 0 {
		return
	}

	if n > b.size {
		n = b.size
	}
	// move the first n elements to the beginning of the buffer
	b.data = b.data[n:]

	// append n "zero" values to the end of the buffer
	var zero T
	for i := uint64(0); i < n; i++ {
		b.data = append(b.data, zero)
	}
}

// ShiftRight shifts all elements to the right by n positions
func (b *Buffer[T]) ShiftRight(n uint64) {
	if b.size == 0 || n == 0 {
		return
	}

	if n > b.size {
		n = b.size
	}

	// Shift elements to the right within the buffer
	copy(b.data[n:], b.data[:b.size-n])

	// Fill the leftmost n positions with zero values
	var zero T
	for i := uint64(0); i < n; i++ {
		b.data[i] = zero
	}
}

// RotateLeft rotates all elements to the left by n positions
func (b *Buffer[T]) RotateLeft(n uint64) {
	if b.size == 0 {
		return
	}

	if n > b.size {
		n = n % b.size
	}
	// move the first n elements to the end of the buffer
	b.data = append(b.data[n:], b.data[:n]...)
}

// RotateRight rotates all elements to the right by n positions
func (b *Buffer[T]) RotateRight(n uint64) {
	if b.size == 0 {
		return
	}

	if n > b.size {
		n = n % b.size
	}
	// move the last n elements to the beginning of the buffer
	b.data = append(b.data[b.size-n:], b.data[:b.size-n]...)
}

// Filter removes elements that don't match the predicate
func (b *Buffer[T]) Filter(predicate func(T) bool) {
	var newData []T
	for i := uint64(0); i < b.size; i++ {
		if predicate(b.data[i]) {
			newData = append(newData, b.data[i])
		}
	}
	b.data = newData
	b.size = uint64(len(newData))
}

// Map creates a new buffer with the results of applying the function to each element
func (b *Buffer[T]) Map(fn func(T) T) (*Buffer[T], error) {
	newBuffer := NewBuffer[T]()
	var i uint64
	for i := uint64(0); i < b.size; i++ {
		err := newBuffer.Append(fn(b.data[i]))
		if err != nil {
			return newBuffer, nil
		}
	}
	newBuffer.capacity = b.capacity
	newBuffer.size = i
	return newBuffer, nil
}

// MapFrom creates a new buffer with the results of applying the function to each element starting from the specified index
func (b *Buffer[T]) MapFrom(start uint64, fn func(T) T) (*Buffer[T], error) {
	if start >= b.size {
		return nil, errors.New(ErrInvalidBuffer)
	}
	newBuffer := NewBuffer[T]()
	var item uint64
	for i := start; i < b.size; i++ {
		err := newBuffer.Append(fn(b.data[i]))
		if err != nil {
			break
		}
		item++
	}
	newBuffer.capacity = b.capacity
	newBuffer.size = item
	return newBuffer, nil
}

// MapRange creates a new buffer with the results of applying the function to each element in the range [start, end]
func (b *Buffer[T]) MapRange(start, end uint64, fn func(T) T) (*Buffer[T], error) {
	if start >= b.size || end > b.size || start > end {
		return nil, errors.New(ErrInvalidBuffer)
	}
	newBuffer := NewBuffer[T]()
	var item uint64
	for i := start; i < end; i++ {
		err := newBuffer.Append(fn(b.data[i]))
		if err != nil {
			break
		}
	}
	newBuffer.capacity = b.capacity
	newBuffer.size = item
	return newBuffer, nil
}

// Reduce reduces the buffer to a single value
func (b *Buffer[T]) Reduce(fn func(T, T) T) (T, error) {
	if b.size == 0 {
		var rVal T
		return rVal, errors.New(ErrBufferEmpty)
	}
	result := b.data[0]
	for i := uint64(1); i < b.size; i++ {
		result = fn(result, b.data[i])
	}
	return result, nil
}

// ReduceFrom reduces the buffer to a single value starting from the specified index
func (b *Buffer[T]) ReduceFrom(start uint64, fn func(T, T) T) (T, error) {
	if start >= b.size {
		var rVal T
		return rVal, errors.New(ErrInvalidBuffer)
	}
	result := b.data[start]
	for i := start + 1; i < b.size; i++ {
		result = fn(result, b.data[i])
	}
	return result, nil
}

// ReduceRange reduces the buffer to a single value in the range [start, end)
func (b *Buffer[T]) ReduceRange(start, end uint64, fn func(T, T) T) (T, error) {
	if start >= b.size || end > b.size || start > end {
		var rVal T
		return rVal, errors.New(ErrInvalidBuffer)
	}
	result := b.data[start]
	for i := start + 1; i < end; i++ {
		result = fn(result, b.data[i])
	}
	return result, nil
}

// ForEach applies the function to each element in the buffer
func (b *Buffer[T]) ForEach(fn func(*T)) {
	for i := uint64(0); i < b.size; i++ {
		fn(&b.data[i])
	}
}

// ForRange applies the function to each element in the buffer in the range [start, end)
func (b *Buffer[T]) ForRange(start, end uint64, fn func(*T)) error {
	if start >= b.size || end > b.size || start > end {
		return errors.New(ErrInvalidBuffer)
	}
	for i := start; i < end; i++ {
		fn(&b.data[i])
	}
	return nil
}

// ForFrom applies the function to each element in the buffer starting from the index
func (b *Buffer[T]) ForFrom(start uint64, fn func(*T)) error {
	if start >= b.size {
		return errors.New(ErrInvalidBuffer)
	}
	for i := start; i < b.size; i++ {
		fn(&b.data[i])
	}
	return nil
}

// Any checks if any element in the buffer matches the predicate
func (b *Buffer[T]) Any(predicate func(T) bool) bool {
	for i := uint64(0); i < b.size; i++ {
		if predicate(b.data[i]) {
			return true
		}
	}
	return false
}

// All checks if all elements in the buffer match the predicate
func (b *Buffer[T]) All(predicate func(T) bool) bool {
	for i := uint64(0); i < b.size; i++ {
		if !predicate(b.data[i]) {
			return false
		}
	}
	return true
}

// FindIndex returns the index of the first element that matches the predicate
func (b *Buffer[T]) FindIndex(predicate func(T) bool) (uint64, error) {
	for i := uint64(0); i < b.size; i++ {
		if predicate(b.data[i]) {
			return i, nil
		}
	}
	return 0, errors.New(ErrValueNotFound)
}

// FindLast returns the last element that matches the predicate
func (b *Buffer[T]) FindLast(predicate func(T) bool) (*T, error) {
	for i := b.size - 1; i > 0; i-- {
		if predicate(b.data[i]) {
			return &b.data[i], nil
		}
	}
	if predicate(b.data[0]) {
		return &b.data[0], nil
	}
	return nil, errors.New(ErrValueNotFound)
}

// FindLastIndex returns the index of the last element that matches the predicate
func (b *Buffer[T]) FindLastIndex(predicate func(T) bool) (uint64, error) {
	for i := b.size - 1; i > 0; i-- {
		if predicate(b.data[i]) {
			return i, nil
		}
	}
	if predicate(b.data[0]) {
		return 0, nil
	}
	return 0, errors.New(ErrValueNotFound)
}

// FindAll returns all elements that match the predicate
func (b *Buffer[T]) FindAll(predicate func(T) bool) *Buffer[T] {
	newBuffer := NewBuffer[T]()
	var i uint64
	for i := uint64(0); i < b.size; i++ {
		if predicate(b.data[i]) {
			err := newBuffer.Append(b.data[i])
			if err != nil {
				break
			}
		}
	}
	newBuffer.capacity = b.capacity
	newBuffer.size = i
	return newBuffer
}

// FindIndices returns the indices of all elements that match the predicate
func (b *Buffer[T]) FindIndices(predicate func(T) bool) []uint64 {
	var indices []uint64
	for i := uint64(0); i < b.size; i++ {
		if predicate(b.data[i]) {
			indices = append(indices, i)
		}
	}
	return indices
}

// LastIndexOf returns the index of the last element with the given value
func (b *Buffer[T]) LastIndexOf(value T) (uint64, error) {
	for i := b.size - 1; i > 0; i-- {
		if b.data[i] == value {
			return i, nil
		}
	}
	if b.data[0] == value {
		return 0, nil
	}
	return 0, errors.New(ErrValueNotFound)
}

// Blit combine/overwrite the values of the in the buffer with the values of another buffer using a function
func (b *Buffer[T]) Blit(other *Buffer[T], f func(T, T) T) error {
	// If both buffers are empty, return nil
	if b.size == 0 && other.size == 0 {
		return nil
	}

	var maxElements uint64
	if other.size < b.size {
		maxElements = other.size
	} else {
		maxElements = b.size
	}

	// Parallelize the blitting process for large buffers
	const minParallelSize = 1024 // Minimum size to consider parallel execution
	if b.size >= minParallelSize {
		numCPU := runtime.NumCPU()
		var wg sync.WaitGroup
		chunkSize := (int(maxElements) + numCPU - 1) / numCPU // Determine chunk size

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
					(*b).data[j] = f((*b).data[j], (*other).data[j])
				}
			}(start, end)
		}
		wg.Wait()
	} else {
		// Single-threaded blitting for small buffers
		for i := uint64(0); i < maxElements; i++ {
			(*b).data[i] = f((*b).data[i], (*other).data[i])
		}
	}

	return nil
}

// BlitFrom combine/overwrite the values of the in the buffer with the values of another buffer starting from the specified index using a function
func (b *Buffer[T]) BlitFrom(start uint64, other *Buffer[T], f func(T, T) T) error {
	if start >= b.size || start >= other.size {
		return errors.New(ErrInvalidBuffer)
	}

	var maxElements uint64
	if other.size-start < b.size-start {
		maxElements = other.size - start
	} else {
		maxElements = b.size - start
	}

	// Parallelize the blitting process for large buffers
	const minParallelSize = 1024 // Minimum size to consider parallel execution
	if maxElements >= minParallelSize {
		numCPU := runtime.NumCPU()
		var wg sync.WaitGroup
		chunkSize := (int(maxElements) + numCPU - 1) / numCPU // Determine chunk size

		wg.Add(numCPU)
		for i := 0; i < numCPU; i++ {
			start := int(start) + i*chunkSize
			end := start + chunkSize
			if end > int(b.size) {
				end = int(b.size)
			}

			go func(start, end int) {
				defer wg.Done()
				for j := start; j < end; j++ {
					(*b).data[j] = f((*b).data[j], (*other).data[j])
				}
			}(start, end)
		}
		wg.Wait()
	} else {
		// Single-threaded blitting for small buffers
		for i := uint64(0); i < maxElements; i++ {
			(*b).data[start+i] = f((*b).data[start+i], (*other).data[start+i])
		}
	}

	return nil
}

// BlitRange combine/overwrite the values of the in the buffer with the values of another buffer in the range [start, end] using a function
func (b *Buffer[T]) BlitRange(start, end uint64, other *Buffer[T], f func(T, T) T) error {
	if start >= b.size || end > b.size || start >= end || start >= other.size || end > other.size {
		return errors.New(ErrInvalidBuffer)
	}

	var maxElements uint64
	if end-start < other.size-start {
		maxElements = end - start
	} else {
		maxElements = other.size - start
	}

	// Parallelize the blitting process for large buffers
	const minParallelSize = 1024 // Minimum size to consider parallel execution
	if maxElements >= minParallelSize {
		numCPU := runtime.NumCPU()
		var wg sync.WaitGroup
		chunkSize := (int(maxElements) + numCPU - 1) / numCPU // Determine chunk size

		wg.Add(numCPU)
		for i := 0; i < numCPU; i++ {
			start := int(start) + i*chunkSize
			end := start + chunkSize
			if end > int(b.size) {
				end = int(b.size)
			}

			go func(start, end int) {
				defer wg.Done()
				for j := start; j < end; j++ {
					(*b).data[j] = f((*b).data[j], (*other).data[j])
				}
			}(start, end)
		}
		wg.Wait()
	} else {
		// Single-threaded blitting for small buffers
		for i := uint64(0); i < maxElements; i++ {
			(*b).data[start+i] = f((*b).data[start+i], (*other).data[start+i])
		}
	}

	return nil
}