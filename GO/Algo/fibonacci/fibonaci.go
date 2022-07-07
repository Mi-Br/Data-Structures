package main

import (
	"Data-structures/go/speed_test"
)

//recursive solution
func fibonaciRec(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return fibonaciRec(n-1) + fibonaciRec(n-2)
}

//					~ O(2^n) ~
// 		       [n]

// 	   [n-1]    	[n-2]

// [n-2]   [n-3]  [n-3]  [n-4]
//

//linear solution
func fibonaci(n int) int {
	sums := []int{}
	sums = append(sums, 0, 1)
	for i := 2; i <= n; i++ {
		s := sums[i-1] + sums[i-2]
		sums = append(sums, s)
	}
	return sums[n]
}

func main() {
	testcase := []int{40}
	speed_test.SpeedTest(fibonaciRec, testcase)
	speed_test.SpeedTest(fibonaci, testcase)

}
