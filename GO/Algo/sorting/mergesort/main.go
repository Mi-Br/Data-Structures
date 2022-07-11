package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomArray(n int) []int {
	// const min = -1000
	const max = 1000
	out := []int{}
	for n > 0 {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(max)
		out = append(out, r)
		n--
	}
	return out
}

func mergesort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := int(len(arr) / 2)
	left := mergesort(arr[:mid])
	right := mergesort(arr[mid:])
	out := []int{}

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			out = append(out, left[0])
			left = left[1:]
		} else {
			out = append(out, right[0])
			right = right[1:]
		}

	}
	if len(left) > 0 {
		out = append(out, left...)
	}

	if len(right) > 0 {
		out = append(out, right...)
	}
	return out
}

func main() {
	testcase := generateRandomArray(50)
	fmt.Println(testcase)
	testcase = mergesort(testcase)
	fmt.Println(testcase)

}
