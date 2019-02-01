package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 3, 4, 10, 40}
	x := 10
	result := binarySearch(arr, 0, len(arr)-1, x)
	if result == -1 {
		fmt.Println("Element is not present in array")
	} else {
		fmt.Printf("Element is present at index %d\n",
			result)
	}
}

func jumpSearch(arr []int, x int) int {
	jump := int(math.Sqrt(float64(len(arr))))

	for arr[min(len(arr)-1, jump)] < x {
		jump += jump
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
