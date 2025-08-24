// Package main demonstrates sorting algorithms: bubble sort and insertion sort
package main

import (
	"fmt"
)

func main() {
	// Create an unsorted integer array to be sorted
	// The [...] syntax lets Go determine the array size based on the elements provided
	var numArr = [...]int{42, 8, 17, 99, 4, 23, 56, 71, 13, 31}

	// Make copies of the original array for each sorting method
	// to demonstrate the different implementations
	arr1 := make([]int, len(numArr))
	arr2 := make([]int, len(numArr))
	arr3 := make([]int, len(numArr))
	arr4 := make([]int, len(numArr))
	
	// Copy the original array to our working arrays
	copy(arr1, numArr[:])
	copy(arr2, numArr[:])
	copy(arr3, numArr[:])
	copy(arr4, numArr[:])

	// Print the original unsorted array
	fmt.Println("Original array:", numArr)

	// Call each bubble sort implementation and print the results
	fmt.Print("Basic bubble sort: ")
	bubbleSort(arr1)

	fmt.Print("Optimized bubble sort: ")
	bubbleSortOptimized(arr2)

	fmt.Print("Best bubble sort: ")
	bubbleSortBest(arr3)
	
	// Call insertion sort and print the results
	fmt.Print("Insertion sort: ")
	InsertionSort(arr4)
	
	// Compare algorithms
	fmt.Println("\nAlgorithm comparison:")
	fmt.Println("- Bubble Sort: Simple to implement but not efficient for large datasets")
	fmt.Println("- Insertion Sort: More efficient than bubble sort for small or nearly sorted arrays")
	fmt.Println("- Both have O(n²) worst-case time complexity, but insertion sort generally performs better in practice")
}