package main

import "fmt"

func main() {
	input := []int{64, 25, 12, 22, 11}

	mergeSort(input, 0, len(input)-1)

	fmt.Println(input)
}

func mergeSort(input []int, l, r int) {
	if l < r {
		mid := l + (r-l)/2
		mergeSort(input, l, mid)
		mergeSort(input, mid+1, r)

		merge(input, l, mid, r)
	}
}

func merge(input []int, l, mid, r int) {
	temp := make([]int, 0, r-l)
	left1, left2 := l, mid+1

	for left1 < left2 && left2 <= r {
		if input[left1] < input[left2] {
			temp = append(temp, input[left1])
			left1++
		} else {
			temp = append(temp, input[left2])
			left2++
		}
	}

	for left1 <= mid {
		temp = append(temp, input[left1])
		left1++
	}

	for left2 <= r {
		temp = append(temp, input[left2])
		left2++
	}

	for i := 0; i < len(temp); i++ {
		input[l] = temp[i]
		l++
	}
}
