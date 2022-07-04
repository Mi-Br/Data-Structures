package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top    *Node
	bottom *Node
	length int
}

func NewNode(val int) *Node {
	return &Node{value: val, next: nil}
}

func NewStack() *Stack {
	return &Stack{top: nil, bottom: nil, length: 0}
}

//lets to see the top node
func peek(s *Stack) {
	println(s.top.value, &s.top)
}

//adds new node to the tack
func push(v int, s *Stack) {
	n := NewNode(v)
	if s.top == nil {
		s.top = n
		s.bottom = n
	} else {
		n.next = s.top
		s.top = n
	}
	s.length++
}

//removes top el
func pop(s *Stack) *Node {
	if !isEmpty(s) {
		n := s.top
		s.top = s.top.next
		s.length--
		return n
	} else {
		return nil
	}
}

func PrintStack(s *Stack) {
	if !isEmpty(s) {
		n := s.top
		fmt.Printf("\nTOP=>")
		for n != nil {
			fmt.Print("{", n.value, "\t", &n, "}")
			n = n.next
		}
	}
}

//checks if stack is empty
func isEmpty(s *Stack) bool {
	if s.length == 0 {
		return true
	} else {
		return false
	}
}
func main() {
	n := NewStack()
	push(10, n)
	push(8, n)
	push(6, n)
	push(4, n)
	PrintStack(n)
	pop(n)
	PrintStack(n)
	peek(n)
}
