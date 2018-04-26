package main

import (
	"fmt"
)

/*

      4
  2	     6
1   3  5   7

*/
func main() {
	root := Tree{value: 4, red: false}

	root.left = &Tree{value: 2, red: false}
	root.left.left = &Tree{value: 1, red: false}
	root.left.right = &Tree{value: 3, red: false}

	root.right = &Tree{value: 6, red: false}
	root.right.left = &Tree{value: 5, red: false}
	root.right.right = &Tree{value: 7, red: false}
}

// In order traversal of the tree for printing
func traverse(root *Tree) {
	if root == nil {
		return
	}

	traverse(root.left)
	fmt.Println(root)
	traverse(root.right)
}

// Tree is a sub-tree in a Red-Black Tree
type Tree struct {
	value int
	red   bool
	left  *Tree
	right *Tree
}

// Contains searches a Red-Black Tree for a value recursively
func (t *Tree) Contains(value int) bool {
	if t == nil {
		return false
	}

	if t.value == value {
		return true
	}

	if t.value > value {
		return t.left.Contains(value)
	}

	if t.value < value {
		return t.right.Contains(value)
	}

	return false
}

func (t *Tree) String() string {
	color := "B"
	if t.red == true {
		color = "R"
	}

	return fmt.Sprintf("%d,%s", t.value, color)
}
