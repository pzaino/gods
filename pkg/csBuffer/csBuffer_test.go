// Description: This file contains tests for the concurrent buffer implementation.
package csBuffer_test

import (
	"sync"
	"testing"

	buffer "github.com/pzaino/gods/pkg/csBuffer"
)

const (
	errUnexpectedErr = "unexpected error during append: %v"
	errExpectedVal   = "expected value %d, got %d"
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
				t.Errorf(errExpectedVal, i, val)
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

// TestConcurrentForFrom tests the ForFrom method of ConcurrentBuffer.
func TestConcurrentForFrom(t *testing.T) {
	cb := buffer.New[int]()
	const numElements = 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf(errUnexpectedErr, err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < numElements; i++ {
		wg.Add(1)
		if i%2 == 0 {
			go incrementElements(t, &wg, cb)
		} else {
			go decrementElements(t, &wg, cb)
		}
	}

	wg.Wait()

	for i := 0; i < numElements; i++ {
		val, err := cb.Get(uint64(i))
		if err != nil {
			t.Errorf("unexpected error during Get: %v", err)
		}
		expectedVal := i
		if val != expectedVal {
			t.Errorf(errExpectedVal, expectedVal, val)
		}
	}
}

func incrementElements(t *testing.T, wg *sync.WaitGroup, cb *buffer.ConcurrentBuffer[int]) {
	defer wg.Done()
	err := cb.ForFrom(uint64(0), increment)
	if err != nil {
		t.Errorf("unexpected error during ForFrom: %v", err)
	}
}

func decrementElements(t *testing.T, wg *sync.WaitGroup, cb *buffer.ConcurrentBuffer[int]) {
	defer wg.Done()
	err := cb.ForFrom(uint64(0), decrement)
	if err != nil {
		t.Errorf("unexpected error during ForFrom: %v", err)
	}
}

func increment(elem *int) error {
	*elem = *elem + 1
	return nil
}

func decrement(elem *int) error {
	*elem = *elem - 1
	return nil
}

// TestConcurrentForRange tests the ForRange method of ConcurrentBuffer.
func TestConcurrentForRange(t *testing.T) {
	cb := buffer.New[int]()
	const numElements = 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf(errUnexpectedErr, err)
		}
	}

	var wg sync.WaitGroup
	const start = uint64(20)
	const end = uint64(80)
	for i := start; i < end; i++ {
		wg.Add(1)
		go func(i uint64) {
			defer wg.Done()
			err := cb.ForRange(start, end, func(elem *int) error {
				*elem = *elem + 1
				return nil
			})
			if err != nil {
				t.Errorf("unexpected error during ForRange: %v", err)
			}
		}(i)
	}

	wg.Wait()

	for i := start; i < end; i++ {
		val, err := cb.Get(i)
		if err != nil {
			t.Errorf("unexpected error during Get: %v", err)
		}
		expectedVal := int(i) + 60
		if val != expectedVal {
			t.Errorf(errExpectedVal, expectedVal, val)
		}
	}
}

// TestConcurrentInsertAt tests concurrent inserts at specific indices.
func TestConcurrentInsertAt(t *testing.T) {
	cb := buffer.New[int]()
	initialSize := 10

	// Initialize the buffer with some data to allow InsertAt to function properly
	for i := 0; i < initialSize; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf("unexpected error during initial append: %v", err)
		}
	}

	var wg sync.WaitGroup
	numGoroutines := 5
	numInsertsPerGoroutine := 2

	// Each goroutine inserts elements at specific calculated indices
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(goroutineIndex int) {
			defer wg.Done()
			for j := 0; j < numInsertsPerGoroutine; j++ {
				// Insert each element at the beginning to ensure no overlap
				index := uint64(goroutineIndex*(initialSize/numGoroutines) + j)
				err := cb.InsertAt(index, -1*(goroutineIndex*numInsertsPerGoroutine+j+1))
				if err != nil {
					t.Errorf("unexpected error during insert at: %v", err)
				}
			}
		}(i)
	}

	wg.Wait()

	// Verify the buffer size after all insertions
	expectedSize := uint64(initialSize + numGoroutines*numInsertsPerGoroutine)
	if cb.Size() != expectedSize {
		t.Errorf("expected buffer size %d, got %d", expectedSize, cb.Size())
	}

	// Create a map to verify that all expected negative values are present
	insertedValues := map[int]bool{
		-1: false, -2: false, -3: false, -4: false, -5: false,
		-6: false, -7: false, -8: false, -9: false, -10: false,
	}

	// Log buffer contents for debugging purposes
	for i := 0; i < int(expectedSize); i++ {
		val, err := cb.Get(uint64(i))
		if err != nil {
			t.Errorf("unexpected error during get: %v", err)
		}
		t.Logf("Value at index %d: %d", i, val)

		// Check if the value is one of the inserted ones
		if _, exists := insertedValues[val]; exists {
			insertedValues[val] = true
		}
	}

	// Verify that all inserted values are present in the buffer
	for val, found := range insertedValues {
		if !found {
			t.Errorf("expected value %d to be inserted but it was not found in the buffer", val)
		}
	}
}

// TestConcurrentPut tests concurrent replacements at specific indices.
func TestConcurrentPut(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < numElements; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := cb.Put(uint64(i), i+1)
			if err != nil {
				t.Errorf("unexpected error during put: %v", err)
			}
		}(i)
	}

	wg.Wait()
	for i := 0; i < numElements; i++ {
		val, err := cb.Get(uint64(i))
		if err != nil {
			t.Errorf("unexpected error during get: %v", err)
		}
		expectedVal := i + 1
		if val != expectedVal {
			t.Errorf(errExpectedVal, expectedVal, val)
		}
	}
}

// TestConcurrentClear tests concurrent clear operations on the buffer.
func TestConcurrentClear(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cb.Clear()
		}()
	}

	wg.Wait()
	if !cb.IsEmpty() {
		t.Error("expected buffer to be empty after clear")
	}
}

// TestConcurrentReverse tests concurrent reversals of the buffer.
func TestConcurrentReverse(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cb.Reverse()
		}()
	}

	wg.Wait()
	// Check if the final state is correctly reversed (or could end up the same if even number of reversals).
	for i := 0; i < numElements; i++ {
		val, err := cb.Get(uint64(i))
		if err != nil {
			t.Errorf("unexpected error during get: %v", err)
		}
		// Since reversals may end in the original order due to even number of reversals,
		// check either possible final state
		if val != i && val != numElements-1-i {
			t.Errorf("expected value %d or %d, got %d", i, numElements-1-i, val)
		}
	}
}

// TestConcurrentContains tests the Contains method under concurrent access.
func TestConcurrentContains(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < numElements; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if !cb.Contains(i) {
				t.Errorf("expected buffer to contain %d", i)
			}
		}(i)
	}

	wg.Wait()
}

// TestConcurrentCopy tests the Copy method under concurrent access.
func TestConcurrentCopy(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			copy := cb.Copy()
			if copy.Size() != cb.Size() {
				t.Errorf("expected copy size %d, got %d", cb.Size(), copy.Size())
			}
			for j := 0; j < int(copy.Size()); j++ {
				val, err := copy.Get(uint64(j))
				if err != nil {
					t.Errorf("unexpected error during get: %v", err)
				}
				originalVal, err := cb.Get(uint64(j))
				if err != nil {
					t.Errorf("unexpected error during get: %v", err)
				}
				if val != originalVal {
					t.Errorf("expected value %d, got %d", originalVal, val)
				}
			}
		}()
	}

	wg.Wait()
}

// TestConcurrentFilter tests the Filter method under concurrent access.
func TestConcurrentFilter(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cb.Filter(func(val int) bool {
				return val%2 == 0 // Keep even numbers
			})
		}()
	}

	wg.Wait()
	if cb.Size() != 50 {
		t.Errorf("expected buffer size 50 after filtering, got %d", cb.Size())
	}
}

// TestConcurrentAny tests the Any method under concurrent access.
func TestConcurrentAny(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !cb.Any(func(val int) bool {
				return val == 50
			}) {
				t.Error("expected buffer to contain 50")
			}
		}()
	}

	wg.Wait()
}

// TestConcurrentAll tests the All method under concurrent access.
func TestConcurrentAll(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(1) // All elements are 1
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !cb.All(func(val int) bool {
				return val == 1
			}) {
				t.Error("expected all elements to be 1")
			}
		}()
	}

	wg.Wait()
}

// TestConcurrentRotateLeft tests concurrent left rotations of the buffer.
func TestConcurrentRotateLeft(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 10
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	rotations := 3
	for i := 0; i < rotations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cb.RotateLeft(1)
		}()
	}

	wg.Wait()

	// Verify rotation effect
	for i := 0; i < numElements; i++ {
		val, err := cb.Get(uint64(i))
		if err != nil {
			t.Errorf("unexpected error during get: %v", err)
		}
		expectedVal := (i + rotations) % numElements
		if val != expectedVal {
			t.Errorf(errExpectedVal, expectedVal, val)
		}
	}
}

// TestConcurrentRotateRight tests concurrent right rotations of the buffer.
func TestConcurrentRotateRight(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 10
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	rotations := 3
	for i := 0; i < rotations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cb.RotateRight(1)
		}()
	}

	wg.Wait()

	// Verify rotation effect
	for i := 0; i < numElements; i++ {
		val, err := cb.Get(uint64(i))
		if err != nil {
			t.Errorf("unexpected error during get: %v", err)
		}
		expectedVal := (i - rotations + numElements) % numElements
		if val != expectedVal {
			t.Errorf(errExpectedVal, expectedVal, val)
		}
	}
}

// TestConcurrentMap tests concurrent map operations.
func TestConcurrentMap(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mappedBuffer, err := cb.Map(func(val int) int {
				return val * 2
			})
			if err != nil {
				t.Errorf("unexpected error during map: %v", err)
			}
			if mappedBuffer.Size() != cb.Size() {
				t.Errorf("expected mapped buffer size %d, got %d", cb.Size(), mappedBuffer.Size())
			}
			for j := 0; j < int(mappedBuffer.Size()); j++ {
				val, err := mappedBuffer.Get(uint64(j))
				if err != nil {
					t.Errorf("unexpected error during get: %v", err)
				}
				expectedVal := j * 2
				if val != expectedVal {
					t.Errorf(errExpectedVal, expectedVal, val)
				}
			}
		}()
	}

	wg.Wait()
}

// TestConcurrentReduce tests concurrent reduce operations.
func TestConcurrentReduce(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(1)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, err := cb.Reduce(func(acc, val int) int {
				return acc + val
			})
			if err != nil {
				t.Errorf("unexpected error during reduce: %v", err)
			}
			expectedResult := numElements
			if result != expectedResult {
				t.Errorf(errExpectedVal, expectedResult, result)
			}
		}()
	}

	wg.Wait()
}

// TestConcurrentShiftLeft tests concurrent left shifts of the buffer.
func TestConcurrentShiftLeft(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 10
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	shifts := 3
	for i := 0; i < shifts; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cb.ShiftLeft(1)
		}()
	}

	wg.Wait()

	// Verify shift effect
	for i := 0; i < numElements-shifts; i++ {
		val, err := cb.Get(uint64(i))
		if err != nil {
			t.Errorf("unexpected error during get: %v", err)
		}
		expectedVal := i + shifts
		if val != expectedVal {
			t.Errorf(errExpectedVal, expectedVal, val)
		}
	}
}

// TestConcurrentShiftRight tests concurrent right shifts of the buffer.
func TestConcurrentShiftRight(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 10
	for i := 0; i < numElements; i++ {
		err := cb.Append(i)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	shifts := 3
	for i := 0; i < shifts; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cb.ShiftRight(1)
		}()
	}

	wg.Wait()

	// Verify shift effect
	for i := shifts; i < numElements; i++ {
		val, err := cb.Get(uint64(i))
		if err != nil {
			t.Errorf("unexpected error during get: %v", err)
		}
		expectedVal := i - shifts
		if val != expectedVal {
			t.Errorf(errExpectedVal, expectedVal, val)
		}
	}
}

// TestConcurrentBlit tests the Blit function under concurrent access.
func TestConcurrentBlit(t *testing.T) {
	cb1 := buffer.New[int]()
	cb2 := buffer.New[int]()
	numElements := 100
	for i := 0; i < numElements; i++ {
		err := cb1.Append(i)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
		err = cb2.Append(i * 2)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Blit (binary combines) the two buffers (cb1 = cb1 OR cb2)
			err := cb1.Blit(cb2, func(a, b int) int {
				return a | b
			})
			if err != nil {
				t.Errorf("unexpected error during blit: %v", err)
			}
		}()
	}

	wg.Wait()

	// Verify blit results
	for i := 0; i < int(cb1.Size()); i++ {
		val, err := cb1.Get(uint64(i))
		if err != nil {
			t.Errorf("unexpected error during get: %v", err)
		}
		expectedVal := i | (i * 2)
		if val != expectedVal {
			t.Errorf(errExpectedVal, expectedVal, val)
		}
	}
}

// TestConcurrentFindAll tests the FindAll method under concurrent access.
func TestConcurrentFindAll(t *testing.T) {
	cb := buffer.New[int]()
	numElements := 100
	for i := 0; i < numElements; i++ {
		err := cb.Append(i % 2)
		if err != nil {
			t.Errorf("unexpected error during append: %v", err)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newBuffer := cb.FindAll(func(val int) bool {
				return val == 1
			})
			if newBuffer.Size() != 50 {
				t.Errorf("expected new buffer size 50, got %d", newBuffer.Size())
			}
		}()
	}

	wg.Wait()
}
