package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(insertionSort([]int{5, 2, 4, 6, 1, 3})) // Average: O(k*n)
	fmt.Println(insertionSort([]int{1, 2, 3, 4, 5, 6})) // Best: O(n)
	fmt.Println(insertionSort([]int{6, 5, 4, 3, 2, 1})) // Worst: O(n*n)

	A := []int{5, 2, 4, 6, 1, 3}
	// ASCENDING ORDER
	fmt.Println(insertionSortFunc(A, func(elem int, key int) bool {
		return elem > key
	}))
	// DESCENDING ORDER
	fmt.Println(insertionSortFunc(A, func(elem int, key int) bool {
		return elem < key
	}))

	integers := []int{4,2,5,1,6,3}
	floats := []float64{3.14, 1.24, 0.35}
	strings := []string{"H", "E", "L", "L", "O"}
	sort.Ints(integers)
	sort.Float64s(floats)
	sort.Strings(strings)
	fmt.Println(integers)
	fmt.Println(floats)
	fmt.Println(strings)
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

// insertionSortFunc sorts a given array/slice using the INSERTION SORT algorithm
// 1. We iterate from left to right
// 2. The left side is always sorted
// 3. We take each element an insert it into its correct position
// 4. We use a sorting function as the second condition of the inner loop
// for better abstraction
func insertionSortFunc(A []int, sortFunc func(elem, key int) bool) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i > -1 && sortFunc(A[i], key) {
			A[i+1] = A[i]
			i -= 1
		}
		A[i+1] = key
	}
	return A
}
