/*
@Copyright:
*/
/*
@Time : 2020/3/12 20:04
@Author : teddy
@File : kmerge_test.go
*/

package heap

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func prepareData(src []int64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(src); i++ {
		src[i] = int64(r.Int31n(100))
	}
	sort.Slice(src, func(i, j int) bool {
		return src[i] < src[j]
	})
}

func TestHeapMerge_Pop(t *testing.T) {
	k := 10

	sortedSlices := make([]*SortedSlice, 0, k)
	lens := []int{2, 8, 7, 8, 9, 6, 4, 3, 5, 6,}
	for i := 0; i < k; i++ {
		s := make([]int64, lens[i])
		prepareData(s)
		sortedSlices = append(sortedSlices, NewSortedSlice(s))
		//log.Println(s)
	}

	merge := NewHeapMerge(sortedSlices)
	merge.Build()

	for {
		s := "pop data:"
		v, err := merge.Pop()
		if err != nil {
			break
		} else {
			s += fmt.Sprintf("%d ", v)
		}
	}
}
