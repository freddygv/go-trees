package treap

import (
	"math"
	"math/rand"
)

// Tree contains a reference to the root of the Red-Black tree
type Tree struct {
	Root *Node
}

// Randomized priorities are in the range of [0 - 2^31)
const maxPriority = math.MaxInt32

// NewTree returns a red-black tree storing the single value given as the black root.
func NewTree(value int) *Tree {
	root := &Node{Value: value, Priority: rand.Intn(maxPriority)}
	return &Tree{Root: root}
}

// Get searches a Treap for a value, returns node ptr and boolean indicating if found
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
	if compare(value, node.Value) < 0 {
		if node.Left == nil {
			node.Left = &Node{Value: value, Priority: rand.Intn(maxPriority), Parent: node}
			return node.Left
		}
		return node.Left.naiveInsert(value)

	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value, Priority: rand.Intn(maxPriority), Parent: node}
			return node.Right
		}
		return node.Right.naiveInsert(value)

	}
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
