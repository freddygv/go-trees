package main

import (
	"fmt"

	"github.com/freddygv/go-trees/redblack"
)

func main() {
	t := redblack.NewTree(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	t.Insert(7)

	arr := make([]*redblack.Node, 0)
	t.ToSlice(&arr)
	fmt.Println(arr)
}
