package main

import (
	"fmt"

	"github.com/freddygv/go-trees/redblack"
)

func main() {
	t := redblack.NewTree(7)
	t.Insert(6)
	t.Insert(5)
	t.Insert(4)
	t.Insert(3)
	t.Insert(2)
	t.Insert(1)

	nodes := make([]*redblack.Node, 0)
	t.ToSlice(&nodes)
	fmt.Println(nodes)
}
