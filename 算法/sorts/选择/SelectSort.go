package main

import "fmt"

type SelectSort struct {
}

func (s *SelectSort) Select(num []int) {
	for i := 0; i < len(num); i++ {
		mid := num[i]
		for j := i + 1; j < len(num); j++ {
			if mid > num[j] {
				num[i] = num[j]
				num[j] = mid
			}
		}
	}
	s.Print(num)
}

func (s *SelectSort) Print(num []int)  {
	for _, v := range num {
		fmt.Printf("%v ", v)
	}
}

func main() {
	selectSort := new(SelectSort)
	var num = []int{1, 3, 2, 6, 5}
	selectSort.Select(num)
}
