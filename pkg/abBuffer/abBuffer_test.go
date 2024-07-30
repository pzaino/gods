package abBuffer_test

import (
	"testing"

	"github.com/pzaino/gods/pkg/abBuffer"
)

func TestNewBuffer(t *testing.T) {
	capacity := uint64(10)
	buf := abBuffer.New[int](capacity)
	if buf.Capacity() != capacity {
		t.Errorf("expected %d, got %d", capacity, buf.Capacity())
	}
	if buf.Size() != 0 {
		t.Errorf("expected 0, got %d", buf.Size())
	}
	if !buf.IsEmpty() {
		t.Error("expected buffer to be empty")
	}
}

func TestAppend(t *testing.T) {
	buf := abBuffer.New[int](3)
	err := buf.Append(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !equal(buf.GetActive(), []int{1}) {
		t.Errorf("expected [1], got %v", buf.GetActive())
	}
	if buf.Size() != 1 {
		t.Errorf("expected size 1, got %d", buf.Size())
	}

	err = buf.Append(2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !equal(buf.GetActive(), []int{1, 2}) {
		t.Errorf("expected [1, 2], got %v", buf.GetActive())
	}

	err = buf.Append(3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !equal(buf.GetActive(), []int{1, 2, 3}) {
		t.Errorf("expected [1, 2, 3], got %v", buf.GetActive())
	}

	err = buf.Append(4)
	if err == nil || err.Error() != "buffer overflow" {
		t.Errorf("expected buffer overflow error, got %v", err)
	}
}

func TestClear(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	buf.Clear()
	if !buf.IsEmpty() {
		t.Error("expected buffer to be empty")
	}
	if !equal(buf.GetActive(), []int{}) {
		t.Errorf("expected [], got %v", buf.GetActive())
	}
}

func TestSwap(t *testing.T) {
	buf := abBuffer.New[int](1)
	err := buf.Append(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	buf.Swap()

	if !buf.IsEmpty() {
		t.Error("expected buffer to be empty")
	}
	if !equal(buf.GetActive(), []int{}) {
		t.Errorf("expected [], got %v", buf.GetActive())
	}
	if !equal(buf.GetInactive(), []int{1}) {
		t.Errorf("expected [1], got %v", buf.GetInactive())
	}
}

func TestGetActive(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	if !equal(buf.GetActive(), []int{1}) {
		t.Errorf("expected [1], got %v", buf.GetActive())
	}
}

func TestGetInactive(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	buf.Swap()
	if !equal(buf.GetInactive(), []int{1}) {
		t.Errorf("expected [1], got %v", buf.GetInactive())
	}
}

func TestSize(t *testing.T) {
	buf := abBuffer.New[int](16)
	if buf.Size() != 0 {
		t.Errorf("expected size 0, got %d", buf.Size())
	}
	_ = buf.Append(1)
	if buf.Size() != 1 {
		t.Errorf("expected size 1, got %d", buf.Size())
	}
}

func TestCapacity(t *testing.T) {
	capacity := uint64(3)
	buf := abBuffer.New[int](capacity)
	if buf.Capacity() != capacity {
		t.Errorf("expected %d, got %d", capacity, buf.Capacity())
	}
}

func TestIsEmpty(t *testing.T) {
	buf := abBuffer.New[int](16)
	if !buf.IsEmpty() {
		t.Error("expected buffer to be empty")
	}
	_ = buf.Append(1)
	if buf.IsEmpty() {
		t.Error("expected buffer to be non-empty")
	}
}

func TestToSlice(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	slice := buf.ToSlice()
	if !equal(slice, []int{1, 2}) {
		t.Errorf("expected [1, 2], got %v", slice)
	}
}

func TestFind(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	index, err := buf.Find(2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if index != 1 {
		t.Errorf("expected 1, got %d", index)
	}

	_, err = buf.Find(3)
	if err == nil || err.Error() != "value not found" {
		t.Errorf("expected 'value not found' error, got %v", err)
	}
}

func TestRemove(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	err := buf.Remove(0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !equal(buf.GetActive(), []int{2}) {
		t.Errorf("expected [2], got %v", buf.GetActive())
	}

	err = buf.Remove(2)
	if err == nil || err.Error() != "value not found" {
		t.Errorf("expected 'value not found' error, got %v", err)
	}
}

func TestInsertAt(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	err := buf.InsertAt(0, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !equal(buf.GetActive(), []int{2, 1}) {
		t.Errorf("expected [2, 1], got %v", buf.GetActive())
	}

	err = buf.InsertAt(3, 3)
	if err == nil || err.Error() != "buffer overflow" {
		t.Errorf("expected invalid buffer error, got %v", err)
	}
}

func TestForEach(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	var sum int
	buf.ForEach(func(v *int) {
		sum += *v
	})
	if sum != 3 {
		t.Errorf("expected sum 3, got %d", sum)
	}
}

func TestForFrom(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	var sum int
	err := buf.ForFrom(1, func(v *int) {
		sum += *v
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if sum != 2 {
		t.Errorf("expected sum 2, got %d", sum)
	}

	err = buf.ForFrom(2, func(v *int) {
		sum += *v
	})
	if err == nil || err.Error() != "invalid buffer" {
		t.Errorf("expected invalid buffer error, got %v", err)
	}
}

func TestForRange(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	var sum int
	err := buf.ForRange(0, 2, func(v *int) {
		sum += *v
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if sum != 3 {
		t.Errorf("expected sum 3, got %d", sum)
	}

	err = buf.ForRange(0, 3, func(v *int) {
		sum += *v
	})
	if err == nil || err.Error() != "invalid buffer" {
		t.Errorf("expected invalid buffer error, got %v", err)
	}
}

func TestMap(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	newBuf, err := buf.Map(func(v int) int {
		return v * 2
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !equal(newBuf.GetActive(), []int{2, 4}) {
		t.Errorf("expected [2, 4], got %v", newBuf.GetActive())
	}
}

func TestMapFrom(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	newBuf, err := buf.MapFrom(1, func(v int) int {
		return v * 2
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !equal(newBuf.GetActive(), []int{4}) {
		t.Errorf("expected [4], got %v", newBuf.GetActive())
	}

	_, err = buf.MapFrom(3, func(v int) int {
		return v * 2
	})
	if err == nil || err.Error() != "invalid buffer" {
		t.Errorf("expected invalid buffer error, got %v", err)
	}
}

func TestMapRange(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	newBuf, err := buf.MapRange(0, 2, func(v int) int {
		return v * 2
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !equal(newBuf.GetActive(), []int{2, 4}) {
		t.Errorf("expected [2, 4], got %v", newBuf.GetActive())
	}

	_, err = buf.MapRange(0, 3, func(v int) int {
		return v * 2
	})
	if err == nil || err.Error() != "invalid buffer" {
		t.Errorf("expected invalid buffer error, got %v", err)
	}
}

func TestFilter(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	buf.Filter(func(v int) bool {
		return v > 1
	})
	if !equal(buf.GetActive(), []int{2}) {
		t.Errorf("expected [2], got %v", buf.GetActive())
	}
}

func TestReduce(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	result, err := buf.Reduce(func(a, b int) int {
		return a + b
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 3 {
		t.Errorf("expected 3, got %d", result)
	}
}

func TestReduceFrom(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	result, err := buf.ReduceFrom(1, func(a, b int) int {
		return a + b
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 2 {
		t.Errorf("expected 2, got %d", result)
	}

	_, err = buf.ReduceFrom(2, func(a, b int) int {
		return a + b
	})
	if err == nil || err.Error() != "invalid buffer" {
		t.Errorf("expected invalid buffer error, got %v", err)
	}
}

func TestReduceRange(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	result, err := buf.ReduceRange(0, 2, func(a, b int) int {
		return a + b
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 3 {
		t.Errorf("expected 3, got %d", result)
	}

	_, err = buf.ReduceRange(0, 3, func(a, b int) int {
		return a + b
	})
	if err == nil || err.Error() != "invalid buffer" {
		t.Errorf("expected invalid buffer error, got %v", err)
	}
}

func TestContains(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	if !buf.Contains(1) {
		t.Error("expected buffer to contain 1")
	}
	if buf.Contains(2) {
		t.Error("expected buffer not to contain 2")
	}
}

func TestAny(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	if !buf.Any(func(v int) bool {
		return v == 1
	}) {
		t.Error("expected buffer to contain an element equal to 1")
	}
	if buf.Any(func(v int) bool {
		return v == 2
	}) {
		t.Error("expected buffer not to contain an element equal to 2")
	}
}

func TestAll(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	if !buf.All(func(v int) bool {
		return v > 0
	}) {
		t.Error("expected all elements to be greater than 0")
	}
	if buf.All(func(v int) bool {
		return v > 1
	}) {
		t.Error("expected not all elements to be greater than 1")
	}
}

func TestLastIndexOf(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(1)
	index, err := buf.LastIndexOf(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if index != 1 {
		t.Errorf("expected 1, got %d", index)
	}

	_, err = buf.LastIndexOf(2)
	if err == nil || err.Error() != "value not found" {
		t.Errorf("expected value not found error, got %v", err)
	}
}

func TestCopy(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	newBuf := buf.Copy()
	if !equal(buf.GetActive(), newBuf.GetActive()) {
		t.Errorf("expected %v, got %v", buf.GetActive(), newBuf.GetActive())
	}
	if !equal(buf.GetInactive(), newBuf.GetInactive()) {
		t.Errorf("expected %v, got %v", buf.GetInactive(), newBuf.GetInactive())
	}
}

func TestCopyInactive(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	buf.Swap()
	_ = buf.Append(1)
	newBuf := buf.CopyInactive()
	if !equal(buf.GetInactive(), newBuf.GetActive()) {
		t.Errorf("expected %v, got %v", buf.GetInactive(), newBuf.GetActive())
	}
}

func TestMerge(t *testing.T) {
	buf1 := abBuffer.New[int](16)
	buf2 := abBuffer.New[int](16)
	_ = buf1.Append(1)
	_ = buf2.Append(2)
	buf1.Merge(buf2)
	if !equal(buf1.GetActive(), []int{1, 2}) {
		t.Errorf("expected [1, 2], got %v", buf1.GetActive())
	}

	buf3 := abBuffer.New[int](16)
	buf1.Merge(buf3)
	if !equal(buf1.GetActive(), []int{1, 2}) {
		t.Errorf("expected [1, 2], got %v", buf1.GetActive())
	}
}

func TestBlit(t *testing.T) {
	buf1 := abBuffer.New[int](16)
	buf2 := abBuffer.New[int](16)
	_ = buf1.Append(1)
	_ = buf2.Append(2)
	err := buf1.Blit(buf2, func(a, b int) int {
		return a + b
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !equal(buf1.GetActive(), []int{3}) {
		t.Errorf("expected [3], got %v", buf1.GetActive())
	}

	buf3 := abBuffer.New[int](16)
	err = buf1.Blit(buf3, func(a, b int) int {
		return a + b
	})
	if err != nil {
		t.Errorf("unexpected error, got %v", err)
	}
	if !equal(buf1.GetActive(), []int{3}) {
		t.Errorf("expected [3], got %v", buf1.GetActive())
	}
}

// Helper function to compare slices
func equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSetActiveA(t *testing.T) {
	buf := abBuffer.New[int](16)
	err := buf.Append(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	buf.SetActiveB()
	buf.SetActiveA()
	expected := []int{1}
	if !equal(buf.GetActive(), expected) {
		t.Errorf("expected %v, got %v", expected, buf.GetActive())
	}
}

func TestToSliceInactive(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	inactive := buf.ToSliceInactive()
	if !equal(inactive, []int{}) {
		t.Errorf("expected [], got %v", inactive)
	}

	buf.Swap()
	inactive = buf.ToSliceInactive()
	if !equal(inactive, []int{1, 2}) {
		t.Errorf("expected [1, 2], got %v", inactive)
	}

	// Clear active buffer (so inactive buffer is not cleared)
	buf.Clear()
	inactive = buf.ToSliceInactive()
	if !equal(inactive, []int{1, 2}) {
		t.Errorf("expected [], got %v", inactive)
	}
}

func TestClearAll(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	buf.ClearAll()
	if !buf.IsEmpty() {
		t.Error("expected buffer to be empty")
	}
	if !equal(buf.GetActive(), []int{}) {
		t.Errorf("expected [], got %v", buf.GetActive())
	}
	if !equal(buf.GetInactive(), []int{}) {
		t.Errorf("expected [], got %v", buf.GetInactive())
	}
}

func TestDestroy(t *testing.T) {
	buf := abBuffer.New[int](16)
	_ = buf.Append(1)
	_ = buf.Append(2)
	buf.Destroy()
	if buf.Size() != 0 {
		t.Errorf("expected size 0, got %d", buf.Size())
	}
	if buf.Capacity() != 0 {
		t.Errorf("expected capacity 0, got %d", buf.Capacity())
	}
	if !buf.IsEmpty() {
		t.Error("expected buffer to be empty")
	}
	if buf.GetActive() != nil {
		t.Error("expected active buffer to be nil")
	}
}
