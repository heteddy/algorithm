/*
@Copyright:
*/
/*
@Time : 2020/3/11 22:47
@Author : teddy
@File : quick_sort_test.go
*/

package quick_sort

import (
	"errors"
	"math/rand"
	"testing"
	"time"
)

type MinInt64Slice struct {
	array []int64
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
func TestMinQuick_Sort(t *testing.T) {
	//random := rand.NewSource(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	s := make([]int64, 0, 30)

	for i := 0; i < cap(s); i++ {
		e := int64(rand.Intn(1000))
		s = append(s, e)
	}

	h := NewQuick(&MinInt64Slice{s})
	h.Sort()
	t.Log(s)
}
