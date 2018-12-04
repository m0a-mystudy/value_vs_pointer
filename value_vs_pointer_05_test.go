package main

import (
	"testing"
)

type ValueVsPointer05 struct {
	Filed1 int
	Filed2 int
	Filed3 int
	Filed4 int
	Filed5 int
}

//go:noinline
func returnValue05(i ValueVsPointer05) ValueVsPointer05 {
	return i
}

//go:noinline
func returnPointer05(i *ValueVsPointer05) *ValueVsPointer05 {
	return i
}

func BenchmarkUsePointerFieldCount05(b *testing.B) {
	v := &ValueVsPointer05{
		Filed1: 1,
		Filed2: 2,
		Filed3: 3,
		Filed4: 4,
		Filed5: 5,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = returnPointer05(v)
	}
}

func BenchmarkUseValueFieldCount05(b *testing.B) {
	v := ValueVsPointer05{
		Filed1: 1,
		Filed2: 2,
		Filed3: 3,
		Filed4: 4,
		Filed5: 5,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = returnValue05(v)
	}
}
