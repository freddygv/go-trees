package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

// List contains a reference to the root of the list
type List struct {
	Root *List
	rnd  *rand.Rand
}

const maxHeight = 31

// New returns an empty Skip List
func New() *List {
	return &List{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Get searches the list for a value, returns node ptr and boolean indicating if found
// TODO:
func (tree *List) Get(value int) (*List, bool) {
	return nil, false
}

// Insert will add a new node to the list with the given value
// TODO:
func (tree *List) Insert(value int) {

}

// TODO:
func (tree *List) toSlice() []*List {
	arr := make([]*List, 0)

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

// TODO:
func (node *Node) String() string {
	return fmt.Sprintf("%v", node.Value)
}

func compare(a, b int) int {
	return a - b
}
