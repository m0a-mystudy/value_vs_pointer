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

func returnValue05(i ValueVsPointer05) ValueVsPointer05 {
	return i
}

func returnPointer05(i *ValueVsPointer05) *ValueVsPointer05 {
	return i
}

func BenchmarkUsePointerFieldCount05(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := returnPointer05(
			&ValueVsPointer05{
				Filed1: 1,
				Filed2: 2,
				Filed3: 3,
				Filed4: 4,
				Filed5: 5,
			},
		)
		if v.Filed1 != 1 || v.Filed2 != 2 || v.Filed3 != 3 || v.Filed4 != 4 ||
			v.Filed5 != 5 {
			b.Fail()
		}
	}
}

func BenchmarkUseValueFieldCount05(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := returnValue05(
			ValueVsPointer05{
				Filed1: 1,
				Filed2: 2,
				Filed3: 3,
				Filed4: 4,
				Filed5: 5,
			},
		)
		if v.Filed1 != 1 || v.Filed2 != 2 || v.Filed3 != 3 || v.Filed4 != 4 ||
			v.Filed5 != 5 {
			b.Fail()
		}
	}
}
