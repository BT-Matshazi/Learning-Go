// Package main implements sorting algorithms including insertion sort
package main

import (
	"fmt"
)

// InsertionSort implements the insertion sort algorithm.
// It works by building a sorted section of the array one element at a time.
// For each element, it finds the appropriate position within the sorted section
// and inserts it there, shifting other elements as necessary.
//
// Time Complexity:
// - Best Case: O(n) - When array is already sorted (minimal comparisons needed)
// - Average Case: O(n²)
// - Worst Case: O(n²) - When array is in reverse order (maximum shifts required)
// Space Complexity: O(1) - Sorts in-place with only a single temporary variable
func InsertionSort(arr []int) {
	n := len(arr)
	
	// Start from the second element (index 1)
	// The first element (index 0) is	 considered already sorted
	for i := 1; i < n; i++ {
		// Store the current element that needs to be inserted
		// into the correct position within the sorted portion
		current := arr[i]
		
		// Start comparing with the elements in the sorted section (to the left)
		j := i - 1
		
		// Move elements of arr[0..i-1] that are greater than current
		// to one position ahead of their current position
		// This creates space to insert the current element
		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j] // Shift element forward
			j--
		}
		
		// After finding the correct position (j+1),
		// insert the current element there
		arr[j+1] = current
	}
	
	// Print the sorted array
	fmt.Println(arr)
}
