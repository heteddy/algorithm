/*
@Copyright:
*/
/*
@Time : 2020/6/17 16:53
@Author : teddy
@File : palindrome
*/

package palindromestring

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	dst := insertDummyChar(s)
	_, ps := longPalindromeStr(dst)
	pStr := ""
	for _idx := 0; _idx < len(ps); _idx++ {
		if _idx%2 != 0 {
			pStr += string(ps[_idx])
		}
	}
	return pStr
}

func insertDummyChar(s string) string {
	var dst string
	for _, c := range s {
		dst += "#"
		dst += string(c)
	}
	dst += "#"
	return dst
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longPalindromeStr(s string) (int, string) {
	length := len(s)

	maxPalindrome := 0
	ps := ""
	for idx := 1; idx < length-1; idx++ {
		// 循环比对前和后
		pre := idx - 1
		next := idx + 1
		_pal := 0
		for pre >= 0 && next < length {
			if s[pre] == s[next] {
				_pal++
			} else {
				break
			}
			pre--
			next++
		}
		//maxPalindrome = max(maxPalindrome, _pal)
		if _pal > maxPalindrome {
			maxPalindrome = _pal
			ps = s[pre+1 : next]
		}
	}
	return maxPalindrome, ps
}
