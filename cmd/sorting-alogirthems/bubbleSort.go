// Package main demonstrates three implementations of bubble sort algorithm
// with increasing optimization levels
package main

import (
	"fmt"
)

// bubbleSort implements the basic bubble sort algorithm.
// This version uses a temporary variable for swapping elements.
//
// Time Complexity:
// - Best Case: O(n²) - Always performs all comparisons
// - Average Case: O(n²)
// - Worst Case: O(n²)
// Space Complexity: O(1) - Sorts in-place
func bubbleSort(arr []int) {
	n := len(arr)

	// Outer loop controls the number of passes through the array
	// We need at most n passes to sort n elements
	for i := 0; i < n; i++ {
		
		// Inner loop compares adjacent elements
		// After i passes, the i largest elements are already in place,
		// so we only need to go up to n-i-1
		for j := 1; j < n-i; j++ {
			
			// Compare adjacent elements (j-1 and j)
			if arr[j-1] > arr[j] {
				// Swap elements using a temporary variable
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
		}
	}

	// Print the sorted array
	fmt.Println(arr)
}

// bubbleSortOptimized implements bubble sort with Go's built-in swap syntax.
// This eliminates the need for a temporary variable when swapping elements.
//
// Time Complexity:
// - Best Case: O(n²) - Still performs all comparisons
// - Average Case: O(n²)
// - Worst Case: O(n²)
// Space Complexity: O(1) - Sorts in-place
func bubbleSortOptimized(arr []int) {
	n := len(arr)
	
	// Outer loop controls the number of passes
	for i := 0; i < n; i++ {
		// Inner loop handles comparisons and swaps
		// We only need to check up to n-i-1 since the last i elements are already sorted
		for j := 1; j < n-i; j++ {
			// Compare adjacent elements
			if arr[j-1] > arr[j] {
				// Go's built-in swap syntax makes swapping more elegant and efficient
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	
	// Print the sorted array
	fmt.Println(arr)
}

// bubbleSortBest implements the most optimized version of bubble sort.
// It includes early termination when the array becomes sorted before completing all passes.
//
// Time Complexity:
// - Best Case: O(n) - When array is already sorted, only one pass is needed
// - Average Case: O(n²)
// - Worst Case: O(n²) - When array is in reverse order
// Space Complexity: O(1) - Sorts in-place
func bubbleSortBest(arr []int) {
	n := len(arr)
	
	// Outer loop for passes through the array
	for i := 0; i < n; i++ {
		// Flag to track if any swaps occur in this pass
		// If no swaps occur, the array is already sorted
		swapped := false
		
		// Inner loop for comparisons and swaps
		// We only need to check up to n-i-1 since the last i elements are already sorted
		for j := 1; j < n-i; j++ {
			// Compare adjacent elements
			if arr[j-1] > arr[j] {
				// Swap using Go's built-in swap syntax
				arr[j-1], arr[j] = arr[j], arr[j-1]
				// Record that we made a swap
				swapped = true
			}
		}
		
		// Early termination optimization:
		// If no elements were swapped in this pass, the array is already sorted
		// This improves best-case time complexity to O(n)
		if !swapped {
			break
		}
	}
	
	// Print the sorted array
	fmt.Println(arr)
}