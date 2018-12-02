package main

import (
	"testing"
)

type ValueVsPointer03 struct {
	Filed1 int
	Filed2 int
	Filed3 int
}

func returnValue03(i ValueVsPointer03) ValueVsPointer03 {
	return i
}

func returnPointer03(i *ValueVsPointer03) *ValueVsPointer03 {
	return i
}

func BenchmarkUsePointerFieldCount03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := returnPointer03(
			&ValueVsPointer03{
				Filed1: 1,
				Filed2: 2,
				Filed3: 3,
			},
		)
		if v.Filed1 != 1 || v.Filed2 != 2 || v.Filed3 != 3 {
			b.Fail()
		}
	}
}

func BenchmarkUseValueFieldCount03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := returnValue03(
			ValueVsPointer03{
				Filed1: 1,
				Filed2: 2,
				Filed3: 3,
			},
		)
		if v.Filed1 != 1 || v.Filed2 != 2 || v.Filed3 != 3 {
			b.Fail()
		}
	}
}
