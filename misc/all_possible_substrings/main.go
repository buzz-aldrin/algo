package main

import "fmt"

var (
	input = "abcd"
)

func main() {
	printAllPossibleSubstrings(input)
}

func printAllPossibleSubstrings(input string) {
	for i := 1; i <= len(input); i++ {
		for j := 0; j < len(input)-i+1; j++ {
			fmt.Println(input[j : i+j])
		}
	}
}
