package main

import (
	"math/rand"

	"github.com/freddygv/go-trees/splay"
)

func main() {
	// This is returning the right result.
	// Why would it grow with a larger N?
	t := splay.NewTree()

	var inputs []int
	var reads []int

	for i := 0; i < 1000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, inputs[0])
	}

	for _, v := range inputs {
		t.Insert(v)
	}

	for _, v := range reads {
		t.Get(v)
	}
}
