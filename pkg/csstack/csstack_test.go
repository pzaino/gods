// Copyright 2024 Paolo Fabio Zaino
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package csstack is a concurrent-safe stack library (LIFO).
package csstack_test

import (
	"sync"
	"testing"

	csstack "github.com/pzaino/gods/pkg/csstack"
)

const (
	errExpectedStackEmpty = "expected stack to be empty"
	errExpectedNoError    = "expected no error, got %v"
	errExpectedSizeX      = "expected size %d, got %d"
)

func runConcurrent(_ *testing.T, n int, fn func()) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn()
		}()
	}
	wg.Wait()
}

func TestCSStackPush(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	runConcurrent(t, 1000, func() {
		cs.Push(1)
	})
	if cs.Size() != 1000 {
		t.Fatalf(errExpectedSizeX, 1000, cs.Size())
	}
}

func TestCSStackIsEmpty(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	if !cs.IsEmpty() {
		t.Fatalf(errExpectedStackEmpty)
	}
	cs.Push(1)
	runConcurrent(t, 1000, func() {
		cs.IsEmpty()
	})
}

func TestCSStackPop(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		_, err := cs.Pop()
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
	if !cs.IsEmpty() {
		t.Fatalf(errExpectedStackEmpty)
	}
}

func TestCSStackToSlice(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		cs.ToSlice()
	})
}

func TestCSStackToStack(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		cs.ToStack()
	})
}

func TestCSStackReverse(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Reverse()
	})
}

func TestCSStackSwap(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	cs.Push(1)
	cs.Push(2)
	runConcurrent(t, 1000, func() {
		err := cs.Swap()
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}

func TestCSStackTop(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	cs.Push(1)
	runConcurrent(t, 1000, func() {
		top, err := cs.Top()
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
		if *top != 1 {
			t.Fatalf("expected top to be 1, got %d", *top)
		}
	})
}

func TestCSStackPeek(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	cs.Push(1)
	runConcurrent(t, 1000, func() {
		peek, err := cs.Peek()
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
		if *peek != 1 {
			t.Fatalf("expected peek to be 1, got %d", *peek)
		}
	})
}

func TestCSStackSize(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	runConcurrent(t, 1000, func() {
		cs.Push(2)
	})
	if cs.Size() != 1000 {
		t.Fatalf(errExpectedSizeX, 1000, cs.Size())
	}
}

func TestCSStackClear(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Clear()
	})
	if cs.Size() != 0 {
		t.Fatalf(errExpectedSizeX, 0, cs.Size())
	}
}

func TestCSStackContains(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	cs.Push(1)
	runConcurrent(t, 1000, func() {
		cs.Contains(1)
	})
}

func TestCSStackCopy(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	cs.Push(1)
	var copy *csstack.CSStack[int]
	runConcurrent(t, 1000, func() {
		copy = cs.Copy()
	})
	if copy.Size() != cs.Size() {
		t.Fatalf(errExpectedSizeX, cs.Size(), copy.Size())
	}
}

func TestCSStackEqual(t *testing.T) {
	cs1 := csstack.NewCSStack[int]()
	cs2 := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs1.Push(i)
		cs2.Push(i)
	}
	runConcurrent(t, 1000, func() {
		if !cs1.Equal(cs2) {
			t.Fatalf("expected stacks to be equal")
		}
	})
}

func TestCSStackString(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	cs.Push(1)
	runConcurrent(t, 1000, func() {
		test := cs.String()
		if test == "" || test == "[]" {
			t.Fatalf("expected string representation of the stack")
		}
	})
}

func TestCSStackPopN(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	if cs.Size() != 1000 {
		t.Fatalf(errExpectedSizeX, 1000, cs.Size())
	}

	runConcurrent(t, 100, func() { // Reduce the number of goroutines to avoid exhausting the stack too quickly
		_, err := cs.PopN(10)
		if err != nil && err.Error() != "Stack has less than n items" {
			t.Fatalf(errExpectedNoError, err)
		}
	})
	if cs.Size() != 0 {
		t.Fatalf(errExpectedSizeX, 0, cs.Size())
	}
}

func TestCSStackPushN(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	runConcurrent(t, 1000, func() {
		cs.PushN(1, 2, 3, 4, 5)
	})
	if cs.Size() != 5000 {
		t.Fatalf(errExpectedSizeX, 5000, cs.Size())
	}
}

func TestCSStackPopAll(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		cs.PopAll()
	})
	if !cs.IsEmpty() {
		t.Fatalf(errExpectedStackEmpty)
	}
}

func TestCSStackPushAll(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	items := []int{1, 2, 3, 4, 5}
	runConcurrent(t, 1000, func() {
		cs.PushAll(items)
	})
	if cs.Size() != 5000 {
		t.Fatalf(errExpectedSizeX, 5000, cs.Size())
	}
}

func TestCSStackFilter(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Filter(func(item int) bool {
			return item%2 == 0
		})
	})
	if cs.Size() != 500 {
		t.Fatalf(errExpectedSizeX, 500, cs.Size())
	}
}

func TestCSStackMap(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Map(func(item int) int {
			return item * 2
		})
	})
}

func TestCSStackReduce(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		_, err := cs.Reduce(func(a, b int) int {
			return a + b
		})
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}

func TestCSStackForEach(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		cs.ForEach(func(item *int) {
			*item = *item + 1
		})
	})
}

func TestCSStackForRange(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		err := cs.ForRange(0, 500, func(item *int) {
			*item = *item + 1
		})
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}

func TestCSStackForFrom(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		err := cs.ForFrom(500, func(item *int) {
			*item = *item + 1
		})
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}

func TestCSStackAny(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Any(func(item int) bool {
			return item == 500
		})
	})
}

func TestCSStackAll(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		cs.All(func(item int) bool {
			return item < 1000
		})
	})
}

func TestCSStackFind(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		_, err := cs.Find(func(item int) bool {
			return item == 500
		})
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}

func TestCSStackFindIndex(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		_, err := cs.FindIndex(func(item int) bool {
			return item == 500
		})
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}

func TestCSStackFindLast(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		_, err := cs.FindLast(func(item int) bool {
			return item == 500
		})
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}

func TestCSStackFindLastIndex(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		_, err := cs.FindLastIndex(func(item int) bool {
			return item == 500
		})
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}

func TestCSStackFindAll(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		cs.FindAll(func(item int) bool {
			return item%2 == 0
		})
	})
}

func TestCSStackFindIndices(t *testing.T) {
	cs := csstack.NewCSStack[int]()
	for i := 0; i < 1000; i++ {
		cs.Push(i)
	}
	runConcurrent(t, 1000, func() {
		cs.FindIndices(func(item int) bool {
			return item%2 == 0
		})
	})
}

func TestNewCSStackFromSlice(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	cs := csstack.NewCSStackFromSlice(items)
	if cs.Size() != len(items) {
		t.Fatalf(errExpectedSizeX, len(items), cs.Size())
	}
	slice := cs.ToSlice()
	for i, item := range items {
		if slice[i] != item {
			t.Fatalf("expected item %d to be %d, got %d", i, item, slice[i])
		}
	}
}
