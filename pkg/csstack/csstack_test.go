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
package csstack

import (
	"strings"
	"testing"
)

// TestCSStack tests the CSStack type.
func TestCSStack(t *testing.T) {
	s := CSStackNew[int]()

	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.IsEmpty() {
		t.Errorf("IsEmpty() = true; want false")
	}

	item, err := s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 3 {
		t.Errorf("Pop() = %v; want 3", *item)
	}

	if s.IsEmpty() {
		t.Errorf("IsEmpty() = true; want false")
	}

	item, err = s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 2 {
		t.Errorf("Pop() = %v; want 2", *item)
	}

	item, err = s.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 1 {
		t.Errorf("Pop() = %v; want 1", *item)
	}

	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}

	item, err = s.Pop()
	if err == nil {
		t.Errorf("Pop() error = nil; want error")
	}
	if item != nil {
		t.Errorf("Pop() = %v; want nil", item)
	}
}

// TestCSStackTop tests the CSStack Top method.
func TestCSStackTop(t *testing.T) {
	s := CSStackNew[int]()

	_, err := s.Top()
	if err == nil {
		t.Errorf("Top() error = nil; want error")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	item, err := s.Top()
	if err != nil {
		t.Errorf("Top() error = %v; want nil", err)
	}
	if *item != 3 {
		t.Errorf("Top() = %v; want 3", *item)
	}
}

// TestCSStackPeek tests the CSStack Peek method.
func TestCSStackPeek(t *testing.T) {
	s := CSStackNew[int]()

	_, err := s.Peek()
	if err == nil {
		t.Errorf("Peek() error = nil; want error")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	item, err := s.Peek()
	if err != nil {
		t.Errorf("Peek() error = %v; want nil", err)
	}
	if *item != 3 {
		t.Errorf("Peek() = %v; want 3", *item)
	}
}

// TestCSStackSize tests the CSStack Size method.
func TestCSStackSize(t *testing.T) {
	s := CSStackNew[int]()

	if size := s.Size(); size != 0 {
		t.Errorf("Size() = %v; want 0", size)
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if size := s.Size(); size != 3 {
		t.Errorf("Size() = %v; want 3", size)
	}
}

// TestCSStackClear tests the CSStack Clear method.
func TestCSStackClear(t *testing.T) {
	s := CSStackNew[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	s.Clear()

	if size := s.Size(); size != 0 {
		t.Errorf("Size() = %v; want 0", size)
	}
}

// TestCSStackContains tests the CSStack Contains method.
func TestCSStackContains(t *testing.T) {
	s := CSStackNew[int]()

	if contains := s.Contains(1); contains {
		t.Errorf("Contains() = true; want false")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if contains := s.Contains(2); !contains {
		t.Errorf("Contains() = false; want true")
	}
}

// TestCSStackCopy tests the CSStack Copy method.
func TestCSStackCopy(t *testing.T) {
	s := CSStackNew[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	stack := s.Copy()

	if size := stack.Size(); size != 3 {
		t.Errorf("Size() = %v; want 3", size)
	}

	item, err := stack.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 3 {
		t.Errorf("Pop() = %v; want 3", *item)
	}

	item, err = stack.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 2 {
		t.Errorf("Pop() = %v; want 2", *item)
	}

	item, err = stack.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v; want nil", err)
	}
	if *item != 1 {
		t.Errorf("Pop() = %v; want 1", *item)
	}

	if !stack.IsEmpty() {
		t.Errorf("IsEmpty() = false; want true")
	}
}

// TestCSStackEqual tests the CSStack Equal method.
func TestCSStackEqual(t *testing.T) {
	s1 := CSStackNew[int]()
	s2 := CSStackNew[int]()

	if equal := s1.Equal(s2); !equal {
		t.Errorf("Equal() = false; want true")
	}

	s1.Push(1)
	s1.Push(2)
	s1.Push(3)

	if equal := s1.Equal(s2); equal {
		t.Errorf("Equal() = true; want false")
	}

	s2.Push(1)
	s2.Push(2)
	s2.Push(3)

	if equal := s1.Equal(s2); !equal {
		t.Errorf("Equal() = false; want true")
	}
}

// TestCSStackString tests the CSStack String method.
func TestCSStackString(t *testing.T) {
	s := CSStackNew[int]()

	if str := s.String(); str != "[]" {
		t.Errorf("String() = %v; want []", str)
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if str := s.String(); str != "[1 2 3]" {
		t.Errorf("String() = %v; want [1 2 3]", str)
	}
}

// TestCSStackConcurrent tests the CSStack type in a concurrent environment.
func TestCSStackConcurrent(t *testing.T) {
	s := CSStackNew[int]()

	const n = 100
	const m = 10

	done := make(chan struct{})
	defer close(done)

	for i := 0; i < m; i++ {
		go func() {
			for j := 0; j < n; j++ {
				s.Push(j)
				item, err := s.Pop()
				if err != nil {
					if strings.ToLower(strings.TrimSpace(err.Error())) != "stack is empty" {
						t.Errorf("Pop() error = %v; want nil", err)
					}
				}
				if err == nil && item == nil {
					t.Errorf("Expected non-nil item")
				}
				item, err = s.Top()
				if err != nil {
					if strings.ToLower(strings.TrimSpace(err.Error())) != "stack is empty" {
						t.Errorf("Top() error = %v; want nil", err)
					}
				}
				if err == nil && item == nil {
					t.Errorf("Top() = nil; want item")
				}
				item, err = s.Peek()
				if err != nil {
					if err.Error() != "stack is empty" {
						t.Errorf("Peek() error = %v; want nil", err)
					}
				}
				if err == nil && item == nil {
					t.Errorf("Peek() = nil; want item")
				}
				s.Size()
				s.Clear()
				s.Contains(j)
				s.Copy()
			}
			done <- struct{}{}
		}()
	}

	for i := 0; i < m; i++ {
		<-done
	}
}
