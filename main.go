package main

import (
	"fmt"
)

func main() {
	fmt.Println(insertionSort([]int{5, 2, 4, 6, 1, 3})) // Average: O(k*n)
	fmt.Println(insertionSort([]int{1, 2, 3, 4, 5, 6})) // Best: O(n)
	fmt.Println(insertionSort([]int{6, 5, 4, 3, 2, 1})) // Worst: O(n*n)
}

// insertionSort sorts a given array/slice using the INSERTION SORT algorithm
// 1. We iterate from left to right
// 2. The left side is always sorted
// 3. We take each element an insert it into its correct position
func insertionSort(A []int) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i > -1 && A[i] > key {
			A[i+1] = A[i]
			i -= 1
		}
		A[i+1] = key
	}
	return A
}
