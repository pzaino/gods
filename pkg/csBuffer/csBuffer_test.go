// Description: This file contains tests for the concurrent buffer implementation.
package csBuffer_test

import (
	"sync"
	"testing"

	buffer "github.com/pzaino/gods/pkg/csBuffer"
)

const (
	errUnexpectedErr = "unexpected error during append: %v"
)

// TestConcurrentAppend tests concurrent appends to the buffer.
func TestConcurrentAppend(t *testing.T) {
	cb := buffer.New[int]()
	var wg sync.WaitGroup
	numGoroutines := 100
	numAppendsPerGoroutine := 10

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < numAppendsPerGoroutine; j++ {
				err := cb.Append(i*numAppendsPerGoroutine + j)
				if err != nil {
					t.Errorf(errUnexpectedErr, err)
				}
			}
		}(i)
	}

	wg.Wait()
	expectedSize := uint64(numGoroutines * numAppendsPerGoroutine)
	if cb.Size() != expectedSize {
		t.Errorf("expected buffer size %d, got %d", expectedSize, cb.Size())
	}
}

// TestConcurrentGet tests concurrent reads from the buffer.
func TestConcurrentGet(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf(errUnexpectedErr, err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < numElements; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val, err := cb.Get(uint64(i))
			if err != nil {
				t.Errorf("unexpected error during get: %v", err)
			}
			if val != i {
				t.Errorf("expected value %d, got %d", i, val)
			}
		}(i)
	}

	wg.Wait()
}

// TestConcurrentRemove tests concurrent removal of elements from the buffer.
func TestConcurrentRemove(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf(errUnexpectedErr, err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < numElements; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := cb.Remove(0) // Always remove the first element
			if err != nil {
				t.Errorf("unexpected error during remove: %v", err)
			}
		}()
	}

	wg.Wait()
	if !cb.IsEmpty() {
		t.Error("expected buffer to be empty after all elements are removed")
	}
}

// TestConcurrentFind tests concurrent find operations in the buffer.
func TestConcurrentFind(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf(errUnexpectedErr, err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < numElements; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			index, err := cb.Find(i)
			if err != nil {
				t.Errorf("unexpected error during find: %v", err)
			}
			if index != uint64(i) {
				t.Errorf("expected index %d, got %d", i, index)
			}
		}(i)
	}

	wg.Wait()
}

// TestConcurrentMerge tests concurrent merge operations.
func TestConcurrentMerge(t *testing.T) {
	cb1 := buffer.New[int]()
	cb2 := buffer.New[int]()

	numElements := 50
	for i := 0; i < numElements; i++ {
		err := cb1.Append(i)
		if err != nil {
			t.Errorf("unexpected error during append to cb1: %v", err)
		}
		err = cb2.Append(i + numElements)
		if err != nil {
			t.Errorf("unexpected error during append to cb2: %v", err)
		}
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		cb1.Merge(cb2)
	}()

	go func() {
		defer wg.Done()
		cb2.Merge(cb1)
	}()

	wg.Wait()

	expectedSize := uint64(numElements * 2)
	if cb1.Size() != expectedSize {
		t.Errorf("expected cb1 size %d after merging twice, got %d", expectedSize, cb1.Size())
	}

	if cb2.Size() != 0 {
		t.Errorf("expected cb2 to be empty after merging into cb1, got size %d", cb2.Size())
	}
}

// TestConcurrentCapacity ensures capacity functions work under concurrency.
func TestConcurrentCapacity(t *testing.T) {
	cb := buffer.New[int]()
	cb.SetCapacity(200)

	var wg sync.WaitGroup
	numGoroutines := 10
	numAppendsPerGoroutine := 10

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < numAppendsPerGoroutine; j++ {
				err := cb.Append(i*numAppendsPerGoroutine + j)
				if err != nil && err.Error() != "buffer overflow" {
					t.Errorf(errUnexpectedErr, err)
				}
			}
		}(i)
	}

	wg.Wait()
	expectedSize := uint64(numGoroutines * numAppendsPerGoroutine)
	if cb.Size() != expectedSize {
		t.Errorf("expected buffer size %d, got %d", expectedSize, cb.Size())
	}
	if cb.Capacity() != 200 {
		t.Errorf("expected buffer capacity 200, got %d", cb.Capacity())
	}
}

// TestConcurrentDestroy ensures that the buffer can be safely destroyed concurrently.
func TestConcurrentDestroy(t *testing.T) {
	cb := buffer.New[int]()
	for i := 0; i < 100; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf(errUnexpectedErr, err)
		}
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		cb.Destroy()
	}()

	go func() {
		defer wg.Done()
		cb.Clear()
	}()

	wg.Wait()
	if !cb.IsEmpty() {
		t.Error("expected buffer to be empty after destroy")
	}
	if cb.Capacity() != 0 {
		t.Errorf("expected buffer capacity to be 0 after destroy, got %d", cb.Capacity())
	}
}
