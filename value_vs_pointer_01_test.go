package main

import (
	"testing"
)

type ValueVsPointer struct {
	Filed1 int
}

//go:noinline
func returnValue(i ValueVsPointer) ValueVsPointer {
	return i
}

//go:noinline
func returnPointer(i *ValueVsPointer) *ValueVsPointer {
	return i
}

func BenchmarkUsePointerFieldCount01(b *testing.B) {
	v := &ValueVsPointer{Filed1: 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = returnPointer(v)
	}
}

func BenchmarkUseValueFieldCount01(b *testing.B) {
	v := ValueVsPointer{Filed1: 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = returnValue(v)
	}
}
