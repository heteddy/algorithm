package probability

import (
	"math/rand"
	"time"
)

func PickMRecursively(input []int, m int) []int {
	if m > len(input) {
		return nil
	}
	ret := make([]int, m)
	for i, v := range input {
		if i < m {
			ret[i] = v
		}
	}

	for i := m; i < len(input); i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		//fmt.Printf("%d ", r.Int31())

		k := r.Int() % i
		// todo： 注意这里是k<m, 不是每次都换
		if k < m {
			ret[k] = input[i]
		}

	}

	return ret

}