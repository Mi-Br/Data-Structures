package main

import (
	"fmt"
	"math/rand"
	"time"
)

var min int = 0
var max int = 100

func generateRandomSortedArray(num int) []int {
	var arr []int

	for num > 0 {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(max) - min
		arr = append(arr, r)
		num--
	}
	return arr
}

func main() {
	fmt.Println(generateRandomSortedArray(100))
}
