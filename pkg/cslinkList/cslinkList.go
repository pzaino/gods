package cslinkList

/*
LinkList is a concurrent-safe linked list.
it's API is identical to the LinkList type in the linkList package.
*/

import (
	"errors"
	"sync"
)

const (
	errIndexOutOfBound = "index out of bounds"
)

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type CSLinkList[T comparable] struct {
	mu   sync.Mutex
	Head *Node[T]
}

func CSLinkListNew[T comparable]() *CSLinkList[T] {
	return &CSLinkList[T]{}
}

func (l *CSLinkList[T]) Append(value T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	newNode := &Node[T]{Value: value}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func (l *CSLinkList[T]) Prepend(value T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	newNode := &Node[T]{Value: value}

	newNode.Next = l.Head
	l.Head = newNode
}

func (l *CSLinkList[T]) DeleteWithValue(value T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.Head == nil {
		return
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}
	}
}

func (l *CSLinkList[T]) ToSlice() []T {
	l.mu.Lock()
	defer l.mu.Unlock()

	var result []T
	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

func (l *CSLinkList[T]) Insert(index int, value T) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if index < 0 {
		return errors.New(errIndexOutOfBound)
	}

	newNode := &Node[T]{Value: value}

	if index == 0 {
		newNode.Next = l.Head
		l.Head = newNode
		return nil
	}

	current := l.Head
	for i := 0; i < index-1; i++ {
		if current == nil {
			return errors.New(errIndexOutOfBound)
		}
		current = current.Next
	}

	if current == nil {
		return errors.New(errIndexOutOfBound)
	}

	newNode.Next = current.Next
	current.Next = newNode

	return nil
}

func (l *CSLinkList[T]) Delete(index int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if index < 0 {
		return errors.New(errIndexOutOfBound)
	}

	if index == 0 {
		if l.Head == nil {
			return errors.New(errIndexOutOfBound)
		}
		l.Head = l.Head.Next
		return nil
	}

	current := l.Head
	for i := 0; i < index-1; i++ {
		if current == nil {
			return errors.New(errIndexOutOfBound)
		}
		current = current.Next
	}

	if current == nil || current.Next == nil {
		return errors.New(errIndexOutOfBound)
	}

	current.Next = current.Next.Next

	return nil
}

func (l *CSLinkList[T]) Find(predicate func(T) bool) (*T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for current := l.Head; current != nil; current = current.Next {
		if predicate(current.Value) {
			return &current.Value, nil
		}
	}
	return nil, errors.New(errIndexOutOfBound)
}

func (l *CSLinkList[T]) FindIndex(predicate func(T) bool) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	index := 0
	for current := l.Head; current != nil; current = current.Next {
		if predicate(current.Value) {
			return index, nil
		}
		index++
	}
	return -1, errors.New(errIndexOutOfBound)
}

func (l *CSLinkList[T]) FindLast(predicate func(T) bool) (*T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var result *T
	for current := l.Head; current != nil; current = current.Next {
		if predicate(current.Value) {
			result = &current.Value
		}
	}
	if result == nil {
		return nil, errors.New(errIndexOutOfBound)
	}
	return result, nil
}

func (l *CSLinkList[T]) FindLastIndex(predicate func(T) bool) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	index := -1
	i := 0
	for current := l.Head; current != nil; current = current.Next {
		if predicate(current.Value) {
			index = i
		}
		i++
	}
	if index == -1 {
		return -1, errors.New(errIndexOutOfBound)
	}
	return index, nil
}

func (l *CSLinkList[T]) FindAll(predicate func(T) bool) *CSLinkList[T] {
	l.mu.Lock()
	defer l.mu.Unlock()

	newList := CSLinkListNew[T]()

	for current := l.Head; current != nil; current = current.Next {
		if predicate(current.Value) {
			newList.Append(current.Value)
		}
	}

	return newList
}

func (l *CSLinkList[T]) FindIndices(predicate func(T) bool) []int {
	l.mu.Lock()
	defer l.mu.Unlock()

	var indices []int
	for i, current := 0, l.Head; current != nil; i, current = i+1, current.Next {
		if predicate(current.Value) {
			indices = append(indices, i)
		}
	}

	return indices
}

func (l *CSLinkList[T]) FindAllIndexes(predicate func(T) bool) []int {
	l.mu.Lock()
	defer l.mu.Unlock()

	var indices []int
	for i, current := 0, l.Head; current != nil; i, current = i+1, current.Next {
		if predicate(current.Value) {
			indices = append(indices, i)
		}
	}

	return indices
}

func (l *CSLinkList[T]) FindAllNodes(predicate func(T) bool) *CSLinkList[T] {
	l.mu.Lock()
	defer l.mu.Unlock()

	newList := CSLinkListNew[T]()

	for current := l.Head; current != nil; current = current.Next {
		if predicate(current.Value) {
			newList.Append(current.Value)
		}
	}

	return newList
}

func (l *CSLinkList[T]) FindAllNodesIndexes(predicate func(T) bool) []int {
	l.mu.Lock()
	defer l.mu.Unlock()

	var indices []int
	for i, current := 0, l.Head; current != nil; i, current = i+1, current.Next {
		if predicate(current.Value) {
			indices = append(indices, i)
		}
	}

	return indices
}
