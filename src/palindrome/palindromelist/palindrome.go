/*
@Copyright:
*/
/*
@Time : 2020/6/16 17:06
@Author : teddy
@File : linked-list
*/

package palindromelist

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针压栈，遍历一遍，当快指针到达终点，慢指针停止压栈

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	var fast, slow = head, head
	count := 0
	stack := make([]*ListNode, 0)
	//node := head
	for slow != nil && fast != nil {
		stack = append(stack, slow)
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			// 这种情况长度是单数，弹出一个
			stack = stack[:len(stack)-1]
			break
		}
		count++
	}
	count -= 1
	//for i := 0; i <= count; i++ {
	//	fmt.Print(stack[i].Val)
	//}
	for slow != nil {
		if slow.Val != stack[count].Val {
			return false
		}
		slow = slow.Next
		count--
	}
	return true
}
