package bit_operation

func Add(a, b uint64) uint64 {
	//此题目需要关注最大值；即边界问题
	if b == 0 {
		return a
	}
	var sum, t uint64
	sum = a ^ b
	t = (a & b) << 1
	// 也可以用循环
	return Add(sum, t)
}
