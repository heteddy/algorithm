/*
@Copyright:
*/
/*
@Time : 2020/6/11 09:51
@Author : teddy
@File : kreverse_test.go
*/

package linked_list

import (
	"strconv"
	"testing"
)

func constructList(n int) *ListNode {
	var l *ListNode
	var tail *ListNode
	for i := 0; i < n; i++ {
		node := ListNode{
			Val:  strconv.Itoa(i),
			Next: nil,
		}
		if tail == nil {
			l = &node
			tail = &node
		} else {
			tail.Next = &node
			tail = &node
		}
	}
	tail.Next = nil
	return l
}

func TestKGroupReverse(t *testing.T) {
	node := constructList(2)
	n := node
	for n != nil {
		t.Log(n.Val)
		n = n.Next
	}
	t.Log("after group reverse")
	node = reverseKGroup(node, 2)

	for node != nil {
		t.Log(node.Val)
		node = node.Next
	}
}
