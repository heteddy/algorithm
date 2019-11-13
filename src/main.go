package main

import (
	//tree "algorithm/tree"
	prime "algorithm/prime"
	probability "algorithm/probability"
	"algorithm/tree"
	"fmt"
	"runtime"
)

func main() {
	//tree8 := tree.NewBinaryTreeNode(nil, nil, 8)
	//tree7 := tree.NewBinaryTreeNode(nil, nil, 7)
	//tree6 := tree.NewBinaryTreeNode(nil, nil, 6)
	//tree5 := tree.NewBinaryTreeNode(nil, nil, 5)
	//tree4 := tree.NewBinaryTreeNode(tree7, nil, 4)
	//tree3 := tree.NewBinaryTreeNode(tree8, tree6, 3)
	//tree2 := tree.NewBinaryTreeNode(tree4, tree5, 2)
	//tree1 := tree.NewBinaryTreeNode(tree2, tree3, 1)
	//
	//fmt.Println("is balance tree", tree1.IsBalance())
	//nodeDistance1, _ := tree1.GetNodeDistance(7, 5)
	//fmt.Println("7 5 父节点, distance", tree1.GetAncestor(7, 5).GetValue(), nodeDistance1)
	//nodeDistance2, _ := tree1.GetNodeDistance(7, 5)
	//fmt.Println("8 4 父节点, distance", tree1.GetAncestor(8, 4).GetValue(), nodeDistance2)
	//
	//distance, _ := tree1.MaxDistance()
	//fmt.Println("max distance", distance)
	//
	//_tree8 := tree.NewBinaryTreeNode(nil, nil, 8)
	//_tree7 := tree.NewBinaryTreeNode(nil, nil, 7)
	//_tree6 := tree.NewBinaryTreeNode(nil, nil, 6)
	//_tree5 := tree.NewBinaryTreeNode(_tree6, _tree7, 5)
	//_tree4 := tree.NewBinaryTreeNode(_tree5, nil, 4)
	//_tree3 := tree.NewBinaryTreeNode(nil, _tree8, 3)
	//_tree2 := tree.NewBinaryTreeNode(_tree4, nil, 2)
	//_tree1 := tree.NewBinaryTreeNode(_tree2, _tree3, 1)
	//
	//fmt.Println("is balance tree", _tree1.IsBalance())
	//nodeDistance3, _ := _tree1.GetNodeDistance(7, 5)
	//fmt.Println("7 5 父节点", _tree1.GetAncestor(7, 5).GetValue(), nodeDistance3)
	//
	//nodeDistance4, _ := _tree1.GetNodeDistance(6, 2)
	//
	//fmt.Println("6 2 父节点", _tree1.GetAncestor(6, 2).GetValue(), nodeDistance4)
	//
	//_distance, _ := _tree1.MaxDistance()
	//fmt.Println("max distance", _distance)

	runtime.GOMAXPROCS(1)
	primers := prime.GeneratePrime(20)
	fmt.Println(primers)
	t := tree.ConstructBinSearchTree(primers)

	fmt.Println("pre order:")
	t.PreOrder()
	fmt.Println("order:")
	t.MiddleOrder()
	fmt.Println("post order:")
	t.PostOrder()

	probability.Shuffle(primers)
	fmt.Println(primers)
	randomM := probability.PickMRecursively(primers, 3)

	fmt.Println(randomM)

}
