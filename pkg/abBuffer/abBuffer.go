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

// Important notes on thsi A/B buffer implementation:
// - The A/B buffer is a double-buffered structure that allows for efficient swapping of buffers.
// - The A/B buffer can be used to store and manipulate data in a buffer-like structure.
// - The "active" buffer in the A/B buffer is the buffer that is currently being used for operations.
// - The "inactive" buffer in the A/B buffer is the buffer that is not currently being used for
//   operations, and therefore can be read safely or passed to other functions.

import (
	"errors"

	"github.com/pzaino/gods/pkg/buffer"
)

const (
	errBufferOverflow = "buffer overflow"
	errInvalidBuffer  = "invalid buffer"
	errBufferEmpty    = "buffer is empty"
	errValueNotFound  = "value not found"
)

// ABBuffer represents a double-buffered structure
type ABBuffer[T comparable] struct {
	A        buffer.Buffer[T]
	B        buffer.Buffer[T]
	active   *buffer.Buffer[T]
	capacity uint64
}

// New creates a new Buffer with a given capacity
func New[T comparable](capacity uint64) *ABBuffer[T] {
	a := buffer.Buffer[T]{}
	b := buffer.Buffer[T]{}
	ab := &ABBuffer[T]{
		A:        a,
		B:        b,
		capacity: capacity,
	}
	ab.active = &ab.A
	return ab
}

// Append adds a new element to the active buffer
func (b *ABBuffer[T]) Append(value T) error {
	if (b.active.Size() >= b.capacity) && (b.capacity != 0) {
		return errors.New(errBufferOverflow)
	}
	err := b.active.Append(value)
	return err
}

// Clear clears the active buffer
func (b *ABBuffer[T]) Clear() {
	b.active.Clear()
}

// ClearAll clears both the active and inactive buffers
func (b *ABBuffer[T]) ClearAll() {
	b.A.Clear()
	b.B.Clear()
	b.active = &b.A
}

// Destroy clears both the active and inactive buffers and sets the active buffer to nil
func (b *ABBuffer[T]) Destroy() {
	b.A.Clear()
	b.B.Clear()
	b.active = nil
	b.capacity = 0
	b = nil
}

// Swap swaps the active buffer with the inactive one
func (b *ABBuffer[T]) Swap() {
	if b.active == &b.A {
		b.active = &b.B
	} else {
		b.active = &b.A
	}
}

// SetActiveA sets the active buffer to A
func (b *ABBuffer[T]) SetActiveA() {
	b.active = &b.A
}

// SetActiveB sets the active buffer to B
func (b *ABBuffer[T]) SetActiveB() {
	b.active = &b.B
}

// GetActive returns the active buffer
func (b *ABBuffer[T]) GetActive() []T {
	return b.active.Values()
}

// GetInactive returns the inactive buffer
func (b *ABBuffer[T]) GetInactive() []T {
	if b == nil {
		return nil
	}
	if b.active == &b.A {
		return b.B.Values()
	}
	return b.A.Values()
}

// Size returns the number of elements in the active buffer
func (b *ABBuffer[T]) Size() uint64 {
	return b.active.Size()
}

// Capacity returns the capacity of the buffer
func (b *ABBuffer[T]) Capacity() uint64 {
	return b.capacity
}

// IsEmpty checks if the active buffer is empty
func (b *ABBuffer[T]) IsEmpty() bool {
	return b.active.IsEmpty()
}

// ToSlice returns the active buffer as a slice
func (b *ABBuffer[T]) ToSlice() []T {
	return b.active.ToSlice()
}

// ToSliceInactive returns the inactive buffer as a slice
func (b *ABBuffer[T]) ToSliceInactive() []T {
	return b.GetInactive()
}

// FetchInactive returns the inactive buffer and clears it in the A/B buffer
func (b *ABBuffer[T]) FetchInactive() []T {
	var inactive *buffer.Buffer[T]
	if b.active == &b.A {
		inactive = &b.B
	} else {
		inactive = &b.A
	}
	data := inactive.ToSlice()
	inactive.Clear()
	return data
}

// Find returns the first index of the given value in the active buffer
func (b *ABBuffer[T]) Find(value T) (uint64, error) {
	return b.active.Find(value)
}

// Remove removes the element at the given index in the active buffer
func (b *ABBuffer[T]) Remove(index uint64) error {
	return b.active.Remove(index)
}

// InsertAt inserts a new element at the given index in the active buffer
func (b *ABBuffer[T]) InsertAt(index uint64, value T) error {
	return b.active.InsertAt(index, value)
}

// ForEach applies the function to all elements in the active buffer
func (b *ABBuffer[T]) ForEach(f func(*T)) {
	b.active.ForEach(f)
}

// ForFrom applies the function to all elements in the active buffer starting from the given index
func (b *ABBuffer[T]) ForFrom(index uint64, f func(*T)) error {
	return b.active.ForFrom(index, f)
}

// ForRange applies the function to all elements in the active buffer in the range [start, end)
func (b *ABBuffer[T]) ForRange(start, end uint64, f func(*T)) error {
	return b.active.ForRange(start, end, f)
}

// Map generates a new buffer by applying the function to all elements in the active buffer
func (b *ABBuffer[T]) Map(f func(T) T) (*ABBuffer[T], error) {
	newBuffer := New[T](b.capacity)
	nb, err := b.active.Map(f)
	if err != nil {
		return nil, err
	}
	newBuffer.A = *nb
	newBuffer.active = &newBuffer.A
	return newBuffer, nil
}

// MapFrom generates a new buffer by applying the function to all elements in the active buffer starting from the given index
func (b *ABBuffer[T]) MapFrom(index uint64, f func(T) T) (*ABBuffer[T], error) {
	if index >= b.active.Size() {
		return nil, errors.New(errInvalidBuffer)
	}

	newBuffer := New[T](b.capacity)
	nb, err := b.active.MapFrom(index, f)
	if err != nil {
		return nil, err
	}
	newBuffer.A = *nb
	newBuffer.active = &newBuffer.A
	return newBuffer, nil
}

// MapRange generates a new buffer by applying the function to all elements in the active buffer in the range [start, end]
func (b *ABBuffer[T]) MapRange(start, end uint64, f func(T) T) (*ABBuffer[T], error) {
	if start >= b.active.Size() || end > b.active.Size() {
		return nil, errors.New(errInvalidBuffer)
	}

	newBuffer := New[T](b.capacity)
	nb, err := b.active.MapRange(start, end, f)
	if err != nil {
		return nil, err
	}
	newBuffer.A = *nb
	newBuffer.active = &newBuffer.A
	return newBuffer, nil
}

// Filter filter the active buffer by removing elements that don't match the predicate
func (b *ABBuffer[T]) Filter(f func(T) bool) {
	b.active.Filter(f)
}

// Reduce reduces the buffer to a single value using the given function and initial value
func (b *ABBuffer[T]) Reduce(f func(T, T) T) (T, error) {
	return b.active.Reduce(f)
}

// ReduceFrom reduces the buffer to a single value starting from the given index using the given function and initial value
func (b *ABBuffer[T]) ReduceFrom(index uint64, f func(T, T) T) (T, error) {
	return b.active.ReduceFrom(index, f)
}

// ReduceRange reduces the buffer to a single value in the range [start, end) using the given function and initial value
func (b *ABBuffer[T]) ReduceRange(start, end uint64, f func(T, T) T) (T, error) {
	return b.active.ReduceRange(start, end, f)
}

// Contains checks if the active buffer contains the given value
func (b *ABBuffer[T]) Contains(value T) bool {
	return b.active.Contains(value)
}

// Any checks if any element in the active buffer matches the predicate
func (b *ABBuffer[T]) Any(f func(T) bool) bool {
	return b.active.Any(f)
}

// All checks if all elements in the active buffer match the predicate
func (b *ABBuffer[T]) All(f func(T) bool) bool {
	return b.active.All(f)
}

// LastIndexOf returns the index of the last element with the given value in the active buffer
func (b *ABBuffer[T]) LastIndexOf(value T) (uint64, error) {
	return b.active.LastIndexOf(value)
}

// Copy creates a new A/B buffer with the same elements as the A/B buffer
// this method does a deep copy of the entire A/B buffer
func (b *ABBuffer[T]) Copy() *ABBuffer[T] {
	newBuffer := New[T](b.capacity)
	newBuffer.A = *b.A.Copy()
	newBuffer.B = *b.B.Copy()
	return newBuffer
}

// CopyActive creates a new buffer with the same elements as the active buffer
// The copied buffer is placed in the A buffer on the new A/B Buffer and A
// buffer is set as the active buffer
func (b *ABBuffer[T]) CopyActive() *ABBuffer[T] {
	newBuffer := New[T](b.capacity)
	if b.active == &b.A {
		newBuffer.A = *b.A.Copy()
	} else {
		newBuffer.A = *b.B.Copy()
	}
	newBuffer.capacity = b.capacity
	newBuffer.active = &newBuffer.A
	return newBuffer
}

// CopyInactive creates a new buffer with the same elements as the inactive buffer
// The copied buffer is placed in the A buffer on the new A/B Buffer and A
// buffer is set as the active buffer
func (b *ABBuffer[T]) CopyInactive() *ABBuffer[T] {
	newBuffer := New[T](b.capacity)
	if b.active == &b.A {
		newBuffer.A = *b.B.Copy()
	} else {
		newBuffer.A = *b.active.Copy()
	}
	newBuffer.capacity = b.capacity
	newBuffer.active = &newBuffer.A
	return newBuffer
}

// Merge merges the active buffer with the active buffer from another A/B buffer
func (b *ABBuffer[T]) Merge(other *ABBuffer[T]) {
	b.active.Merge(other.active)
}

// Blit overwrite the values of the active buffer with the values of the other buffer using the "blitting" function
func (b *ABBuffer[T]) Blit(other *ABBuffer[T], f func(T, T) T) error {
	return b.active.Blit(other.active, f)
}
