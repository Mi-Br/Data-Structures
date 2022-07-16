package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var valid bool = true
	if root.Left == nil && root.Right == nil {
		return valid
	}
	if root.Left != nil {
		valid = valid && isValidBST(root.Left) && root.Left.Val < root.Val
	}
	if root.Right != nil {
		valid = valid && isValidBST(root.Right) && root.Right.Val > root.Val
	}
	return valid
}

func main() {

}
