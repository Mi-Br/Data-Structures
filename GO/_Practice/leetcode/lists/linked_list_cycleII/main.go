// Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously
//following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is
// connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}
	// //snode - slow pointer
	// //fnode - fast node pointer
	var snode, fnode *ListNode = head, head
	for fnode != nil && fnode.Next != nil {
		snode = snode.Next
		fnode = fnode.Next.Next

		if fnode == snode {
			h := head
			for h != snode {
				h = h.Next
				snode = snode.Next
			}
			return snode
		}
	}
	return nil
}
func main() {

}
