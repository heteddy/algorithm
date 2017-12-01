/*
	最大背包问题：
		问题描述：
			小偷带着一个背包，到商场偷东西，背包的容量确定，如果小偷想保证所偷价值最大化，应该选择偷哪些商品
		解题思路：
			1）采用暴力方法，采用组合的方式，1件商品 Cn1  2件商品 Cn2......如果有n个商品，复杂度为2的n次方
			2）动态规划，先解决小背包问题，然后再解决当前给定的背包：
				2.1 创建一个二维数组( 横坐标i：背包重量，纵坐标j：商品(价值和重量) )保存 当前的价值和选择的商品
				2.2 流程：当前的背包能放下的前提下，执行
						cell[i-1][j]  表示：少一件商品时背包装下的价值
						max( cell[i-1][j], 当前商品的价值+剩余空间的价值 )  //注意是当前物品而不是上一个表格的值
						剩余空间的价值怎么计算  Cell[i-1][j-当前商品重量]
*/
package max_package

import (
	"fmt"
	"math"
)

type GoodsList []Goods

// 拆分背包的重量，从最小的背包重量开始，可以优化的是，取商品的最小重量为最小背包重量的粒度；
// 比如最小商品重量为5，如果我们还从1还是动态规划，<5的部分没有任何意义
//
type GoodsTable [][]GoodsList

//使用list保存物品列表
type Goods struct {
	Name   string
	Value  int //"假设商品的价值都是整数"
	Weight int //"假设商品的重量为整数"
}

func max(p ...int) int {
	var max = 0
	for _, i := range p {
		if i > max {
			max = i
		}
	}
	return max
}

func New(cap int) GoodsList {
	l := make(GoodsList, 0, cap)
	return l
}

func (gs GoodsList) MinWeight() int {
	var min = math.MaxInt32
	for _, g := range gs {
		if g.Weight < min {
			min = g.Weight
		}
	}
	return min
}

func (gs GoodsList) SumValue() (sum int) {
	for _, g := range gs {
		sum += g.Value
	}
	return
}

func (gs GoodsList) MaxPackage(weight int) GoodsList {
	length := len(gs)
	//最后一个单元格的元素，就是最终的结果

	//用来保存中间结果
	gt := make(GoodsTable, 0, length)
	//直接从最小的，还是从最小的商品重量开始
	//

	for goodsIndex, g := range gs {
		//这个是每一行的数据
		goodsLine := make([]GoodsList, 0, length)
		for currentWeight := 1; currentWeight < weight+1; currentWeight++ {
			weightIndex := currentWeight - 1
			//这里是当前表格的数据
			goodsWeightItem := make(GoodsList, 0)

			//小于当前的重量，可以放到当前的表格中
			if g.Weight <= currentWeight {

				//如果是第一件商品
				if goodsIndex == 0 {
					goodsWeightItem = append(goodsWeightItem, g)
				} else { //goodsIndex > 0
					//其他商品在当前容积下的最大价值（上一行对应列）
					lastMaxValue := gt[goodsIndex-1][weightIndex].SumValue()

					leftWeight := currentWeight - g.Weight
					leftWeightIndex := leftWeight - 1
					//当前商品+剩余商品的价值
					if leftWeightIndex >= 0 { //装入该商品，背包还可以放入其他的东西
						currentGoodsAndLeftValue := g.Value + gt[goodsIndex-1][leftWeightIndex].SumValue()
						if currentGoodsAndLeftValue > lastMaxValue {
							// 当前的价值大，则加入当前的goods 列表中
							goodsWeightItem = append(goodsWeightItem, g)
							for _, g := range gt[goodsIndex-1][leftWeightIndex] {
								goodsWeightItem = append(goodsWeightItem, g)
							}
						} else {
							// 之前的价值大，直接复制之前的 goods
							for _, g := range gt[goodsIndex-1][weightIndex] {
								goodsWeightItem = append(goodsWeightItem, g)
							}
						}
					} else { //背包已经满
						goodsWeightItem = append(goodsWeightItem, g)
					}

				}
			} else { // goods重量大于当前的背包的重量，不能放入背包，要么使用上一个商品，要么就是nil
				if goodsIndex > 0 {
					lastGoods := gt[goodsIndex-1][weightIndex]
					for _, g := range lastGoods {
						goodsWeightItem = append(goodsWeightItem, g)
					}
				} else {

				}
			}
			goodsLine = append(goodsLine, goodsWeightItem)
		}
		gt = append(gt, goodsLine)
	}
	fmt.Println(gt)
	return gt[weight-1][length-1]
}

func (gt GoodsTable) String() string {
	var output string
	for lineIndex, lineGoodsList := range gt {
		output += fmt.Sprintf("line%d:", lineIndex)
		for itemIndex, itemGoodsList := range lineGoodsList {
			output += fmt.Sprintf("item%d:(", itemIndex)
			for _, g := range itemGoodsList {
				output += fmt.Sprintf(" %s", g.Name)
			}
			output += " )\t"
		}
		output += "\n"
	}
	return output
}
