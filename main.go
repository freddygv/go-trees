package main

import (
	"fmt"
)

/*

      4
  2      6
1   3  5   7

*/
func main() {
	root := Tree{value: 4, red: false}

	root.left = &Tree{value: 2, red: false, parent: &root}
	root.left.left = &Tree{value: 1, red: false, parent: root.left}
	root.left.right = &Tree{value: 3, red: false, parent: root.left}

	root.right = &Tree{value: 6, red: false, parent: &root}
	root.right.left = &Tree{value: 5, red: false, parent: root.right}
	root.right.right = &Tree{value: 7, red: false, parent: root.right}
}

// In order traversal of the tree for printing
func traverse(root *Tree) {
	if root == nil {
		return
	}

	traverse(root.left)
	fmt.Printf("%v ", root)
	traverse(root.right)
}

// Tree is a sub-tree in a Red-Black Tree
type Tree struct {
	value  int
	red    bool
	left   *Tree
	right  *Tree
	parent *Tree
}

func (tree *Tree) rightRotate() {
	left := tree.left
	parent := tree.parent

	// Promote the left node over the current root
	if parent != nil && parent.value > tree.value {
		parent.left = left

	} else if parent != nil && parent.value <= tree.value {
		parent.right = left

	}
	left.parent = parent

	// Hand over the right child of the left node
	tree.left = left.right

	// Swap parent/child relationship
	left.right = tree
	tree.parent = left
}

func (tree *Tree) leftRotate() {
	right := tree.right
	parent := tree.parent

	// Promote the right node over the current root
	if parent.value > tree.value {
		parent.left = right
	} else {
		parent.right = right
	}
	right.parent = parent

	// Hand over the left child of the right node
	tree.right = right.left

	// Swap parent/child relationship
	right.left = tree
	tree.parent = right
}

// Contains searches a Red-Black Tree for a value recursively
func (tree *Tree) Contains(value int) bool {
	if tree == nil {
		return false
	}

	if tree.value == value {
		return true
	}

	if tree.value > value {
		return tree.left.Contains(value)
	}

	if tree.value < value {
		return tree.right.Contains(value)
	}

	return false
}

func (tree *Tree) String() string {
	color := "B"
	if tree.red == true {
		color = "R"
	}

	return fmt.Sprintf("%d,%s", tree.value, color)
}
