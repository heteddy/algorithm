package tree

/*
给定一个有序数组，编写算法创建最小的二叉查找树
*/

func ConstructBinSearchTree(input []int) *BinaryTree {
	if input == nil {
		return nil
	}
	length := len(input)
	var middle int
	if length > 2 {
		middle = length / 2

		left := ConstructBinSearchTree(input[:middle])
		right := ConstructBinSearchTree(input[middle+1 : length])
		node := NewBinaryTreeNode(left, right, input[middle])
		return node

	} else if length == 2 {
		left := NewBinaryTreeNode(nil, nil, input[0])
		node := NewBinaryTreeNode(left, nil, input[1])
		return node
	} else {
		return NewBinaryTreeNode(nil, nil, input[0])
	}

}
