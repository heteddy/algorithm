package main

import (
	"algorithm/knapsack/max-package"
	"log"
)

func main() {
	goodsList := max_package.New(4)

	guitar := max_package.Goods{Name: "guitar", Value: 1500, Weight: 1}
	laptop := max_package.Goods{Name: "laptop", Value: 2000, Weight: 3}
	sound := max_package.Goods{Name: "sound", Value: 3000, Weight: 4}
	iphone := max_package.Goods{Name: "iphone", Value: 2000, Weight: 1}

	goodsList = append(goodsList, guitar, laptop, sound, iphone)
	resultGoods := goodsList.MaxPackage(4)
	var sum int
	for i, g := range resultGoods {
		log.Printf("%02d goods:%s %d %d", i, g.Name, g.Weight, g.Value)
		sum += g.Value
	}
	log.Println("sum", sum)
}
