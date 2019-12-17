package bit_operation

func Add(a, b uint64) uint64 {
	//此题目需要关注最大值；即边界问题
	var sum, t uint64

	for {
		sum = a ^ b
		t = (a & b) << 1

		if t != 0 {
			a = t
			b = sum
		} else {
			return sum
		}
	}
}
