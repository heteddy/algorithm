/*
@Copyright:
*/
/*
@Time : 2020/6/16 16:12
@Author : teddy
@File : symmetric
*/

package symmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func childSymmetric(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	//这里是对称的定义
	return childSymmetric(left.Left, right.Right) && childSymmetric(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return childSymmetric(root.Left, root.Right)
}
