package consistentHash

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestConsistentHashService_Get(t *testing.T) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	service := NewConsistentService(100)
	for i := 0; i < 7; i++ {
		service.Add(NewNode(i, "192.168.1."+strconv.Itoa(i), 2379, "node_"+strconv.Itoa(i), 1))
	}

	ipMap := make(map[string]int, 0)
	for i := 0; i < 1000; i++ {
		si := fmt.Sprintf("key%d_%d", i, rand.Int31())
		// 拿到node
		k := service.Get(si)
		if _, ok := ipMap[k.String()]; ok {
			ipMap[k.String()] += 1
		} else {
			ipMap[k.String()] = 1
		}
	}

	for k, v := range ipMap {
		t.Log("Node:", k, " count:", v)
	}
}
