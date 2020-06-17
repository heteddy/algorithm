/*
@Copyright:
*/
/*
@Time : 2020/6/17 17:08
@Author : teddy
@File : palindrome_test.go
*/

package palindromestring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	t.Log(longestPalindrome("abba"))
	t.Log(longestPalindrome("abcba"))
	t.Log(longestPalindrome("babad"))
}
