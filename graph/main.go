package main

import (
	"algorithm/graph/Dijkstra"
	"log"
)

func main(){
	p := make([]Dijkstra.WeightPath,0)
	p = append(p,Dijkstra.WeightPath{"yuepu","changpian",5})
	p = append(p,Dijkstra.WeightPath{"yuepu","haibao",0})
	p = append(p,Dijkstra.WeightPath{"haibao","jita",30})
	p = append(p,Dijkstra.WeightPath{"haibao","jiazigu",35})
	p = append(p,Dijkstra.WeightPath{"changpian","jita",15})
	p = append(p,Dijkstra.WeightPath{"changpian","jiazigu",20})
	p = append(p,Dijkstra.WeightPath{"jita","gangqin",20})
	p = append(p,Dijkstra.WeightPath{"jiazigu","gangqin",10})
	graph := Dijkstra.New(p)
	log.Println(graph)
	log.Println(graph.GetAllNode())
	graph.ShortPath(Dijkstra.GraphNode("yuepu"),Dijkstra.GraphNode("gangqin"))
	
}
