/*
@Copyright:
*/
/*
@Time : 2020/6/10 16:32
@Author : teddy
@File : rain.go
*/

package rain

import (
	"strconv"
)

// 记录每个点的左右两边的最大最小值
type bucket struct {
	maxLeft, maxRight int
	height            int
}

func (b bucket) String() string {
	return "maxLeft:" + strconv.Itoa(b.maxLeft) + " height:" + strconv.Itoa(b.height) + " maxRight:" + strconv.Itoa(b.maxRight) + "\n"
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func trapRainWater(input []int) int {
	length := len(input)
	buckets := make([]*bucket, 0, len(input))
	maxLeft, maxRight := 0, 0
	count := 0
	for _, i := range input {
		b := &bucket{
			maxLeft:  maxLeft,
			maxRight: 0,
			height:   i,
		}
		buckets = append(buckets, b)
		maxLeft = max(maxLeft, i)
	}
	// 找到右侧的最大值，并计算水量
	for idx, _ := range input {
		_idx := idx + 1
		maxRight = max(maxRight, input[length-_idx])
		buckets[length-_idx].maxRight = maxRight

		point := buckets[length-_idx]

		m := min(point.maxRight, point.maxLeft)
		if m > point.height {
			count += m - point.height
		}
	}
	//log.Println(buckets)
	return count
}

func trapRainWater_DoublePointer(height []int) int {
	// 定义两个指针，一个从左向右移动，一个从右向左移动
	count := 0
	length := len(height)
	left, right := 0, length-1
	maxLeft, maxRight := 0, 0
	for ; left < right; {
		maxLeft = max(maxLeft, height[left])
		maxRight = max(maxRight, height[right])

		if maxLeft < maxRight {
			count += maxLeft - height[left]
			left++
		} else {
			count += maxRight - height[right]
			right--
		}
	}
	return count
}
