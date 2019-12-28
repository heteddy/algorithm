package heap

import (
	"log"
	"sort"
)

type TopKInt struct {
	kHeap  *Heap
	k      int
	topMax bool
}

func NewTopKInt(k int, topMax bool) *TopKInt {
	s := sort.IntSlice(make([]int, 0, k))
	var slice HeapSortable
	if topMax {
		slice = &MinIntSlice{s}
	} else {
		slice = &MaxIntSlice{s}
	}
	return &TopKInt{
		kHeap:  NewHeap(slice),
		k:      k,
		topMax: topMax,
	}
}

func (t *TopKInt) Insert(value int) {
	if t.kHeap.Len() < t.k {

		t.kHeap.Append(value)

		if t.kHeap.Len() == t.k {
			t.kHeap.Sort()
		}
	} else {
		if top, err := t.kHeap.Top(); err == nil {
			// 检查是topMax 还是TopMin，
			if t.topMax {
				if value > top {
					// 替换0
					t.kHeap.array.Replace(0, value)
					t.kHeap.adjust(0)
				}
			} else {
				if value < top {
					t.kHeap.array.Replace(0, value)
					t.kHeap.adjust(0)
				}
			}
		} else {
			log.Println(err)
		}
	}
}
