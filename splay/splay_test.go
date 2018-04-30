package splay

import (
	"testing"
)

func TestInsert(t *testing.T) {
	t.Run("1 to 7 in order traversal", func(t *testing.T) {
		t.Parallel()

		tree := NewTree()
		tree.Insert(1)
		tree.Insert(2)
		tree.Insert(3)
		tree.Insert(4)
		tree.Insert(5)
		tree.Insert(6)
		tree.Insert(7)

		result := tree.toSlice()

		for i := 1; i < 8; i++ {
			if result[i-1].Value != i {
				t.Fatalf("failed to find: %v", i)
			}
		}
	})

	t.Run("7 to 1 in order traversal", func(t *testing.T) {
		t.Parallel()

		tree := NewTree()
		tree.Insert(7)
		tree.Insert(6)
		tree.Insert(5)
		tree.Insert(4)
		tree.Insert(3)
		tree.Insert(2)
		tree.Insert(1)

		result := tree.toSlice()

		for i := 1; i < 8; i++ {
			if result[i-1].Value != i {
				t.Fatalf("failed to find: %v", i)
			}
		}
	})
}

func TestGet(t *testing.T) {
	t.Run("1 to 7 contained", func(t *testing.T) {
		t.Parallel()

		tree := NewTree()
		tree.Insert(1)
		tree.Insert(2)
		tree.Insert(3)
		tree.Insert(4)
		tree.Insert(5)
		tree.Insert(6)
		tree.Insert(7)

		for i := 1; i < 8; i++ {
			if _, ok := tree.Get(i); !ok {
				t.Fatalf("failed to find: %v", i)
			}
		}
	})

	t.Run("1 to 7 and two gets", func(t *testing.T) {
		t.Parallel()

		tree := NewTree()
		tree.Insert(1)
		tree.Insert(2)
		tree.Insert(3)
		tree.Insert(4)
		tree.Insert(5)
		tree.Insert(6)
		tree.Insert(7)

		root, _ := tree.Get(1)

		if root != tree.Root {
			t.Fatalf("failed to splay: %v", root.Value)
		}

		result := []*Node{
			root,
			root.Right.Left.Left,
			root.Right.Left.Left.Right,
			root.Right.Left,
			root.Right.Left.Right,
			root.Right,
			root.Right.Right,
		}

		for i := 1; i < 8; i++ {
			if result[i-1].Value != i {
				t.Fatalf("expected: '%v', got: '%v'", i, result[i-1])
			}
		}

		n, ok := tree.Get(8)
		if n != nil {
			t.Fatalf("expected nil pointer for missing value")
		}
		if ok {
			t.Fatalf("expecting false for Get result")
		}
		if tree.Root.Value != 7 {
			t.Fatalf("expecting next closest value at root")
		}
	})
}
