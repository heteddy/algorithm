/*
@Copyright:
*/
/*
@Time : 2020/6/9 23:14
@Author : teddy
@File : profit_test.go
*/

package stocks

import "testing"

func TestMaxProfit(t *testing.T) {
	prices := []int{
		7, 1, 5, 3, 6, 3, 8, 5, 10, 12, 6, 18,
	}
	t.Log(SingleTransaction(prices))
	t.Log(MultiTransaction(prices))
	t.Log(TupleTransaction(prices))

}

func TestKTransaction(t *testing.T) {
	prices := []int{
		7, 1, 5, 3, 6, 3, 8, 5, 10, 12, 6, 18,
	}
	t.Log("k=2", KTransaction(prices, 2))
	t.Log("k=3", KTransaction(prices, 3))
	t.Log("k=4", KTransaction(prices, 4))
	t.Log("k=5", KTransaction(prices, 5))
}
