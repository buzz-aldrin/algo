package main

import (
	"fmt"
)

var (
	input = "abc"
)

func main() {
	printAllPermutations([]rune(input), 0, len(input)-1)
}

func printAllPermutations(str []rune, b, e int) {
	if b == e {
		fmt.Println(string(str))
		return
	}

	for i := b; i <= e; i++ {
		str[b], str[i] = str[i], str[b]
		printAllPermutations(str, b+1, e)
		str[b], str[i] = str[i], str[b]
	}
}
