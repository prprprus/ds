// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package linkedliststack

import "testing"

func verifyElements(result []interface{}, stack *Stack) bool {
	values := stack.Values()
	for i, v := range result {
		if values[i] != v {
			return false
		}
	}
	return true
}

func TestNew(t *testing.T) {
	stack := New()
	if stack.Size() != 0 {
		t.Error("TestNew error")
	}
}

// Stack Interface

func TestPush(t *testing.T) {
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	result := []interface{}{3, 2, 1}
	if !verifyElements(result, stack) {
		t.Error("TestPush error")
	}
}

func TestPop(t *testing.T) {
	// case1: stack has some elements
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	v, err := stack.Pop()
	if err != nil || v.(int) != 3 {
		t.Error("case1 error: stack has some elements")
	}
	v, err = stack.Pop()
	if err != nil || v.(int) != 2 {
		t.Error("case1 error: stack has some elements")
	}
	v, err = stack.Pop()
	if err != nil || v.(int) != 1 {
		t.Error("case1 error: stack has some elements")
	}

	// case2: stack has one element
	stack = New()
	stack.Push(1)
	v, err = stack.Pop()
	if v.(int) != 1 || err != nil {
		t.Error("case2 error: stack has one element")
	}
	v, err = stack.Pop()
	if v != nil || err == nil {
		t.Error("case2 error: stack has one element")
	}

	// case3: stack has no element
	stack = New()
	v, err = stack.Pop()
	if v != nil || err == nil {
		t.Error("case3 error: stack has no element")
	}
}

func TestPeek(t *testing.T) {
	// case1: stack has some elements
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	v, err := stack.Peek()
	if v.(int) != 3 || err != nil {
		t.Error("case1 error: stack has some elements")
	}
	v, err = stack.Peek()
	if v.(int) != 3 || err != nil {
		t.Error("case1 error: stack has some elements")
	}

	// case2: stack has one element
	stack = New()
	stack.Push(1)
	v, err = stack.Peek()
	if v.(int) != 1 || err != nil {
		t.Error("case1 error: stack has one element")
	}
	v, err = stack.Peek()
	if v.(int) != 1 || err != nil {
		t.Error("case1 error: stack has one element")
	}

	// case3: stack has no element
	stack = New()
	v, err = stack.Peek()
	if v != nil || err == nil {
		t.Error("case3 error: stack has no element")
	}
}

func TestIndexInRange(t *testing.T) {
	// case1: stack has some elements, check with first element index
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if !stack.indexInRange(0) {
		t.Error("case1 error: stack has some elements, check with first element index")
	}

	// case2: stack has some elements, check with last element index
	stack = New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if !stack.indexInRange(2) {
		t.Error("case2 error: stack has some elements, check with last element index")
	}

	// case3: stack has some elements, check with medium element index
	stack = New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if !stack.indexInRange(1) {
		t.Error("case3 error: stack has some elements, check with medium element index")
	}

	// case4: stack has some elements. check with not exists index
	stack = New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.indexInRange(-1) {
		t.Error("case4 error: stack has some elements. check with not exists index")
	}
	if stack.indexInRange(3) {
		t.Error("case4 error: stack has some elements. check with not exists index")
	}

	// case5: stack has one element
	stack = New()
	stack.Push(1)
	if !stack.indexInRange(0) {
		t.Error("case5 error: stack has one element")
	}
	if stack.indexInRange(1) {
		t.Error("case5 error: stack has one element")
	}

	// case6: stack has no element
	stack = New()
	if stack.indexInRange(0) {
		t.Error("case6 error: stack has no element")
	}
}

// Container Interface

func TestEmpty(t *testing.T) {
	// case1: stack has some elements
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Empty() {
		t.Error("case1 error: stack has some elements")
	}

	// case2: stack has one element
	stack = New()
	stack.Push(1)
	if stack.Empty() {
		t.Error("case2 error: stack has one element")
	}

	// case3: stack has no element
	stack = New()
	if !stack.Empty() {
		t.Error("case3 error: stack has no element")
	}
}

func TestSize(t *testing.T) {
	// case1: stack has some elements
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Error("case1 error: stack has some elements")
	}

	// case2: stack has one element
	stack = New()
	stack.Push(1)
	if stack.Size() != 1 {
		t.Error("case2 error: stack has one element")
	}

	// case3: stack has no element
	stack = New()
	if stack.Size() != 0 {
		t.Error("case3 error: stack has no element")
	}
}

func TestClear(t *testing.T) {
	// case1: stack has some elements
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Clear()
	if stack.Size() != 0 {
		t.Error("case1 error: stack has some elements")
	}

	// case2: stack has one element
	stack = New()
	stack.Push(1)
	stack.Clear()
	if stack.Size() != 0 {
		t.Error("case2 error: stack has one element")
	}

	// case3: stack has no element
	stack = New()
	stack.Clear()
	if stack.Size() != 0 {
		t.Error("case3 error: stack has no element")
	}
}

func TestValues(t *testing.T) {
	// case1: stack has some elements
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	result := []interface{}{3, 2, 1}
	if !verifyElements(result, stack) {
		t.Error("case1 error: stack has some elements")
	}

	// case2: stack has one element
	stack = New()
	stack.Push(1)
	result = []interface{}{1}
	if !verifyElements(result, stack) {
		t.Error("case2 error: stack has one element")
	}

	// case3: stack has no element
	stack = New()
	values := stack.Values()
	if len(values) != 0 {
		t.Error("case3 error: stack has no element")
	}
}

// Iterator Interface

func TestIterator(t *testing.T) {
	// case1: stack has some elements
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	iterator := stack.Iterator()
	if iterator.index != -1 {
		t.Error("case1 error: stack has some elements")
	}

	// case2: stack has one element
	stack = New()
	stack.Push(1)
	iterator = stack.Iterator()
	if iterator.index != -1 {
		t.Error("case2 error: stack has one element")
	}

	// case3: stack has no element
	stack = New()
	iterator = stack.Iterator()
	if iterator.index != -1 {
		t.Error("case3 error: stack has no element")
	}
}

func TestNextAndValue(t *testing.T) {
	// case1: stack has some elements
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	iterator := stack.Iterator()
	result := []interface{}{3, 2, 1}
	i := 0
	for iterator.Next() {
		if iterator.Value() != result[i] {
			t.Error("case1 error: stack has some elements")
		}
		i++
	}

	// case2: stack has one element
	stack = New()
	stack.Push(1)
	iterator = stack.Iterator()
	result = []interface{}{1}
	i = 0
	for iterator.Next() {
		if iterator.Value() != result[i] {
			t.Error("case2 error: stack has one element")
		}
		i++
	}

	// case3: stack has no element
	stack = New()
	iterator = stack.Iterator()
	if iterator.Next() {
		t.Error("case3 error: stack has no element")
	}
}

func TestBegin(t *testing.T) {
	// case1: stack has some elements
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	iterator := stack.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.index != -1 {
		t.Error("case1 error: stack has some elements")
	}

	// case2: stack has one element
	stack = New()
	stack.Push(1)
	iterator = stack.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.index != -1 {
		t.Error("case2 error: stack has one element")
	}

	// case3: stack has no element
	stack = New()
	iterator = stack.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.index != -1 {
		t.Error("case3 error: stack has no element")
	}
}

func TestIndex(t *testing.T) {
	// case1: stack has some elements
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	iterator := stack.Iterator()
	i := 0
	for iterator.Next() {
		if iterator.Index() != i {
			t.Error("case1 error: stack has some elements")
		}
		i++
	}

	// case2: stack has one element
	stack = New()
	stack.Push(1)
	iterator = stack.Iterator()
	i = 0
	for iterator.Next() {
		if iterator.Index() != i {
			t.Error("case2 error: stack has one element")
		}
		i++
	}

	// case3: stack has no element
	stack = New()
	iterator = stack.Iterator()
	if iterator.index != -1 {
		t.Error("case3 error: stack has no element")
	}
}

// Benchmark Test

// Push

func benchmarkPush(b *testing.B, stack *Stack, size int) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			stack.Push(i)
		}
	}
}

func BenchmarkGetPush0(b *testing.B) {
	b.StopTimer()
	stack := New()
	size := 0
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkPush100(b *testing.B) {
	b.StopTimer()
	stack := New()
	size := 100
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkPush1000(b *testing.B) {
	b.StopTimer()
	stack := New()
	size := 1000
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkPush10000(b *testing.B) {
	b.StopTimer()
	stack := New()
	size := 10000
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkPush100000(b *testing.B) {
	b.StopTimer()
	stack := New()
	size := 100000
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

// Pop

func benchmarkPop(b *testing.B, stack *Stack, size int) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			stack.Pop()
		}
	}
}

func BenchmarkGetPop0(b *testing.B) {
	b.StopTimer()
	stack := New()
	size := 0
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkPop100(b *testing.B) {
	b.StopTimer()
	stack := New()
	size := 100
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkPop1000(b *testing.B) {
	b.StopTimer()
	stack := New()
	size := 1000
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkPop10000(b *testing.B) {
	b.StopTimer()
	stack := New()
	size := 10000
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkPop100000(b *testing.B) {
	b.StopTimer()
	stack := New()
	size := 100000
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

// Peek

func benchmarkPeek(b *testing.B, stack *Stack, size int) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			stack.Peek()
		}
	}
}

func BenchmarkGetPeek0(b *testing.B) {
	b.StopTimer()
	stack := New()
	size := 0
	b.StartTimer()
	benchmarkPeek(b, stack, size)
}

func BenchmarkPeek100(b *testing.B) {
	b.StopTimer()
	stack := New()
	size := 100
	b.StartTimer()
	benchmarkPeek(b, stack, size)
}

func BenchmarkPeek1000(b *testing.B) {
	b.StopTimer()
	stack := New()
	size := 1000
	b.StartTimer()
	benchmarkPeek(b, stack, size)
}

func BenchmarkPeek10000(b *testing.B) {
	b.StopTimer()
	stack := New()
	size := 10000
	b.StartTimer()
	benchmarkPeek(b, stack, size)
}

func BenchmarkPeek100000(b *testing.B) {
	b.StopTimer()
	stack := New()
	size := 100000
	b.StartTimer()
	benchmarkPeek(b, stack, size)
}
