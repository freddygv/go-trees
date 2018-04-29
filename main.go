package main

import "github.com/freddygv/go-trees/treap"

func main() {
	tree := treap.NewTree(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
}
