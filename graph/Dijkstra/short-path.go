package Dijkstra

import (
	"container/list"
	"fmt"
	"log"
)

//import "log"

type WeightGraph map[GraphNode][]WeightPath
type GraphNode string
type WeightPath struct {
	From   GraphNode
	To     GraphNode
	Weight int
}

func New(path []WeightPath) WeightGraph {
	graph := make(WeightGraph)
	for _, item := range path {
		if pathList, existed := graph[item.From]; existed {
			pathList = append(pathList, item)
			graph[item.From] = pathList
		} else {
			graph[item.From] = make([]WeightPath, 0)
			pathList = append(graph[item.From], item)
			graph[item.From] = pathList
		}
	}
	return graph
}

func (graph WeightGraph) GetAllNode() *list.List {

	nodeList := list.New()
	nodeSet := make(map[GraphNode]bool)

	for k, v := range graph {
		nodeSet[k] = true
		for _, nodePath := range v {
			nodeSet[nodePath.To] = true
		}
	}
	for node, _ := range nodeSet {
		nodeList.PushBack(node)
	}
	return nodeList
}

func (graph WeightGraph) ShortPath(startNode, endNode GraphNode) {
	allNodes := graph.GetAllNode()
	if allNodes.Len() <= 1 {
		return
	}
	nodesButStartList := list.New()
	nodesButEndList := list.New()

	for item := allNodes.Front(); item != nil; item = item.Next() {
		if item.Value.(GraphNode) == startNode {
			nodesButEndList.PushFront(item.Value.(GraphNode))
		} else if item.Value.(GraphNode) == endNode {
			nodesButStartList.PushBack(item.Value.(GraphNode))
		} else {
			nodesButEndList.PushBack(item.Value.(GraphNode))
			nodesButStartList.PushFront(item.Value.(GraphNode))
		}
	}
	var noStart, noEnd string
	for item := nodesButStartList.Front(); item != nil; item = item.Next() {
		noStart += fmt.Sprintf(" %s", item.Value.(GraphNode))
	}
	log.Println(noStart+"\n", "--------------------------")

	for item := nodesButEndList.Front(); item != nil; item = item.Next() {
		noEnd += fmt.Sprintf(" %s", item.Value.(GraphNode))
	}
	log.Println(noEnd+"\n", "--------------------------")
	return
}
