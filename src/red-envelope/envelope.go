/*
@Copyright:
*/
/*
@Time : 2020/7/20 18:20
@Author : teddy
@File : envelope.go
*/

package red_envelope

import (
	"math/rand"
	"sort"
	"time"
)
// 思路： 把amount看成一条线段，从中随机选择person-1个点，剪段之后分给所有的persons，线段的长度就是每个人的红包
func envelope(persons int, amount int) []int {
	if persons < 1 {
		return nil
	}
	if persons == 1 {
		return []int{amount,}
	}
	sections := make(map[int]bool, 0)
	result := make([]int, 0, persons)

	for ; len(sections) < persons-1; {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		n := r.Intn(amount)
		if _, existed := sections[n]; !existed {
			sections[n] = true
		}
	}
	tmp := make([]int, 0, persons-1)
	for k, _ := range sections {
		tmp = append(tmp, k)
	}
	sort.Ints(tmp)

	start := 0
	for _, v := range tmp {
		result = append(result, v-start)
		start = v
	}
	return result
}
