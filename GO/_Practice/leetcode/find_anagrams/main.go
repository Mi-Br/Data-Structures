package main

import (
	"fmt"
)

func main() {

	s := "abab"
	p := "ab"

	fmt.Println(findAnagrams(s, p))

}

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}
	var out []int = []int{}

	//sliding window ,len of p
	p0 := 0
	p1 := len(p) - 1

	// initate maps for window range (sMap )

	sMap := MakeHash(s[:p1])
	pMap := MakeHash(p)

	//while window edge in range of s
	for p1 <= len(s)-1 {
		addKey(rune(s[p1]), sMap)
		if isTheSameHash(sMap, pMap) {
			out = append(out, p0)
		}
		//remove p0 from sMap
		removeKey(rune(s[p0]), sMap)
		p0++
		p1++
	}
	return out
}

func MakeHash(s string) map[rune]int {
	var hash map[rune]int = map[rune]int{}
	for _, r := range s {

		if v, ok := hash[r]; ok {
			hash[r] = v + 1
		} else {
			hash[r] = 1
		}
	}
	return hash
}

func isTheSameHash(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}

	return true
}

func addKey(r rune, m map[rune]int) {
	if _, ok := m[r]; ok {
		m[r] += 1
	} else {
		m[r] = 1
	}
}

func removeKey(r rune, m map[rune]int) {
	if v, ok := m[r]; ok {
		if v > 1 {
			m[r] -= 1
		} else {
			delete(m, r)
		}
	}
}
