package heap

import (
	"errors"
	"fmt"
)

type HeapSortableSlice interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Append(v int)
	IndexOf(int) (int, error)
	Replace(int, int)
	Pop() (int, error)
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

func (s *MaxIntSlice) IndexOf(i int) (int, error) {
	if len(s.array) > i {
		return s.array[i], nil
	}
	return 0, errors.New("out of index")
}

func (s *MaxIntSlice) Pop() (int, error) {
	if len(s.array) > 0 {
		v := s.array[0]
		s.array = s.array[1:]
		return v, nil
	}
	return 0, errors.New("no element")
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
func (s *MinIntSlice) IndexOf(i int) (int, error) {
	if len(s.array) > i {
		return s.array[i], nil
	}
	return 0, errors.New("heap is empty")
}

func (s *MinIntSlice) Pop() (int, error) {
	if len(s.array) > 0 {
		v := s.array[0]
		s.array = s.array[1:]
		return v, nil
	}
	return 0, errors.New("no element")
}
func (s *MinIntSlice) String() string {
	return fmt.Sprint(s.array)
}

// 大根堆还是小根堆主要看 Compare的结果
// 默认按照大根堆来实现；如果是小根堆，只要覆盖less的实现就行了
type Heap struct {
	slice HeapSortableSlice
}

func NewHeap(p HeapSortableSlice) *Heap {
	return &Heap{slice: p}
}

func NewEmptyHeap() *Heap {
	return &Heap{slice: nil}
}

func (h *Heap) Append(v int) {
	h.slice.Append(v)
}

func (h *Heap) Sort() {
	// 从非叶子节点开始排序

	for index := h.slice.Len() / 2; index >= 0; index-- {
		h.adjust(index)
	}
}

func (h *Heap) Pop() (int, error) {
	if v, err := h.slice.Pop(); err == nil {
		h.Sort()
		return v, nil
	} else {
		return 0, nil
	}
}

func (h *Heap) Len() int {
	return h.slice.Len()
}

func (h *Heap) adjust(index int) {
	childIndex := 2*index + 1
	// 下标应该比长度小
	if childIndex >= h.slice.Len() {
		return
	}

	if childIndex+1 < h.slice.Len() && h.slice.Less(childIndex, childIndex+1) {
		childIndex++
	}

	if h.slice.Less(index, childIndex) {
		h.slice.Swap(index, childIndex)
		// 一旦交换了之后，后面的节点要重新调整顺序
		h.adjust(childIndex)
	}
}
