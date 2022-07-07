package main

import "fmt"

func bubbleSort(ar []int) []int {
	swapedElements := true //if we havent swapped elements - means array is in order
	for swapedElements {   //iterate for as long as we find elements to swap
		swapedElements = false //reset
		for i := 0; i < len(ar)-1; i++ {
			for j := i + 1; j < len(ar); j++ {
				if ar[i] > ar[j] {
					ar[i], ar[j] = ar[j], ar[i]
					swapedElements = true
				}
			}
		}
	}
	return ar
}

func bubbleSort2(ar []int) []int {
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar)-2; j++ {
			if ar[j+1] < ar[j] {
				ar[j+1], ar[j] = ar[j], ar[j+1]
			}
		}
	}
	return ar
}

func main() {
	testCase := []int{9, 8, 6, 4}
	testCase = bubbleSort2(testCase)
	fmt.Println(testCase)
}
