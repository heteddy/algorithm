/*
@Copyright:
*/
/*
@Time : 2020/5/4 10:42
@Author : teddy
@File : kv
*/

package skiplist

import (
	"bytes"
	"math/rand"
	"sync"
)

type bytesComparer struct {
}

func (bytesComparer) Compare(a, b []byte) int {
	return bytes.Compare(a, b)
}

type KVSkipList struct {
	cmp bytesComparer
	rnd *rand.Rand

	mu     sync.RWMutex
	kvData []byte
	// Node data:
	// [0]         : KV offset
	// [1]         : Key length
	// [2]         : Val length
	// [3]         : Height
	// [3..height] : Next nodes
	nodeData  []int
	prevNode  [12]int
	maxHeight int
	len       int
	kvSize    int
}

func (kv *KVSkipList) Get(key []byte) (value []byte, err error) {
	return
}

func (kv *KVSkipList) Put(key []byte, value []byte) (err error) {
	return
}
func (kv *KVSkipList) Find(key []byte) (rkey, value []byte, err error) {
	return
}
func (kv *KVSkipList) Contains(key []byte) bool {
	return false
}
