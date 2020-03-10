package heap

import (
	"math/rand"
	"testing"
	"time"
)

func TestMinHeap_Sort(t *testing.T) {
	//random := rand.NewSource(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	s := make([]int64, 0, 15)

	for i := 0; i < cap(s); i++ {
		e := int64(rand.Intn(1000))
		s = append(s, e)
	}

	h := NewHeap(&MinInt64Slice{s})
	h.Sort()
	t.Log(s)
}
