package main

import "fmt"

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

// Example 1:

// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

// Example 2:

// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.

// Constraints:
//     1 <= prices.length <= 105
//     0 <= prices[i] <= 104

//[7,1,5,3,6,4]

//start left, find min
//find max right from min

//if doesnt exist find new min

// how would  I know this is max profit?

// Kane's algorithm

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	min := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i] // if smaller - we adjust min
		} else { //if larger - we check for profit
			if profit < prices[i]-min {
				profit = prices[i] - min
			}
		}

	}
	return profit
}

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	glo_max_profit, local_max_profit := 0, 0
	for i := 0; i < len(prices)-1; i++ {
		current_profit := prices[i+1] - prices[i]
		// fmt.Println("comparing:", prices[i], prices[i+1])
		// local_max_profit is the profit so far before the current day
		local_max_profit = Max(0, local_max_profit+current_profit)
		glo_max_profit = Max(local_max_profit, glo_max_profit)

		fmt.Println("current: ", current_profit, "\tlocal  :", local_max_profit, "\tglobal ", glo_max_profit)
	}
	return glo_max_profit
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

	fmt.Println(maxProfit2([]int{100, 2, 3, 6}))

}
