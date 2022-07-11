package main

import (
	"testing"
)

func generateTestCase(n []int) *ListNode {
	var out *ListNode
	i := len(n) - 1
	for i > 0 {
		val := n[i]
		if out != nil {
			out.Next = &ListNode{Val: val, Next: nil}
		} else {
			out = &ListNode{Val: val}
		}
		i--
	}
	return out
}

func TestFindMid(t *testing.T) {
	// Input: head = [1,2,3,4,5]
	// Output: [3,4,5]
	// Explanation: The middle node of the list is node 3.
	testcase := generateTestCase([]int{1, 2, 3, 4})
	result := FindMidNode(testcase)
	if result.Val != 3 {
		t.Errorf("Finding mid in [1,2,3,4] FAILED. Expected %d, got %d", 3, result.Val)
	} else {
		t.Logf("Finding mid in [1,2,3,4] PASSED. Expected %d, got %d", 3, result.Val)
	}

	testcase = generateTestCase([]int{3, 4, 5})
	result = FindMidNode(testcase)

	if result.Val != 4 {
		t.Errorf("Finding mid in [3,4,5] FAILED. Expected %d, got %d", 4, result.Val)
	} else {
		t.Logf("Finding mid in [3,4,5] PASSED. Expected %d, got %d", 4, result.Val)
	}
}
