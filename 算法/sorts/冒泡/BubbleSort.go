package main

import "fmt"

type BubbleSort struct {
}

func (b *BubbleSort) BubbleSorts(num []int) {
	for i := 0; i < len(num)-1; i++ {
		for j := 0; j < len(num)-1-i; j++ {
			if num[j] > num[j+1] {
				mid := num[j]
				num[j] = num[j+1]
				num[j+1] = mid
			}
		}
	}
	b.Print(num)
}
func (b *BubbleSort) Print(num []int) {
	for _, v := range num {
		fmt.Printf("%v ", v)
	}
}

func main() {
	selectSort := new(BubbleSort)
	var num = []int{1, 3, 2, 6, 5}
	selectSort.BubbleSorts(num)
}
