// Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4

// Example 2:

// Input: nums = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1
package main

import "fmt"

func search_recurs(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	} else {
		m := int(len(nums) / 2)
		if nums[m] == target {
			return m
		} else if target < nums[m] {
			return search_recurs(nums[:m], target)
		} else if target > nums[m] {
			indx := search_recurs(nums[m:], target)
			if indx != -1 {
				return m + indx
			} else {
				return -1
			}
		}
	}
	return -1
}

func search_iterative(nums []int, target int) int {
	l := 0
	h := len(nums) - 1

	for l < h {
		m := l + int((h-l+1)/2) // should always use upper mid, otherwise it will get stuck with 2 elements. mid will keep reseting to l and l will always be < h
		fmt.Println(l, m, h)
		if nums[m] == target {
			return m
		}
		if target < nums[m] {
			h = m - 1
		}
		if target > nums[m] {
			l = m
		}
	}
	return -1
}

func main() {
	// fmt.Println(search_recurs([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println(search_iterative([]int{-1, 0, 3, 5, 9, 12}, 2))
}
