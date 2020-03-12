/*
@Copyright:

@Time : 2020/3/12 15:38
@Author : teddy
@File : merge.go
输入：k路有序数组或者链表
输出：归并排序生成一个有序链表
构造一个k路归并树，每一路为一个有序队列或者链表，利用堆排序(小顶堆) 归并排序
*/

package heap

import (
	"errors"
	"fmt"
	"log"
)

type Iterator struct {
	slice []int64
	index int
}

type SortedSlice struct {
	slice []int64
	Iterator
}

func NewSortedSlice(slice []int64) *SortedSlice {
	return &SortedSlice{
		slice: slice,
		Iterator: Iterator{
			slice: slice,
			index: 0,
		},
	}
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.slice)-1
}

func (i *Iterator) Next() {
	i.index++
}

func (i *Iterator) Value() int64 {
	return i.slice[i.index]
}

type HeapMerge struct {
	nodes []*SortedSlice
}

func NewHeapMerge(sources []*SortedSlice) *HeapMerge {
	// 需要保证
	return &HeapMerge{nodes: sources}
}

func (h *HeapMerge) Build() {
	for index := len(h.nodes) / 2; index >= 0; index-- {
		h.adjust(index, len(h.nodes))
	}
	//h.Print()
}

func (h *HeapMerge) Pop() (int64, error) {
	var value int64
	var err error

	if len(h.nodes) > 0 {
		value = h.nodes[0].Value()
		err = nil

		if h.nodes[0].HasNext() {
			h.nodes[0].Next() //不需要获取值
			h.adjust(0, len(h.nodes))
		} else { // 顶部的node(slice)已经为空
			if len(h.nodes) >= 1 {
				// 移除为已经合并完成的slice
				h.nodes = h.nodes[1:]
				h.adjust(0, len(h.nodes))
			} else {
				return 0, errors.New("merge complete")
			}
		}
	} else {
		return 0, errors.New("merge complete")
	}
	return value, err
}

func (h *HeapMerge) Print() {
	s := "heap merge:"
	for _, n := range h.nodes {
		s += fmt.Sprintf("%d ", n.Value())
	}
	log.Println(s)
}

func (h *HeapMerge) adjust(start, end int) {
	childIndex := 2*start + 1
	// 下标应该比长度小
	if childIndex >= end {
		return
	}
	if childIndex+1 < end && h.nodes[childIndex+1].Value() < h.nodes[childIndex].Value() {
		childIndex++
	}

	if h.nodes[childIndex].Value() < h.nodes[start].Value() {
		h.nodes[start], h.nodes[childIndex] = h.nodes[childIndex], h.nodes[start]

		// 一旦交换了之后，后面的节点要重新调整顺序
		h.adjust(childIndex, end)
	}
}
