package maxArray

import (
	"errors"
	"fmt"

	Queue "./stack/Queue"
)

type IntArrayElement struct {
	value int
	index int
}

func (e IntArrayElement) LessThan(other Queue.Element) bool {
	if _other, ok := other.(IntArrayElement); ok {
		return e.value < _other.value
	}
	return false
}
func MaxArray(array []int, w int) error {
	if len(array) < w {
		return errors.New("窗口太大")
	}
	linkedQueue := Queue.NewLinkedQueue()
	// _input 是int 类型
	for index, _input := range array {
		input := IntArrayElement{value: _input, index: index}
		for {
			if linkedQueue.Empty() {
				linkedQueue.AppendEnd(input)
				if index >= w-1 {
					if v, err := linkedQueue.Front(); err == nil {
						fmt.Println(v)
					}
				}
				break
			} else if v, err := linkedQueue.End(); err == nil {
				if v.LessThan(input) {
					linkedQueue.PopEnd()
					continue
				} else {
					linkedQueue.AppendEnd(input)
					// 检查队列头的下标是否过期,即窗口是否已经移走
					// 过期最大值移走，放在appendEnd前后都一样的
					if _v, err := linkedQueue.Front(); err == nil {
						if v, ok := _v.(IntArrayElement); ok {
							//front 的index
							if index > w && v.index == (index-w) {
								linkedQueue.PopFront()
							}
						}
					}
					if index >= w-1 {
						if v, err := linkedQueue.Front(); err == nil {
							fmt.Println(v)
						}
					}
					break
				}
			} else {
				break
			}
		}
	}
	return nil
}
