package Dijkstra

import (
	"container/list"
	"fmt"
	"log"
	"strconv"
)

//import "log"

type WeightGraph map[GraphNode][]WeightPath
type GraphNode string

type WeightPath struct {
	From   GraphNode
	To     GraphNode
	Weight int
}

type StepPath struct {
	NodeList  *list.List
	SumWeight int
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

/*
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
*/
func (graph WeightGraph) ShortPath(startNode, endNode GraphNode) (ret bool) {
	stepPathList := list.New()

	if ok := graph.initStepPathList(stepPathList, startNode); ok {
		return graph.calculateShortPath(stepPathList, endNode)
	}
	return
}

func (graph WeightGraph) initStepPathList(stepPathList *list.List, start GraphNode) (ret bool) {
	currentWeightPath, existed := graph[start]
	if !existed {
		//这种情况下没有路
		return
	} else {
		ret = true
	}
	for _, p := range currentWeightPath {
		nodeList := list.New()
		nodeList.PushBack(p.From)
		nodeList.PushBack(p.To)
		s := StepPath{NodeList: nodeList, SumWeight: p.Weight}
		stepPathList.PushBack(s)
	}
	delete(graph, start)
	return
}

func (graph WeightGraph) calculateShortPath(stepPathList *list.List, end GraphNode) bool {

	//2.遍历当前的stepPathList，获取其中最短的一条路径
	//2.1 如果当前的最短路径已经是终点，而且没有其他的路径待计算，则返回当前的就是最短的，结束
	//3.沿着最短路径继续向前进，获取下一个路径，如果这是一个终点（非end，但是没有路径），将当前的的路径从stepPathList中删除
	//3.1.如果有下一条路径则修改stepPathList并重新计算SumWeight，然后从graph中删除当前的路径
	minItem := stepPathList.Front()
	min := minItem.Value.(StepPath).SumWeight

	for item := stepPathList.Front(); item != nil; item = item.Next() {
		if item.Value.(StepPath).SumWeight < min {
			minItem = item
			min = minItem.Value.(StepPath).SumWeight
		}
	}
	//2.1如果已经是终点,并且没有其他的路径需要计算，则返回最短的一条路劲
	//查找结束
	if minItem.Value.(StepPath).NodeList.Back().Value.(GraphNode) == end && len(graph) == 0 {
		item := minItem.Value.(StepPath).NodeList
		var output string
		for i := item.Front(); i != nil; i = i.Next() {
			output += fmt.Sprintf("%s ", i.Value.(GraphNode))
		}
		output += ":" + strconv.Itoa(minItem.Value.(StepPath).SumWeight)
		log.Println(output)
		return true
	}
	//取出minItem的最后链表的最后一项，即当前的终点，然后加入
	candidateStart := minItem.Value.(StepPath).NodeList.Back().Value.(GraphNode)
	//3.1如果不存在下一跳路径，则删除当前的路径
	if currentWeightPath, existed := graph[candidateStart]; !existed {
		stepPathList.Remove(minItem)
		//如果没有路可选，说明从起点到终点没有一条可达的路径
		if stepPathList.Len() == 0 {
			return false
		} else {
			return graph.calculateShortPath(stepPathList, end)
		}

	} else {
		delete(graph, candidateStart)
		stepPathList.Remove(minItem)
		for _, p := range currentWeightPath {
			newList := list.New()
			for mi := minItem.Value.(StepPath).NodeList.Front(); mi != nil; mi = mi.Next() {
				newList.PushBack(mi.Value.(GraphNode))
			}
			newList.PushBack(p.To)
			s := StepPath{NodeList: newList, SumWeight: p.Weight + minItem.Value.(StepPath).SumWeight}

			stepPathList.PushBack(s)
		}
		return graph.calculateShortPath(stepPathList, end)
	}
}
