package main

import "fmt"

func URLfy(str string, l int) string {
	out := ""
	replace := []rune("%20")

	for k := range str {
		k = len(str) - 1 - k
	}

	return out
}

func main() {

	str := "Mr John Smith "
	r := "%20"
	str
	fmt.Println(len(str))
	// fmt.Println(URLfy(str, 14))

}
