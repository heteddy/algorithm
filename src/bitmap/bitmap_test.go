package bitmap

import (
	"math/rand"
	"testing"
	"time"
)

//func TestBitMap_Put(t *testing.T) {
//	var maxNum = 20
//	var addedNum = make(map[uint64]bool)
//	var bitmap = New(maxNum)
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	num := uint64(r.Int31n(int32(maxNum)))
//	//这里put成功表示不存在，与map表示的概念相反
//	bitmap.Put(num)
//	_, existed := addedNum[num]
//	if !existed {
//		addedNum[num] = true
//	}
//}

func TestBitMap_Exist(t *testing.T) {
	var maxNum = 20
	var randomInt uint64
	var bitmap = New(maxNum)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomInt = uint64(r.Int31n(int32(maxNum)))

	bitmap.Put(randomInt)
	if !bitmap.Exist(randomInt) {
		t.Error("bitmap existed failure")
	}
}
