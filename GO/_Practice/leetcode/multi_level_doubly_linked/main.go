package main

// excersize https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/
// You are given a doubly linked list, which contains nodes that have a next pointer, a previous pointer, and an additional child pointer. This child pointer may or may not point to a separate doubly linked list, also containing these special nodes. These child lists may have one or more children of their own, and so on, to produce a multilevel data structure as shown in the example below.

// Given the head of the first level of the list, flatten the list so that all the nodes appear in a single-level, doubly linked list. Let curr be a node with a child list. The nodes in the child list should appear after curr and before curr.next in the flattened list.

// Return the head of the flattened list. The nodes in the list must have all of their child pointers set to null.
// * Definition for a Node.
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

//implementation using pointers
func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	c := root
	for c != nil {
		if c.Child != nil {
			tmp := c.Next
			f := flatten(c.Child)
			c.Child = nil
			c.Next = f
			f.Prev = c
			for c.Next != nil {
				c = c.Next
			}
			c.Next = tmp
			if tmp != nil {
				tmp.Prev = c
			}
			c = c.Next
		}
		if c != nil {
			c = c.Next
		}
	}
	return root
}

//implementation using queue to store nodes
func flatten(root *Node) *Node {
	q := traverse_recurs(root)
	if len(q) < 1 {
		return nil
	}
	for i := 0; i < len(q)-1; i++ {
		cur, next := q[i], q[i+1]
		cur.Next = next
		next.Prev = cur
		cur.Child = nil
	}
	return q[0]
}

func traverse_recurs(n *Node) []*Node {
	q := []*Node{}
	c := n
	for c != nil {
		q = append(q, c)
		if c.Child != nil {
			q = append(q, traverse_recurs(c.Child)...)
		}
		c = c.Next
	}
	return q
}

func main() {

}
