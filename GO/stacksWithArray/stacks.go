package main

import "fmt"

type Stack struct {
	data   []interface{}
	length int
}

func push(v interface{}, s *Stack) {
	s.data = append(s.data, v)
	s.length++
}

func NewStack() *Stack {
	s := Stack{}
	return &s
}

func peek(s *Stack) {
	if s.length > 0 {
		fmt.Println(s.length, " ", s.data[len(s.data)-1])
	} else {
		fmt.Println("EMPTY")
	}

}

func pop(s *Stack) interface{} {
	if s.length == 0 {
		return nil
	}
	if s.length == 1 {
		el := s.data[0]
		s.data = append([]interface{}{})
		s.length--
		return el
	} else {
		el := s.data[len(s.data)-1]
		s.data = append(s.data[:len(s.data)-1])
		s.length--
		return el
	}

}

func main() {
	s := NewStack()
	push("google", s)
	push("facebook", s)
	push("NETFLIX", s)
	peek(s)
	pop(s)
	peek(s)
	pop(s)
	peek(s)
	pop(s)
	peek(s)
}
