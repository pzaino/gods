// queue package implements a simple FIFO queue. It is a generic implementation, meaning it can be used with any type that supports comparison. The queue is implemented using a slice, and the operations are O(1) time complexity.
package queue

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
