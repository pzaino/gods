package circularLinkList_test

import (
	"testing"

	"github.com/pzaino/gods/pkg/circularLinkList" // Adjust the import path as necessary
)

func TestAppend(t *testing.T) {
	list := circularLinkList.NewCircularLinkList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	expected := []int{1, 2, 3}
	actual := list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf("expected list length %d, got %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf("expected value %d, got %d", v, actual[i])
		}
	}
}

func TestPrepend(t *testing.T) {
	list := circularLinkList.NewCircularLinkList[int]()
	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)

	expected := []int{3, 2, 1}
	actual := list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf("expected length %d, got %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf("expected %d, got %d", v, actual[i])
		}
	}
}

func TestDeleteWithValue(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})
	list.DeleteWithValue(3)

	expected := []int{1, 2, 4}
	actual := list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf("expected length %d, got %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf("expected %d, got %d", v, actual[i])
		}
	}
}

func TestFind(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})
	node, err := list.Find(3)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if node == nil || node.Value != 3 {
		t.Fatalf("expected to find node with value 3")
	}
}

func TestReverse(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})
	list.Reverse()

	expected := []int{4, 3, 2, 1}
	actual := list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf("expected length %d, got %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf("expected %d, got %d", v, actual[i])
		}
	}
}

func TestSize(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})

	expected := 4
	actual := list.Size()

	if expected != actual {
		t.Fatalf("expected size %d, got %d", expected, actual)
	}
}

func TestGetFirst(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})
	node := list.GetFirst()

	if node == nil || node.Value != 1 {
		t.Fatalf("expected to get first node with value 1")
	}
}

func TestGetLast(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})
	node := list.GetLast()

	if node == nil || node.Value != 4 {
		t.Fatalf("expected to get last node with value 4")
	}
}

func TestGetAt(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4})
	node, err := list.GetAt(2)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if node == nil || node.Value != 3 {
		t.Fatalf("expected to get node with value 3 at index 2")
	}
}

func TestInsertAt(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 4, 5})
	err := list.InsertAt(2, 3)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []int{1, 2, 3, 4, 5}
	actual := list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf("expected length %d, got %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf("expected %d, got %d", v, actual[i])
		}
	}
}

func TestDeleteAt(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4, 5})
	err := list.DeleteAt(2)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []int{1, 2, 4, 5}
	actual := list.ToSlice()

	if len(expected) != len(actual) {
		t.Fatalf("expected length %d, got %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf("expected %d, got %d", v, actual[i])
		}
	}
}

func TestClear(t *testing.T) {
	list := circularLinkList.NewCircularLinkListFromSlice([]int{1, 2, 3, 4, 5})
	list.Clear()

	expected := 0
	actual := list.Size()

	if expected != actual {
		t.Fatalf("expected size %d, got %d", expected, actual)
	}
}
