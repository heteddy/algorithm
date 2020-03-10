package bitmap

import "log"

type BitMap []uint8

func New(len int) BitMap {
	size := len >> 3
	if len&0x7 != 0 {
		size += 1
	}
	bm := make([]byte, size, size)
	return bm
}

func (bitmap BitMap) arrayIndex(item uint64) (index uint64, position uint64) {
	index = item >> 3
	position = item & 0x7
	return
}

// Put: 返回值仅表示可能的冲突或者重复风险
// 放入时如果对应位置已经被设置过返回false，如果对应位置为0则返回true；
func (bitmap BitMap) Put(item uint64) bool {
	index, position := bitmap.arrayIndex(item)
	defer func() {
		log.Println(index, position, item, bitmap)
	}()
	if bitmap[index]&(1<<position) != 0 {
		return false
	} else {
		bitmap[index] |= 1 << position
		return true
	}
}

//给定一个数，如果对应位已经是1 则已经输出true，否则输出false
func (bitmap BitMap) Exist(item uint64) bool {
	index, position := bitmap.arrayIndex(item)
	if bitmap[index]&(1<<uint8(position)) != 0 {
		return true
	}
	return false
}

// 计算bitmap占用了多少
func (bitmap BitMap) Count() int {
	count := 0
	for _, num := range bitmap {
		for num != 0 {
			num = num & (num - 1)
			count++
		}
	}
	return count
}
