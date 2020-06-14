/*
@Copyright:
*/
/*
@Time : 2020/6/9 23:02
@Author : teddy
@File : profit.go
*/

package stocks

//import "math"

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func maxInt32(a, b int32) int32 {
//	if a > b {
//		return a
//	}
//	return b
//}

/*


题解：
贪心思想：记录前面的最小价格，将这个最小价格作为买入价格，然后将当前的价格作为售出价格，查看当前收益是不是最大收益。
*/
func SingleTransaction(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var maxProfit, last = 0, 0

	for _idx := 0; _idx < len(prices)-1; _idx++ {
		last = maxInt(0, last+prices[_idx+1]-prices[_idx])
		maxProfit = maxInt(last, maxProfit)
	}
	return maxProfit
}

/*

*/
func MultiTransaction(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var maxProfit = 0

	for _idx := 0; _idx < len(prices)-1; _idx++ {
		if prices[_idx+1]-prices[_idx] > 0 {
			maxProfit += prices[_idx+1] - prices[_idx]
		}
	}
	return maxProfit
}

/*

*/
func TupleTransaction(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	/*
	  对于每一天，需要维护四个状态:
	      firstBuy:   考虑在该天第一次买入股票可获得的最大收益
	      firstSell:  考虑在该天第一次卖出股票可获得的最大收益
	      secondBuy:  考虑在该天第二次买入股票可获得的最大收益
	      secondSell: 考虑在该天第二次卖出股票可获得的最大收益

	  分别对四个变量进行相应的更新, 最后secSell就是最大，收益值(secSell >= fstSell)
	*/
	//firstBuy, firstSell, secondBuy, secondSell := math.MinInt32, 0, 0, 0
	var firstBuy, secondBuy, firstSell, secondSell int
	firstBuy, secondBuy = -1*prices[0], -1*prices[0]
	firstSell, secondSell = 0, 0
	for idx := 0; idx < len(prices); idx++ {
		firstBuy = maxInt(firstBuy, 0-prices[idx])
		firstSell = maxInt(firstSell, firstBuy+prices[idx])
		secondBuy = maxInt(secondBuy, firstSell-prices[idx])
		secondSell = maxInt(secondSell, secondBuy+prices[idx])
	}
	return secondSell
}

/*可以进行K次的股票交易

leetcode 188 买卖股票的最佳时机 IV（中等）
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

    输入: [2,4,1], k = 2
    输出: 2
    解释: 在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。

示例 2:

    输入: [3,2,6,5,0,3], k = 2
    输出: 7
    解释: 在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
    随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。
*/

type TransactionInfo struct {
	buy  int
	sell int
}

func KTransaction(prices []int, k int) int {
	// 交易需要买卖各一次，因此大于一半就表示任意次数
	if len(prices) == 0 {
		return 0
	}
	if k > len(prices)/2 {
		return MultiTransaction(prices)
	} else {
		//保存k次
		transactions := make([]TransactionInfo, k)

		for idx, _ := range transactions {
			transactions[idx].buy = -1 * prices[0]
			//transactions[idx].sell = 0
		}

		for pidx := 0; pidx < len(prices); pidx++ {
			for tidx := 0; tidx < k; tidx++ {
				if tidx == 0 {
					transactions[tidx].buy = maxInt(transactions[tidx].buy, 0-prices[pidx])
					transactions[tidx].sell = maxInt(transactions[tidx].sell, transactions[tidx].buy+prices[pidx])
				} else {
					transactions[tidx].buy = maxInt(transactions[tidx].buy, transactions[tidx-1].sell-prices[pidx])
					transactions[tidx].sell = maxInt(transactions[tidx].sell, transactions[tidx].buy+prices[pidx])
				}
			}
			//log.Println(transactions)
		}
		return transactions[k-1].sell
	}
}


/*
需要冷冻期的股票交易

leetcode 309 最佳买卖股票时机含冷冻期（中等）
给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

示例:

    输入: [1,2,3,0,2]
    输出: 3
    解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
*/