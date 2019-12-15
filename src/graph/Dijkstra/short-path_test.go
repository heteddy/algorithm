package Dijkstra

import (
	"log"
	"testing"
)

func TestWeightGraph_ShortPath(t *testing.T) {
	p := make([]WeightPath, 0)
	p = append(p, WeightPath{"yuepu", "changpian", 5})
	p = append(p, WeightPath{"yuepu", "haibao", 0})
	p = append(p, WeightPath{"haibao", "jita", 30})
	p = append(p, WeightPath{"haibao", "jiazigu", 35})
	p = append(p, WeightPath{"changpian", "jita", 15})
	p = append(p, WeightPath{"changpian", "jiazigu", 20})
	p = append(p, WeightPath{"jita", "gangqin", 20})
	p = append(p, WeightPath{"jiazigu", "gangqin", 10})
	graph := New(p)

	if ok := graph.ShortPath(GraphNode("yuepu"), GraphNode("gangqin")); ok {
		log.Println("success")
	} else {
		log.Println("failure")
	}
}