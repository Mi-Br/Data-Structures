// Given an array containing 0s and 1s, if you are allowed to replace no more than ‘k’ 0s with 1s,
// find the length of the longest contiguous subarray having all 1s.
// Input: Array=[0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1], k=2
// Output: 6
// Explanation: Replace the '0' at index 5 and 8 to have the longest contiguous subarray of 1s having length 6.

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestSubbaray(arr []int, k int) int {
	maxSeq := 0
	currentSeq := 0
	end := 0

	for start := 0; start < len(arr); start++ {
		if arr[start] == 1 {
			currentSeq++
		}

		// I need to shrink the window (move end) when I am out of K

		if (start - end + 1 - currentSeq) > k {
			if arr[end] == 1 {
				currentSeq--
			}
			end++
		}
		fmt.Println(end, start, currentSeq)
		maxSeq = max(maxSeq, start-end+1)
	}
	return maxSeq
}

func main() {

	tc1 := []int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}
	fmt.Println(longestSubbaray(tc1, 3))
}
