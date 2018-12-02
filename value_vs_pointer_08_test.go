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

func returnValue08(i ValueVsPointer08) ValueVsPointer08 {
	return i
}

func returnPointer08(i *ValueVsPointer08) *ValueVsPointer08 {
	return i
}

func BenchmarkUsePointerFieldCount08(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := returnPointer08(
			&ValueVsPointer08{
				Filed1: 1,
				Filed2: 2,
				Filed3: 3,
				Filed4: 4,
				Filed5: 5,
				Filed6: 6,
				Filed7: 7,
				Filed8: 8,
			},
		)
		if v.Filed1 != 1 || v.Filed2 != 2 || v.Filed3 != 3 || v.Filed4 != 4 ||
			v.Filed5 != 5 || v.Filed6 != 6 || v.Filed7 != 7 || v.Filed8 != 8 {
			b.Fail()
		}
	}
}

func BenchmarkUseValueFieldCount08(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := returnValue08(
			ValueVsPointer08{
				Filed1: 1,
				Filed2: 2,
				Filed3: 3,
				Filed4: 4,
				Filed5: 5,
				Filed6: 6,
				Filed7: 7,
				Filed8: 8,
			},
		)
		if v.Filed1 != 1 || v.Filed2 != 2 || v.Filed3 != 3 || v.Filed4 != 4 ||
			v.Filed5 != 5 || v.Filed6 != 6 || v.Filed7 != 7 || v.Filed8 != 8 {

			b.Fail()
		}
	}
}
