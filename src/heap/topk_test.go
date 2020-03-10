package heap

import (
	"math/rand"
	"testing"
	"time"
)

func TestTopKInt_Insert(t *testing.T) {
	k, samples := 10, 1000

	rand.Seed(time.Now().UnixNano())

	topMax10 := NewTopKInt(k)

	s := make([]int64, 0, samples)

	for i := 0; i < samples; i++ {
		sample := int64(rand.Intn(10000))
		s = append(s, sample)
		topMax10.Insert(sample)
	}

	results := make([]int64, 0, samples)

	for i := 0; i < k; i++ {
		if v, err := topMax10.kHeap.Pop(); err == nil {
			results = append(results, v)
		}

	}

}
