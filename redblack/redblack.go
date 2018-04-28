package redblack

import (
	"fmt"
)

// Tree contains a reference to the root of the Red-Black tree
type Tree struct {
	Root *Node
}

// NewTree returns a red-black tree storing the single value given as the black root.
func NewTree(value int) *Tree {
	return &Tree{Root: newNode(value, false, nil)}
}

// Contains searches a Red-Black Tree for a value recursively
func (tree *Tree) Contains(value int) bool {
	root := tree.Root

	for root != nil {
		if value == root.Value {
			return true
		}
		if value < root.Value {
			root = root.Left
		}
		if value > root.Value {
			root = root.Right
		}
	}

	return false
}

// Insert will add a new node to the tree with the given value
func (tree *Tree) Insert(value int) {
	current := tree.Root.naiveInsert(value)
	if current.Parent.Parent == nil {
		return
	}

	// Loop until reaching root (nil parent pointer), black node, or black parent
	for current.Parent != nil && current.red && current.Parent.red {
		parent := current.Parent
		grandparent := parent.Parent

		if parent == grandparent.Left {
			uncle := grandparent.Right
			if uncle.red {
				// Case 1A: Re-color and move up
				uncle.red = false
				parent.red = false
				grandparent.red = true
				current = grandparent
				continue

			} else if current == parent.Right {
				// Case 2A: Zigzag from grandparent to current, Left then Right
				current = parent
				current.leftRotate()

			}
			if current.Parent != nil && current == current.Parent.Left {
				// Case 3A: Straight from grandparent, Left then Left
				parent = current.Parent
				grandparent = parent.Parent
				grandparent.rightRotate()

				// Parent becomes black root and grandparent becomes red sibling
				parent.red = false
				grandparent.red = true
				current = parent
			}
		} else { // Reverse Left and Right
			uncle := grandparent.Left
			if uncle.red {
				// Case 1B: Re-color and move up
				uncle.red = false
				parent.red = false
				grandparent.red = true
				current = grandparent
				continue

			} else if current == parent.Left {
				// Case 2B: Zigzag from grandparent to current, Right then Left
				current = parent
				current.rightRotate()

			}
			if current.Parent != nil && current == current.Parent.Right {
				// Case 3B: Straight from grandparent, Right then Right
				parent = current.Parent
				grandparent = parent.Parent
				grandparent.leftRotate()

				// Parent becomes black root and grandparent becomes red sibling
				parent.red = false
				grandparent.red = true
				current = parent
			}
		}
	}
	// Re-color root if needed
	if current.Parent == nil {
		current.red = false
		tree.Root = current
	}
}

func (tree *Tree) toSlice(arr *[]*Node) {
	tree.Root.flatten(arr)
}

// Node is a sub-tree in a Red-Black tree
type Node struct {
	Value  int
	red    bool
	Left   *Node
	Right  *Node
	Parent *Node
}

// newNode adds a new red node with two empty black leaves
func newNode(value int, red bool, parent *Node) *Node {
	node := Node{Value: value, red: red, Parent: parent}

	l := Node{Parent: &node}
	node.Left = &l

	r := Node{Parent: &node}
	node.Right = &r

	return &node
}

// Naive BST insertion for a given value (new nodes are always red)
func (node *Node) naiveInsert(value int) *Node {
	if value < node.Value {
		if node.Left.isLeaf() {
			node.Left = newNode(value, true, node)
			return node.Left
		}
		return node.Left.naiveInsert(value)

	} else {
		if node.Right.isLeaf() {
			node.Right = newNode(value, true, node)
			return node.Right
		}
		return node.Right.naiveInsert(value)

	}
}

// isLeaf checks if a node is a child-less black sentinel
func (node *Node) isLeaf() bool {
	if node.Left == nil && node.Right == nil && node.red == false {
		return true
	}
	return false
}

func (node *Node) rightRotate() {
	Left := node.Left
	parent := node.Parent

	// Promote node to be its grandparent's child
	if parent != nil && parent.Value > node.Value {
		parent.Left = Left

	} else if parent != nil && parent.Value <= node.Value {
		parent.Right = Left

	}
	Left.Parent = parent

	// Hand over the Right child of the Left node
	node.Left = Left.Right

	// Swap parent/child relationship
	Left.Right = node
	node.Parent = Left
}

func (node *Node) leftRotate() {
	Right := node.Right
	parent := node.Parent

	// Promote node to be its grandparent's child
	if parent != nil && parent.Value > node.Value {
		parent.Left = Right

	} else if parent != nil && parent.Value <= node.Value {
		parent.Right = Right

	}
	Right.Parent = parent

	// Hand over the Left child of the Right node
	node.Right = Right.Left

	// Swap parent/child relationship
	Right.Left = node
	node.Parent = Right
}

// In order traversal to flatten tree into slice
func (node *Node) flatten(arr *[]*Node) {
	if node == nil {
		return
	}

	node.Left.flatten(arr)

	if !node.isLeaf() {
		*arr = append(*arr, node)
	}

	node.Right.flatten(arr)
}

// String representation of a black node with value 7 is: 7,B
func (node *Node) String() string {
	color := "B"
	if node.red == true {
		color = "R"
	}

	return fmt.Sprintf("%d,%s", node.Value, color)
}
