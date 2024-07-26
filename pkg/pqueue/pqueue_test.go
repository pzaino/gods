package pqueue_test

import (
	"fmt"
	"testing"

	"github.com/pzaino/gods/pkg/pqueue"
)

func TestNewPriorityQueue(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	if pq == nil {
		t.Fatal("Expected new priority queue to be non-nil")
	}
	if !pq.IsEmpty() {
		t.Fatal("Expected new priority queue to be empty")
	}
}

func TestEnqueueAndDequeue(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	if pq.Size() != 2 {
		t.Fatal("Expected priority queue size to be 2")
	}

	val, err := pq.Dequeue()
	if err != nil {
		t.Fatal(err)
	}
	if val != 20 {
		t.Fatalf("Expected dequeued value to be 20, got %d", val)
	}

	val, err = pq.Dequeue()
	if err != nil {
		t.Fatal(err)
	}
	if val != 10 {
		t.Fatalf("Expected dequeued value to be 10, got %d", val)
	}

	if !pq.IsEmpty() {
		t.Fatal("Expected priority queue to be empty after dequeueing all elements")
	}
}

func TestPeek(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	val, err := pq.Peek()
	if err != nil {
		t.Fatal(err)
	}
	if val != 10 {
		t.Fatalf("Expected peek value to be 10, got %d", val)
	}
}

func TestClear(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Clear()
	if !pq.IsEmpty() {
		t.Fatal("Expected priority queue to be empty after clear")
	}
}

func TestContains(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	if !pq.Contains(10) {
		t.Fatal("Expected priority queue to contain value 10")
	}
	if pq.Contains(20) {
		t.Fatal("Expected priority queue to not contain value 20")
	}
}

func TestEquals(t *testing.T) {
	pq1 := pqueue.NewPriorityQueue[int]()
	pq2 := pqueue.NewPriorityQueue[int]()
	pq1.Enqueue(10, 1)
	pq2.Enqueue(10, 1)
	if !pq1.Equals(pq2) {
		t.Fatal("Expected priority queues to be equal")
	}
	pq2.Enqueue(20, 2)
	if pq1.Equals(pq2) {
		t.Fatal("Expected priority queues to be not equal")
	}
}

func TestCopy(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	copy := pq.Copy()
	if !copy.Equals(pq) {
		t.Fatal("Expected copy to be equal to original")
	}
	copy.Enqueue(20, 2)
	if copy.Equals(pq) {
		t.Fatal("Expected copy to not be equal to original after modification")
	}
}

func TestString(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	str := pq.String(func(val int) string {
		return fmt.Sprintf("%d", val)
	})
	expected := "[10]"
	if str != expected {
		t.Fatalf("Expected string representation to be %s, got %s", expected, str)
	}
}

func TestMap(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	mapped := pq.Map(func(val int) int {
		return val * 2
	})
	expectedValues := []int{40, 20}
	for i, val := range mapped.Values() {
		if val != expectedValues[i] {
			t.Fatalf("Expected mapped value %d, got %d", expectedValues[i], val)
		}
	}
}

func TestFilter(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.Filter(func(val int) bool {
		return val > 10
	})
	if pq.Size() != 1 || pq.Contains(10) {
		t.Fatal("Expected priority queue to contain only values greater than 10")
	}
}

func TestReduce(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	result := pq.Reduce(func(a, b int) int {
		return a + b
	}, 0)
	if result != 30 {
		t.Fatalf("Expected reduced result to be 30, got %d", result)
	}
}

func TestForEach(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.ForEach(func(val *int) {
		*val += 5
	})
	expectedValues := []int{25, 15}
	for i, val := range pq.Values() {
		if val != expectedValues[i] {
			t.Fatalf("Expected for each modified value %d, got %d", expectedValues[i], val)
		}
	}
}

func TestAny(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	if !pq.Any(func(val int) bool { return val > 15 }) {
		t.Fatal("Expected any to return true for values greater than 15")
	}
	if pq.Any(func(val int) bool { return val > 25 }) {
		t.Fatal("Expected any to return false for values greater than 25")
	}
}

func TestAll(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	if !pq.All(func(val int) bool { return val > 5 }) {
		t.Fatal("Expected all to return true for values greater than 5")
	}
	if pq.All(func(val int) bool { return val > 15 }) {
		t.Fatal("Expected all to return false for values greater than 15")
	}
}

func TestIndexOf(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	index := pq.IndexOf(10)
	if index != 1 {
		t.Fatalf("Expected index of 10 to be 1, got %d", index)
	}
	index = pq.IndexOf(30)
	if index != -1 {
		t.Fatalf("Expected index of 30 to be -1, got %d", index)
	}
}

func TestLastIndexOf(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.Enqueue(10, 3)
	index := pq.LastIndexOf(10)
	if index != 1 {
		t.Fatalf("Expected last index of 10 to be 0, got %d", index)
	}
}

func TestFindIndex(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	index := pq.FindIndex(func(val int) bool { return val > 15 })
	if index != 0 {
		t.Fatalf("Expected find index for value greater than 15 to be 0, got %d", index)
	}
	index = pq.FindIndex(func(val int) bool { return val > 25 })
	if index != -1 {
		t.Fatalf("Expected find index for value greater than 25 to be -1, got %d", index)
	}
}

func TestFindLastIndex(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.Enqueue(10, 3)
	index := pq.FindLastIndex(func(val int) bool { return val == 10 })
	if index != 1 {
		t.Fatalf("Expected find last index for value 10 to be 0, got %d", index)
	}
}

func TestFindAll(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.Enqueue(10, 3)
	result := pq.FindAll(func(val int) bool { return val == 10 })
	if result.Size() != 2 {
		t.Fatalf("Expected find all to return 2 elements, got %d", result.Size())
	}
}

func TestFindLast(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	val, err := pq.FindLast(func(val int) bool { return val == 10 })
	if err != nil {
		t.Fatal(err)
	}
	if val != 10 {
		t.Fatalf("Expected find last value to be 10, got %d", val)
	}
	_, err = pq.FindLast(func(val int) bool { return val == 30 })
	if err == nil {
		t.Fatal("Expected find last to return error for non-existing value")
	}
}

func TestFindAllIndexes(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	pq.Enqueue(10, 3)
	indexes := pq.FindAllIndexes(func(val int) bool { return val == 10 })
	if len(indexes) != 2 {
		t.Fatalf("Expected find all indexes to return 2 elements, got %d", len(indexes))
	}
}

func TestValues(t *testing.T) {
	pq := pqueue.NewPriorityQueue[int]()
	pq.Enqueue(10, 1)
	pq.Enqueue(20, 2)
	values := pq.Values()
	expectedValues := []int{20, 10}
	for i, val := range values {
		if val != expectedValues[i] {
			t.Fatalf("Expected value %d, got %d", expectedValues[i], val)
		}
	}
}
