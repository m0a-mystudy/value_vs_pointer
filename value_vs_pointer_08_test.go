package main

import (
	"testing"
)

type ValueVsPointer08 struct {
	Filed1 int
	Filed2 int
	Filed3 int
	Filed4 int
	Filed5 int
	Filed6 int
	Filed7 int
	Filed8 int
}

//go:noinline
func returnValue08(i ValueVsPointer08) ValueVsPointer08 {
	return i
}

//go:noinline
func returnPointer08(i *ValueVsPointer08) *ValueVsPointer08 {
	return i
}

func BenchmarkUsePointerFieldCount08(b *testing.B) {
	v := &ValueVsPointer08{
		Filed1: 1,
		Filed2: 2,
		Filed3: 3,
		Filed4: 4,
		Filed5: 5,
		Filed6: 6,
		Filed7: 7,
		Filed8: 8,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = returnPointer08(v)
	}
}

func BenchmarkUseValueFieldCount08(b *testing.B) {
	v := ValueVsPointer08{
		Filed1: 1,
		Filed2: 2,
		Filed3: 3,
		Filed4: 4,
		Filed5: 5,
		Filed6: 6,
		Filed7: 7,
		Filed8: 8,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = returnValue08(v)
	}
}
