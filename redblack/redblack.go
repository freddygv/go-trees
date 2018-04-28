package redblack

import (
	"fmt"
)

// Tree contains a reference to the root of the Red-Black tree
type Tree struct {
	root *Node
}

// NewTree returns a red-black tree storing the single value given as the black root.
func NewTree(value int) *Tree {
	return &Tree{root: newNode(value, false, nil)}
}

// Contains searches a Red-Black Tree for a value recursively
func (tree *Tree) Contains(value int) bool {
	root := tree.root

	for root != nil {
		if value == root.value {
			return true
		}
		if value < root.value {
			root = root.left
		}
		if value > root.value {
			root = root.right
		}
	}

	return false
}

// Insert will add a new node to the tree with the given value
func (tree *Tree) Insert(value int) {
	current := tree.root.naiveInsert(value)
	if current.parent.parent == nil {
		return
	}

	// Loop until reaching root (nil parent pointer), black node, or black parent
	for current.parent != nil && current.red && current.parent.red {
		parent := current.parent
		grandparent := parent.parent

		if parent == grandparent.left {
			uncle := grandparent.right
			if uncle.red {
				// Case 1A: Re-color and move up
				uncle.red = false
				parent.red = false
				grandparent.red = true
				current = grandparent

			} else if current == parent.right {
				// Case 2A: Zigzag from grandparent to current, left then right
				current = parent
				current.leftRotate()

			}
			if current.parent != nil && current == current.parent.left {
				// Case 3A: Straight from grandparent, left then left
				parent = current.parent
				grandparent = parent.parent
				grandparent.rightRotate()

				// Parent becomes black root and grandparent becomes red sibling
				parent.red = false
				grandparent.red = true
				current = parent
			}
		} else { // Reverse left and right
			uncle := grandparent.left
			if uncle.red {
				// Case 1B: Re-color and move up
				uncle.red = false
				parent.red = false
				grandparent.red = true
				current = grandparent
				continue

			} else if current == parent.left {
				// Case 2B: Zigzag from grandparent to current, right then left
				current = parent
				current.rightRotate()

			}
			if current.parent != nil && current == current.parent.right {
				// Case 3B: Straight from grandparent, right then right
				parent = current.parent
				grandparent = parent.parent
				grandparent.leftRotate()

				// Parent becomes black root and grandparent becomes red sibling
				parent.red = false
				grandparent.red = true
				current = parent
			}
		}
	}
	// Re-color root if needed
	if current.parent == nil {
		current.red = false
		tree.root = current
	}
}

func (tree *Tree) ToSlice(arr *[]*Node) {
	tree.root.flatten(arr)
}

// Node is a sub-tree in a Red-Black tree
type Node struct {
	value  int
	red    bool
	left   *Node
	right  *Node
	parent *Node
}

// newNode adds a new red node with two empty black leaves
func newNode(value int, red bool, parent *Node) *Node {
	node := Node{value: value, red: red, parent: parent}

	l := Node{parent: &node}
	node.left = &l

	r := Node{parent: &node}
	node.right = &r

	return &node
}

// Naive BST insertion for a given value (new nodes are always red)
func (root *Node) naiveInsert(value int) *Node {
	if value < root.value {
		if root.left.isLeaf() {
			root.left = newNode(value, true, root)
			return root.left
		}
		return root.left.naiveInsert(value)

	} else {
		if root.right.isLeaf() {
			root.right = newNode(value, true, root)
			return root.right
		}
		return root.right.naiveInsert(value)

	}
}

// isLeaf checks if a node is a child-less black sentinel
func (node *Node) isLeaf() bool {
	if node.left == nil && node.right == nil && node.red == false {
		return true
	}
	return false
}

func (node *Node) rightRotate() {
	left := node.left
	parent := node.parent

	// Promote node to be its grandparent's child
	if parent != nil && parent.value > node.value {
		parent.left = left

	} else if parent != nil && parent.value <= node.value {
		parent.right = left

	}
	left.parent = parent

	// Hand over the right child of the left node
	node.left = left.right

	// Swap parent/child relationship
	left.right = node
	node.parent = left
}

func (node *Node) leftRotate() {
	right := node.right
	parent := node.parent

	// Promote node to be its grandparent's child
	if parent != nil && parent.value > node.value {
		parent.left = right

	} else if parent != nil && parent.value <= node.value {
		parent.right = right

	}
	right.parent = parent

	// Hand over the left child of the right node
	node.right = right.left

	// Swap parent/child relationship
	right.left = node
	node.parent = right
}

// In order traversal to flatten tree into slice
func (root *Node) flatten(arr *[]*Node) {
	if root == nil {
		return
	}

	root.left.flatten(arr)

	if !root.isLeaf() {
		*arr = append(*arr, root)
	}

	root.right.flatten(arr)
}

// String representation of a black node with value 7 is: 7,B
func (node *Node) String() string {
	color := "B"
	if node.red == true {
		color = "R"
	}

	return fmt.Sprintf("%d,%s", node.value, color)
}
