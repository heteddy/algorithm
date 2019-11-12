package main

import (
	"algorithm/sort/select-sort"
	"log"
	"algorithm/sort/quick-sort"
	"sort"
)


func main() {
	
	selectInput := []int{20, 12, 34, 56, 12, 45, 89, 2, 4}
	quickInput := make([]int,len(selectInput))
	copy(quickInput,selectInput)
	
	select_sort.SelectSort(selectInput)
	log.Println("select-sort output:",selectInput)
	
	log.Println("quick-sort output:",quick_sort.BasicQuickSort(quickInput))
	
	stringSortInput := quick_sort.Sequence{"hello world","helloTeddy","hello Teddy","helloteddy","hello teddy"}
	//only accept interface as a arg
	sort.Sort(stringSortInput)
	log.Println(stringSortInput)
}
