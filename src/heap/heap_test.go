package heap

import (
	"math/rand"
	"testing"
	"time"
)

func TestMinHeap_Sort(t *testing.T) {
	//random := rand.NewSource(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	s := make([]int64, 0, 8)

	for i := 0; i < cap(s); i++ {
		e := int64(rand.Intn(100))
		s = append(s, e)
	}

	h := NewHeap(&MinInt64Slice{s})
	h.Sort()
	t.Log(s)
}
