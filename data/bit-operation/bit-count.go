package bit_operation

//一种解题思路是，将给定的数，循环除以2，直到不能整除或者值为1
//另外一种方式是找规律 2的幂的数 2(10)，4(100)，8(1000)，16(10000),32(100000)
//可以看到这种数二进制表示只有最高位为1 ，其余都是0
//另外一个规律是这些数，n&n-1一定是0；因此代码非常简单
func IsPowerOf2(num int) bool {
	if num&(num-1) != 0 {
		return false
	}
	return true
}

//统计二进制位为1的数量
func BitCount(num int) (count int) {
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return
}
