package stairs

import (
	"fmt"
	//"log"
)

func StairsWithoutCache(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return StairsWithoutCache(n-1) + StairsWithoutCache(n-2)

}

type CachedFunction func(int) int
type CachingFunc func(function CachedFunction) CachedFunction

//优化1： 使用备忘录算法保存中间结果,避免重复计算
func cache(f CachedFunction) CachedFunction {
	results := make(map[string]int)

	return func(i int) int {
		args := fmt.Sprint(i)
		//如果这里的参数有很多个
		//for _,v := range xs{
		//	args += fmt.Sprint(v)
		//}
		if value, found := results[args]; found {
			//log.Println("cached", i, " : ", value)
			return value
		} else {
			value := f(i)
			results[args] = value
			return value
		}
	}
}

var StairCounter CachedFunction

func init() {
	StairCounter = cache(func(n int) int {
		if n == 1 {
			return 1
		}
		if n == 2 {
			return 1
		}
		return StairCounter(n-1) + StairCounter(n-2)
	})
}

//优化2：既然每个都依赖于其前面的计算结果，我们就直接从第一个开始计算，直到计算到输入参数为止，这也是动态规划算法的常用解决思路

func SequenceStairCounter(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	var result int
	var preValue = 1
	var prePreValue = 1
	temp := 3

	for temp <= n {
		result = preValue + prePreValue
		prePreValue = preValue
		preValue = result

		temp++
	}
	return result
}
