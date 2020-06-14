/*
@Copyright:
*/
/*
@Time : 2020/6/12 15:00
@Author : teddy
@File : merge_seqence.go
*/

package merge

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	from string
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}
	s := ""
	n := l
	for ; n != nil; {
		s += strconv.Itoa(n.Val) + ":" + n.from
		n = n.Next
		if n != nil {
			s += " -> "
		}
	}
	return s
}

func merge2List(list1, list2 *ListNode) *ListNode {
	var head, tail *ListNode
	aNode, bNode := list1, list2

	if aNode.Val < bNode.Val {
		head = aNode
		tail = aNode
		aNode = aNode.Next
	} else {
		head = bNode
		tail = bNode
		bNode = bNode.Next
	}

	for ; aNode != nil && bNode != nil; {
		if aNode.Val < bNode.Val {
			tail.Next = aNode
			tail = aNode
			aNode = aNode.Next
			//log.Println("move anode",tail.from)
		} else {
			tail.Next = bNode
			tail = bNode
			bNode = bNode.Next
			//log.Println("move bnode",tail.from)
		}

	}
	if aNode == nil {
		tail.Next = bNode
	} else if bNode == nil {
		tail.Next = aNode
	}
	return head
}

type Tuple struct {
	from, to int
}

type TupleGenerator struct {
	input    int
	interval int
	round    int
	index    int
	step     int
}

func NewTupleGenerator(n int) *TupleGenerator {
	_log2 := int(math.Log2(float64(n)))
	if int(math.Exp2(float64(_log2))) < n {
		_log2++
	}
	return &TupleGenerator{
		input:    n,
		interval: 1,
		round:    _log2,
		index:    1,
		step:     0,
	}
}

func (t *TupleGenerator) Generate() ([]Tuple, error) {
	tuple := make([]Tuple, 0, 1)
	if t.index <= t.round {
		fmt.Printf("第%d轮：", t.index)
		// step 呈幂的方式递进
		step := int(math.Exp2(float64(t.index)))
		for i := 0; step*i < t.input; i++ {
			if step*i+t.interval < t.input {
				//fmt.Fprintf("(%d,%d)", i*step, step*i+t.interval)
				tuple = append(tuple, Tuple{i * step, step*i + t.interval})
			} else {
				fmt.Printf("(%d)", i*step)
			}
		}
		t.interval *= 2
		//fmt.Println("")
	} else {
		return nil, errors.New("end")
	}
	t.index++
	return tuple, nil
}

func generateMergeTuple(n int) {
	/*
		第1轮：(0,1)(2,3)(4,5)(6,7)(8,9)(10,11)(12,13)(14,15)(16)
		第2轮：(0,2)(4,6)(8,10)(12,14)(16)
		第3轮：(0,4)(8,12)(16)
		第4轮：(0,8)(16)
		第5轮：(0,16)
		第6轮：(0)
	*/
	interval := 1

	_log2 := int(math.Log2(float64(n)))
	if int(math.Exp2(float64(_log2))) < n {
		_log2++
	}
	//
	fmt.Println(_log2)
	for round := 1; round <= _log2; round++ {
		fmt.Printf("第%d轮：", round)
		// step 呈幂的方式递进
		step := int(math.Exp2(float64(round)))
		for i := 0; step*i < n; i++ {
			if step*i+interval < n {
				fmt.Printf("(%d,%d)", i*step, step*i+interval)
			} else {
				fmt.Printf("(%d)", i*step)
			}
		}
		interval *= 2
		fmt.Println("")
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 1 {
		return lists[0]
	}
	// 实际上直接按序合并就可以了，首先0，1 然后合并的结果
	gen := NewTupleGenerator(length)
	for {
		if tuples, err := gen.Generate(); err != nil {
			break
		} else {
			for _, _tuple := range tuples {
				lists[_tuple.from] = merge2List(lists[_tuple.from], lists[_tuple.to])
			}
		}
	}
	return lists[0]
}
