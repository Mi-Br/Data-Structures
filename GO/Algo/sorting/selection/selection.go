package main

import "fmt"

func selectionSort(ar []int) []int {
	for i := 0; i < len(ar)-1; i++ {
		minIndx := i
		for j := i + 1; j < len(ar); j++ {
			if ar[minIndx] > ar[j] {
				minIndx = j
			}
		}
		if minIndx != i { // means we found new min at the end of the loop
			ar[i], ar[minIndx] = ar[minIndx], ar[i]
		}
	}
	return ar
}

func main() {
	testcase := []int{1, 2, 3, 4, 1, 3, 0}
	fmt.Println(testcase)
	testcase = selectionSort(testcase)
	fmt.Println(testcase)
}
