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
// func (bTree *BinaryTree) GetNodeDistance(node1, node2 int) (int, int) {
// 	var found = 0
// 	var dist = 0
// 	var lDist, lFound int
// 	var rDist, rFound int

// 	if bTree.value == node2 || bTree.value == node1 {
// 		found = 1
// 		dist = 1
// 	}
// 	fmt.Printf("%2d enter: dist=%d found=%d \n", bTree.value, dist, found)
// 	if bTree.left != nil {
// 		lDist, lFound = bTree.left.GetNodeDistance(node1, node2)
// 		found += lFound
// 		if lFound == 2 {
// 			return lDist, found
// 		}
// 	}

// 	if bTree.right != nil && found < 2 {
// 		rDist, rFound = bTree.right.GetNodeDistance(node1, node2)
// 		found += rFound
// 		if rFound == 2 {
// 			return rDist, found
// 		}
// 	}
// 	// 当前的节点为目的节点 distance
// 	fmt.Printf("%2d ret: %d,%d \n", bTree.value, dist+lDist+rDist, found)
// 	if found == 1 {
// 		return dist + lDist + rDist + 1, found
// 	}
// 	return dist + lDist + rDist, found

// }
