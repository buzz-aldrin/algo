package main

import "fmt"

func main() {
	input := []int{64, 25, 12, 22, 11}

	quickSort(input, 0, len(input)-1)

	fmt.Println(input)
}

func quickSort(arr []int, x, y int) {
	if x < y {
		p := partition(arr, x, y)
		quickSort(arr, x, p-1)
		quickSort(arr, p+1, y)
	}
}

func partition(arr []int, x, y int) int {
	i := x - 1
	pivot := arr[y]

	for j := i; j <= y-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i] = pivot

	return i
}
