
package tree


/*
查找两个二叉树节点的最近公共祖先
*/
// GetAncestor 给定两个树节点的值，找到对应节点的父节点；
// 首先判断当前节点是否为找的节点，是就返回当前节点，
// 找左子树，找到就直接返回，右子树，找到就返回
// 如果 left is nil, right is nil;左右都找不到，那么没有
// 如果left not nil，right is nil从右子树查找
// 如果left is not nil, right is not nil,说明当前节点就是父节点
func (bTree *BinaryTree) GetAncestor(node1, node2 int) *BinaryTree {
	if bTree.value == node1 || bTree.value == node2 {
		return bTree
	}
	var left, right *BinaryTree
	if bTree.left != nil {
		left = bTree.left.GetAncestor(node1, node2)
	}
	if bTree.right != nil {
		right = bTree.right.GetAncestor(node1, node2)
	}
	if left == nil && right == nil {
		return nil
	}
	if left != nil && right != nil {
		// 左子树找到一个，右子树找到一个，因此当前就是父节点
		return bTree
	} else {
		// 1. node1是node2的父节点，
		// 2. 只找到其中一个节点比如
		if left != nil {
			return left
		} else {
			return right
		}
	}
}
