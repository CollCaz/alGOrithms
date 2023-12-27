package main

import (
	Sorts "algos/BubbleSort"
	"fmt"
)

func main() {
	unsorted := []int{9, 7, 5, 3, 1, 8, 6, 0, 4, 2}
	bubble_sorted := Sorts.BubbleSort(unsorted)
	fmt.Println("bubble sort:", bubble_sorted)
}
