/*
	单链表反转
		问题描述：给定一个单链表，反转每个元素，使得表头变表尾，表尾成为表头
		解题思路：1.表头变表尾，使用压栈的方式；递归恰好可以自动压栈，因此可以用递归解题
				  2. 使用头插方式，从链表的第二个元素开始，依次插入到头部，知道最后一个元素
*/
package reverse

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type LinkedListNode struct {
	Value interface{}
	Next  *LinkedListNode
}

func (node *LinkedListNode) Reverse() (next, newHeader *LinkedListNode) {
	if node.Next == nil {
		return node, node
	}
	originalNext, newHeader := node.Next.Reverse()
	originalNext.Next = node
	node.Next = nil
	return node, newHeader
}

func (node *LinkedListNode) ReverseInsert() (newHeader *LinkedListNode) {
	var currentNode *LinkedListNode
	currentNode = node
	nextNode := node.Next
	//断开跟后面的联系
	node.Next = nil
	item := nextNode
	for item != nil {
		//保存原来链表的下一个元素
		tempNext := item.Next
		//插入元素到当前链表的前面
		item.Next = currentNode
		//修改当前的元素指针
		currentNode = item
		//迭代到下一个元素
		newHeader = item
		item = tempNext
	}
	log.Println("return block")
	return newHeader
}

func NewRandomIntList(length int) *LinkedListNode {
	var header *LinkedListNode
	var currentNode *LinkedListNode
	for length > 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		num := r.Intn(100)
		//firstNode
		if currentNode == nil {
			header = &LinkedListNode{Value: num, Next: nil}
			currentNode = header
		} else {
			temp := &LinkedListNode{Value: num, Next: nil}
			currentNode.Next = temp
			currentNode = temp
		}
		log.Println(num)
		length--
	}
	return header
}

func (node *LinkedListNode) String() string {
	var output string
	for temp := node; temp != nil; temp = temp.Next {
		output += fmt.Sprintf("%d->", temp.Value.(int))
	}
	return output[:len(output)-2]
}
