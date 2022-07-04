package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	bst := MyBST{}
	bst.Insert(10)
	bst.Insert(9)
	bst.Insert(5)
	bst.Insert(20)
	// bst.String()
	fmt.Println(bst.Lookup(10))
}
