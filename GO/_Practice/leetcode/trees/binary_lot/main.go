
// https://leetcode.com/problems/binary-tree-level-order-traversal/
// Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int

	var q = make([]*TreeNode, 1)
	q[0] = root
	level := 0

	for len(q) > 0 {

		cur_len := len(q)
		for i := 0; i < cur_len; i++ {
			node := q[0]
			q = q[1:]

			if len(res) <= level { //add first element
				fmt.Println("first el goes: ", node.Val)
				fmt.Println(len(res), level)
				res = append(res, []int{node.Val})
			} else {
				fmt.Println("remaining el goes: ", node.Val)
				fmt.Println(len(res), level)
				res[level] = append(res[level], node.Val)
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		level++

	}

	return res

}