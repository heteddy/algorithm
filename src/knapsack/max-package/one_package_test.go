package max_package

import (
	"testing"
)

func TestGoodsList_MaxPackage(t *testing.T) {
	goodsList := New(4)

	guitar := Goods{Name: "guitar", Value: 1500, Weight: 1}
	laptop := Goods{Name: "laptop", Value: 2000, Weight: 3}
	sound := Goods{Name: "sound", Value: 3000, Weight: 4}
	iphone := Goods{Name: "iphone", Value: 2000, Weight: 1}

	goodsList = append(goodsList, guitar, laptop, sound, iphone)
	resultGoods := goodsList.MaxPackage(4)
	var sum int
	for i, g := range resultGoods {
		t.Logf("%02d goods:%s %d %d", i, g.Name, g.Weight, g.Value)
		sum += g.Value
	}
	t.Log("sum", sum)
}