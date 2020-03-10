/*
@Copyright:
*/
/*
@Time : 2020/3/9 16:05
@Author : teddy
@File : bloom
*/

package bloomfilter

import (
	"algorithm/bitmap"
	"github.com/spaolacci/murmur3"
	"hash"
	"hash/fnv"
)

type BloomKey interface {
	ToBytes() ([]byte, error)
}

type Filter struct {
	bitmap bitmap.BitMap

	k      int    // hash定位的次数
	m      int    /*每个item(可以认为字节)占用的bit数*/
	n      uint64 /*bloom slots总的item数*/
	//factor uint32 /*每个item(可以认为字节)占用的bit数*/
	fnv  hash.Hash64 // 默认使用与wiredtiger一样的 fnv和city
	murmur hash.Hash64
}

func NewBloomFilter(k int, m int) *Filter {

	f := &Filter{
		bitmap: bitmap.New(m),
		k:      k,
		m:      m,
		n:      0,
		fnv:    fnv.New64(),
		murmur:   murmur3.New64(),
	}
	return f
}

func (filter *Filter) calHash(key BloomKey) (h1 uint64, h2 uint64, err error) {
	var bs []byte
	if bs, err = key.ToBytes(); err != nil {
		return
	} else {
		filter.fnv.Reset()
		filter.fnv.Write(bs)
		h1 = filter.fnv.Sum64()

		filter.murmur.Reset()
		filter.murmur.Write(bs)
		h2 = filter.murmur.Sum64()
	}
	return
}

func (filter *Filter) Add(key BloomKey) error {
	h1, h2, err := filter.calHash(key)

	if err != nil {
		return err
	}
	filter.n++

	for i := 0; i < filter.k; i++ {
		bitsIdx := h1 % (uint64(filter.m))
		filter.bitmap.Put(bitsIdx)
		h1 += h2
	}
	return nil
}

func (filter *Filter) Exist(key BloomKey) bool {
	h1, h2, err := filter.calHash(key)
	if err != nil {
		// 计算错误
		return false
	}
	for i := 0; i < filter.k; i++ {
		bitsIdx := h1 % (uint64(filter.m))
		if !filter.bitmap.Exist(bitsIdx) {
			return false
		}
		h1 += h2
	}
	return true
}
