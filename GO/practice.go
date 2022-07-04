package main

import (
	"fmt"
	"strings"
)

func main() {

}

func compareArrays(arr1, arr2 []string) bool { //O(n^2)
	for _, v := range arr1 { // O(n)
		for _, y := range arr2 { // O(n)
			if v == y {
				return true
			}
		}
	}
	return false
}

func compareArrays2(arr1, arr2 []string) bool { //O(a+b) => O(a + b)
	var arrMap = map[string]bool{}
	for _, v := range arr1 { //O(a)
		arrMap[v] = true
	}

	for _, v := range arr2 { //O(b)
		if arrMap[v] {
			return true
		}
	}
	return false
}

func twoSum(nums []int, target int) []int {
	var outp []int

	var dict = map[int]int{}
	for i, v := range nums {
		dict[target-v] = i
		fmt.Println(target-v, " ", i)
	}
	for i, v := range nums {
		fmt.Println(i, " ", v)
		if _, ok := dict[v]; ok {
			fmt.Println(":: ", dict[v], " of:::", v)
			outp = append(outp, i, dict[v])
			// fmt.Println("adding ", i, dict[v])
			break
		}
	}
	return outp
}

func reverseString(str string) []string {
	inp := strings.Split(str, "")
	//brute force solutin
	var outp []string
	j := 0
	for i := len(inp) - 1; i >= 0; i-- { //O(n) , space O(2*n) = O(n)
		outp = append(outp, inp[i])
		j++
	}

	fmt.Println(inp)
	fmt.Println(outp)
	// for _, v := range outp {
	// 	fmt.Println("%s", v)
	// }

	return outp
}

func reverseString2(str string) string {

	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
