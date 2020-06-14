/*
@Copyright:
*/
/*
@Time : 2020/6/10 17:33
@Author : teddy
@File : rain_test.go
*/

package rain

import "testing"

func TestRainCount(t *testing.T) {
	i := []int{0, 1, 0, 2, 5, 3, 1, 3, 2, 1, 2, 1, 8, 6, 1, 2}
	t.Log(trapRainWater(i))
	t.Log(trapRainWater_DoublePointer(i))
}
