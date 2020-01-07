package skiplist

import (
	"math/rand"
	"testing"
	"time"
)

func TestSkipList_Insert(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	skipList := NewSkipList(32)

	for i := 0; i < 30; i++ {
		skipList.Insert(rand.Uint64() & 0xFFFF)
	}
	skipList.PrintList()
}

func TestSkipList_Delete(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	skipList := NewSkipList(32)

	deleteItems := make([]uint64, 0)

	for i := 0; i < 30; i++ {
		key := rand.Uint64() & 0xFFFFFFFF
		if i%6 == 0 {
			deleteItems = append(deleteItems, key)
		}
		skipList.Insert(key)
	}
	skipList.PrintList()

	for i := len(deleteItems); i > 0; i-- {
		skipList.Delete(deleteItems[i-1])
	}
	skipList.PrintList()
}

func TestSkipList_Search(t *testing.T) {

}
