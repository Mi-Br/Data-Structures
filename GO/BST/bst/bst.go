package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type MyBST struct {
	root *Node
}

func (t *MyBST) Insert(val int) {
	newNode := Node{value: val, left: nil, right: nil}
	if t.root == nil {
		t.root = &newNode
	} else {
		t.root.Insert(&newNode)
	}
}

func (cn *Node) Insert(nn *Node) {
	if nn.value < cn.value {
		//move left
		if cn.left == nil {
			cn.left = nn
		} else {
			cn.left.Insert(nn)
		}
	}
	if nn.value > cn.value {
		//move left
		if cn.right == nil {
			cn.right = nn
		} else {
			cn.right.Insert(nn)
		}
	}
}

func (b *MyBST) bfs_print() {
	var stack []*Node
	stack = append(stack, b.root)
	for len(stack) > 0 {
		fmt.Println(stack[0].value)
		if stack[0].left != nil {
			stack = append(stack, stack[0].left)
		}
		if stack[0].right != nil {
			stack = append(stack, stack[0].right)
		}
		stack = stack[1:]
	}
}

func isValidBST(tree *Node) bool {
	return validateBst(tree, math.MinInt64, math.MaxInt64)
}

func validateBst(tree *Node, minValue, maxValue int) bool {
	if tree == nil {
		return true
	}
	if tree.value <= minValue || tree.value >= maxValue {
		return false
	}
	if tree.left != nil && !validateBst(tree.left, minValue, tree.value) {
		return false
	}
	if tree.right != nil && !validateBst(tree.right, tree.value, maxValue) {
		return false
	}

	return true
}

func dfs_print(n *Node) {
	fmt.Println(n.value)
	if n.left != nil {
		dfs_print(n.left)
	}
	if n.right != nil {
		dfs_print(n.right)
	}
	if n.left == nil && n.right == nil {
		return
	}
}

func dfs_in_order_recurs(n *Node) {
	if n.left == nil && n.right == nil {
		fmt.Println(n.value)
		return
	} else {
		if n.left != nil {
			dfs_in_order_recurs(n.left)
		}
		fmt.Println(n.value)
		if n.right != nil {
			dfs_in_order_recurs(n.right)
		}
		return
	}

}

func dfs_pre_order_recurs(n *Node) {
	if n != nil {
		fmt.Println(n.value)
		if n.left != nil {
			dfs_pre_order_recurs(n.left)
		}
		if n.right != nil {
			dfs_pre_order_recurs(n.right)
		}
	}
}
func dfs_post_order_recurs(n *Node) {
	if n.left == nil && n.right == nil {
		fmt.Println(n.value)
	} else {
		if n.left != nil {
			dfs_post_order_recurs(n.left)
		}
		if n.right != nil {
			dfs_post_order_recurs(n.right)
		}
		fmt.Println(n.value)
	}

}
func (n *Node) Lookup(val int) bool {
	if n.value == val {
		return true
	}
	if val < n.value && n.left != nil {
		return n.left.Lookup(val)
	}
	if val > n.value && n.right != nil {
		return n.right.Lookup(val)
	}
	return false
}

func (t *MyBST) Lookup(val int) bool {
	if t.root == nil {
		return false
	}
	return t.root.Lookup(val)
}
func (t *MyBST) PrintBST() {
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(b))
	}
}

func (bst *MyBST) String() {
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.right, level)
		fmt.Printf(format+"%d\n", n.value)
		stringify(n.left, level)
	}
}

func generateRandomList(n int) *MyBST {
	const max int = 100
	bst := MyBST{}

	for n > 0 {
		rand.Seed(time.Now().UnixNano())
		val := rand.Intn(max)
		bst.Insert(val)
		n--
	}
	return &bst
}

func main() {
	bst := generateRandomList(9)
	bst.String()
	// bst.bfs_print()
	// dfs_print(bst.root)
	// dfs_in_order_recurs(bst.root)
	// dfs_pre_order_recurs(bst.root)
	dfs_post_order_recurs(bst.root)
	fmt.Println(isValidBST(bst.root))
}
