package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	h := head //will become tail
	//while current pointer.next != nil
	nodeCurrent := h
	nodeNext := nodeCurrent.Next
	for nodeNext != nil {
		tmp := nodeNext.Next
		nodeNext.Next = nodeCurrent
		nodeCurrent = nodeNext
		nodeNext = tmp
	}
	h.Next = nil
	return nodeCurrent
	// swap head and tail
	//  [ B ] --> [ A ] --- >
}

const min = -1000
const max = +1000

func generateLinkedList(numberOfNodes int) (head *ListNode) {
	head = &ListNode{}
	pointer := head
	for numberOfNodes > 0 {
		rand.Seed(time.Now().UnixNano())
		rval := rand.Intn(max) + min
		newNode := ListNode{Val: rval, Next: nil}
		pointer.Next = &newNode
		pointer = pointer.Next
		numberOfNodes--
	}
	return head.Next
}
func traverseList(l *ListNode) {
	p := l
	fmt.Print("O->")
	for p != nil {
		fmt.Print(p.Val, "-->")
		p = p.Next
	}
}

func main() {
	randList := generateLinkedList(3)
	traverseList(randList)
	randList = reverseList(randList)
	fmt.Println(randList)
	traverseList(randList)
}
