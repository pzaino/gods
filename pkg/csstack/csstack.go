// CSStack is a concurrent-safe stack library.
package csstack

/*
 The  CSStack  type is a concurrent-safe stack that uses a  sync.Mutex  to protect the  items  field. The  Push ,  IsEmpty , and  Pop  methods are protected by the mutex. The  ToSlice  method is not necessary for a stack, but it is useful for debugging purposes.
 The  CSStack  type is similar to the  Stack  type, but it is safe to use concurrently.
 To use the  CSStack  type, you can import the  stack  package and create a new stack with the  CSStackNew  function:
*/

import (
	"errors"
	"fmt"
	"sync"

	stack "github.com/pzaino/gods/pkg/stack"
)

const (
	errItemNotFound = "item not found"
	errStackIsEmpty = "stack is empty"
)

// CSStack is a concurrent-safe stack.
type CSStack[T comparable] struct {
	mu    sync.Mutex
	items []T
}

// CSStackNew creates a new CSStack.
func CSStackNew[T comparable]() *CSStack[T] {
	return &CSStack[T]{}
}

// Push adds an item to the stack.
func (s *CSStack[T]) Push(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, item)
}

// IsEmpty checks if the stack is empty.
func (s *CSStack[T]) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items) == 0
}

// Pop removes and returns the top item from the stack.
func (s *CSStack[T]) Pop() (*T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.items) == 0 {
		return nil, errors.New(errStackIsEmpty)
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return &item, nil
}

// ToSlice returns the stack as a slice.
func (s *CSStack[T]) ToSlice() []T {
	s.mu.Lock()
	defer s.mu.Unlock()

	items := make([]T, len(s.items))
	copy(items, s.items)
	return items
}

// ToStack converts a CSStack to a Stack.
func (s *CSStack[T]) ToStack() *stack.Stack[T] {
	ns := stack.New[T]()
	for _, item := range s.ToSlice() {
		s.Push(item)
	}
	return ns
}

// Reverse reverses the stack.
func (s *CSStack[T]) Reverse() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := 0; i < len(s.items)/2; i++ {
		j := len(s.items) - i - 1
		s.items[i], s.items[j] = s.items[j], s.items[i]
	}
}

// Swap swaps the top two items on the stack.
func (s *CSStack[T]) Swap() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.items) < 2 {
		return errors.New("CSStack has less than 2 items")
	}

	s.items[len(s.items)-1], s.items[len(s.items)-2] = s.items[len(s.items)-2], s.items[len(s.items)-1]
	return nil
}

// Top returns the top item from the stack without removing it.
func (s *CSStack[T]) Top() (*T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.items) == 0 {
		return nil, errors.New(errStackIsEmpty)
	}

	item := s.items[len(s.items)-1]
	return &item, nil
}

// Peek is a wrapper around Top (for who's more used to use Peek).
func (s *CSStack[T]) Peek() (*T, error) {
	return s.Top()
}

// Size returns the number of items in the stack.
func (s *CSStack[T]) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return css_size(s)
}

func css_size[T comparable](s *CSStack[T]) int {
	return len(s.items)
}

// Clear removes all items from the stack.
func (s *CSStack[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = s.items[:0]
}

// Contains checks if the stack contains an item.
func (s *CSStack[T]) Contains(item T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, v := range s.items {
		if v == item {
			return true
		}
	}
	return false
}

// Copy returns a new CSStack with the same items.
func (s *CSStack[T]) Copy() *CSStack[T] {
	s.mu.Lock()
	defer s.mu.Unlock()

	stack := CSStackNew[T]()
	for _, item := range s.items {
		stack.Push(item)
	}
	return stack
}

// Equal checks if two stacks are equal.
func (s *CSStack[T]) Equal(other *CSStack[T]) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	other.mu.Lock()
	defer other.mu.Unlock()

	if css_size(s) != css_size(other) {
		return false
	}

	if css_size(s) == 0 && css_size(other) == 0 {
		return true
	}

	for i, v := range s.items {
		if v != other.items[i] {
			return false
		}
	}
	return true
}

// String returns a string representation of the stack.
func (s *CSStack[T]) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	return fmt.Sprintf("%v", s.items)
}

// PopN removes and returns the top n items from the stack.
func (s *CSStack[T]) PopN(n int) ([]T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.items) < n {
		return nil, errors.New("CSStack has less than n items")
	}

	items := s.items[len(s.items)-n:]
	s.items = s.items[:len(s.items)-n]
	return items, nil
}

// PushN adds multiple items to the stack.
func (s *CSStack[T]) PushN(items ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, items...)
}

// PopAll removes and returns all items from the stack.
func (s *CSStack[T]) PopAll() []T {
	s.mu.Lock()
	defer s.mu.Unlock()

	items := make([]T, len(s.items))
	copy(items, s.items)
	s.items = s.items[:0]
	return items
}

// PushAll adds multiple items to the stack.
func (s *CSStack[T]) PushAll(items []T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, items...)
}

// Filter removes items from the stack that don't match the predicate.
func (s *CSStack[T]) Filter(predicate func(T) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var items []T
	for _, item := range s.items {
		if predicate(item) {
			items = append(items, item)
		}
	}
	s.items = items
}

// Map creates a new stack with the results of applying the function to each item.
func (s *CSStack[T]) Map(fn func(T) T) *CSStack[T] {
	s.mu.Lock()
	defer s.mu.Unlock()

	stack := CSStackNew[T]()
	for _, item := range s.items {
		stack.Push(fn(item))
	}
	return stack
}

// Reduce reduces the stack to a single value.
func (s *CSStack[T]) Reduce(fn func(T, T) T) (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.items) == 0 {
		var rVal T
		return rVal, errors.New(errStackIsEmpty)
	}

	result := s.items[0]
	for i := 1; i < len(s.items); i++ {
		result = fn(result, s.items[i])
	}
	return result, nil
}

// ForEach applies the function to each item in the stack.
func (s *CSStack[T]) ForEach(fn func(T)) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, item := range s.items {
		fn(item)
	}
}

// Any checks if any item in the stack matches the predicate.
func (s *CSStack[T]) Any(predicate func(T) bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, item := range s.items {
		if predicate(item) {
			return true
		}
	}
	return false
}

// All checks if all items in the stack match the predicate.
func (s *CSStack[T]) All(predicate func(T) bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, item := range s.items {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Find returns the first item that matches the predicate.
func (s *CSStack[T]) Find(predicate func(T) bool) (*T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, item := range s.items {
		if predicate(item) {
			return &item, nil
		}
	}
	return nil, errors.New(errItemNotFound)
}

// FindIndex returns the index of the first item that matches the predicate.
func (s *CSStack[T]) FindIndex(predicate func(T) bool) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, item := range s.items {
		if predicate(item) {
			return i, nil
		}
	}
	return -1, errors.New(errItemNotFound)
}

// FindLast returns the last item that matches the predicate.
func (s *CSStack[T]) FindLast(predicate func(T) bool) (*T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := len(s.items) - 1; i >= 0; i-- {
		if predicate(s.items[i]) {
			return &s.items[i], nil
		}
	}
	return nil, errors.New(errItemNotFound)
}

// FindLastIndex returns the index of the last item that matches the predicate.
func (s *CSStack[T]) FindLastIndex(predicate func(T) bool) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := len(s.items) - 1; i >= 0; i-- {
		if predicate(s.items[i]) {
			return i, nil
		}
	}
	return -1, errors.New(errItemNotFound)
}

// FindAll returns all items that match the predicate.
func (s *CSStack[T]) FindAll(predicate func(T) bool) []T {
	s.mu.Lock()
	defer s.mu.Unlock()

	var items []T
	for _, item := range s.items {
		if predicate(item) {
			items = append(items, item)
		}
	}
	return items
}

// FindIndices returns the indices of all items that match the predicate.
func (s *CSStack[T]) FindIndices(predicate func(T) bool) []int {
	s.mu.Lock()
	defer s.mu.Unlock()

	var indices []int
	for i, item := range s.items {
		if predicate(item) {
			indices = append(indices, i)
		}
	}
	return indices
}
