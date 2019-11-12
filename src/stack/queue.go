/*
使用双端队列
创建一个双端队列管理最大值
*/

package Queue

import (
	"errors"
	"fmt"
)

type Element interface {
	LessThan(Element) bool
}
type IntElement int

func (i IntElement) LessThan(other Element) bool {
	if v, ok := other.(IntElement); ok {
		return i < v
	}
	return false

}

type Queue interface {
	AppendEnd(e Element)
	InsertFront(e Element)
	PopEnd() (Element, error)
	PopFront() (Element, error)
	Front() (Element, error)
	End() (Element, error)
	Empty() bool
	Size() int
	Clear()
}

type LinkedQueueNode struct {
	value    Element
	previous *LinkedQueueNode
	next     *LinkedQueueNode
}
type LinkedQueue struct {
	head   *LinkedQueueNode
	tail   *LinkedQueueNode
	length int
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}
func (q *LinkedQueue) AppendEnd(e Element) {
	fmt.Println("LinkedQueue", e)
	node := &LinkedQueueNode{
		value:    e,
		previous: nil,
		next:     nil,
	}
	if q.length == 0 {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		node.previous = q.tail
		q.tail = node
	}
	q.length++
}

func (q *LinkedQueue) InsertFront(e Element) {
	node := &LinkedQueueNode{
		value:    e,
		previous: nil,
		next:     nil,
	}

	if q.length == 0 {
		q.head = node
		q.tail = node
	} else {
		q.head.previous = node
		node.next = q.head
		q.head = node
	}
	q.length++
}

func (q *LinkedQueue) Front() (Element, error) {
	if q.length > 0 {
		return q.head.value, nil
	}
	return nil, errors.New("empty")
}
func (q *LinkedQueue) End() (Element, error) {
	if q.length > 0 {
		return q.tail.value, nil
	}
	return nil, errors.New("empty")
}
func (q *LinkedQueue) PopEnd() (Element, error) {
	if q.length > 0 {
		node := q.tail
		q.tail = node.previous
		q.length--
		return node.value, nil
	} else {
		return nil, errors.New("empty")
	}
}
func (q *LinkedQueue) PopFront() (Element, error) {
	if q.length > 0 {
		node := q.head
		q.head = node.next
		q.length--
		return node.value, nil
	} else {
		return nil, errors.New("empty")
	}
}
func (q *LinkedQueue) Empty() bool {
	return q.length == 0
}

func (q *LinkedQueue) Size() int {
	return q.length
}
func (q *LinkedQueue) Clear() {
	q.head = nil
	q.tail = nil
	q.length = 0
}
