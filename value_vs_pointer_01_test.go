package main

import (
	"testing"
)

type ValueVsPointer struct {
	Filed1 int
}

func returnValue(i ValueVsPointer) ValueVsPointer {
	return i
}

func returnPointer(i *ValueVsPointer) *ValueVsPointer {
	return i
}

func BenchmarkUsePointerFieldCount01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := returnPointer(&ValueVsPointer{Filed1: 1})
		if v.Filed1 != 1 {
			b.Fail()
		}
	}
}

func BenchmarkUseValueFieldCount01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := returnValue(ValueVsPointer{Filed1: 1})
		if v.Filed1 != 1 {
			b.Fail()
		}
	}
}
