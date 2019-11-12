package tree

type BinaryTree struct {
	value int // value可以设置为一个接口，用于比较以及获取 值
	left  *BinaryTree
	right *BinaryTree
}

func NewBinaryTreeNode(left *BinaryTree, right *BinaryTree, value int) *BinaryTree {
	return &BinaryTree{
		value,
		left,
		right,
	}
}

func (bTree *BinaryTree) GetValue() int {
	return bTree.value
}

/*
	tree8 := tree.NewBinaryTreeNode(nil, nil, 8)
	tree7 := tree.NewBinaryTreeNode(nil, nil, 7)
	tree6 := tree.NewBinaryTreeNode(nil, nil, 6)
	tree5 := tree.NewBinaryTreeNode(nil, nil, 5)
	tree4 := tree.NewBinaryTreeNode(tree7, nil, 4)
	tree3 := tree.NewBinaryTreeNode(tree8, tree6, 3)
	tree2 := tree.NewBinaryTreeNode(tree4, tree5, 2)
	tree1 := tree.NewBinaryTreeNode(tree2, tree3, 1)

	fmt.Println("is balance tree", tree1.IsBalance())

	fmt.Println("7 5 父节点", tree1.GetAncestor(7, 5).GetValue())
	fmt.Println("8 4 父节点", tree1.GetAncestor(8, 4).GetValue())
	_tree8 := tree.NewBinaryTreeNode(nil, nil, 8)
	_tree7 := tree.NewBinaryTreeNode(nil, nil, 7)
	_tree6 := tree.NewBinaryTreeNode(nil, nil, 6)
	_tree5 := tree.NewBinaryTreeNode(_tree6, _tree7, 5)
	_tree4 := tree.NewBinaryTreeNode(_tree5, nil, 4)
	_tree3 := tree.NewBinaryTreeNode(nil, _tree8, 3)
	_tree2 := tree.NewBinaryTreeNode(_tree4, nil, 2)
	_tree1 := tree.NewBinaryTreeNode(_tree2, _tree3, 1)

	fmt.Println("is balance tree", _tree1.IsBalance())

	fmt.Println("7 5 父节点", _tree1.GetAncestor(7, 5).GetValue())
	fmt.Println("6 2 父节点", _tree1.GetAncestor(6, 2).GetValue())
*/
