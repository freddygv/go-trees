package treap

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Tree contains a reference to the root of the tree
type Tree struct {
	Root *Node
	rnd  *rand.Rand
}

// Randomized priorities are in the range of [0 - 2^31)
const maxPriority = math.MaxInt32

// NewTree returns a treap Tree storing the single value given as the root.
func NewTree(value int) *Tree {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Tree{
		Root: &Node{Value: value, Priority: r.Intn(maxPriority)},
		rnd:  r,
	}
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

// Insert will add a new node to the tree with the given value
func (tree *Tree) Insert(value int) {
	current := tree.naiveInsert(value)

	// Bubble up while the current node's priority is lower than its parent's
	for current.Parent != nil && compare(current.Priority, current.Parent.Priority) < 0 {
		if current == current.Parent.Left {
			current.Parent.rightRotate()
		} else {
			current.Parent.leftRotate()
		}

		if current.Parent == nil {
			tree.Root = current
		}
	}
}

// Naive BST insertion for a given value
func (tree *Tree) naiveInsert(value int) *Node {
	root := tree.Root

	var inserted *Node
	for inserted == nil {
		if compare(value, root.Value) < 0 {
			if root.Left == nil {
				root.Left = &Node{Value: value, Priority: tree.rnd.Intn(maxPriority), Parent: root}
				inserted = root.Left
			} else {
				root = root.Left
			}

		} else {
			// Duplicate values placed on the right
			if root.Right == nil {
				root.Right = &Node{Value: value, Priority: tree.rnd.Intn(maxPriority), Parent: root}
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

// Node is a sub-tree
type Node struct {
	Value    int
	Priority int
	Left     *Node
	Right    *Node
	Parent   *Node
}

func (node *Node) rightRotate() {
	Left := node.Left
	parent := node.Parent

	// Promote node to be its grandparent's child
	// If the current node is the parent's left child, make node.Left the parent's left child
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

// In order traversal to flatten tree into slice
func (node *Node) flatten(arr *[]*Node) {
	if node == nil {
		return
	}

	node.Left.flatten(arr)

	*arr = append(*arr, node)

	node.Right.flatten(arr)
}

func (node *Node) String() string {
	return fmt.Sprintf("(%d, %d)", node.Value, node.Priority)
}

func compare(a, b int) int {
	return a - b
}
