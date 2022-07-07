package main

import "fmt"

func isSubsequence(s string, t string) bool {

	if (len(s)) == 0 {
		return true
	}

	n := len(t)
	j := 0
	for i := 0; i < n; i++ {
		if s[j] == t[i] {
			j++
		}
		if j == len(s) {
			return true
		}
	}
	return false
}

func main() {

	fmt.Println(isSubstring("b", "c"))
	fmt.Println(isSubstring("axc", "ahbgdc"))
	fmt.Println(isSubstring("abc", "ahbgdc"))
}
