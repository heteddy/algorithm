package skiplist

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

// 经验值  1/4的升级概率
const SKIPLIST_P = 4

type Node struct {
	key   uint64
	level int
	// 指向后面的指针，与层数相关
	forward []*Node
}

type SkipList struct {
	maxLevel int
	header   *Node
}

func NewSkipList(maxLevel int) *SkipList {
	next := make([] *Node, maxLevel, maxLevel)

	node := &Node{
		key:     0,
		level:   maxLevel,
		forward: next,
	}

	return &SkipList{
		maxLevel: maxLevel,
		header:   node,
	}
}

func (s *SkipList) createNode(key uint64, level int) *Node {

	next := make([] *Node, s.maxLevel, s.maxLevel)

	return &Node{
		key:     key,
		level:   level,
		forward: next,
	}
}

func (s *SkipList) insertAfter(newNode, afterNode *Node, level int) {
	newNode.forward[level], afterNode.forward[level] = afterNode.forward[level], newNode
}

func (s *SkipList) Insert(key uint64) {
	// 创建一个node
	// 按层次查找
	newNode := s.createNode(key, s.randomLevel(key))
	node := s.header
	for i := s.maxLevel - 1; i >= 0; i-- {
		if i >= newNode.level {
			// 如果当前层的下一个元素不为nil，并且下一个元素的key < 当前key；向后移动
			for ; node.forward[i] != nil && node.forward[i].key < key; {
				node = node.forward[i]
			}
		} else { // 当前的level需要插入
			for ; node.forward[i] != nil && node.forward[i].key <= key; {
				node = node.forward[i]
			}
			if node.forward[i] == nil || node.forward[i].key > key {
				s.insertAfter(newNode, node, i)
			}
		}

	}

}
func (s *SkipList) Delete(key uint64) {
	node := s.header
	for j := s.maxLevel - 1; j >= 0; j-- {
		// TODO 不比较头结点
		for ; node.forward[j] != nil && node.forward[j].key < key; {
			// 向后移动
			node = node.forward[j]
		}
		if node.forward[j] == nil {
			//进入下一层
			continue
		}
		// 下一个节点为待删除节点
		if node.forward[j].key == key {
			// 删除当前层
			node.forward[j] = node.forward[j].forward[j]
			//进入下一层
		}
	}
}

// 给定一个key，查找是否包含这个key的node
func (s *SkipList) Search(key uint64) *Node {
	node := s.header
	// 先从最上面一层开始找
	for j := s.maxLevel - 1; j >= 0; j-- {
		for ; node.forward[j] != nil && node.forward[j].key < key; {
			// 向后移动
			node = node.forward[j]
		}
		if node.forward[j] == nil {
			//进入下一层
			continue
		}
		// 下一个节点为待删除节点
		if node.forward[j].key == key {
			// 删除当前层
			return node.forward[j]
		}

	}
	return nil
}

// randomLevel 基于概率，计算当前的层数；默认是1层，抛硬币的方式决定是否晋升到上层
// 层数如何使用
func (s *SkipList) randomLevel(key uint64) int {
	level := 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for (r.Uint32()&0xFFFF)%SKIPLIST_P == 0 {
		// %4 25%的概率提升到下一级
		level += 1
		r.Seed(time.Now().UnixNano())
	}
	if level > s.maxLevel {
		level = s.maxLevel
	}
	return level
}

func (s *SkipList) PrintList() {
	for i := s.maxLevel - 1; i >= 0; i-- {
		fmt.Printf("level[%2d]:", i)
		node := s.header
		for {
			// 打印当前
			fmt.Printf("%d->", node.key)
			if node.forward[i] != nil {
				node = node.forward[i]
			} else {
				break
			}
		}
		fmt.Println("nil")
	} //end for

}
