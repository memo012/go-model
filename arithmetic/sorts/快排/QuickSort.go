package main

type QuickSort struct {
}

func (q *QuickSort) QuickSorts(num []int, left int, right int) {
	if num == nil {
		return
	}
	temp := num[left]
	i := left
	j := right
	for i != j {
		for i < j && num[j] >= temp {
			j--
		}
		for i < j && num[i] <= temp {
			i++
		}
		if i < j {
			tmp := num[i]
			num[i] = num[j]
			num[j] = tmp
		}
	}
	num[left] = num[i]
	num[i] = temp
	q.QuickSorts(num, left, i - 1)
	q.QuickSorts(num, i + 1, right)
}

func main() {
	//q := new(QuickSort)
}
