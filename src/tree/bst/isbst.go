/*
@Copyright:
*/
/*
@Time : 2020/6/15 19:59
@Author : teddy
@File : isbst.go
*/

package bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBSTWithBoundary(root *TreeNode, max, min *TreeNode) bool {
	leftValid, rightValid := true, true
	if root == nil {
		return true
	}
	if max != nil {
		if root.Val >= max.Val {
			return false
		}
	}
	if min != nil {
		if root.Val <= min.Val {
			return false
		}
	}
	if root.Left != nil {
		if root.Val <= root.Left.Val {
			return false
		}

		leftValid = isValidBSTWithBoundary(root.Left, root, min)
	}
	if !leftValid {
		return false
	}
	if root.Right != nil {
		if root.Val >= root.Right.Val {
			return false
		}
		rightValid = isValidBSTWithBoundary(root.Right, max, root)
	}
	return leftValid && rightValid
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		return isValidBSTWithBoundary(root, nil, nil)
	}
}
