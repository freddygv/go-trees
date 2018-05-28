package skiplist

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"
)

// List exports a reference to the head of the list
type List struct {
	Head   *Node
	tail   *Node
	height int8
	rnd    *rand.Rand
}

const (
	maxHeight = 31
	minValue  = math.MinInt64
	maxValue  = math.MaxInt64
)

// New returns an empty Skip List
func New() *List {
	root := Node{
		Value:     minValue,
		Level:     maxHeight,
		Shortcuts: make([]*Node, maxHeight+1),
	}
	tail := Node{Value: maxValue}

	// Create shortcuts from root to tail at all (maxHeight) levels
	for i := 0; i <= maxHeight; i++ {
		root.Shortcuts[i] = &tail
	}

	return &List{
		Head: &root,
		tail: &tail,
		rnd:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Get searches the list for a target, returns node ptr and boolean indicating if found.
func (list *List) Get(target interface{}) (*Node, bool) {
	current := list.Head

	for i := list.height; i >= 0; i-- {
		neighbor := current.Shortcuts[i]

		// Advance while the neighbor does not overshoot target
		for neighbor != nil && compare(neighbor.Value, target) <= 0 {
			current = neighbor
			neighbor = current.Shortcuts[i]
		}

		if compare(target, current.Value) == 0 {
			return current, true
		}
	}

	return nil, false
}

// Insert will add a new node to the list with the given value
func (list *List) Insert(value interface{}) {
	var level int8

	// Generate random int32 and take least significant bits as coinflips to increment level
	r := list.rnd.Int31()
	for r&1 == 1 {
		level++
		r >>= 1
	}

	if level > list.height {
		list.height = level
	}

	new := Node{
		Value:     value,
		Level:     level,
		Shortcuts: make([]*Node, maxHeight+1),
	}

	current := list.Head

	for i := list.height; i >= 0; i-- {
		neighbor := current.Shortcuts[i]

		// Advance while the neighbor does not overshoot target
		for compare(neighbor.Value, value) <= 0 {
			current = neighbor
			neighbor = current.Shortcuts[i]
		}

		if i > level {
			continue
		}

		// Insert node at current level by updating neighbor shortcuts
		new.Shortcuts[i] = neighbor
		current.Shortcuts[i] = &new
	}
}

func (list *List) toSlice() []*Node {
	arr := make([]*Node, 0)

	current := list.Head
	next := current.Shortcuts[0]

	// Continue until tail, which has no Shortcuts
	for next.Shortcuts != nil {
		arr = append(arr, next)
		current = next
		next = current.Shortcuts[0]
	}

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
		err := fmt.Errorf("compare expected: (int, int), got: (%v, %v)", reflect.TypeOf(a), reflect.TypeOf(b))
		panic(err)
	}
	return intA - intB
}
