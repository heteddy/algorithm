package main

import (
	"algorithm/data/bitmap"
	"log"
	"math/rand"
	"time"
)

func main() {
	var count int
	bitmap := bitmap.New()

	bitmap.Init()
	addedNum := make(map[int32]bool)
	for count < 15 {

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		num := r.Int31n(15)
		//这里put成功表示不存在，与map表示的概念相反
		ok := bitmap.Put(int32(num))
		_, existed := addedNum[num]
		if !existed {
			addedNum[num] = true
		}
		//如果存在，ok为false existed 为true
		if existed == ok {
			log.Fatalf("% 30s", "error")
		}
		count++
	}
	log.Fatalf("% 30s", "complete!!!")
}
