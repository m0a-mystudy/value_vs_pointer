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

func returnValue04(i ValueVsPointer04) ValueVsPointer04 {
	return i
}

func returnPointer04(i *ValueVsPointer04) *ValueVsPointer04 {
	return i
}

func BenchmarkUsePointerFieldCount04(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := returnPointer04(
			&ValueVsPointer04{
				Filed1: 1,
				Filed2: 2,
				Filed3: 3,
				Filed4: 4,
			},
		)
		if v.Filed1 != 1 || v.Filed2 != 2 || v.Filed3 != 3 || v.Filed4 != 4 {
			b.Fail()
		}
	}
}

func BenchmarkUseValueFieldCount04(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := returnValue04(
			ValueVsPointer04{
				Filed1: 1,
				Filed2: 2,
				Filed3: 3,
				Filed4: 4,
			},
		)
		if v.Filed1 != 1 || v.Filed2 != 2 || v.Filed3 != 3 || v.Filed4 != 4 {
			b.Fail()
		}
	}
}
