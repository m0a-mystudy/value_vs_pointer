package main

import (
	"testing"
)

type ValueVsPointer03 struct {
	Filed1 int
	Filed2 int
	Filed3 int
}

//go:noinline
func returnValue03(i ValueVsPointer03) ValueVsPointer03 {
	return i
}

//go:noinline
func returnPointer03(i *ValueVsPointer03) *ValueVsPointer03 {
	return i
}

func BenchmarkUsePointerFieldCount03(b *testing.B) {
	v := &ValueVsPointer03{
		Filed1: 1,
		Filed2: 2,
		Filed3: 3,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = returnPointer03(v)
	}
}

func BenchmarkUseValueFieldCount03(b *testing.B) {
	v := ValueVsPointer03{
		Filed1: 1,
		Filed2: 2,
		Filed3: 3,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = returnValue03(v)
	}
}
