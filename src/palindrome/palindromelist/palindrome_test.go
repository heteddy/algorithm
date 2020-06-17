/*
@Copyright:
*/
/*
@Time : 2020/6/16 17:36
@Author : teddy
@File : palindrome_test.go
*/

package palindromelist

import "testing"

func TestIsPalindrome(t *testing.T) {
	// 1->2->3>2->1
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}
	t.Log(isPalindrome(head))
}
