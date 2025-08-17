	package main

	import (
		// "errors"
		"fmt"
	)

	/*
		Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

		You may assume that each input would have exactly one solution, and you may not use the same element twice.

		You can return the answer in any order.

		Example 1:

		Input: nums = [2,7,11,15], target = 9
		Output: [0,1]
		Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
		Example 2:

		Input: nums = [3,2,4], target = 6
		Output: [1,2]
		Example 3:

		Input: nums = [3,3], target = 6
		Output: [0,1]
	*/ 
	func main() {
		nums := []int{2,11,7,15}
		var target int = 9

		fmt.Println("SUM TWO: ", twoSum(nums, target))
	}

	func oldTwoSum(nums []int, target int) []int {
		for i := 0; i < len(nums); i++ {  //for i := range nums {} for a modern approach
			for j := 0; j < len(nums); j++ {	//for j := range nums {} for a modern approach
				if(i != j && nums[i] + nums[j] == target){
					return []int{i,j}
				}
			}
		}

		return []int{0,0}
	}

	
func twoSum(nums []int, target int) []int {
    // Create a map to store number -> index mapping
    numMap := make(map[int]int)
    
    for i, num := range nums {
        // Calculate what number we need to find
        complement := target - num
        
        // Check if the complement exists in our map
        if index, found := numMap[complement]; found {
            // Found the pair! Return indices
            return []int{index, i}
        }
        
        // Store current number and its index for future lookups
        numMap[num] = i
    }
    
    // No solution found (shouldn't happen with valid input)
    return []int{}
}