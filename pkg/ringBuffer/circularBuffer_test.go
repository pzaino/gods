package ringBuffer_test

import (
	"testing"

	cBuf "github.com/pzaino/gods/pkg/ringBuffer"
)

const (
	//errCircularBufferEmpty = "ring buffer is empty"
	//errExpectedError       = "Expected error, got %v"
	errExpectedNoError = "Expected no error, got %v"
)

func TestNewCircularBuffer(t *testing.T) {
	buffer := cBuf.New[byte](4)

	if buffer.Size() != 0 {
		t.Errorf("Expected buffer size to be 0, got %d", buffer.Size())
	}

	if buffer.Capacity() != 4 {
		t.Errorf("Expected buffer capacity to be 4, got %d", buffer.Capacity())
	}
}

func TestAppendAndRemove(t *testing.T) {
	buffer := cBuf.New[byte](4)

	// Append elements
	buffer.Append(1)
	buffer.Append(2)
	buffer.Append(3)

	if buffer.Size() != 3 {
		t.Errorf("Expected buffer size to be 3, got %d", buffer.Size())
	}

	// Remove elements
	val, err := buffer.Remove()
	if err != nil {
		t.Fatalf(errExpectedNoError, err)
	}
	if val != 1 {
		t.Errorf("Expected removed value to be 1, got %d", val)
	}

	val, err = buffer.Remove()
	if err != nil {
		t.Fatalf(errExpectedNoError, err)
	}
	if val != 2 {
		t.Errorf("Expected removed value to be 2, got %d", val)
	}

	val, err = buffer.Remove()
	if err != nil {
		t.Fatalf(errExpectedNoError, err)
	}
	if val != 3 {
		t.Errorf("Expected removed value to be 3, got %d", val)
	}

	// Test underflow
	_, err = buffer.Remove()
	if err == nil {
		t.Errorf("Expected error on removing from empty buffer")
	}
}

func TestOverwriteOnFullBuffer(t *testing.T) {
	buffer := cBuf.New[byte](3)

	buffer.Append(1)
	buffer.Append(2)
	buffer.Append(3)

	// Buffer is full, next append should overwrite the oldest element
	buffer.Append(4)

	if buffer.Size() != 3 {
		t.Errorf("Expected buffer size to be 3, got %d", buffer.Size())
	}

	val, err := buffer.Remove()
	if err != nil {
		t.Fatalf(errExpectedNoError, err)
	}
	if val != 2 {
		t.Errorf("Expected removed value to be 2, got %d", val)
	}

	buffer.Append(5)
	val, err = buffer.Remove()
	if err != nil {
		t.Fatalf(errExpectedNoError, err)
	}
	if val != 3 {
		t.Errorf("Expected removed value to be 3, got %d", val)
	}

	val, err = buffer.Remove()
	if err != nil {
		t.Fatalf(errExpectedNoError, err)
	}
	if val != 4 {
		t.Errorf("Expected removed value to be 4, got %d", val)
	}
}

func TestGet(t *testing.T) {
	buffer := cBuf.New[byte](4)

	buffer.Append(1)
	buffer.Append(2)
	buffer.Append(3)

	val, err := buffer.Get(0)
	if err != nil {
		t.Fatalf(errExpectedNoError, err)
	}
	if val != 1 {
		t.Errorf("Expected value to be 1, got %d", val)
	}

	val, err = buffer.Get(1)
	if err != nil {
		t.Fatalf(errExpectedNoError, err)
	}
	if val != 2 {
		t.Errorf("Expected value to be 2, got %d", val)
	}

	val, err = buffer.Get(2)
	if err != nil {
		t.Fatalf(errExpectedNoError, err)
	}
	if val != 3 {
		t.Errorf("Expected value to be 3, got %d", val)
	}

	// Test out of range
	_, err = buffer.Get(3)
	if err == nil {
		t.Errorf("Expected error when accessing out of range index")
	}
}

func TestToSlice(t *testing.T) {
	buffer := cBuf.New[int](3)

	buffer.Append(1)
	buffer.Append(2)
	buffer.Append(3)

	slice := buffer.ToSlice()
	expected := []int{1, 2, 3}

	for i, val := range slice {
		if val != expected[i] {
			t.Errorf("Expected slice[%d] to be %d, got %d", i, expected[i], val)
		}
	}

	// Overwrite an element and test again
	buffer.Append(4)
	slice = buffer.ToSlice()
	expected = []int{2, 3, 4}

	for i, val := range slice {
		if val != expected[i] {
			t.Errorf("Expected slice[%d] to be %d, got %d", i, expected[i], val)
		}
	}
}

func TestIsEmptyAndIsFull(t *testing.T) {
	buffer := cBuf.New[int](3)

	if !buffer.IsEmpty() {
		t.Errorf("Expected buffer to be empty")
	}

	buffer.Append(1)
	buffer.Append(2)
	buffer.Append(3)

	if !buffer.IsFull() {
		t.Errorf("Expected buffer to be full")
	}

	_, err := buffer.Remove()
	if err != nil {
		t.Fatalf(errExpectedNoError, err)
	}
	if buffer.IsFull() {
		t.Errorf("Expected buffer to not be full after removing an element")
	}

	_, err = buffer.Remove()
	if err != nil {
		t.Fatalf(errExpectedNoError, err)
	}
	_, err = buffer.Remove()
	if err != nil {
		t.Fatalf(errExpectedNoError, err)
	}
	if !buffer.IsEmpty() {
		t.Errorf("Expected buffer to be empty after removing all elements")
	}
}

func TestClear(t *testing.T) {
	buffer := cBuf.New[int](4)

	buffer.Append(1)
	buffer.Append(2)
	buffer.Append(3)

	buffer.Clear()

	if !buffer.IsEmpty() {
		t.Errorf("Expected buffer to be empty after clear")
	}

	if buffer.Size() != 0 {
		t.Errorf("Expected buffer size to be 0 after clear, got %d", buffer.Size())
	}
}

func TestForEach(t *testing.T) {
	buffer := cBuf.New[int](4)

	buffer.Append(1)
	buffer.Append(2)
	buffer.Append(3)

	var sum int
	buffer.ForEach(func(val int) {
		sum += val
	})

	expectedSum := 1 + 2 + 3
	if sum != expectedSum {
		t.Errorf("Expected sum to be %d, got %d", expectedSum, sum)
	}
}

func TestContains(t *testing.T) {
	buffer := cBuf.New[int](3)

	buffer.Append(1)
	buffer.Append(2)
	buffer.Append(3)

	if !buffer.Contains(2) {
		t.Errorf("Expected buffer to contain value 2")
	}

	if buffer.Contains(4) {
		t.Errorf("Expected buffer to not contain value 4")
	}

	buffer.Append(4)

	if !buffer.Contains(4) {
		t.Errorf("Expected buffer to contain value 4 after append")
	}

	if buffer.Contains(1) {
		t.Errorf("Expected buffer to not contain value 1 after overwrite")
	}
}
