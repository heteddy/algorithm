/*
@Copyright:
*/
/*
@Time : 2020/6/10 20:13
@Author : teddy
@File : kreverse.go
*/

package linked_list

type ListNode struct {
	Val  string
	Next *ListNode
}

func reverseGroup(group []*ListNode) (head, tail *ListNode) {
	length := len(group)
	if length == 0 {
		return
	}

	head = group[length-1]

	tail = group[0]
	tail.Next = nil

	for idx := length - 1; idx > 0; idx-- {
		group[idx].Next = group[idx-1]
	}
	return
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}
	group := make([]*ListNode, k)
	var _head, tail *ListNode
	n := head
	groupIdx := 0
	for ; n != nil; n = n.Next {
		if groupIdx < k {
			group[groupIdx] = n
			groupIdx++
		} else if groupIdx == k {
			// 执行翻转 拼接
			groupHead, groupTail := reverseGroup(group)
			// 保证都不是nil
			if _head == nil || tail == nil {
				_head = groupHead
				tail = groupTail
			} else {
				tail.Next = groupHead
				tail = groupTail
			}
			groupIdx = 0
			// 放入group
			group[groupIdx] = n
			groupIdx++
		}

	}
	if groupIdx == k {
		groupHead, groupTail := reverseGroup(group)
		// 保证都不是nil
		if _head == nil || tail == nil {
			_head = groupHead
			tail = groupTail
		} else {
			tail.Next = groupHead
			tail = groupTail
		}
	} else if tail != nil {
		// 剩余的部分保持就好
		tail.Next = group[0]
	}
	return _head
}
