package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	if q.Size() != 0 {
		t.Errorf("Queue should be empty")
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Errorf("Queue should have 3 elements")
	}

	item, err := q.Peek()
	if err != nil {
		t.Errorf("Peek should not return an error")
	}
	if item != 1 {
		t.Errorf("Peek should return 1")
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue should not return an error")
	}
	if item != 1 {
		t.Errorf("Dequeue should return 1")
	}

	if q.Size() != 2 {
		t.Errorf("Queue should have 2 elements")
	}

	if q.IsEmpty() {
		t.Errorf("Queue should not be empty")
	}

	q.Clear()
	if !q.IsEmpty() {
		t.Errorf("Queue should be empty")
	}
}
