package main

import (
	"fmt"
)

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

	traverse(&root, 0)
}

func traverse(node *RBNode, indent int) {
	if node == nil {
		return
	}

	traverse(node.left, indent+4)
	print(node)
	traverse(node.right, indent+4)
}

func print(node *RBNode) {
	color := "B"
	if node.red == true {
		color = "R"
	}

	fmt.Printf("%d,%s ", node.value, color)
}
