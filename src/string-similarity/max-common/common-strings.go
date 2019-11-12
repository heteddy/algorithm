/*
	1.最大公共子序列问题
	输入：2个字符串
	输出：最大公共子序列：比如 poll和pull最大子序列是3（公共子串长度为2）

	解题思路：
		采用动态规划方式
			创建结果二维数组，
			把字符串拆分成一个个的网格（2 维度），然后比较每个字符；
			如果相同则值为  紧邻斜对角值+1 即value（横坐标-1， 纵坐标-1） + 1（紧邻斜对角的值就是各后退一步的公共子串）
			否则为max(v[i-1]j,v[i][j-1])  因为是序列所以有累计效果

	2.最大公共子串问题，给定两个字符串，判定这两个字符串的最大公共子串
	解题方法：
		采用动态规划方式，把字符串拆分成一个个的网格（2 维度），然后比较每个字符；
			如果相同则值为  紧邻斜对角值+1 即value（横坐标-1， 纵坐标-1） + 1（紧邻斜对角的值就是各后退一步的公共子串）
			否则为0
		由于公共子串可以出现在2个输入字符串的任何位置，因此最大值有可能出现动态规划矩阵的任一位置；
		所以最大子串的长度，就是求动态规划矩阵中最大值
*/

package max_common

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
