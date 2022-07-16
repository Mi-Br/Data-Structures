package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//basic validation if lists are not empty
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil && list2 != nil {
		return list2
	}
	if list2 == nil && list1 != nil {
		return list1
	}

	var a, b *ListNode //sliding pointers for both lists

	//make sure that my  list "a" holds smallest starting number
	if list1.Val < list2.Val {
		a, b = list1, list2
	} else {
		b, a = list1, list2
	}

	A := a // store start of a "list" for return later
	for a.Next != nil && b != nil {
		if a.Next.Val > b.Val {
			//insert b before a.Next
			freeNode := b
			b = b.Next
			tmp := a.Next
			a.Next = freeNode
			freeNode.Next = tmp
			a = a.Next //move pointer to the next node
		} else {
			a = a.Next
		}
	}
	if b != nil { //if by the end of list a we stil have remaining b list items - add them to the end
		a.Next = b
	}
	return A //original pointer to list "a"
}

func main() {
	//Test case 1 runs out of memory on leetcode ^^
	//[1,2,4]
	//[1,3,4]
	l1 := &ListNode{Val: -10, Next: &ListNode{Val: -10, Next: &ListNode{Val: -9, Next: &ListNode{Val: -4, Next: nil}}}}
	l2 := &ListNode{Val: -7, Next: nil}

	merged := mergeTwoLists(l1, l2)

	for merged != nil {
		fmt.Print(merged.Val, "---->")
		merged = merged.Next
	}

}
