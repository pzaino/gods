package linkList

import "errors"

const (
	errIndexOutOfBound = "index out of bounds"
)

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type LinkList[T comparable] struct {
	Head *Node[T]
}

func New[T comparable]() *LinkList[T] {
	return &LinkList[T]{}
}

func (l *LinkList[T]) Append(value T) {
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

func (l *LinkList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value}

	newNode.Next = l.Head
	l.Head = newNode
}

func (l *LinkList[T]) DeleteWithValue(value T) {
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
		current = current.Next
	}
}

func (l *LinkList[T]) ToSlice() []T {
	var result []T

	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result
}

func (l *LinkList[T]) IsEmpty() bool {
	return l.Head == nil
}

func (l *LinkList[T]) Find(value T) (*Node[T], error) {
	current := l.Head
	for current != nil {
		if current.Value == value {
			return current, nil
		}
		current = current.Next
	}

	return nil, errors.New("value not found")
}

func (l *LinkList[T]) Reverse() {
	var prev *Node[T]
	current := l.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev
}

func (l *LinkList[T]) Size() int {
	size := 0
	current := l.Head
	for current != nil {
		size++
		current = current.Next
	}

	return size
}

func (l *LinkList[T]) GetFirst() *Node[T] {
	return l.Head
}

func (l *LinkList[T]) GetLast() *Node[T] {
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	return current
}

func (l *LinkList[T]) GetAt(index int) (*Node[T], error) {
	if index < 0 {
		return nil, errors.New(errIndexOutOfBound)
	}

	current := l.Head
	for i := 0; i < index; i++ {
		if current == nil {
			return nil, errors.New(errIndexOutOfBound)
		}
		current = current.Next
	}

	if current == nil {
		return nil, errors.New(errIndexOutOfBound)
	}

	return current, nil
}

func (l *LinkList[T]) InsertAt(index int, value T) error {
	if index < 0 {
		return errors.New(errIndexOutOfBound)
	}

	if index == 0 {
		l.Prepend(value)
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

	newNode := &Node[T]{Value: value}
	newNode.Next = current.Next
	current.Next = newNode

	return nil
}

func (l *LinkList[T]) DeleteAt(index int) error {
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

func (l *LinkList[T]) Remove(value T) {
	l.DeleteWithValue(value)
}

func (l *LinkList[T]) Clear() {
	l.Head = nil
}

func (l *LinkList[T]) Copy() *LinkList[T] {
	newList := New[T]()

	current := l.Head
	for current != nil {
		newList.Append(current.Value)
		current = current.Next
	}

	return newList
}

func (l *LinkList[T]) Merge(list *LinkList[T]) {
	current := list.Head
	for current != nil {
		l.Append(current.Value)
		current = current.Next
	}
}

func (l *LinkList[T]) Map(f func(T) T) {
	current := l.Head
	for current != nil {
		current.Value = f(current.Value)
		current = current.Next
	}
}

func (l *LinkList[T]) Filter(f func(T) bool) {
	if l.Head == nil {
		return
	}

	if !f(l.Head.Value) {
		l.Head = l.Head.Next
	}

	current := l.Head
	for current.Next != nil {
		if f(current.Next.Value) {
			current = current.Next
		} else {
			current.Next = current.Next.Next
		}
	}
}

func (l *LinkList[T]) Reduce(f func(T, T) T, initial T) T {
	result := initial

	current := l.Head
	for current != nil {
		result = f(result, current.Value)
		current = current.Next
	}

	return result
}

func (l *LinkList[T]) ForEach(f func(T)) {
	current := l.Head
	for current != nil {
		f(current.Value)
		current = current.Next
	}
}

func (l *LinkList[T]) Any(f func(T) bool) bool {
	current := l.Head
	for current != nil {
		if f(current.Value) {
			return true
		}
		current = current.Next
	}

	return false
}

func (l *LinkList[T]) All(f func(T) bool) bool {
	current := l.Head
	for current != nil {
		if !f(current.Value) {
			return false
		}
		current = current.Next
	}

	return true
}

func (l *LinkList[T]) Contains(value T) bool {
	current := l.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}

	return false
}

func (l *LinkList[T]) IndexOf(value T) int {
	current := l.Head
	index := 0
	for current != nil {
		if current.Value == value {
			return index
		}
		current = current.Next
		index++
	}

	return -1
}

func (l *LinkList[T]) LastIndexOf(value T) int {
	current := l.Head
	index := -1
	i := 0
	for current != nil {
		if current.Value == value {
			index = i
		}
		current = current.Next
		i++
	}

	return index
}

func (l *LinkList[T]) FindIndex(f func(T) bool) int {
	current := l.Head
	index := 0
	for current != nil {
		if f(current.Value) {
			return index
		}
		current = current.Next
		index++
	}

	return -1
}

func (l *LinkList[T]) FindLastIndex(f func(T) bool) int {
	current := l.Head
	index := -1
	i := 0
	for current != nil {
		if f(current.Value) {
			index = i
		}
		current = current.Next
		i++
	}

	return index
}

func (l *LinkList[T]) FindAll(f func(T) bool) *LinkList[T] {
	newList := New[T]()

	current := l.Head
	for current != nil {
		if f(current.Value) {
			newList.Append(current.Value)
		}
		current = current.Next
	}

	return newList
}

func (l *LinkList[T]) FindLast(f func(T) bool) (*Node[T], error) {
	var result *Node[T]

	current := l.Head
	for current != nil {
		if f(current.Value) {
			result = current
		}
		current = current.Next
	}

	if result == nil {
		return nil, errors.New("value not found")
	}

	return result, nil
}

func (l *LinkList[T]) FindAllIndexes(f func(T) bool) []int {
	var result []int

	current := l.Head
	index := 0
	for current != nil {
		if f(current.Value) {
			result = append(result, index)
		}
		current = current.Next
		index++
	}

	return result
}
