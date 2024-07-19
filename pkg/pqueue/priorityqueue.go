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

// Package queue provides a non-concurrent-safe priority queue (FIFO).
package pqueue

import "errors"

// PriorityQueue is a thread-safe priority queue
type PriorityQueue[T comparable] struct {
	data *priorityQueue[T]
}

type priorityQueue[T comparable] struct {
	items   []T
	compare func(a, b T) bool
}

// NewPriorityQueue creates a new PriorityQueue
func NewPriorityQueue[T comparable](c func(T, T) bool) *PriorityQueue[T] {
	pq := &priorityQueue[T]{}
	pq.compare = c
	return &PriorityQueue[T]{data: pq}
}

// Insert adds an element to the priority queue
func (pq *PriorityQueue[T]) Insert(elem T) {
	pq.data.insert(elem)
}

// Pop removes and returns the element with the highest priority
func (pq *PriorityQueue[T]) Pop() (T, error) {
	return pq.data.pop()
}

// Peek returns the element with the highest priority without removing it
func (pq *PriorityQueue[T]) Peek() (T, error) {
	return pq.data.peek()
}

// Size returns the number of elements in the priority queue
func (pq *PriorityQueue[T]) Size() int {
	return pq.data.size()
}

// IsEmpty returns true if the priority queue is empty
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.data.isEmpty()
}

// Clear removes all elements from the priority queue
func (pq *PriorityQueue[T]) Clear() {
	pq.data.clear()
}

// Values returns all elements in the priority queue
func (pq *PriorityQueue[T]) Values() []T {
	return pq.data.values()
}

func (pq *priorityQueue[T]) insert(elem T) {
	pq.items = append(pq.items, elem)
}

func (pq *priorityQueue[T]) pop() (T, error) {
	if pq.isEmpty() {
		var rVal T
		return rVal, errors.New("priority queue is empty")
	}
	idx := 0
	for i := 1; i < len(pq.items); i++ {
		if pq.compare(pq.items[i], pq.items[idx]) {
			idx = i
		}
	}
	elem := pq.items[idx]
	pq.items = append(pq.items[:idx], pq.items[idx+1:]...)
	return elem, nil
}

func (pq *priorityQueue[T]) peek() (T, error) {
	if pq.isEmpty() {
		var rVal T
		return rVal, errors.New("priority queue is empty")
	}
	idx := 0
	for i := 1; i < len(pq.items); i++ {
		if pq.compare(pq.items[i], pq.items[idx]) {
			idx = i
		}
	}
	return pq.items[idx], nil
}

func (pq *priorityQueue[T]) size() int {
	return len(pq.items)
}

func (pq *priorityQueue[T]) isEmpty() bool {
	return len(pq.items) == 0
}

func (pq *priorityQueue[T]) clear() {
	pq.items = []T{}
}

func (pq *priorityQueue[T]) values() []T {
	return pq.items
}
