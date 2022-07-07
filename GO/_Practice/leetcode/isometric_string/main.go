package main

import "fmt"

func isIsomorphic(s, t string) bool {
	return isIsomorphicOneWay(s, t) && isIsomorphicOneWay(t, s)
}

func isIsomorphicOneWay(s, t string) bool {
	sDict := genDictionary(s, t)
	for i, c := range s {
		if sDict[string(c)] != string(t[i]) {
			return false
		}
	}
	return true
}

func genDictionary(s, t string) map[string]string {
	dict := make(map[string]string)
	for i, c := range s {
		dict[string(c)] = string(t[i])
	}
	return dict
}

func main() {
	fmt.Println("test1 passed: ", isIsomorphic("egg", "add") == true)
	fmt.Println("test1 passed: ", isIsomorphic("foo", "bar") == false)
	fmt.Println("test1 passed: ", isIsomorphic("paper", "title") == true)
	fmt.Println("test1 passed: ", isIsomorphic("badc", "baba") == false)
}
