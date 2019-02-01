package main

import "fmt"

func main(){
	input := []int{64, 25, 12, 22, 11}

	bubble(input)

	fmt.Println(input)
}

func bubble(input []int){
	for i := 0; i < len(input) - 1 ; i++{
		for j := i + 1; j < len(input); j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
}