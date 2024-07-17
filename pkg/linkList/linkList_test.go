package linkList

import (
	"testing"
)

const (
	errListNotEmpty = "Expected list to be empty, but it was not"
	errListIsEmpty  = "Expected list to not be empty, but it was"
)

func TestNew(t *testing.T) {
	list := New[int]()
	if list == nil {
		t.Error("Expected list to be initialized, but got nil")
	}
}

func TestAppend(t *testing.T) {
	list := New[int]()
	list.Append(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list to have 1 item, but got %v", list.Size())
	}
}

func TestPrepend(t *testing.T) {
	list := New[int]()
	list.Prepend(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
	if list.Size() != 1 {
		t.Errorf("Expected list to have 1 item, but got %v", list.Size())
	}
}

func TestRemove(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestRemoveEmpty(t *testing.T) {
	list := New[int]()
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestIsEmpty(t *testing.T) {
	list := New[int]()
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
}

func TestIsEmptyAfterAppend(t *testing.T) {
	list := New[int]()
	list.Append(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
}

func TestIsEmptyAfterPrepend(t *testing.T) {
	list := New[int]()
	list.Prepend(1)
	if list.IsEmpty() {
		t.Error(errListIsEmpty)
	}
}

func TestIsEmptyAfterRemove(t *testing.T) {
	list := New[int]()
	list.Append(1)
	list.Remove(1)
	if !list.IsEmpty() {
		t.Error(errListNotEmpty)
	}
	if list.Size() != 0 {
		t.Errorf("Expected list to have 0 items, but got %v", list.Size())
	}
}
