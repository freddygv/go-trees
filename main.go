package main

import (
	"fmt"
)

/*

         7,B
  3,B          18,R
	     10,B       22,B
      8,R   11,R       26,R

*/
func main() {
	root := NewTree(7)

	root.left = newNode(3, root)
	root.left.red = false

	root.right = newNode(18, root)

	root.right.left = newNode(10, root.right)
	root.right.left.red = false
	root.right.left.left = newNode(8, root.right.left)
	root.right.left.right = newNode(11, root.right.left)

	root.right.right = newNode(22, root.right)
	root.right.right.red = false
	root.right.right.right = newNode(26, root.right.right)

	root.naiveInsert(15)

	traverse(root)
}

// In order traversal of the tree for printing
func traverse(root *Tree) {
	if root == nil {
		return
	}

	traverse(root.left)

	if !root.isLeaf() {
		fmt.Printf("%v ", root)
	}

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

// NewTree returns a red-black tree storing the single value given as the black root.
func NewTree(value int) *Tree {
	return &Tree{value: value}
}

// Insert will add a new node to the tree with the given value
func (tree *Tree) Insert(value int) {
	inserted := tree.naiveInsert(value)

	current := inserted
	// Loop until reaching root or a black node
	for current.parent != nil && current.red == true {
		parent := current.parent
		grandparent := parent.parent

		if parent == grandparent.left {
			uncle := grandparent.right
			if uncle.red {
				fmt.Println("Case 1: Re-color and move up")
				uncle.red = false
				parent.red = false
				current = grandparent

			} else if current == parent.right {
				fmt.Println("Case 2: Zigzag from GP to current, left then right")
				current = parent
				current.leftRotate()

				if current == parent.left {
					fmt.Println("Case 3: Straight from GP, left then left")
					parent := current.parent
					grandparent := parent.parent
					grandparent.rightRotate()

					// Parent becomes black root and GP becomes red sibling
					parent.red = false
					grandparent.red = true
					current = parent
				}
			}
		} else { // Reverse left and right
			uncle := grandparent.left
			if uncle.red {
				fmt.Println("Case 1: Re-color and move up")
				uncle.red = false
				parent.red = false
				current = grandparent

			} else if current == parent.left {
				fmt.Println("Case 2: Zigzag from GP to current, right then left")
				current = parent
				current.rightRotate()

				if current == parent.right {
					fmt.Println("Case 3: Straight from GP, right then right")
					parent := current.parent
					grandparent := parent.parent
					grandparent.leftRotate()

					// Parent becomes black root and GP becomes red sibling
					parent.red = false
					grandparent.red = true
					current = parent
				}
			}
		}
	}
	// Re-color root if needed
	if current.parent == nil {
		current.red = false
	}
}

// Naive BST insertion for a given value (new nodes are always red)
func (tree *Tree) naiveInsert(value int) *Tree {
	if value < tree.value {
		if tree.left.isLeaf() {
			tree.left = newNode(value, tree)
			return tree.left
		}
		tree.left.naiveInsert(value)

	} else {
		if tree.right.isLeaf() {
			tree.right = newNode(value, tree)
			return tree.right
		}
		tree.right.naiveInsert(value)

	}
	return nil
}

// isLeaf checks if a node is a child-less black sentinel
func (tree *Tree) isLeaf() bool {
	if tree.left == nil && tree.right == nil && tree.red == false {
		return true
	}
	return false
}

// newNode adds a new red node with two empty black leaves
func newNode(value int, parent *Tree) *Tree {
	node := Tree{value: value, red: true, parent: parent}

	l := Tree{parent: &node}
	node.left = &l

	r := Tree{parent: &node}
	node.right = &r

	return &node
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

// String representation of a black node with value 7 is: 7,B
func (tree *Tree) String() string {
	color := "B"
	if tree.red == true {
		color = "R"
	}

	return fmt.Sprintf("%d,%s", tree.value, color)
}
