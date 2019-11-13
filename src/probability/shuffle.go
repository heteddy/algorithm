package probability

import (
	"math/rand"
	"time"
)

func Shuffle(input []int) {

	length := len(input)

	for i := 0; i < length; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		if i > 0 {
			k := r.Int() % i
			temp := input[i]
			input[i] = input[k]
			input[k] = temp
		}

	}
}
