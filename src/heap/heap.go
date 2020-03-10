package heap

import "errors"

type HeapSortableSlice interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Append(v int64)
	IndexOf(int) (int64, error)
	Pop() (int64, error)
	Replace(int, int64)
	GetSlice() []int64
}

type MinInt64Slice struct {
	array []int64
}

func (s *MinInt64Slice) Replace(i int, v int64) {
	s.array[i] = v
}
func (s *MinInt64Slice) GetSlice() []int64 {
	return s.array
}
func (s *MinInt64Slice) Len() int {
	return len(s.array)
}

func (s *MinInt64Slice) Less(i, j int) bool {
	return s.array[i] < s.array[j]
}

func (s *MinInt64Slice) Swap(i, j int) {
	s.array[i], s.array[j] = s.array[j], s.array[i]
}

func (s *MinInt64Slice) Append(v int64) {
	s.array = append(s.array, v)
}
func (s *MinInt64Slice) IndexOf(i int) (int64, error) {
	if len(s.array) > i {
		return s.array[i], nil
	}
	return 0, errors.New("heap is empty")
}

func (s *MinInt64Slice) Pop() (int64, error) {
	if len(s.array) > 0 {
		v := s.array[0]
		s.array = s.array[1:]
		return v, nil
	}
	return 0, errors.New("no element")
}

// 大根堆还是小根堆主要看 Compare的结果
// 默认按照大根堆来实现；如果是小根堆，只要覆盖less的实现就行了
type Heap struct {
	slice HeapSortableSlice
}

func NewHeap(p HeapSortableSlice) *Heap {
	return &Heap{slice: p}
}

func (h *Heap) GetSortable() HeapSortableSlice {
	return h.slice
}

func NewEmptyHeap() *Heap {
	return &Heap{slice: nil}
}

func (h *Heap) Append(v int64) {
	h.slice.Append(v)
}

func (h *Heap) Build() {
	// 从非叶子节点开始排序
	for index := h.slice.Len() / 2; index >= 0; index-- {
		h.adjust(index, h.slice.Len())
	}
}
func (h *Heap) Sort() {
	// 从非叶子节点开始排序
	h.Build()
	//将堆顶的元素放入最后的位置，
	//依次创建堆并且放入指定位置
	//s := h.slice
	for j := h.slice.Len() - 1; j > 0; j-- {
		h.slice.Swap(0, j)
		h.adjust(0, j)
	}
}

func (h *Heap) Pop() (int64, error) {
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

func (h *Heap) adjust(start, end int) {
	childIndex := 2*start + 1
	// 下标应该比长度小
	if childIndex >= end {
		return
	}

	if childIndex+1 < end && h.slice.Less(childIndex, childIndex+1) {
		childIndex++
	}

	if h.slice.Less(start, childIndex) {
		h.slice.Swap(start, childIndex)
		// 一旦交换了之后，后面的节点要重新调整顺序
		h.adjust(childIndex, end)
	}
}
