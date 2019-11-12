package bitmap

import (
	//"math"
	"log"
)

type BitMap []uint8

//如果是40亿个int32 则直接创建1<<32-1个bit，那么就是
func New() BitMap {
	//len := math.MaxInt32>>5 + 1
	len := 2
	return make([]byte, len, len)

}

func (bitmap BitMap) Init() {
	for _, v := range bitmap {
		v &= uint8(0)
	}
}
func (bitmap BitMap) arrayIndex(item int32) (index int32, position int32) {
	index = item >> 3
	position = item & 0x7
	return
}

//初始化为全0，当放入新的数的时候，如果对应位是0，则改为1，如果对应是1，则输入已经重复
func (bitmap BitMap) Put(item int32) bool {
	index, position := bitmap.arrayIndex(item)
	log.Println(item, bitmap)
	if bitmap[index]&(1<<uint8(position)) != 0 {
		return false
	} else {
		bitmap[index] |= 1 << uint8(position)
		return true
	}

}

//给定一个数，如果对应位已经是1 则已经输出true，否则输出false
func (bitmap BitMap) Contain(item int32) bool {
	index, position := bitmap.arrayIndex(item)
	if bitmap[index]&(1<<uint8(position)) != 0 {
		return true
	}
	return false
}
