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

// Package csBuffer provides a thread-safe wrapper around the Buffer type.
package csBuffer

import (
	"sync"

	buffer "github.com/pzaino/gods/pkg/buffer"
)

// ConcurrentBuffer is a thread-safe wrapper around the Buffer type.
type ConcurrentBuffer[T comparable] struct {
	b  *buffer.Buffer[T]
	mu sync.RWMutex
}

// New creates a new ConcurrentBuffer.
func New[T comparable]() *ConcurrentBuffer[T] {
	return &ConcurrentBuffer[T]{b: buffer.New[T]()}
}

// NewWithCapacity creates a new ConcurrentBuffer with the given capacity.
func NewWithCapacity[T comparable](capacity uint64) *ConcurrentBuffer[T] {
	return &ConcurrentBuffer[T]{b: buffer.NewWithCapacity[T](capacity)}
}

// NewWithSize creates a new ConcurrentBuffer with the given size.
func NewWithSize[T comparable](size uint64) *ConcurrentBuffer[T] {
	return &ConcurrentBuffer[T]{b: buffer.NewWithSize[T](size)}
}

// NewWithSizeAndCapacity creates a new ConcurrentBuffer with the given size and capacity.
func NewWithSizeAndCapacity[T comparable](size, capacity uint64) *ConcurrentBuffer[T] {
	return &ConcurrentBuffer[T]{b: buffer.NewWithSizeAndCapacity[T](size, capacity)}
}

// Append adds an element to the end of the buffer.
func (cb *ConcurrentBuffer[T]) Append(elem T) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.b.Append(elem)
}

// InsertAt adds an element at the given index.
func (cb *ConcurrentBuffer[T]) InsertAt(index uint64, elem T) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.b.InsertAt(index, elem)
}

// Put replaces the element at the given index.
func (cb *ConcurrentBuffer[T]) Put(index uint64, elem T) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.b.Put(index, elem)
}

// Get returns the element at the given index.
func (cb *ConcurrentBuffer[T]) Get(index uint64) (T, error) {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.Get(index)
}

// Remove removes the element at the given index.
func (cb *ConcurrentBuffer[T]) Remove(index uint64) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.b.Remove(index)
}

// Clear removes all elements from the buffer.
func (cb *ConcurrentBuffer[T]) Clear() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.b.Clear()
}

// Destroy removes all elements from the buffer and sets the capacity to 0.
func (cb *ConcurrentBuffer[T]) Destroy() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.b.Destroy()
}

// Values returns all elements in the buffer.
func (cb *ConcurrentBuffer[T]) Values() []T {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.Values()
}

// Size returns the number of elements in the buffer.
func (cb *ConcurrentBuffer[T]) Size() uint64 {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.Size()
}

// Capacity returns the capacity of the buffer.
func (cb *ConcurrentBuffer[T]) Capacity() uint64 {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.Capacity()
}

// SetCapacity sets the capacity of the buffer.
func (cb *ConcurrentBuffer[T]) SetCapacity(capacity uint64) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.b.SetCapacity(capacity)
}

// Contains returns true if the buffer contains the given element.
func (cb *ConcurrentBuffer[T]) Contains(value T) bool {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.Contains(value)
}

// IsEmpty returns true if the buffer is empty.
func (cb *ConcurrentBuffer[T]) IsEmpty() bool {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.IsEmpty()
}

// IsFull returns true if the buffer is full.
func (cb *ConcurrentBuffer[T]) IsFull() bool {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.IsFull()
}

// Find returns the index of the first element with the given value.
func (cb *ConcurrentBuffer[T]) Find(value T) (uint64, error) {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.Find(value)
}

// Reverse reverses the buffer.
func (cb *ConcurrentBuffer[T]) Reverse() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.b.Reverse()
}

// Equals returns true if the buffer is equal to another buffer.
func (cb *ConcurrentBuffer[T]) Equals(other *ConcurrentBuffer[T]) bool {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	return cb.b.Equals(other.b)
}

// Copy returns a new buffer with copied elements.
func (cb *ConcurrentBuffer[T]) Copy() *ConcurrentBuffer[T] {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	newBuffer := cb.b.Copy()
	return &ConcurrentBuffer[T]{b: newBuffer}
}

// Merge appends all elements from another buffer.
func (cb *ConcurrentBuffer[T]) Merge(other *ConcurrentBuffer[T]) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	cb.b.Merge(other.b)
}

// PopN removes and returns the last n elements.
func (cb *ConcurrentBuffer[T]) PopN(n uint64) ([]T, error) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.b.PopN(n)
}

// PushN adds multiple elements to the end of the buffer.
func (cb *ConcurrentBuffer[T]) PushN(items ...T) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.b.PushN(items...)
}

// ShiftLeft shifts all elements to the left by n positions.
func (cb *ConcurrentBuffer[T]) ShiftLeft(n uint64) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.b.ShiftLeft(n)
}

// ShiftRight shifts all elements to the right by n positions.
func (cb *ConcurrentBuffer[T]) ShiftRight(n uint64) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.b.ShiftRight(n)
}

// RotateLeft rotates all elements to the left by n positions.
func (cb *ConcurrentBuffer[T]) RotateLeft(n uint64) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.b.RotateLeft(n)
}

// RotateRight rotates all elements to the right by n positions.
func (cb *ConcurrentBuffer[T]) RotateRight(n uint64) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.b.RotateRight(n)
}

// Filter removes elements that don't match the predicate.
func (cb *ConcurrentBuffer[T]) Filter(predicate func(T) bool) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.b.Filter(predicate)
}

// Map creates a new buffer with the results of applying the function to each element.
func (cb *ConcurrentBuffer[T]) Map(fn func(T) T) (*ConcurrentBuffer[T], error) {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	mappedBuffer, err := cb.b.Map(fn)
	if err != nil {
		return nil, err
	}
	return &ConcurrentBuffer[T]{b: mappedBuffer}, nil
}

// Reduce reduces the buffer to a single value.
func (cb *ConcurrentBuffer[T]) Reduce(fn func(T, T) T) (T, error) {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.Reduce(fn)
}

// ForEach applies the function to each element in the buffer.
func (cb *ConcurrentBuffer[T]) ForEach(fn func(*T) error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.b.ForEach(fn)
}

// ForFrom applies the function to each element in the buffer starting from the given index.
func (cb *ConcurrentBuffer[T]) ForFrom(start uint64, fn func(*T) error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.b.ForFrom(start, fn)
}

// ForRange applies the function to each element in the buffer within the given range.
func (cb *ConcurrentBuffer[T]) ForRange(start, end uint64, fn func(*T) error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.b.ForRange(start, end, fn)
}

// Any checks if any element in the buffer matches the predicate.
func (cb *ConcurrentBuffer[T]) Any(predicate func(T) bool) bool {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.Any(predicate)
}

// All checks if all elements in the buffer match the predicate.
func (cb *ConcurrentBuffer[T]) All(predicate func(T) bool) bool {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.All(predicate)
}

// FindIndex returns the index of the first element that matches the predicate.
func (cb *ConcurrentBuffer[T]) FindIndex(predicate func(T) bool) (uint64, error) {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.FindIndex(predicate)
}

// FindLast returns the last element that matches the predicate.
func (cb *ConcurrentBuffer[T]) FindLast(predicate func(T) bool) (*T, error) {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.FindLast(predicate)
}

// FindLastIndex returns the index of the last element that matches the predicate.
func (cb *ConcurrentBuffer[T]) FindLastIndex(predicate func(T) bool) (uint64, error) {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.FindLastIndex(predicate)
}

// FindAll returns all elements that match the predicate.
func (cb *ConcurrentBuffer[T]) FindAll(predicate func(T) bool) *ConcurrentBuffer[T] {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	newBuffer := cb.b.FindAll(predicate)
	return &ConcurrentBuffer[T]{b: newBuffer}
}

// FindIndices returns the indices of all elements that match the predicate.
func (cb *ConcurrentBuffer[T]) FindIndices(predicate func(T) bool) []uint64 {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.FindIndices(predicate)
}

// LastIndexOf returns the index of the last element with the given value.
func (cb *ConcurrentBuffer[T]) LastIndexOf(value T) (uint64, error) {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.b.LastIndexOf(value)
}

// Blit combines/overwrites the values in the buffer with the values of another buffer using a function.
func (cb *ConcurrentBuffer[T]) Blit(other *ConcurrentBuffer[T], f func(T, T) T) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	return cb.b.Blit(other.b, f)
}
