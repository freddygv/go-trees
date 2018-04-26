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
	root := NewTree(4)
	root.Insert(2)
	root.Insert(1)
	root.Insert(3)
	root.Insert(6)
	root.Insert(5)
	root.Insert(7)
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

// NewTree returns a red-black tree storing the single value given
func NewTree(value int) *Tree {
	return &Tree{value: value}
}

// Insert will add a new node to the tree with the given value
func (tree *Tree) Insert(value int) {
	if value < tree.value {
		if tree.left == nil {
			tree.left = &Tree{value: value, parent: tree}
			return
		}
		tree.left.Insert(value)

	} else {
		if tree.right == nil {
			tree.right = &Tree{value: value, parent: tree}
			return
		}
		tree.right.Insert(value)

	}
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

	if value == tree.value {
		return true
	}

	if value < tree.value {
		return tree.left.Contains(value)
	}

	if value > tree.value {
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
