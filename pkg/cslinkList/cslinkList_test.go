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

// Package cslinkList provides a concurrent-safe linked list.
package cslinkList_test

import (
	"sync"
	"testing"

	cslinkList "github.com/pzaino/gods/pkg/cslinkList"
)

const (
	errExpectedNoError = "expected no error, got %v"
	errExpectedSizeX   = "expected size %d, got %d"
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

func TestCSLinkListFromSlice(t *testing.T) {
	cs := cslinkList.NewCSLinkListFromSlice[int]([]int{1, 2, 3, 4, 5})
	if cs.Size() != 5 {
		t.Fatalf(errExpectedSizeX, 5, cs.Size())
	}
}

func TestCSLinkListAppend(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	runConcurrent(t, 1000, func() {
		cs.Append(1)
	})
	if cs.Size() != 1000 {
		t.Fatalf(errExpectedSizeX, 1000, cs.Size())
	}
}

func TestCSLinkListPrepend(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	runConcurrent(t, 1000, func() {
		cs.Prepend(1)
	})
	if cs.Size() != 1000 {
		t.Fatalf(errExpectedSizeX, 1000, cs.Size())
	}
}

func TestCSLinkListDeleteWithValue(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 100, func() {
		cs.DeleteWithValue(500)
	})
	if cs.Contains(500) {
		t.Fatalf("expected value 500 to be deleted")
	}
}

func TestCSLinkListRemove(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i * 2)
	}
	runConcurrent(t, 100, func() {
		cs.Remove(500)
	})
	if cs.Contains(500) {
		t.Fatalf("expected value 500 to be deleted")
	}
}

func TestCSLinkListToSlice(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.ToSlice()
	})
}

func TestCSLinkListIsEmpty(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	if !cs.IsEmpty() {
		t.Fatalf("expected list to be empty")
	}
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		cs.IsEmpty()
	})
}

func TestCSLinkListFind(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		_, err := cs.Find(1)
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}

func TestCSLinkListReverse(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Reverse()
	})
}

func TestCSLinkListSize(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	runConcurrent(t, 1000, func() {
		cs.Append(2)
	})
	if cs.Size() != 1000 {
		t.Fatalf(errExpectedSizeX, 1000, cs.Size())
	}
}

func TestCSLinkListGetFirst(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		cs.GetFirst()
	})
}

func TestCSLinkListGetLast(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		cs.GetLast()
	})
}

func TestCSLinkListGetAt(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		_, err := cs.GetAt(500)
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}

func TestCSLinkListInsertAt(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		err := cs.InsertAt(500, 999)
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}

func TestCSLinkListDeleteAt(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		err := cs.DeleteAt(500)
		if err != nil && err.Error() != "index out of bounds" {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}

func TestCSLinkListClear(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Clear()
	})
	if cs.Size() != 0 {
		t.Fatalf(errExpectedSizeX, 0, cs.Size())
	}
}

func TestCSLinkListCopy(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	cs.Append(1)
	var copy *cslinkList.CSLinkList[int]
	runConcurrent(t, 1000, func() {
		copy = cs.Copy()
	})
	if copy.Size() != cs.Size() {
		t.Fatalf(errExpectedSizeX, cs.Size(), copy.Size())
	}
}

func TestCSLinkListMerge(t *testing.T) {
	cs1 := cslinkList.NewCSLinkList[int]()
	cs2 := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs1.Append(i)
		cs2.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs1.Merge(cs2)
	})
	if cs1.Size() != 2000 {
		t.Fatalf(errExpectedSizeX, 2000, cs1.Size())
	}
	if cs2.Size() != 0 {
		t.Fatalf(errExpectedSizeX, 0, cs2.Size())
	}
}

func TestCSLinkListMap(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Map(func(item int) int {
			return item * 2
		})
	})
}

func TestCSLinkListFilter(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Filter(func(item int) bool {
			return item%2 == 0
		})
	})
}

func TestCSLinkListReduce(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Reduce(func(a, b int) int {
			return a + b
		}, 0)
	})
}

func TestCSLinkListForEach(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.ForEach(func(item *int) {
			*item = *item + 1
		})
	})
}

func TestCSLinkListForRange(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
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

func TestCSLinkListForFrom(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
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

func TestCSLinkListAny(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Any(func(item int) bool {
			return item == 500
		})
	})
}

func TestCSLinkListAll(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.All(func(item int) bool {
			return item < 1000
		})
	})
}

func TestCSLinkListContains(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		cs.Contains(1)
	})
}

func TestCSLinkListIndexOf(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		_, err := cs.IndexOf(500)
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}

func TestCSLinkListLastIndexOf(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		_, err := cs.LastIndexOf(500)
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}

func TestCSLinkListFindIndex(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
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

func TestCSLinkListFindLastIndex(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
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

func TestCSLinkListFindAll(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.FindAll(func(item int) bool {
			return item%2 == 0
		})
	})
}

func TestCSLinkListFindLast(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
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

func TestCSLinkListFindAllIndexes(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.FindAllIndexes(func(item int) bool {
			return item%2 == 0
		})
	})
}

func TestCSLinkListMapFrom(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		_, err := cs.MapFrom(500, func(item int) int {
			return item * 2
		})
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}

func TestCSLinkListMapRange(t *testing.T) {
	cs := cslinkList.NewCSLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		_, err := cs.MapRange(0, 500, func(item int) int {
			return item * 2
		})
		if err != nil {
			t.Fatalf(errExpectedNoError, err)
		}
	})
}
