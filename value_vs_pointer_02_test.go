package main

import (
	"testing"
)

type ValueVsPointer02 struct {
	Filed1 int
	Filed2 int
}

func returnValue02(i ValueVsPointer02) ValueVsPointer02 {
	return i
}

func returnPointer02(i *ValueVsPointer02) *ValueVsPointer02 {
	return i
}

func BenchmarkUsePointerFieldCount02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := returnPointer02(
			&ValueVsPointer02{
				Filed1: 1,
				Filed2: 2,
			},
		)
		if v.Filed1 != 1 || v.Filed2 != 2 {
			b.Fail()
		}
	}
}

func BenchmarkUseValueFieldCount02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := returnValue02(
			ValueVsPointer02{
				Filed1: 1,
				Filed2: 2,
			},
		)
		if v.Filed1 != 1 || v.Filed2 != 2 {
			b.Fail()
		}
	}
}
