package redblack

import (
	"testing"
)

func TestInsert(t *testing.T) {
	t.Run("1 to 7", func(t *testing.T) {
		t.Parallel()

		tree := NewTree(1)
		tree.Insert(2)
		tree.Insert(3)
		tree.Insert(4)
		tree.Insert(5)
		tree.Insert(6)
		tree.Insert(7)

		expect := []*Node{
			&Node{Value: 1, red: false},
			&Node{Value: 2, red: false},
			&Node{Value: 3, red: false},
			&Node{Value: 4, red: true},
			&Node{Value: 5, red: true},
			&Node{Value: 6, red: false},
			&Node{Value: 7, red: true},
		}

		result := make([]*Node, 0)
		tree.ToSlice(&result)

		for i := 0; i < len(expect); i++ {
			if expect[i].Value != result[i].Value || expect[i].red != result[i].red {
				t.Fatalf("expected: %v, got: %v", expect[i], result[i])
			}
		}
	})

	t.Run("7 to 1", func(t *testing.T) {
		t.Parallel()

		tree := NewTree(7)
		tree.Insert(6)
		tree.Insert(5)
		tree.Insert(4)
		tree.Insert(3)
		tree.Insert(2)
		tree.Insert(1)

		expect := []*Node{
			&Node{Value: 1, red: true},
			&Node{Value: 2, red: false},
			&Node{Value: 3, red: true},
			&Node{Value: 4, red: true},
			&Node{Value: 5, red: false},
			&Node{Value: 6, red: false},
			&Node{Value: 7, red: false},
		}

		result := make([]*Node, 0)
		tree.ToSlice(&result)

		for i := 0; i < len(expect); i++ {
			if expect[i].Value != result[i].Value || expect[i].red != result[i].red {
				t.Fatalf("expected: %v, got: %v", expect[i], result[i])
			}
		}
	})

	t.Run("3 case insert A", func(t *testing.T) {
		t.Parallel()

		tree := NewTree(22)

		root := tree.Root
		root.Right = newNode(26, false, root)

		root.Left = newNode(8, true, root)

		root.Left.Left = newNode(7, false, root.Left)
		root.Left.Left.Left = newNode(3, true, root.Left.Left)

		root.Left.Right = newNode(15, false, root.Left)
		root.Left.Right.Left = newNode(11, true, root.Left.Right)
		root.Left.Right.Right = newNode(18, true, root.Left.Right)

		tree.Insert(10)

		expect := []*Node{
			&Node{Value: 3, red: true},
			&Node{Value: 7, red: false},
			&Node{Value: 8, red: true},
			&Node{Value: 10, red: true},
			&Node{Value: 11, red: false},
			&Node{Value: 15, red: false},
			&Node{Value: 18, red: false},
			&Node{Value: 22, red: true},
			&Node{Value: 26, red: false},
		}

		result := make([]*Node, 0)
		tree.ToSlice(&result)

		for i := 0; i < len(expect); i++ {
			if expect[i].Value != result[i].Value || expect[i].red != result[i].red {
				t.Fatalf("expected: %v, got: %v", expect[i], result[i])
			}
		}
	})

	t.Run("3 case insert B", func(t *testing.T) {
		t.Parallel()

		tree := NewTree(7)

		root := tree.Root
		root.Left = newNode(3, false, root)

		root.Right = newNode(18, true, root)

		root.Right.Left = newNode(10, false, root.Right)
		root.Right.Left.Left = newNode(8, true, root.Right.Left)
		root.Right.Left.Right = newNode(11, true, root.Right.Left)

		root.Right.Right = newNode(22, false, root.Right)
		root.Right.Right.Right = newNode(26, true, root.Right.Right)

		tree.Insert(15)

		expect := []*Node{
			&Node{Value: 3, red: false},
			&Node{Value: 7, red: true},
			&Node{Value: 8, red: false},
			&Node{Value: 10, red: false},
			&Node{Value: 11, red: false},
			&Node{Value: 15, red: true},
			&Node{Value: 18, red: true},
			&Node{Value: 22, red: false},
			&Node{Value: 26, red: true},
		}

		result := make([]*Node, 0)
		tree.ToSlice(&result)

		for i := 0; i < len(expect); i++ {
			if expect[i].Value != result[i].Value || expect[i].red != result[i].red {
				t.Fatalf("expected: %v, got: %v", expect[i], result[i])
			}
		}
	})
}

func TestContains(t *testing.T) {
	t.Run("1 to 7", func(t *testing.T) {
		t.Parallel()

		tree := NewTree(1)
		tree.Insert(2)
		tree.Insert(3)
		tree.Insert(4)
		tree.Insert(5)
		tree.Insert(6)
		tree.Insert(7)

		expect := []int{1, 2, 3, 4, 5, 6, 7}

		result := make([]*Node, 0)
		tree.ToSlice(&result)

		for i := 0; i < len(expect); i++ {
			if !tree.Contains(expect[i]) {
				t.Fatalf("failed to find: %d", expect[i])
			}
		}
	})
}
