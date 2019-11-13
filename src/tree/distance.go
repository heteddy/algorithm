package tree

import (
	"math"
)

// 节点间的最大距离：所有节点的最大距离
// 两个节点的最小距离，给定2个节点，找出一条最小路径
// 返回 最大距离，最深的层次
func (bTree *BinaryTree) MaxDistance() (int, int) {

	var leftDis, rightDis int
	var leftMaxLevel, rightMaxLevel int

	if bTree.left != nil {
		leftDis, leftMaxLevel = bTree.left.MaxDistance()
	}
	if bTree.right != nil {
		rightDis, rightMaxLevel = bTree.right.MaxDistance()
	}
	// 当前节点到叶子节点的深度
	maxLevel := int(math.Max(float64(leftMaxLevel), float64(rightMaxLevel))) + 1
	// 当前节点max distance
	currentMaxDistance := leftMaxLevel + rightMaxLevel + 1

	return int(math.Max(float64(currentMaxDistance), math.Max(float64(leftDis), float64(rightDis)))), maxLevel
}

// // 返回值 1.距离，2.找到node个数
func (bTree *BinaryTree) GetNodeDistance(node1, node2 int) (dist int, found int) {
	var lDist, lFound int
	var rDist, rFound int

	if bTree.left != nil {
		lDist, lFound = bTree.left.GetNodeDistance(node1, node2)
		found += lFound
		dist += lDist
	}
	if bTree.right != nil && found != 2 {
		rDist, rFound = bTree.right.GetNodeDistance(node1, node2)
		found += rFound
		dist += rDist
	}
	// 这种情况是左右子树
	if found == 2 {
		return
	}

	if bTree.value == node2 || bTree.value == node1 {
		found = rFound + lFound + 1
		// 这种情况，2个节点都在一边；其中一个节点是另外一个节点的父节点
		if found == 2 {
			dist = int(math.Max(float64(lDist), float64(rDist)))
		} else {
			// 仅仅是当前的节点 found = 1
			dist = lDist + rDist
		}
	}
	// 当仅有一个节点，说明需要继续向上查找；因此需要增加 距离(节点层次需要递增)
	if found == 1 {
		dist += 1
	}
	return
}
