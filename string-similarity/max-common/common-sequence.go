package max_common

//import "log"

func max(p ...int) int {
	var max = 0
	for _, i := range p {
		if i > max {
			max = i
		}
	}
	return max
}

func CommonSequence(s1, s2 string) int {
	lenS1 := len(s1)
	lenS2 := len(s2)
	var resultArray [][]int

	for i := 0; i < lenS1; i++ {
		tmp := make([]int, 0, lenS2)
		for j := 0; j < lenS2; j++ {
			if s1[i] == s2[j] {
				if i > 0 && j > 0 {
					tmp = append(tmp, resultArray[i-1][j-1]+1)
				} else {
					tmp = append(tmp, 1)
				}
			} else {
				var left, up int
				if i > 0 {
					left = resultArray[i-1][j]
				}
				if j > 0 {
					//注意：下面这种写法，会到值index out of range
					//因为resultArray[i]还没有添加进去
					//up = resultArray[i][j-1]
					up = tmp[len(tmp)-1]
				}
				maxValue := max(left, up)
				tmp = append(tmp, max(maxValue))
			}
		}
		resultArray = append(resultArray, tmp)
	}

	return resultArray[lenS1-1][lenS2-1]
}
