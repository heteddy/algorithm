/*
@Copyright:
*/
/*
@Time : 2020/6/15 20:00
@Author : teddy
@File : isbst_test.go
*/

package bst

import "testing"

func constructBST() *TreeNode {
	//[10,5,15,null,null,6,20]
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   20,
				Left:  nil,
				Right: nil,
			},
		},
	}
	return root
}

func TestIsValidBST(t *testing.T) {
	r := constructBST()
	t.Log(isValidBST(r))
}
