
package tree

import (
	"math"
)

func (bTree *BinaryTree) IsBalance() bool {
	if bTree.getHeight(1) < 0 {
		return false
	}
	return true
}
/*
判断一棵树是否为平衡二叉树
*/
//GetHeight 获取树的高度，返回值>0 表示高度，<0 说明已经不平衡，可以直接退出
//
func (bTree *BinaryTree) getHeight(level int) int {
	if level < 0 {
		return -1
	}
	if bTree == nil {
		return level
	}
	var left = level
	var right = level
	// var err errors.new()
	if bTree.left != nil {
		left = bTree.left.getHeight(level + 1)
		// 快速返回失败
		if left < 0 {
			return -1
		}
	}
	if bTree.right != nil {
		right = bTree.right.getHeight(level + 1)
		if right < 0 {
			return -1
		}
	}
	_delta := int(math.Abs(float64(left - right)))
	if _delta > 1 {
		return -1
	} else {
		return int(math.Max(float64(left), float64(right)))
	}
}
