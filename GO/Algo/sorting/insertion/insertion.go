package main

import "fmt"

//### Go through the sorted portion of the list and insert it in the right order so that the sorted portion of the list remains sorted ####

func insertSort(ar []int) []int {
	var n = len(ar)

	for i := 0; i < n; i++ {
		j := i
		for j > 0 {
			if ar[j-1] > ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
			j--
		}
	}
	return ar
}

func main() {
	testcase := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Println(testcase)
	fmt.Println(insertSort(testcase))

}
