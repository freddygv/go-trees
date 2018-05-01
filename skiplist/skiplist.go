package skiplist

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// List exports a reference to the head of the list
type List struct {
	Head *Node
	tail *Node
	rnd  *rand.Rand
}

const (
	maxHeight = 31
	minValue  = math.MinInt32
	maxValue  = math.MaxInt32
)

// New returns an empty Skip List
func New() *List {
	root := &Node{
		Value:     minValue,
		Level:     maxHeight,
		Shortcuts: make([]*Node, maxHeight),
	}
	tail := &Node{
		Value:     maxValue,
		Level:     maxHeight,
		Shortcuts: make([]*Node, maxHeight),
	}
	return &List{
		Head: root,
		tail: tail,
		rnd:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Get searches the list for a target, returns node ptr and boolean indicating if found.
// If the target value is not found, a pointer to the previous node is returned
func (list *List) Get(target int) (*Node, bool) {
	level := maxHeight
	current := list.Head

	for level >= 0 {
		neighbor := current.Shortcuts[level]

		// Advance while the neighbor does not overshoot target
		for compare(neighbor.Value, target) <= 0 {
			current = neighbor
			neighbor = current.Shortcuts[level]
		}

		if compare(target, current.Value) == 0 {
			return current, true
		}
		level--
	}

	return current, false
}

// Insert will add a new node to the list with the given value
func (list *List) Insert(value int) {
	var level int8

	// Generate random int32 and take least significant bits as coinflips to increment level
	r := list.rnd.Int31()
	for r&1 == 1 {
		level++
		r >>= 1
	}

	// TODO: Update shortcuts
	// previous, _ := list.Get(value)
	// inserted := Node{
	// 	Value:     value,
	// 	Level:     level,
	// 	Shortcuts: make([]*Node, maxHeight),
	// }
}

// TODO:
func (list *List) toSlice() []*Node {
	arr := make([]*Node, 0)

	return arr
}

// Node is a sub-list
type Node struct {
	Value     interface{}
	Level     int8
	Shortcuts []*Node
}

func (node *Node) String() string {
	return fmt.Sprintf("%v,%v", node.Value, node.Level)
}

// TODO: Don't panic
func compare(a, b interface{}) int {
	intA, okA := a.(int)
	intB, okB := b.(int)
	if !okA || !okB {
		panic("item is not an integer")
	}
	return intA - intB
}
