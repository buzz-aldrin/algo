package main

import "fmt"

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

func binarySearch(arr []int, l, r int, x int) int {
	if l <= r {
		mid := l + (r-l)/2
		if arr[mid] == x {
			return mid
		}

		if arr[mid] > x {
			return binarySearch(arr, l, mid-1, x)
		} else {
			return binarySearch(arr, mid+1, r, x)
		}
	}
	return -1
}
