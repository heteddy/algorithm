package prime

func GeneratePrime(max int) []int {
	ret := make([]int, 0)
	ret = append(ret, 2)
	for i := 3; i < max; i += 2 {
		var isPrime = true
		for _, v := range ret {
			if i%v == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			ret = append(ret, i)
		}
	}
	return ret
}
