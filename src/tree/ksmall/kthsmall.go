/*
@Copyright:
*/
/*
@Time : 2020/6/14 17:32
@Author : teddy
@File : kthsmall.go
*/

package ksmall

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	// 使用一个栈顺序遍历直到第k个元素
	stack := make([]*TreeNode, 0, k)
	count := 0
	node := root
	// 如果达不到k个怎么办
	for {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		// node为nil，从栈顶去一个元素
		if len(stack) > 0 {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			count++
			if count < k {
			} else {
				v := current.Val
				return v
			}
			current = current.Right
			node = current
		} else {
			return 0
		}
	}
}
