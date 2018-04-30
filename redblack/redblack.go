package redblack

import (
	"fmt"
)

// Tree contains a reference to the root of the Red-Black tree
type Tree struct {
	Root *Node
}

// NewTree returns an empty red-black tree reference.
func NewTree() *Tree {
	return &Tree{}
}

// Get searches a Red-Black Tree for a value, returns node ptr and boolean indicating if found
func (tree *Tree) Get(value int) (*Node, bool) {
	root := tree.Root

	for root != nil {
		if compare(value, root.Value) == 0 {
			return root, true
		}
		if compare(value, root.Value) < 0 {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return nil, false
}

// Insert will add a new node to the tree with the given value
func (tree *Tree) Insert(value int) {
	current := tree.naiveInsert(value)
	if current.Parent == nil || current.Parent.Parent == nil {
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

// Naive BST insertion for a given value (new nodes are always red)
func (tree *Tree) naiveInsert(value int) *Node {
	var inserted *Node

	root := tree.Root
	if root == nil {
		inserted = newNode(value, true, nil)
		tree.Root = inserted
	}

	for inserted == nil {
		if compare(value, root.Value) < 0 {
			if root.Left.isLeaf() {
				root.Left = newNode(value, true, root)
				inserted = root.Left
			} else {
				root = root.Left
			}

		} else {
			// Duplicate values placed on the right
			if root.Right.isLeaf() {
				root.Right = newNode(value, true, root)
				inserted = root.Right
			} else {
				root = root.Right
			}

		}
	}
	return inserted
}

func (tree *Tree) toSlice() []*Node {
	arr := make([]*Node, 0)
	tree.Root.flatten(&arr)
	return arr
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
	if parent != nil && compare(node.Value, parent.Value) < 0 {
		parent.Left = Left

	} else if parent != nil && compare(node.Value, parent.Value) >= 0 {
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
	if parent != nil && compare(node.Value, parent.Value) < 0 {
		parent.Left = Right

	} else if parent != nil && compare(node.Value, parent.Value) >= 0 {
		parent.Right = Right

	}
	Right.Parent = parent

	// Hand over the Left child of the Right node
	node.Right = Right.Left

	// Swap parent/child relationship
	Right.Left = node
	node.Parent = Right
}

func compare(a, b int) int {
	return a - b
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

	return fmt.Sprintf("%v,%s", node.Value, color)
}
