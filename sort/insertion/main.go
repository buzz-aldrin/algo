package main

import "fmt"

// arr[] = 64 25 12 22 11

func main() {
	input := []int{64, 25, 12, 22, 11}

	insertion(input)

	fmt.Println(input)
}

func insertion(input []int) {
	for i := 1; i < len(input); i++ {
		temp := input[i]
		j := i - 1
		for ; j >= 0 && input[j] > temp; j-- {
			input[j+1] = input[j]
		}
		input[j+1] = temp
	}
}
