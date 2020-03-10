package heap

import (
	"log"
)

type TopKInt struct {
	kHeap *Heap
	k     int
	//topMax bool
}

func NewTopKInt(k int) *TopKInt {
	var slice HeapSortableSlice
	slice = &MinInt64Slice{make([]int64, 0, k)}

	return &TopKInt{
		kHeap: NewHeap(slice),
		k:     k,
	}
}

func (t *TopKInt) Insert(value int64) {

	if t.kHeap.Len() < t.k {
		t.kHeap.Append(value)
		if t.kHeap.Len() == t.k {
			t.kHeap.Sort()
		}
	} else {
		if top, err := t.kHeap.slice.IndexOf(0); err == nil {
			// 检查是topMax 还是TopMin，
			if value > top {
				// 替换0
				t.kHeap.slice.Replace(0, value)
				t.kHeap.Build()
			}
		} else {
			log.Panic(err)
		}
	}
}
