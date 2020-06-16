package ancestor

/*
查找两个二叉树节点的最近公共祖先
*/
// GetAncestor 给定两个树节点的值，找到对应节点的父节点；
// 首先判断当前节点是否为找的节点，是就返回当前节点，
// 找左子树，找到就直接返回，右子树，找到就返回
// 如果 left is nil, right is nil;左右都找不到，那么没有
// 如果left not nil，right is nil从右子树查找
// 如果left is not nil, right is not nil,说明当前节点就是父节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}

	var left, right *TreeNode

	if root.Left != nil {
		left = lowestCommonAncestor(root.Left, p, q)
	}
	if root.Right != nil {
		right = lowestCommonAncestor(root.Right, p, q)
	}
	if left == nil && right == nil {
		return nil
	}
	if left != nil && right != nil {
		return root
	} else {
		if left != nil {
			return left
		}
		return right
	}

}
