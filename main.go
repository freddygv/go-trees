package main

import (
	"fmt"
)

// RBNode is a node in a Red-Black Tree
type RBNode struct {
	value int
	red   bool
	left  *RBNode
	right *RBNode
}

/*

	  4
  2		 6
1   3  5   7

*/
func main() {
	root := RBNode{value: 4, red: false}

	root.left = &RBNode{value: 2, red: false}
	root.left.left = &RBNode{value: 1, red: false}
	root.left.right = &RBNode{value: 3, red: false}

	root.right = &RBNode{value: 6, red: false}
	root.right.left = &RBNode{value: 5, red: false}
	root.right.right = &RBNode{value: 7, red: false}

	traverse(&root)
}

func traverse(node *RBNode) {
	if node == nil {
		return
	}

	traverse(node.left)
	print(node)
	traverse(node.right)
}

func print(node *RBNode) {
	color := "B"
	if node.red == true {
		color = "R"
	}

	fmt.Printf("%d,%s ", node.value, color)
}
