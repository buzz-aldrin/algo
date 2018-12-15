package main

import "fmt"

// arr[] = 64 25 12 22 11

func main() {
	input := []int{64, 25, 12, 22, 11}

	selection(input)

	fmt.Println(input)
}

func selection(input []int) {
	for i := 0; i < len(input)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[minIndex] {
				minIndex = j
			}
		}
		input[i], input[minIndex] = input[minIndex], input[i]
	}
}
