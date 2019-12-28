package heap

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestTopKInt_Insert(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	topMax10 := NewTopKInt(10,true)

	s := make([]int,0,1000)

	for i:= 0; i< 1000; i++ {
		sample := rand.Intn(1000)
		s= append(s, sample)

		topMax10.Insert(sample)
	}
	t.Log(topMax10.kHeap.array)
	s2 := sort.IntSlice(s)
	sort.Sort(s2)
	t.Log(s[990:])
}