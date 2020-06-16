package tree

import "errors"

// https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/submissions/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func checkHeight(root *TreeNode) (int, error) {
	var err error
	var leftHeight, rightHeight int
	if root == nil {
		return 0, nil
	}
	if root.Left != nil {
		leftHeight, err = checkHeight(root.Left)
	}
	if err != nil {
		return 0, err
	} else {
		leftHeight += 1
	}
	if root.Right != nil {
		rightHeight, err = checkHeight(root.Right)
	}
	if err != nil {
		return 0, err
	} else {
		rightHeight++
	}
	if max(leftHeight, rightHeight)-min(leftHeight, rightHeight) <= 1 {
		return max(leftHeight, rightHeight), nil
	} else {
		return 0, errors.New("not balance")
	}
}

func isBalanced(root *TreeNode) bool {
	if _, err := checkHeight(root); err == nil {
		return true
	}
	return false
}
