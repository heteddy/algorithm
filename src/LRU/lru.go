package LRU

import (
	"fmt"
	"log"
)

/*
使用 HashMap 存储 key，这样可以做到 save 和 get key的时间都是 O(1)，而 HashMap 的 Val 指向双向链表实现的 LRU 的 Node 节点
*/

type List struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	Next  *Node
	Pre   *Node
	key   string
	Value interface{}
}

func (l *List) IsEmpty() bool {
	return l.head == nil && l.tail == nil
}

func NewList() *List {
	return &List{
		head: nil,
		tail: nil,
	}
}

func (l *List) InsertToFront(n *Node) {

	if l.IsEmpty() {
		l.head = n
		l.tail = n
	} else {
		n.Next = l.head
		l.head.Pre = n
		l.head = n
		n.Pre = nil
	}

	l.length++

}

func (l *List) MoveNodeToFront(n *Node) {
	if l.length == 1 {
		log.Println("只有一个节点，应该是头结点直接返回")
		return
	}
	// 位于开头
	if l.head == n {
		log.Println("头结点直接返回")
		return
	}
	// 最后一个节点
	if l.tail == n {
		log.Println("尾结点,")
		previousNode := n.Pre
		previousNode.Next = nil

		l.tail = previousNode
		n.Next = l.head
		l.head.Pre = n
		l.head = n
		n.Pre = nil
		return
	}

	// 多个节点 而且位于中间位置
	previousNode := n.Pre
	previousNode.Next = n.Next
	n.Next.Pre = previousNode

	n.Next = l.head
	l.head.Pre = n
	l.head = n
	n.Pre = nil
}

func (l *List) RemoveLast() {
	last := l.tail
	if last != nil {
		if last.Pre != nil {
			last.Pre.Next = nil
			l.tail = last.Pre
			l.length--
		} else {
			log.Println("这里错误，pre不能为空")
		}
	}
}

type LRUCache struct {
	List
	keyTable map[string]*Node
	capacity int
}

func NewLRUCache(length int) *LRUCache {
	return &LRUCache{
		List:     *NewList(),
		keyTable: make(map[string]*Node, length),
		capacity: length,
	}
}

func (c *LRUCache) Get(key string) interface{} {
	if _v, ok := c.keyTable[key]; ok {
		c.List.MoveNodeToFront(_v)
		return _v.Value
	} else {
		return nil
	}
}

func (c *LRUCache) Save(key string, v interface{}) {

	if _v, ok := c.keyTable[key]; ok {
		_v.Value = v
		// 移到最前
		c.List.MoveNodeToFront(_v)
	} else {
		if c.length == c.capacity {
			// 删除最后一个value？最后一个key怎么办
			lastKey := c.tail.key
			c.List.RemoveLast()
			//如何删除最后一个key
			delete(c.keyTable, lastKey)
		}

		node := &Node{
			Next:  nil,
			Pre:   nil,
			key:   key,
			Value: v,
		}
		c.List.InsertToFront(node)
		c.keyTable[key] = node
	}

}

func (c *LRUCache) String() string {
	node := c.head
	tail := c.tail
	var str string
	for {
		str += fmt.Sprintf("%s:%d; ", node.key, node.Value.(int))
		if node != tail {
			node = node.Next
		} else {
			break
		}
	}
	return str
}
