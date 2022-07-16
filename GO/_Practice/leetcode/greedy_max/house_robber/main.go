package main

import "fmt"

//[1,2,3,1]
//rob 1, then 3 = 4

//[2,7,9,3,1]
//  291  73  91

// max rob of array is

// if I have [2,1] - return 0 as its immediate alarm

// if I have [371] - return either 3,1 or 7 (max)
// if I have  [4,3,2,4] its same as max problem for

func robhouseRecurs(n []int) int {
	if len(n) == 0 {
		return 0
	}
	if len(n) == 1 {
		return n[0]
	}
	if len(n) == 2 {
		if n[0] > n[1] {
			return n[0]
		} else {
			return n[1]
		}
	}
	return max(n[0]+robhouseRecurs(n[2:]), robhouseRecurs(n[1:]))
}

func rob(n []int) int {
	if len(n) == 0 {
		return 0
	}
	if len(n) == 1 {
		return n[0]
	}

	dp := make([]int, len(n))
	dp[0] = n[0]
	dp[1] = n[1]
	for i := 2; i < len(n); i++ {
		dp[i] = max(n[i]+dp[i-2], dp[i-1])
	}
	return (dp[len(n)-1])
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func rob2(nums []int) int {
	prevMax := 0
	currMax := 0

	for i := 0; i < len(nums); i++ {
		temp := currMax
		if prevMax+nums[i] > currMax {
			currMax = prevMax + nums[i]
		}
		prevMax = temp
		fmt.Println(prevMax, currMax, nums[i])
	}
	return currMax
}

func main() {
	fmt.Println(rob([]int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}))
	// fmt.Println(robhouse([]int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}))
	fmt.Println(rob2([]int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}))
}
