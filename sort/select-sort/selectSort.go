package select_sort

//import "log"

func SelectSort(list []int) {
	var index int
	for i, _:= range list {
		min := list[i]
		for j, item := range list[i+1:] {
			if item < min {
				min = item
				index = i + j + 1
			}
		}
		if index != i {
			list[i], list[index] = list[index], list[i]
		}
		
	}
	return
}
