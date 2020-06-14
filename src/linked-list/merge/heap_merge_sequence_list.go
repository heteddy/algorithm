/*
@Copyright:
*/
/*
@Time : 2020/6/13 21:15
@Author : teddy
@File : merge_sequence_list2.go
*/

package merge

type Heap struct {
	slice []*ListNode
}

// 堆排序和构建堆过程相反；从小到大的排序，需要构建大顶堆，然后与最后一个交换；
func (h *Heap) Build() {

	if len(h.slice) <= 1 {
		return
	}

	// 调整左右子树
	for i := (len(h.slice) - 1) / 2; i >= 0; i-- {
		h.adjust(i, len(h.slice))
	}
	// 到这里已经构建完成，如果要整个有序就要循环把最大值落位

}

func (h *Heap) Pop() (*ListNode, bool) {
	if len(h.slice) == 0 {
		return nil, true
	}
	if len(h.slice) == 1 {
		node := h.slice[0]
		h.slice = h.slice[1:]
		return node, true
	} else {
		if h.slice[0].Next != nil {
			node := h.slice[0]
			h.slice[0] = h.slice[0].Next
			h.adjust(0, len(h.slice))
			return node, false
		} else {
			// 堆顶元素
			node := h.slice[0]
			h.slice = h.slice[1:]
			if len(h.slice) > 1 {
				h.Build()
			}
			return node, false
		}
	}
	//return nil
}

// 构建小顶堆
func (h *Heap) adjust(start, end int) {
	// 不包括end
	leftChildIndex := 2*start + 1
	rightChildIndex := leftChildIndex + 1

	smallIndex := leftChildIndex

	if leftChildIndex >= end {
		return
	}
	if rightChildIndex < end && h.slice[rightChildIndex].Val < h.slice[leftChildIndex].Val {
		// 左右孩子中较小的节点
		smallIndex = rightChildIndex
	}
	// 父节点与左右孩子中较小的一个比较，如果子节点小，那么交换
	if h.slice[smallIndex].Val < h.slice[start].Val {
		h.slice[smallIndex], h.slice[start] = h.slice[start], h.slice[smallIndex]
		h.adjust(smallIndex, end)
	}
}

func HeapMergeKLists(lists []*ListNode) *ListNode {
	temp := make([]*ListNode, 0, len(lists))
	// 去掉空链表
	for _, l := range lists {
		if l != nil {
			temp = append(temp, l)
		}
	}
	if len(temp) == 0 {
		return nil
	}
	if len(temp) == 1 {
		return temp[0]
	}
	h := Heap{slice: temp}
	h.Build()
	var head, tail *ListNode
	var node *ListNode
	var complete bool

	for {
		node, complete = h.Pop()
		if node == nil {
			break
		}
		if head == nil {
			head = node
		} else {
			tail.Next = node
		}
		tail = node
		if complete {
			break
		}
	}
	return head
}
