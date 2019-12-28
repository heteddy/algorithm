package heap

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestTopKInt_Insert(t *testing.T) {
	k, samples := 10, 1000

	rand.Seed(time.Now().UnixNano())

	topMax10 := NewTopKInt(k, true)

	s := make([]int, 0, samples)

	for i := 0; i < samples; i++ {
		sample := rand.Intn(10000)
		s = append(s, sample)
		topMax10.Insert(sample)
	}

	results := make([]int, 0, samples)

	for i := 0; i < k; i++ {
		if v, err := topMax10.kHeap.Pop(); err == nil {
			results = append(results, v)
		}

	}
	s2 := sort.IntSlice(s)
	sort.Sort(s2)

	for i := 0; i < k; i++ {
		if s[samples-10+i] == results[i] {
		} else {
			t.FailNow()
		}
	}
}
