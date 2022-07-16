package main

import "fmt"

// Given the head of a singly linked list, return the middle node of the linked list.
// If there are two middle nodes, return the second middle node.

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(l *ListNode) *ListNode {
	if l.Next == nil {
		return l
	}
	len := 0 //len
	p := l   //pointer
	for p != nil {
		p = p.Next
		len++
	}
	mid := 0
	// if len%2 == 0 {
	// 	mid = int(len/2) + 1
	// } else {
	// 	mid = int(len / 2)
	// }
	mid = int(len/2) + 1

	fmt.Println("len:", len, "mid: ", mid)
	// fmt.Println(mid)
	i := 1
	for i < mid {
		// fmt.Println(mid, l.Val)
		l = l.Next
		i++
	}
	// fmt.Println(mid)
	return l
}

func main() {
	fmt.Println(middleNode(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}).Val)
	// fmt.Println(middleNode(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}).Val)
}

// [][x][]

// [][]x[x][]
