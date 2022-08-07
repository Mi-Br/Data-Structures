package main

import "fmt"

type set struct {
	key   string
	value interface{}
}

type Hash struct {
	data [][]set
	cap  int
}

func MakeHash(size int) Hash {
	var out Hash = Hash{}
	out.data = make([][]set, size)
	out.cap = size
	return out
}

func (h *Hash) _hashValue(v string) int {
	hash := 0
	for i, v := range v {
		hash = (hash + int(v)*i) % h.cap
	}
	// fmt.Println(hash)
	return hash
}

func (h *Hash) Set(k string, val interface{}) {
	index := h._hashValue(k)
	newset := set{key: k, value: val}
	if h.data[index] == nil {
		h.data[index] = []set{}
		h.data[index] = append(h.data[index], newset)
	} else {
		for i, el := range h.data[index] {
			if el.value == k {
				h.data[index][i].value = val
				return
			}
		}
		h.data[index] = append(h.data[index], newset)
	}
}

func (h *Hash) Get(k string) (val interface{}, ok bool) {
	ok = false
	index := h._hashValue(k)

	if h.data[index] == nil {
		return "", false
	} else {
		for i := 0; i < len(h.data[index]); i++ {
			if h.data[index][i].key == k {
				val = h.data[index][i].value
				ok = true
				return val, ok
			}
		}
	}
	return "", ok
}

func (h *Hash) Display() {
	fmt.Println()
	for r := 0; r < len(h.data); r++ {
		if h.data[r] != nil {
			fmt.Printf("%v:", r)
			for c := 0; c < len(h.data[r]); c++ {
				fmt.Printf("|%v -> {%v:%v}", c, h.data[r][c].key, h.data[r][c].value)
			}
			fmt.Println()
		}
	}
}

func (h *Hash) Remove(k string) {
	if _, ok := h.Get(k); ok {
		index := h._hashValue(k)
		for r := 0; r < len(h.data[index]); r++ {
			if h.data[index][r].key == k {
				len := len(h.data[index])
				//copy last element to the index r
				h.data[index][r] = h.data[index][len-1]
				h.data[index][len-1] = set{} // empty set
				//truncate slice
				h.data[index] = h.data[index][:len-1]
			}
		}
	}

}

func (h *Hash) Keys() []string {
	out := []string{}

	for r := 0; r < len(h.data); r++ {
		for j := 0; j < len(h.data[r]); j++ {
			out = append(out, h.data[r][j].key)
		}
	}
	return out
}

func main() {
	H := MakeHash(50)
	H.Set("John", 34)
	H.Set("Michail", 37)
	H.Set("Michail2", true)
	H.Set("Michail3", "medium rich")
	H.Set("Michail4", "rich")

	v, ok := H.Get("Michail")
	fmt.Println(v, ok)
	// H.Display()

	// H.Remove("Michail")
	// v, ok = H.Get("Michail")
	// // fmt.Println(v, ok)
	// H.Display()
	fmt.Println(H.Keys())
}
