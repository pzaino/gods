package stack

import (
	"testing"
)

func TestNew(t *testing.T) {
	stack := New[int]()
	if stack == nil {
		t.Error("Expected stack to be initialized, but got nil")
	}
}

func TestPush(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	if stack.IsEmpty() {
		t.Error("Expected stack to not be empty, but it was")
	}
	if len(stack.items) != 1 {
		t.Errorf("Expected stack to have 1 item, but got %v", len(stack.items))
	}
}

func TestPop(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	item, err := stack.Pop()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if item == nil {
		t.Error("Expected item to not be nil, but it was")
	} else if *item != 1 {
		t.Errorf("Expected item to be 1, but got %v", *item)
	}
}

func TestPopEmpty(t *testing.T) {
	stack := New[int]()
	_, err := stack.Pop()
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}

func TestIsEmpty(t *testing.T) {
	stack := New[int]()
	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty, but it was not")
	}
}

func TestIsEmptyAfterPush(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	if stack.IsEmpty() {
		t.Error("Expected stack to not be empty, but it was")
	}
}

func TestIsEmptyAfterPop(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	item, err := stack.Pop()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if item == nil {
		t.Error("Expected item to not be nil, but it was")
	}
	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty, but it was not")
	}
}
