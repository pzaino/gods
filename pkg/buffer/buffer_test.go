package buffer_test

import (
	"testing"

	buffer "github.com/pzaino/gods/pkg/buffer"
)

const (
	errUnexpectedErr  = "unexpected error: %v"
	errExpectedLength = "expected length %v, got %v"
	errExpectedErr    = "expected error %v, got %v"
	errExpectedValue  = "expected value %v, got %v"
	errBlitterErr     = "Blitter should not return an error, got %v"
)

// Helper function to create a buffer with given elements and capacity
func createBufferWithElements(t *testing.T, elements []int, capacity uint64) *buffer.Buffer[int] {
	b := buffer.NewBuffer[int]()
	b.SetCapacity(capacity)
	for _, elem := range elements {
		err := b.Append(elem)
		if err != nil {
			t.Errorf(errUnexpectedErr, err)
		}
	}
	return b
}

// TestNewBuffer tests the creation of a new buffer
func TestNewBuffer(t *testing.T) {
	b := buffer.NewBuffer[int]()
	if b == nil {
		t.Error("NewBuffer should not return nil")
	}
	if !b.IsEmpty() {
		t.Error("NewBuffer should create an empty buffer")
	}
}

// TestIsEmpty tests the IsEmpty method
func TestIsEmpty(t *testing.T) {
	b := buffer.NewBuffer[int]()
	if !b.IsEmpty() {
		t.Error("IsEmpty should return true for an empty buffer")
	}
	err := b.Append(1)
	if err != nil {
		t.Errorf(errUnexpectedErr, err)
	}
	if b.IsEmpty() {
		t.Error("IsEmpty should return false for a non-empty buffer")
	}
}

// TestIsFull tests the IsFull method
func TestIsFull(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	if !b.IsFull() {
		t.Error("IsFull should return true when buffer is full")
	}
	b = createBufferWithElements(t, []int{1, 2}, 3)
	if b.IsFull() {
		t.Error("IsFull should return false when buffer is not full")
	}
}

// TestAppend tests the Append method
func TestAppend(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2}, 3)
	err := b.Append(3)
	if err != nil {
		t.Errorf(errUnexpectedErr, err)
	}
	if !b.IsFull() {
		t.Error("Buffer should be full after appending to its capacity")
	}
	err = b.Append(4)
	if err == nil {
		t.Error("Append should return an error when the buffer is full")
	}
	if err.Error() != buffer.ErrBufferOverflow {
		t.Errorf(errExpectedErr, buffer.ErrBufferOverflow, err)
	}
}

// TestGet tests the Get method
func TestGet(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	elem, err := b.Get(1)
	if err != nil {
		t.Errorf("Get should not return an error, got %v", err)
	}
	if elem != 2 {
		t.Errorf("Expected element 2, got %v", elem)
	}
	_, err = b.Get(3)
	if err == nil {
		t.Error("Get should return an error for an out-of-bounds index")
	}
	if err.Error() != buffer.ErrValueNotFound {
		t.Errorf(errExpectedErr, buffer.ErrValueNotFound, err)
	}
}

// TestSet tests the Set method
func TestSet(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	err := b.Set(1, 4)
	if err != nil {
		t.Errorf("Set should not return an error, got %v", err)
	}
	elem, _ := b.Get(1)
	if elem != 4 {
		t.Errorf("Expected element 4, got %v", elem)
	}
	err = b.Set(3, 5)
	if err == nil {
		t.Error("Set should return an error for an out-of-bounds index")
	}
	if err.Error() != buffer.ErrValueNotFound {
		t.Errorf(errExpectedErr, buffer.ErrValueNotFound, err)
	}
}

// TestRemove tests the Remove method
func TestRemove(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	err := b.Remove(1)
	if err != nil {
		t.Errorf("Remove should not return an error, got %v", err)
	}
	if b.Size() != 2 {
		t.Errorf("Expected size 2, got %v", b.Size())
	}
	elem, _ := b.Get(1)
	if elem != 3 {
		t.Errorf("Expected element 3, got %v", elem)
	}
	err = b.Remove(2)
	if err == nil {
		t.Error("Remove should return an error for an out-of-bounds index")
	}
	if err.Error() != buffer.ErrValueNotFound {
		t.Errorf(errExpectedErr, buffer.ErrValueNotFound, err)
	}
}

// TestClear tests the Clear method
func TestClear(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	b.Clear()
	if !b.IsEmpty() {
		t.Error("Clear should empty the buffer")
	}
}

// TestValues tests the Values method
func TestValues(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	values := b.Values()
	if len(values) != 3 {
		t.Errorf("Expected length 3, got %v", len(values))
	}
	for i, v := range values {
		if v != i+1 {
			t.Errorf(errExpectedValue, i+1, v)
		}
	}
}

// TestSize tests the Size method
func TestSize(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	if b.Size() != 3 {
		t.Errorf("Expected size 3, got %v", b.Size())
	}
}

// TestCapacity tests the Capacity method
func TestCapacity(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	if b.Capacity() != 3 {
		t.Errorf("Expected capacity 3, got %v", b.Capacity())
	}
}

// TestSetCapacity tests the SetCapacity method
func TestSetCapacity(t *testing.T) {
	b := buffer.NewBuffer[int]()
	b.SetCapacity(5)
	if b.Capacity() != 5 {
		t.Errorf("Expected capacity 5, got %v", b.Capacity())
	}
}

// TestEquals tests the Equals method
func TestEquals(t *testing.T) {
	b1 := createBufferWithElements(t, []int{1, 2, 3}, 3)
	b2 := createBufferWithElements(t, []int{1, 2, 3}, 3)
	if !b1.Equals(b2) {
		t.Error("Buffers should be equal")
	}
	b3 := createBufferWithElements(t, []int{1, 2}, 3)
	if b1.Equals(b3) {
		t.Error("Buffers should not be equal")
	}
}

// TestToSlice tests the ToSlice method
func TestToSlice(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	slice := b.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected length 3, got %v", len(slice))
	}
	for i, v := range slice {
		if v != i+1 {
			t.Errorf(errExpectedValue, i+1, v)
		}
	}
}

// TestReverse tests the Reverse method
func TestReverse(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	b.Reverse()
	expected := []int{3, 2, 1}
	for i, v := range b.Values() {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}
}

// TestFind tests the Find method
func TestFind(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	index, err := b.Find(2)
	if err != nil {
		t.Errorf("Find should not return an error, got %v", err)
	}
	if index != 1 {
		t.Errorf("Expected index 1, got %v", index)
	}
	_, err = b.Find(4)
	if err == nil {
		t.Error("Find should return an error for a non-existent value")
	}
	if err.Error() != buffer.ErrValueNotFound {
		t.Errorf(errExpectedErr, buffer.ErrValueNotFound, err)
	}
}

// TestContains tests the Contains method
func TestContains(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	if !b.Contains(2) {
		t.Error("Buffer should contain value 2")
	}
	if b.Contains(4) {
		t.Error("Buffer should not contain value 4")
	}
}

// TestCopy tests the Copy method
func TestCopy(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 4)
	copy := b.Copy()
	if !b.Equals(copy) {
		t.Error("Copy should create an equal buffer")
	}
	err := copy.Append(4)
	if err != nil {
		t.Errorf(errUnexpectedErr, err)
	}
	if b.Equals(copy) {
		t.Error("Modifying copy should not affect original buffer")
	}
}

// TestMerge tests the Merge method
func TestMerge(t *testing.T) {
	b1 := createBufferWithElements(t, []int{1, 2}, 3)
	b2 := createBufferWithElements(t, []int{3}, 3)
	b1.Merge(b2)
	if b1.Size() != 3 {
		t.Errorf("Expected size 3, got %v", b1.Size())
	}
	if !b1.Contains(3) {
		t.Error("Buffer should contain merged elements")
	}
	if !b2.IsEmpty() {
		t.Error("Merged buffer should be empty")
	}
}

// TestPopN tests the PopN method
func TestPopN(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	values, err := b.PopN(2)
	if err != nil {
		t.Errorf("PopN should not return an error, got %v", err)
	}
	expected := []int{2, 3}
	for i, v := range values {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}
	if b.Size() != 1 {
		t.Errorf("Expected size 1, got %v", b.Size())
	}
	_, err = b.PopN(2)
	if err == nil {
		t.Error("PopN should return an error when popping more elements than present")
	}
	if err.Error() != buffer.ErrBufferEmpty {
		t.Errorf(errExpectedErr, buffer.ErrBufferEmpty, err)
	}
}

// TestPushN tests the PushN method
func TestPushN(t *testing.T) {
	b := createBufferWithElements(t, []int{1}, 3)
	err := b.PushN(2, 3)
	if err != nil {
		t.Errorf("PushN should not return an error, got %v", err)
	}
	if !b.IsFull() {
		t.Error("Buffer should be full after pushing elements to its capacity")
	}
	err = b.PushN(4)
	if err == nil {
		t.Error("PushN should return an error when pushing beyond capacity")
	}
	if err.Error() != buffer.ErrBufferOverflow {
		t.Errorf(errExpectedErr, buffer.ErrBufferOverflow, err)
	}
}

// TestShiftLeft tests the ShiftLeft method
func TestShiftLeft(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	b.ShiftLeft(2)
	expected := []int{3, 0, 0}
	for i, v := range b.Values() {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}
}

// TestShiftRight tests the ShiftRight method
func TestShiftRight(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	b.ShiftRight(2)
	expected := []int{0, 0, 1}
	for i, v := range b.Values() {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}
}

// TestFilter tests the Filter method
func TestFilter(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3, 4, 5}, 5)
	b.Filter(func(x int) bool {
		return x%2 == 0
	})
	expected := []int{2, 4}
	values := b.Values()
	if len(values) != len(expected) {
		t.Errorf(errExpectedLength, len(expected), len(values))
	}
	for i, v := range values {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}
}

// TestMap tests the Map method
func TestMap(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	newB, err := b.Map(func(x int) int {
		return x * 2
	})
	if err != nil {
		t.Errorf("Map should not return an error, got %v", err)
	}

	expected := []int{2, 4, 6}
	values := newB.Values()
	if len(values) != len(expected) {
		t.Errorf(errExpectedLength, len(expected), len(values))
	}
	for i, v := range values {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}
}

// TestMapFrom tests the MapFrom method
func TestMapFrom(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	newB, err := b.MapFrom(1, func(x int) int {
		return x * 2
	})
	if err != nil {
		t.Errorf("MapFrom should not return an error, got %v", err)
	}
	expected := []int{4, 6}
	values := newB.Values()
	if len(values) != len(expected) {
		t.Errorf(errExpectedLength, len(expected), len(values))
	}
	for i, v := range values {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}
	_, err = b.MapFrom(3, func(x int) int {
		return x * 2
	})
	if err == nil {
		t.Error("MapFrom should return an error for an out-of-bounds index")
	}
	if err.Error() != buffer.ErrInvalidBuffer {
		t.Errorf(errExpectedErr, buffer.ErrInvalidBuffer, err)
	}
}

// TestMapRange tests the MapRange method
func TestMapRange(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	newB, err := b.MapRange(1, 3, func(x int) int {
		return x * 2
	})
	if err != nil {
		t.Errorf("MapRange should not return an error, got %v", err)
	}
	expected := []int{4, 6}
	values := newB.Values()
	if len(values) != len(expected) {
		t.Errorf(errExpectedLength, len(expected), len(values))
	}
	for i, v := range values {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}
	_, err = b.MapRange(3, 3, func(x int) int {
		return x * 2
	})
	if err == nil {
		t.Error("MapRange should return an error for an out-of-bounds index")
	}
	if err.Error() != buffer.ErrInvalidBuffer {
		t.Errorf(errExpectedErr, buffer.ErrInvalidBuffer, err)
	}
}

// TestReduce tests the Reduce method
func TestReduce(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	result, err := b.Reduce(func(x, y int) int {
		return x + y
	})
	if err != nil {
		t.Errorf("Reduce should not return an error, got %v", err)
	}
	if result != 6 {
		t.Errorf("Expected result 6, got %v", result)
	}
	b.Clear()
	_, err = b.Reduce(func(x, y int) int {
		return x + y
	})
	if err == nil {
		t.Error("Reduce should return an error for an empty buffer")
	}
	if err.Error() != buffer.ErrBufferEmpty {
		t.Errorf(errExpectedErr, buffer.ErrBufferEmpty, err)
	}
}

// TestReduceFrom tests the ReduceFrom method
func TestReduceFrom(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	result, err := b.ReduceFrom(1, func(x, y int) int {
		return x + y
	})
	if err != nil {
		t.Errorf("ReduceFrom should not return an error, got %v", err)
	}
	if result != 5 {
		t.Errorf("Expected result 5, got %v", result)
	}
	_, err = b.ReduceFrom(3, func(x, y int) int {
		return x + y
	})
	if err == nil {
		t.Error("ReduceFrom should return an error for an out-of-bounds index")
	}
	if err.Error() != buffer.ErrInvalidBuffer {
		t.Errorf(errExpectedErr, buffer.ErrInvalidBuffer, err)
	}
}

// TestReduceRange tests the ReduceRange method
func TestReduceRange(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	result, err := b.ReduceRange(1, 3, func(x, y int) int {
		return x + y
	})
	if err != nil {
		t.Errorf("ReduceRange should not return an error, got %v", err)
	}
	if result != 5 {
		t.Errorf("Expected result 5, got %v", result)
	}
	_, err = b.ReduceRange(3, 3, func(x, y int) int {
		return x + y
	})
	if err == nil {
		t.Error("ReduceRange should return an error for an out-of-bounds index")
	}
	if err.Error() != buffer.ErrInvalidBuffer {
		t.Errorf(errExpectedErr, buffer.ErrInvalidBuffer, err)
	}
}

// TestForEach tests the ForEach method
func TestForEach(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	b.ForEach(func(x *int) {
		*x *= 2
	})
	expected := []int{2, 4, 6}
	values := b.Values()
	for i, v := range values {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}
}

// TestForRange tests the ForRange method
func TestForRange(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	err := b.ForRange(1, 3, func(x *int) {
		*x *= 2
	})
	if err != nil {
		t.Errorf("ForRange should not return an error, got %v", err)
	}
	expected := []int{1, 4, 6}
	values := b.Values()
	for i, v := range values {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}
	err = b.ForRange(3, 3, func(x *int) {
		*x *= 2
	})
	if err == nil {
		t.Error("ForRange should return an error for an out-of-bounds index")
	}
	if err.Error() != buffer.ErrInvalidBuffer {
		t.Errorf(errExpectedErr, buffer.ErrInvalidBuffer, err)
	}
}

// TestForFrom tests the ForFrom method
func TestForFrom(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	err := b.ForFrom(1, func(x *int) {
		*x *= 2
	})
	if err != nil {
		t.Errorf("ForFrom should not return an error, got %v", err)
	}
	expected := []int{1, 4, 6}
	values := b.Values()
	for i, v := range values {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}
	err = b.ForFrom(3, func(x *int) {
		*x *= 2
	})
	if err == nil {
		t.Error("ForFrom should return an error for an out-of-bounds index")
	}
	if err.Error() != buffer.ErrInvalidBuffer {
		t.Errorf(errExpectedErr, buffer.ErrInvalidBuffer, err)
	}
}

// TestAny tests the Any method
func TestAny(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	if !b.Any(func(x int) bool {
		return x == 2
	}) {
		t.Error("Any should return true if any element matches the predicate")
	}
	if b.Any(func(x int) bool {
		return x == 4
	}) {
		t.Error("Any should return false if no element matches the predicate")
	}
}

// TestAll tests the All method
func TestAll(t *testing.T) {
	b := createBufferWithElements(t, []int{2, 4, 6}, 0)
	if !b.All(func(x int) bool {
		return x%2 == 0
	}) {
		t.Error("All should return true if all elements match the predicate")
	}
	err := b.Append(3)
	if err != nil {
		t.Errorf(errUnexpectedErr, err)
	}

	if b.All(func(x int) bool {
		return x%2 == 0
	}) {
		t.Error("All should return false if any element does not match the predicate")
	}
}

// TestFindIndex tests the FindIndex method
func TestFindIndex(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3}, 3)
	index, err := b.FindIndex(func(x int) bool {
		return x == 2
	})
	if err != nil {
		t.Errorf("FindIndex should not return an error, got %v", err)
	}
	if index != 1 {
		t.Errorf("Expected index 1, got %v", index)
	}
	_, err = b.FindIndex(func(x int) bool {
		return x == 4
	})
	if err == nil {
		t.Error("FindIndex should return an error for a non-existent value")
	}
	if err.Error() != buffer.ErrValueNotFound {
		t.Errorf(errExpectedErr, buffer.ErrValueNotFound, err)
	}
}

// TestFindLast tests the FindLast method
func TestFindLast(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3, 2}, 4)
	value, err := b.FindLast(func(x int) bool {
		return x == 2
	})
	if err != nil {
		t.Errorf("FindLast should not return an error, got %v", err)
	}
	if *value != 2 {
		t.Errorf("Expected value 2, got %v", *value)
	}
	_, err = b.FindLast(func(x int) bool {
		return x == 4
	})
	if err == nil {
		t.Error("FindLast should return an error for a non-existent value")
	}
	if err.Error() != buffer.ErrValueNotFound {
		t.Errorf(errExpectedErr, buffer.ErrValueNotFound, err)
	}
}

// TestFindLastIndex tests the FindLastIndex method
func TestFindLastIndex(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3, 2}, 4)
	index, err := b.FindLastIndex(func(x int) bool {
		return x == 2
	})
	if err != nil {
		t.Errorf("FindLastIndex should not return an error, got %v", err)
	}
	if index != 3 {
		t.Errorf("Expected index 3, got %v", index)
	}
	_, err = b.FindLastIndex(func(x int) bool {
		return x == 4
	})
	if err == nil {
		t.Error("FindLastIndex should return an error for a non-existent value")
	}
	if err.Error() != buffer.ErrValueNotFound {
		t.Errorf(errExpectedErr, buffer.ErrValueNotFound, err)
	}
}

// TestFindAll tests the FindAll method
func TestFindAll(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3, 2}, 4)
	newB := b.FindAll(func(x int) bool {
		return x == 2
	})
	expected := []int{2, 2}
	values := newB.Values()
	if len(values) != len(expected) {
		t.Errorf(errExpectedLength, len(expected), len(values))
	}
	for i, v := range values {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}
}

// TestFindIndices tests the FindIndices method
func TestFindIndices(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3, 2}, 4)
	indices := b.FindIndices(func(x int) bool {
		return x == 2
	})
	expected := []uint64{1, 3}
	if len(indices) != len(expected) {
		t.Errorf(errExpectedLength, len(expected), len(indices))
	}
	for i, v := range indices {
		if v != expected[i] {
			t.Errorf("Expected index %v, got %v", expected[i], v)
		}
	}
}

// TestLastIndexOf tests the LastIndexOf method
func TestLastIndexOf(t *testing.T) {
	b := createBufferWithElements(t, []int{1, 2, 3, 2}, 4)
	index, err := b.LastIndexOf(2)
	if err != nil {
		t.Errorf("LastIndexOf should not return an error, got %v", err)
	}
	if index != 3 {
		t.Errorf("Expected index 3, got %v", index)
	}
	_, err = b.LastIndexOf(4)
	if err == nil {
		t.Error("LastIndexOf should return an error for a non-existent value")
	}
	if err.Error() != buffer.ErrValueNotFound {
		t.Errorf(errExpectedErr, buffer.ErrValueNotFound, err)
	}
}

// TestBlit tests the Blitter method
func TestBlit(t *testing.T) {
	b1 := createBufferWithElements(t, []int{1, 2, 3}, 3)
	b2 := createBufferWithElements(t, []int{4, 5, 6}, 3)

	err := b1.Blit(b2, func(a, b int) int {
		return a + b
	})
	if err != nil {
		t.Errorf(errBlitterErr, err)
	}

	expected := []int{5, 7, 9}
	for i, v := range b1.Values() {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}

	// Test with different buffer sizes
	b3 := createBufferWithElements(t, []int{1, 2, 3, 4}, 4)
	b4 := createBufferWithElements(t, []int{5, 6, 7}, 3)

	err = b3.Blit(b4, func(a, b int) int {
		return a + b
	})
	if err != nil {
		t.Errorf(errBlitterErr, err)
	}

	expected = []int{6, 8, 10, 4}
	for i, v := range b3.Values() {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}

	// Test with different buffer capacities
	b5 := createBufferWithElements(t, []int{1, 2, 3}, 3)
	b6 := createBufferWithElements(t, []int{4, 5, 6, 7}, 4)

	err = b5.Blit(b6, func(a, b int) int {
		return a + b
	})
	if err != nil {
		t.Errorf(errBlitterErr, err)
	}

	expected = []int{5, 7, 9}
	for i, v := range b5.Values() {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}

	// Test with different buffer sizes and capacities
	b7 := createBufferWithElements(t, []int{1, 2, 3, 4}, 4)
	b8 := createBufferWithElements(t, []int{5, 6, 7, 8}, 4)

	err = b7.Blit(b8, func(a, b int) int {
		return a + b
	})
	if err != nil {
		t.Errorf(errBlitterErr, err)
	}

	expected = []int{6, 8, 10, 12}
	for i, v := range b7.Values() {
		if v != expected[i] {
			t.Errorf(errExpectedValue, expected[i], v)
		}
	}
}
