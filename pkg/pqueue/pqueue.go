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

// Package pqueue provides a non-concurrent-safe priority queue.
package pqueue

import (
	"errors"
	"strings"
)

// Element represents an element in the priority queue with a value and a priority.
type Element[T comparable] struct {
	Value    T
	Priority int
}

// PriorityQueue is a priority queue data structure
type PriorityQueue[T comparable] struct {
	data []Element[T]
}

// New creates a new PriorityQueue
func NewPriorityQueue[T comparable]() *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

// IsEmpty returns true if the priority queue is empty
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.data) == 0
}

// Enqueue adds an element to the priority queue
func (pq *PriorityQueue[T]) Enqueue(value T, priority int) {
	element := Element[T]{Value: value, Priority: priority}
	pq.data = append(pq.data, element)
	pq.upHeap(len(pq.data) - 1)
}

// Dequeue removes and returns the highest priority element in the queue
func (pq *PriorityQueue[T]) Dequeue() (T, error) {
	if pq.IsEmpty() {
		var rVal T
		return rVal, errors.New("priority queue is empty")
	}

	element := pq.data[0]
	lastIndex := len(pq.data) - 1
	pq.data[0] = pq.data[lastIndex]
	pq.data = pq.data[:lastIndex]
	pq.downHeap(0)

	return element.Value, nil
}

// Peek returns the highest priority element in the queue without removing it
func (pq *PriorityQueue[T]) Peek() (T, error) {
	if pq.IsEmpty() {
		var rVal T
		return rVal, errors.New("priority queue is empty")
	}
	return pq.data[0].Value, nil
}

// Size returns the number of elements in the priority queue
func (pq *PriorityQueue[T]) Size() int {
	return len(pq.data)
}

// Clear removes all elements from the priority queue
func (pq *PriorityQueue[T]) Clear() {
	pq.data = []Element[T]{}
}

// Values returns all elements in the priority queue
func (pq *PriorityQueue[T]) Values() []T {
	values := make([]T, len(pq.data))
	for i, element := range pq.data {
		values[i] = element.Value
	}
	return values
}

// Contains returns true if the priority queue contains the given element
func (pq *PriorityQueue[T]) Contains(value T) bool {
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
	copy := NewPriorityQueue[T]()
	copy.data = append(copy.data, pq.data...)
	return copy
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

// Helper functions for heap operations

func (pq *PriorityQueue[T]) upHeap(index int) {
	if len(pq.data) == 0 {
		return
	}
	if index < 0 {
		return
	}
	if index >= len(pq.data) {
		return
	}

	element := pq.data[index]
	for index > 0 {
		parent := (index - 1) / 2
		if pq.data[parent].Priority >= element.Priority {
			break
		}
		pq.data[index] = pq.data[parent]
		index = parent
	}
	pq.data[index] = element
}

func (pq *PriorityQueue[T]) downHeap(index int) {
	if len(pq.data) == 0 {
		return
	}
	if index < 0 {
		return
	}
	if index >= len(pq.data) {
		return
	}
	element := pq.data[index]
	lastIndex := len(pq.data) - 1
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

// Map creates a new priority queue with the results of applying the function to each element
func (pq *PriorityQueue[T]) Map(f func(T) T) *PriorityQueue[T] {
	newQueue := NewPriorityQueue[T]()
	for i := 0; i < len(pq.data); i++ {
		newQueue.Enqueue(f(pq.data[i].Value), pq.data[i].Priority)
	}
	return newQueue
}

// Filter removes elements from the priority queue that don't match the predicate
func (pq *PriorityQueue[T]) Filter(f func(T) bool) {
	var newData []Element[T]
	for i := 0; i < len(pq.data); i++ {
		if f(pq.data[i].Value) {
			newData = append(newData, pq.data[i])
		}
	}
	pq.data = newData
}

// Reduce reduces the priority queue to a single value
func (pq *PriorityQueue[T]) Reduce(f func(T, T) T, initial T) T {
	result := initial
	for i := 0; i < len(pq.data); i++ {
		result = f(result, pq.data[i].Value)
	}
	return result
}

// ForEach applies the function to all the elements in the priority queue
func (pq *PriorityQueue[T]) ForEach(f func(*T)) {
	for i := 0; i < len(pq.data); i++ {
		f(&pq.data[i].Value)
	}
}

// Any checks if any element in the priority queue matches the predicate
func (pq *PriorityQueue[T]) Any(f func(T) bool) bool {
	for i := 0; i < len(pq.data); i++ {
		if f(pq.data[i].Value) {
			return true
		}
	}
	return false
}

// All checks if all elements in the priority queue match the predicate
func (pq *PriorityQueue[T]) All(f func(T) bool) bool {
	for i := 0; i < len(pq.data); i++ {
		if !f(pq.data[i].Value) {
			return false
		}
	}
	return true
}

// IndexOf returns the index of the first element with the given value
func (pq *PriorityQueue[T]) IndexOf(value T) int {
	for i := 0; i < len(pq.data); i++ {
		if pq.data[i].Value == value {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last element with the given value
func (pq *PriorityQueue[T]) LastIndexOf(value T) int {
	index := -1
	for i := 0; i < len(pq.data); i++ {
		if pq.data[i].Value == value {
			index = i
		}
	}
	return index
}

// FindIndex returns the index of the first element that matches the predicate
func (pq *PriorityQueue[T]) FindIndex(f func(T) bool) int {
	for i := 0; i < len(pq.data); i++ {
		if f(pq.data[i].Value) {
			return i
		}
	}
	return -1
}

// FindLastIndex returns the index of the last element that matches the predicate
func (pq *PriorityQueue[T]) FindLastIndex(f func(T) bool) int {
	index := -1
	for i := 0; i < len(pq.data); i++ {
		if f(pq.data[i].Value) {
			index = i
		}
	}
	return index
}

// FindAll returns all elements that match the predicate
func (pq *PriorityQueue[T]) FindAll(f func(T) bool) *PriorityQueue[T] {
	newQueue := NewPriorityQueue[T]()
	for i := 0; i < len(pq.data); i++ {
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
	for i := 0; i < len(pq.data); i++ {
		if f(pq.data[i].Value) {
			result = pq.data[i].Value
			found = true
		}
	}
	if !found {
		return result, errors.New("value not found")
	}
	return result, nil
}

// FindAllIndexes returns the indexes of all elements that match the predicate
func (pq *PriorityQueue[T]) FindAllIndexes(f func(T) bool) []int {
	var result []int
	for i := 0; i < len(pq.data); i++ {
		if f(pq.data[i].Value) {
			result = append(result, i)
		}
	}
	return result
}
