package main

import (
	"testing"
)

type ValueVsPointer04 struct {
	Filed1 int
	Filed2 int
	Filed3 int
	Filed4 int
}

//go:noinline
func returnValue04(i ValueVsPointer04) ValueVsPointer04 {
	return i
}

//go:noinline
func returnPointer04(i *ValueVsPointer04) *ValueVsPointer04 {
	return i
}

func BenchmarkUsePointerFieldCount04(b *testing.B) {
	v := &ValueVsPointer04{
		Filed1: 1,
		Filed2: 2,
		Filed3: 3,
		Filed4: 4,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = returnPointer04(v)
	}
}

func BenchmarkUseValueFieldCount04(b *testing.B) {
	v := ValueVsPointer04{
		Filed1: 1,
		Filed2: 2,
		Filed3: 3,
		Filed4: 4,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = returnValue04(v)
	}
}
