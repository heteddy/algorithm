package max_common

//import "log"

//import "unicode/utf8"

/*
假设字符串都是ASICII码字符
*/
func CommonString(s1, s2 string) int {
	lenS1 := len(s1)
	lenS2 := len(s2)
	var resultArray [][]int

	for i := 0; i < lenS1; i++ {
		tmp := make([]int, 0, lenS2)
		for j := 0; j < lenS2; j++ {
			if s1[i] == s2[j] {
				if i > 0 && j > 0 {
					last := resultArray[i-1][j-1]
					tmp = append(tmp, last+1)
				} else {
					tmp = append(tmp, 1)
				}
			} else {
				tmp = append(tmp, 0)
			}
		}
		resultArray = append(resultArray, tmp)
	}

	var max int
	for i := 0; i < lenS1; i++ {
		for j := 0; j < lenS2; j++ {
			if resultArray[i][j] > max {
				max = resultArray[i][j]
			}
		}
	}

	return max
}
