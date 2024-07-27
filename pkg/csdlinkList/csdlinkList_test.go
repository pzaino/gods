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

// Package circularLinkList provides a non-concurrent-safe circular linked list.
package csdlinkList_test

import (
	"sync"
	"testing"

	csdlinkList "github.com/pzaino/gods/pkg/csdlinkList"
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

func TestCSDLinkListAppend(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	runConcurrent(t, 1000, func() {
		cs.Append(1)
	})
	if cs.Size() != 1000 {
		t.Fatalf("expected the size to be 1000, got %d", cs.Size())
	}
}

func TestCSDLinkListPrepend(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	runConcurrent(t, 1000, func() {
		cs.Prepend(1)
	})
	if cs.Size() != 1000 {
		t.Fatalf("expected size to be 1000, got %d", cs.Size())
	}
}

func TestCSDLinkListInsert(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	runConcurrent(t, 1000, func() {
		err := cs.Insert(1)
		if err != nil {
			t.Fatalf("unexpected error: %v ", err)
		}
	})
	if cs.Size() != 1000 {
		t.Fatalf("expected size  1000, got %d", cs.Size())
	}
}

func TestCSDLinkListInsertAfter(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	cs.Append(0)
	runConcurrent(t, 1000, func() {
		cs.InsertAfter(0, 1)
	})
}

func TestCSDLinkListInsertBefore(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		cs.InsertBefore(1, 0)
	})
}

func TestCSDLinkListInsertAt(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	runConcurrent(t, 1000, func() {
		err := cs.InsertAt(0, 1)
		if err != nil && err.Error() != "index out of bounds" {
			t.Fatalf("unexpected error:  %v", err)
		}
	})
}

func TestCSDLinkListDeleteWithValue(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.DeleteWithValue(500)
	})
	if cs.Contains(500) {
		t.Fatalf("expected value 500 to be deleted")
	}
}

func TestCSDLinkListRemoveAt(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		err := cs.RemoveAt(500)
		if err != nil && err.Error() != "index out of bounds" {
			t.Fatalf("unexpected error:  %v ", err)
		}
	})
}

func TestCSDLinkListDelete(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Delete(500)
	})
	if cs.Contains(500) {
		t.Fatalf("expected value 500 to be deleted")
	}
}

func TestCSDLinkListDeleteLast(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.DeleteLast()
	})
	if cs.Size() != 0 {
		t.Fatalf("expected list size 0, got %d", cs.Size())
	}
}

func TestCSDLinkListDeleteFirst(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.DeleteFirst()
	})
	if cs.Size() != 0 {
		t.Fatalf("expected  size 0, got %d", cs.Size())
	}
}

func TestCSDLinkListToSlice(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		test := cs.ToSlice()
		if len(test) != 1000 {
			t.Fatalf("expected size 1000 , got %d", len(test))
		}
	})
}

func TestCSDLinkListToSliceReverse(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		test := cs.ToSliceReverse()
		if len(test) != 1000 {
			t.Fatalf("expected size  1000 , got %d", len(test))
		}
		if test[0] != 999 {
			t.Fatalf("expected first element to be 999, got %d", test[0])
		}
	})
}

func TestCSDLinkListToSliceFrom(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		test := cs.ToSliceFromIndex(50)
		if len(test) != 950 {
			t.Fatalf("expected size 950, got %d", len(test))
		}
		if test[0] != 50 {
			t.Fatalf("expected first element to be 50, got %d", test[0])
		}
	})
}

func TestCSDLinkListToSliceReverseFromIndex(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		test := cs.ToSliceReverseFromIndex(50)
		if len(test) != 950 {
			t.Fatalf("expected size 50, got %d", len(test))
		}
		if test[0] != 949 {
			t.Fatalf("expected first element to be 950, got %d", test[0])
		}
	})
}

func TestCSDLinkListFind(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		_, err := cs.Find(1)
		if err != nil {
			t.Fatalf("unexpected test error: %v", err)
		}
	})
}

func TestCSDLinkListReverse(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Reverse()
	})
}

func TestCSDLinkListSize(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()

	runConcurrent(t, 1000, func() {
		cs.Append(2)
	})

	if cs.Size() != 1000 {
		t.Fatalf("expected the size to be 1000, got %d", cs.Size())
	}
}

func TestCSDLinkListGetFirst(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		cs.GetFirst()
	})
}

func TestCSDLinkListGetLast(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		cs.GetLast()
	})
}

func TestCSDLinkListGetAt(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		_, err := cs.GetAt(500)
		if err != nil {
			t.Fatalf("unexpected test error:  %v", err)
		}
	})
}

func TestCSDLinkListClear(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Clear()
	})
	if cs.Size() != 0 {
		t.Fatalf("expected the size to be 0, got %d", cs.Size())
	}
}

func TestCSDLinkListContains(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		cs.Contains(1)
	})
}

func TestCSDLinkListForEach(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.ForEach(func(item *int) {
			*item = *item + 1
		})
	})
}

func TestCSDLinkListForEachReverse(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 1000; i > 0; i-- {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.ForEachReverse(func(item *int) {
			*item = *item + 1
		})
	})
}

func TestCSDLinkListForRange(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		cs.ForRange(0, 1, func(item *int) {
			*item = *item + 1
		})
	})
}

func TestCSDLinkListForFrom(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		cs.ForFrom(0, func(item *int) {
			*item = *item + 1
		})
	})
}

func TestCSDLinkListForReverseFrom(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		cs.ForReverseFrom(0, func(item *int) {
			*item = *item + 1
		})
	})
}

func TestCSDLinkListForReverseRange(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		cs.ForReverseRange(0, 1, func(item *int) {
			*item = *item + 1
		})
	})
}

func TestCSDLinkListAny(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Any(func(item int) bool {
			return item == 500
		})
	})
}

func TestCSDLinkListAll(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.All(func(item int) bool {
			return item < 1000
		})
	})
}

func TestCSDLinkListIndexOf(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.IndexOf(500)
	})
}

func TestCSDLinkListLastIndexOf(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.LastIndexOf(500)
	})
}

func TestCSDLinkListFilter(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Filter(func(item int) bool {
			return item%2 == 0
		})
	})
}

func TestCSDLinkListMap(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Map(func(item int) int {
			return item * 2
		})
	})
}

func TestCSDLinkListMapFrom(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		cs.MapFrom(0, func(item int) int {
			return item * 2
		})
	})
}

func TestCSDLinkListMapRange(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	cs.Append(1)
	runConcurrent(t, 1000, func() {
		cs.MapRange(0, 1, func(item int) int {
			return item * 2
		})
	})
}

func TestCSDLinkListReduce(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Reduce(func(a, b int) int {
			return a + b
		})
	})
}

func TestCSDLinkListCopy(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	cs.Append(1)
	var copy *csdlinkList.CSDLinkList[int]
	runConcurrent(t, 1000, func() {
		copy = cs.Copy()
	})
	if copy.Size() != cs.Size() {
		t.Fatalf("expected size %d, got %d", cs.Size(), copy.Size())
	}
}

func TestCSDLinkListMerge(t *testing.T) {
	cs1 := csdlinkList.NewCSDLinkList[int]()
	cs2 := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs1.Append(i)
		cs2.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs1.Merge(cs2)
	})
	if cs1.Size() != 2000 {
		t.Fatalf("expected size 2000, got %d", cs1.Size())
	}
	if cs2.Size() != 0 {
		t.Fatalf("expected size 0, got %d", cs2.Size())
	}
}

func TestCSDLinkListReverseCopy(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.ReverseCopy()
	})
}

func TestCSDLinkListReverseMerge(t *testing.T) {
	cs1 := csdlinkList.NewCSDLinkList[int]()
	cs2 := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs1.Append(i)
		cs2.Append(i)
	}
	if cs1.Size() != 1000 {
		t.Fatalf("expected size 1000, got %d", cs1.Size())
	}
	if cs2.Size() != 1000 {
		t.Fatalf("expected size 1000, got %d", cs2.Size())
	}
	runConcurrent(t, 1000, func() {
		cs1.ReverseMerge(cs2)
	})
	if cs1.Size() != 2000 {
		t.Fatalf("expected size 2000, got %d", cs1.Size())
	}
	if cs2.Size() != 0 {
		t.Fatalf("expected size 0, got %d", cs2.Size())
	}
}

func TestCSDLinkListEqual(t *testing.T) {
	cs1 := csdlinkList.NewCSDLinkList[int]()
	cs2 := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs1.Append(i)
		cs2.Append(i)
	}
	runConcurrent(t, 1000, func() {
		if !cs1.Equal(cs2) {
			t.Fatalf("expected lists to be equal")
		}
	})
}

func TestCSDLinkListSwap(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		err := cs.Swap(0, 999)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestCSDLinkListSort(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 1000; i > 0; i-- {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.Sort(func(a, b int) bool { return a < b })
	})
}

func TestCSDLinkListFindAll(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.FindAll(func(item int) bool {
			return item%2 == 0
		})
	})
}

func TestCSDLinkListFindLast(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		_, err := cs.FindLast(func(item int) bool {
			return item == 500
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestCSDLinkListFindLastIndex(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.FindLastIndex(func(item int) bool {
			return item == 500
		})
	})
}

func TestCSDLinkListFindIndex(t *testing.T) {
	cs := csdlinkList.NewCSDLinkList[int]()
	for i := 0; i < 1000; i++ {
		cs.Append(i)
	}
	runConcurrent(t, 1000, func() {
		cs.FindIndex(func(item int) bool {
			return item == 500
		})
	})
}
