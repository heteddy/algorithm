/*
@Copyright:
*/
/*
@Time : 2020/6/12 15:56
@Author : teddy
@File : merge_sequence_test.go
*/

package merge

import (
	"math/rand"
	"testing"
	"time"
)

func constructList(n int, suffix string) *ListNode {
	var l *ListNode
	var tail *ListNode

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := r.Intn(10)
	for i := 0; i < n; i++ {
		node := ListNode{
			Val:  s + i,
			from: suffix,
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

func TestMerge2List(t *testing.T) {
	node1 := constructList(16, "node1")
	node2 := constructList(11, "node2")
	n1, n2 := node1, node2

	t.Log("log1")
	t.Log(n1)
	t.Log("log2")
	t.Log(n2)

	head := merge2List(node1, node2)

	t.Log("after merge")
	t.Log(head)
}

func TestMergeKList(t *testing.T) {
	node1 := constructList(16, "node1")
	node2 := constructList(11, "node2")
	node3 := constructList(12, "node3")
	node4 := constructList(5, "node4")
	n1, n2, n3, n4 := node1, node2, node3, node4
	lists := []*ListNode{
		n1, n2, n3, n4,
	}
	for _, l := range lists {
		t.Log(l)
	}

	t.Log("after merge")
	head := mergeKLists(lists)
	t.Log(head)
}



func TestGenerateMergeTuple(t *testing.T) {
	generateMergeTuple(15)
	generateMergeTuple(16)
	generateMergeTuple(17)
}

func TestTupleGenerator_Generate(t *testing.T) {
	g := NewTupleGenerator(15)
	for {
		if tuple, err := g.Generate(); err != nil {
			break
		} else {
			t.Log(tuple)
		}
	}
}
