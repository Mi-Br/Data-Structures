//https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/submissions/
// Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val == root.Val || q.Val == root.Val { // one of nodes is root
		return root
	} else if (p.Val < root.Val && q.Val > root.Val) || (q.Val < root.Val && p.Val > root.Val) { //one of nodes on left side , another one of the right side means ancestor is root
		return root
	} else {
		// left or right
		if q.Val < root.Val && p.Val < root.Val {
			return lowestCommonAncestor(root.Left, p, q)
		} else {
			return lowestCommonAncestor(root.Right, p, q)
		}
	}

}

func main() {

}
