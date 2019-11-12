package min

/**/

import (
	"errors"
	"fmt"
)

type Stack struct {
	body      []int
	min       []int
	length    int
	bodyCount int
	minCount  int
}

func NewStack(length int) *Stack {
	return &Stack{
		make([]int, length, length),
		make([]int, length, length),
		length,
		0,
		0,
	}
}
func (stack *Stack) Full() bool {
	return stack.bodyCount >= stack.length
}
func (stack *Stack) Empty() bool {
	return stack.bodyCount == 0
}
func (stack *Stack) Push(element int) (ret bool) {
	if stack.Empty() {
		stack.body[stack.bodyCount] = element
		stack.min[stack.bodyCount] = element
		stack.bodyCount++
		stack.minCount++
		return true

	} else if !stack.Full() {
		stack.body[stack.bodyCount] = element
		stack.bodyCount++
		if element <= stack.min[stack.minCount-1] {
			stack.min[stack.minCount] = element
			stack.minCount++
		}
		return true
	} else {

		fmt.Println("栈满")
		return false
	}
}

func (stack *Stack) Pop() (int, error) {
	if stack.Empty() {
		return -1, errors.New("empty")
	} else {
		stack.bodyCount--
		ret := stack.body[stack.bodyCount]

		min := stack.min[stack.minCount-1]
		if ret == min {
			stack.minCount--
		}
		return ret, nil
	}
}

func (stack *Stack) GetMin() (int, error) {
	if stack.Empty() {
		return -1, errors.New("empty")
	} else {
		ret := stack.min[stack.minCount-1]
		return ret, nil
	}
}
