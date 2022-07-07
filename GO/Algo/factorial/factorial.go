package main

import (
	"fmt"
	"time"
)

//recursive factorial
func factorialRec(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorialRec(n-1)
	}
}

//recursive with memoisation
func factorialMemo(n int) int {
	var memo map[int]int = make(map[int]int)
	return factMemoHelper(n, memo)
}

func factMemoHelper(n int, memo map[int]int) int {
	//basecase
	if n == 1 { //--> O(1)
		return 1
	}
	//check if we already calculated res for n*(n-1) and if so, get the value  --> O(1)
	if val, ok := memo[n]; ok {
		return n * val
	}
	//calculate
	res := n * factMemoHelper(n-1, memo) //--> O(n)
	memo[n] = res
	return res
}

func main() {
	speedTest(factorialRec)
	speedTest(factorialMemo)
}

func speedTest(f func(int) int) {
	fmt.Println("--------------------")
	testcase := []int{18}
	for _, v := range testcase {
		t0 := time.Now()
		res := f(v)
		elapsed := time.Since(t0) * time.Microsecond
		fmt.Println("inp: ", v, ": res:  ", res, "\t. ExecTime: \t", elapsed)
	}
	fmt.Println("--------------------")
}
