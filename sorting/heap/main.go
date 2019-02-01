package main

import "fmt"

var (
	input = []int{12, 11, 13, 5, 6, 7}
)

func main() {
	heapSortAscending(input, len(input))
	fmt.Println(input)

	heapSortDescending(input, len(input))
	fmt.Println(input)
}

func minHeapify(input []int, size, rootIndex int) {
	smallest := rootIndex
	left := 2*rootIndex + 1
	right := 2*rootIndex + 2

	if left < size && input[left] < input[smallest] {
		smallest = left
	}

	if right < size && input[right] < input[smallest] {
		smallest = right
	}

	if smallest != rootIndex {
		input[smallest], input[rootIndex] = input[rootIndex], input[smallest]

		minHeapify(input, size, smallest)
	}
}

func heapSortAscending(input []int, size int) {
	for i := (size / 2) - 1; i >= 0; i-- {
		maxHeapify(input, size, i)
	}

	// One by one extract an element from heap
	for i := size - 1; i >= 0; i-- {
		// Move current root to end
		input[0], input[i] = input[i], input[0]

		// call max heapify on the reduced heap
		maxHeapify(input, i, 0)
	}
}

func maxHeapify(input []int, size, rootIndex int) {
	largest := rootIndex
	left := 2*rootIndex + 1
	right := 2*rootIndex + 2

	if left < size && input[left] > input[largest] {
		largest = left
	}

	if right < size && input[right] > input[largest] {
		largest = right
	}

	if largest != rootIndex {
		input[largest], input[rootIndex] = input[rootIndex], input[largest]

		minHeapify(input, size, largest)
	}
}

func heapSortDescending(input []int, size int) {
	for i := (size / 2) - 1; i >= 0; i-- {
		minHeapify(input, size, i)
	}

	// One by one extract an element from heap
	for i := len(input) - 1; i >= 0; i-- {
		// Move current root to end
		input[0], input[i] = input[i], input[0]

		// call max heapify on the reduced heap
		minHeapify(input, i, 0)
	}
}
