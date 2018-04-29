package treap

import (
	"math"
	"math/rand"
)

// Tree contains a reference to the root of the Red-Black tree
type Tree struct {
	Root *Node
}

// NewTree returns a red-black tree storing the single value given as the black root.
func NewTree(value int) *Tree {
	root := &Node{Value: value, Priority: rand.Intn(math.MaxInt32)}
	return &Tree{Root: root}
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

// Node is a sub-tree in a Red-Black tree
type Node struct {
	Value    int
	Priority int
	Left     *Node
	Right    *Node
	Parent   *Node
}

// Naive BST insertion for a given value
func (node *Node) naiveInsert(value int) *Node {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value, Priority: rand.Intn(math.MaxInt32), Parent: node}
			return node.Left
		}
		return node.Left.naiveInsert(value)

	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value, Priority: rand.Intn(math.MaxInt32), Parent: node}
			return node.Right
		}
		return node.Right.naiveInsert(value)

	}
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
