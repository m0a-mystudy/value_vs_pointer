package main

import (
	"testing"
)

type ValueVsPointer02 struct {
	Filed1 int
	Filed2 int
}

//go:noinline
func returnValue02(i ValueVsPointer02) ValueVsPointer02 {
	return i
}

//go:noinline
func returnPointer02(i *ValueVsPointer02) *ValueVsPointer02 {
	return i
}

func BenchmarkUsePointerFieldCount02(b *testing.B) {
	v := &ValueVsPointer02{
		Filed1: 1,
		Filed2: 2,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = returnPointer02(v)
	}
}

func BenchmarkUseValueFieldCount02(b *testing.B) {
	v := ValueVsPointer02{
		Filed1: 1,
		Filed2: 2,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = returnValue02(v)
	}
}
