package quick_sort

func BasicQuickSort(list []int) []int {
	length := len(list)
	if length < 2 {
		return list
	}
	
	min := list[0]
	
	left := make([]int, 0)
	right := make([]int, 0)
	
	for _, item := range list[1:] {
		if item <= min {
			left = append(left, item)
		} else {
			right = append(right, item)
		}
	}
	newSlice := append(BasicQuickSort(left), min)
	return append(newSlice, BasicQuickSort(right)...)
}

/*
base on sort package, should implement function
*/

type Sequence []string

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


