// Package circularLinkList provides a non-concurrent-safe circular linked list.
package circularLinkList

import "errors"

const (
	errIndexOutOfBound = "index out of bounds"
)

// Node represents a node in the circular linked list
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// CircularLinkList represents a circular linked list
type CircularLinkList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

// NewCircularLinkList creates a new CircularLinkList
func NewCircularLinkList[T comparable]() *CircularLinkList[T] {
	return &CircularLinkList[T]{}
}

// NewCircularLinkListFromSlice creates a new CircularLinkList from a slice
func NewCircularLinkListFromSlice[T comparable](items []T) *CircularLinkList[T] {
	l := NewCircularLinkList[T]()
	for i := 0; i < len(items); i++ {
		l.Append(items[i])
	}
	return l
}

// Append adds a new node to the end of the list
func (l *CircularLinkList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		newNode.Next = newNode
		return
	}

	l.Tail.Next = newNode
	newNode.Next = l.Head
	l.Tail = newNode
}

// Prepend adds a new node to the beginning of the list
func (l *CircularLinkList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		newNode.Next = newNode
		return
	}

	newNode.Next = l.Head
	l.Head = newNode
	l.Tail.Next = newNode
}

// DeleteWithValue deletes the first node with the given value
func (l *CircularLinkList[T]) DeleteWithValue(value T) {
	if l.Head == nil {
		return
	}

	// Special case: head node needs to be deleted
	if l.Head.Value == value {
		if l.Head == l.Tail {
			l.Head = nil
			l.Tail = nil
			return
		}
		l.Head = l.Head.Next
		l.Tail.Next = l.Head
		return
	}

	current := l.Head
	for current.Next != l.Head {
		if current.Next.Value == value {
			if current.Next == l.Tail {
				l.Tail = current
			}
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

// ToSlice returns the list as a slice
func (l *CircularLinkList[T]) ToSlice() []T {
	var result []T

	if l.Head == nil {
		return result
	}

	current := l.Head
	for {
		result = append(result, current.Value)
		current = current.Next
		if current == l.Head {
			break
		}
	}

	return result
}

// IsEmpty checks if the list is empty
func (l *CircularLinkList[T]) IsEmpty() bool {
	return l.Head == nil
}

// Find returns the first node with the given value
func (l *CircularLinkList[T]) Find(value T) (*Node[T], error) {
	if l.Head == nil {
		return nil, errors.New("value not found")
	}

	current := l.Head
	for {
		if current.Value == value {
			return current, nil
		}
		current = current.Next
		if current == l.Head {
			break
		}
	}

	return nil, errors.New("value not found")
}

// Reverse reverses the list
func (l *CircularLinkList[T]) Reverse() {
	if l.Head == nil {
		return
	}

	var prev, next *Node[T]
	current := l.Head
	l.Tail = l.Head

	for {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
		if current == l.Head {
			break
		}
	}

	l.Head.Next = prev
	l.Head = prev
}

// Size returns the number of nodes in the list
func (l *CircularLinkList[T]) Size() int {
	size := 0

	if l.Head == nil {
		return size
	}

	current := l.Head
	for {
		size++
		current = current.Next
		if current == l.Head {
			break
		}
	}

	return size
}

// GetFirst returns the first node in the list
func (l *CircularLinkList[T]) GetFirst() *Node[T] {
	return l.Head
}

// GetLast returns the last node in the list
func (l *CircularLinkList[T]) GetLast() *Node[T] {
	return l.Tail
}

// GetAt returns the node at the given index
func (l *CircularLinkList[T]) GetAt(index int) (*Node[T], error) {
	if index < 0 {
		return nil, errors.New(errIndexOutOfBound)
	}

	if l.Head == nil {
		return nil, errors.New(errIndexOutOfBound)
	}

	current := l.Head
	for i := 0; i < index; i++ {
		current = current.Next
		if current == l.Head {
			return nil, errors.New(errIndexOutOfBound)
		}
	}

	return current, nil
}

// InsertAt inserts a new node at the given index
func (l *CircularLinkList[T]) InsertAt(index int, value T) error {
	if index < 0 {
		return errors.New(errIndexOutOfBound)
	}

	if index == 0 {
		l.Prepend(value)
		return nil
	}

	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
		if current == l.Head {
			return errors.New(errIndexOutOfBound)
		}
	}

	newNode := &Node[T]{Value: value}
	newNode.Next = current.Next
	current.Next = newNode

	if current == l.Tail {
		l.Tail = newNode
	}

	return nil
}

// DeleteAt deletes the node at the given index
func (l *CircularLinkList[T]) DeleteAt(index int) error {
	if index < 0 {
		return errors.New(errIndexOutOfBound)
	}

	if index == 0 {
		if l.Head == nil {
			return errors.New(errIndexOutOfBound)
		}
		if l.Head == l.Tail {
			l.Head = nil
			l.Tail = nil
			return nil
		}
		l.Head = l.Head.Next
		l.Tail.Next = l.Head
		return nil
	}

	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
		if current == l.Head {
			return errors.New(errIndexOutOfBound)
		}
	}

	if current.Next == l.Head {
		return errors.New(errIndexOutOfBound)
	}

	if current.Next == l.Tail {
		l.Tail = current
	}

	current.Next = current.Next.Next

	return nil
}

// Clear removes all nodes from the list
func (l *CircularLinkList[T]) Clear() {
	l.Head = nil
	l.Tail = nil
}
