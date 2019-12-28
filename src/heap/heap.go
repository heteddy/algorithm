package heap

import (
	"errors"
	"fmt"
)

type HeapSortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Append(v int)
	Top() (int, error)
	Replace(int, int)
}

type MaxIntSlice struct {
	array []int
}

func (s *MaxIntSlice) Replace(i int, value int) {
	s.array[i] = value
}
func (s *MaxIntSlice) Len() int {
	return len(s.array)
}

func (s *MaxIntSlice) Less(i, j int) bool {
	return s.array[i] < s.array[j]
}

func (s *MaxIntSlice) Swap(i, j int) {
	s.array[i], s.array[j] = s.array[j], s.array[i]
}

func (s *MaxIntSlice) Append(v int) {
	s.array = append(s.array, v)
}

func (s *MaxIntSlice) Top() (int, error) {
	if len(s.array) > 0 {
		return s.array[0], nil
	}
	return 0, errors.New("heap is empty")
}

func (s *MaxIntSlice) String() string {
	return fmt.Sprint(s.array)
}

type MinIntSlice struct {
	array []int
}

func (s *MinIntSlice) Replace(i int, value int) {
	s.array[i] = value
}

func (s MinIntSlice) Len() int {
	return len(s.array)
}

func (s MinIntSlice) Less(i, j int) bool {
	return s.array[i] > s.array[j]
}

func (s MinIntSlice) Swap(i, j int) {
	s.array[i], s.array[j] = s.array[j], s.array[i]
}

func (s *MinIntSlice) Append(v int) {
	s.array = append(s.array, v)
}
func (s *MinIntSlice) Top() (int, error) {
	if len(s.array) > 0 {
		return s.array[0], nil
	}
	return 0, errors.New("heap is empty")
}

func (s *MinIntSlice) String() string {
	return fmt.Sprint(s.array)
}

// 大根堆还是小根堆主要看 Compare的结果
// 默认按照大根堆来实现；如果是小根堆，只要覆盖less的实现就行了
type Heap struct {
	array HeapSortable
}

func NewHeap(p HeapSortable) *Heap {
	return &Heap{array: p}
}

func NewEmptyHeap() *Heap {
	return &Heap{array: nil}
}

func (h *Heap) Append(v int) {
	h.array.Append(v)
}

func (h *Heap) Sort() {
	// 从非叶子节点开始排序

	for index := h.array.Len() / 2; index >= 0; index-- {
		h.adjust(index)
	}
}

func (h *Heap) Len() int {
	return h.array.Len()
}

func (h *Heap) Top() (int, error) {
	return h.array.Top()
}

func (h *Heap) adjust(index int) {
	childIndex := 2*index + 1
	// 下标应该比长度小
	if childIndex >= h.array.Len() {
		return
	}

	if childIndex+1 < h.array.Len() && h.array.Less(childIndex, childIndex+1) {
		childIndex++
	}

	if h.array.Less(index, childIndex) {
		h.array.Swap(index, childIndex)
		// 一旦交换了之后，后面的节点要重新调整顺序
		h.adjust(childIndex)
	}
}
