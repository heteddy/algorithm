/*
@Copyright:
*/
/*
@Time : 2020/6/14 09:06
@Author : teddy
@File : heap_merge_sequence_list_test.go
*/

package merge

import "testing"

func TestHeapMergeKList(t *testing.T) {
	node1 := constructList(16, "node1")
	node2 := constructList(11, "node2")
	node3 := constructList(12, "node3")
	node4 := constructList(5, "node4")
	n1, n2, n3, n4 := node1, node2, node3, node4
	lists := []*ListNode{
		n1, n2, n3, n4,
	}
	for idx, l := range lists {
		t.Log("#", idx, ":", l)
	}

	t.Log("after merge")
	head := HeapMergeKLists(lists)
	t.Log(head)
}

func TestHeapMergeKListEmpty(t *testing.T) {

	lists := []*ListNode{
		nil, nil,
	}
	for idx, l := range lists {
		t.Log("#", idx, ":", l)
	}

	t.Log("after merge")
	head := HeapMergeKLists(lists)
	t.Log(head)
}
