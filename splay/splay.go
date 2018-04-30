package splay

import (
	"fmt"
)

// Tree contains a reference to the root of the tree
type Tree struct {
	Root *Node
}

// NewTree returns an empty splay Tree.
func NewTree() *Tree {
	return &Tree{}
}

// Get searches a Treap for a value, returns node ptr and boolean indicating if found
func (tree *Tree) Get(value int) (*Node, bool) {
	current := tree.Root
	parent := current.Parent

	for current != nil {
		if compare(value, current.Value) == 0 {
			// Splay the inserted node to make it the root
			tree.splay(current)
			return current, true
		}

		parent = current
		if compare(value, current.Value) < 0 {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	// Splay the parent of where the node would have been
	tree.splay(parent)
	return nil, false
}

// Insert will add a new node to the tree with the given value
func (tree *Tree) Insert(value int) {
	current := tree.naiveInsert(value)
	tree.splay(current)
}

// Naive BST insertion for a given value
func (tree *Tree) naiveInsert(value int) *Node {
	var inserted *Node

	root := tree.Root
	if root == nil {
		inserted = &Node{Value: value}
		tree.Root = inserted
	}

	for inserted == nil {
		if compare(value, root.Value) < 0 {
			if root.Left == nil {
				root.Left = &Node{Value: value, Parent: root}
				inserted = root.Left
			} else {
				root = root.Left
			}

		} else {
			// Duplicate values placed on the right
			if root.Right == nil {
				root.Right = &Node{Value: value, Parent: root}
				inserted = root.Right
			} else {
				root = root.Right
			}
		}
	}
	return inserted
}

func (tree *Tree) splay(node *Node) {
	if node == tree.Root {
		return
	}

	for node.Parent != nil {
		root := tree.Root
		parent := node.Parent

		// Zig (left child of root || right child of root)
		if parent == root {
			if node == root.Left {
				root.rightRotate()
				break
			}
			root.leftRotate()
			break
		}

		grandparent := parent.Parent
		if parent == grandparent.Left {
			if node == parent.Right {
				// Zig-zag (right child of left child)
				parent.leftRotate()
				grandparent.rightRotate()

			} else {
				// Zig-zig (left child of left child)
				grandparent.rightRotate()
				parent.rightRotate()

			}
		} else {
			if node == parent.Left {
				// Zig-zag (left child of right child)
				parent.rightRotate()
				grandparent.leftRotate()

			} else {
				// Zig-zig (right child of right child)
				grandparent.leftRotate()
				parent.leftRotate()

			}
		}
	}
	tree.Root = node
}

func (tree *Tree) toSlice() []*Node {
	arr := make([]*Node, 0)
	tree.Root.flatten(&arr)
	return arr
}

// Node is a sub-tree
type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
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
	if Left.Right != nil {
		Left.Right.Parent = node
	}

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
	if Right.Left != nil {
		Right.Left.Parent = node
	}

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
	return fmt.Sprintf("%v", node.Value)
}

func compare(a, b int) int {
	return a - b
}
