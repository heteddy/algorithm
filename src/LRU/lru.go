package LRU

/*
使用 HashMap 存储 key，这样可以做到 save 和 get key的时间都是 O(1)，而 HashMap 的 Value 指向双向链表实现的 LRU 的 Node 节点

*/

type List struct {
	head *Node
	tail *Node

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
	n.Next = l.head
	n.Pre = nil
	l.head = n

	l.length++

}

func (l *List) MoveNodeToFront(n *Node) {
	previousNode := n.Pre
	nextNode := n.Next
	if previousNode != nil {
		previousNode.Next = n.Next

	}
	if nextNode != nil {
		nextNode.Pre = previousNode
	}
	n.Next = l.head
	n.Pre = nil
	l.head = n
}

func (l *List) RemoveLast() {
	last := l.tail
	if last != nil {
		if last.Pre != nil {
			last.Pre.Next = nil
			l.length--
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
