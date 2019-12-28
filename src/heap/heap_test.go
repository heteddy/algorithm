package heap

import (
	"math/rand"
	"testing"
	"time"
)

func TestMaxHeap_Sort(t *testing.T) {
	//random := rand.NewSource(time.Now().UnixNano())

	rand.Seed(time.Now().UnixNano())
	s := make([]int, 0, 15)

	for i := 0; i < cap(s); i++ {
		e := rand.Intn(1000)
		s = append(s, e)
	}

	h := NewHeap(&MaxIntSlice{s})
	h.Sort()
	t.Log(s)
}

func TestMinHeap_Sort(t *testing.T) {
	//random := rand.NewSource(time.Now().UnixNano())

	rand.Seed(time.Now().UnixNano())
	s := make([]int, 0, 15)

	for i := 0; i < cap(s); i++ {
		e := rand.Intn(1000)
		s = append(s, e)
	}

	h := NewHeap(&MinIntSlice{s})
	h.Sort()
	t.Log(s)
}
