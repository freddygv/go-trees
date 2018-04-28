package main

import "testing"

func TestInsert(t *testing.T) {
	t.Run("1 to 7", func(t *testing.T) {
		tree := NewTree(1)
		tree.Insert(2)
		tree.Insert(3)
		tree.Insert(4)
		tree.Insert(5)
		tree.Insert(6)
		tree.Insert(7)

		expect := []*Node{
			&Node{value: 1, red: false},
			&Node{value: 2, red: false},
			&Node{value: 3, red: false},
			&Node{value: 4, red: true},
			&Node{value: 5, red: true},
			&Node{value: 6, red: false},
			&Node{value: 7, red: true},
		}

		result := make([]*Node, 0)
		flatten(tree.root, &result)

		for i := 0; i < len(expect); i++ {
			if expect[i].value != result[i].value || expect[i].red != result[i].red {
				t.Fatalf("expected: %v, got: %v", expect[i], result[i])
			}
		}
	})

	t.Run("3 case insert", func(t *testing.T) {
		tree := NewTree(7)

		root := tree.root
		root.left = newNode(3, false, root)
		root.left.red = false

		root.right = newNode(18, true, root)

		root.right.left = newNode(10, false, root.right)
		root.right.left.red = false
		root.right.left.left = newNode(8, true, root.right.left)
		root.right.left.right = newNode(11, true, root.right.left)

		root.right.right = newNode(22, false, root.right)
		root.right.right.red = false
		root.right.right.right = newNode(26, true, root.right.right)

		tree.Insert(15)

		expect := []*Node{
			&Node{value: 3, red: false},
			&Node{value: 7, red: true},
			&Node{value: 8, red: false},
			&Node{value: 10, red: false},
			&Node{value: 11, red: false},
			&Node{value: 15, red: true},
			&Node{value: 18, red: true},
			&Node{value: 22, red: false},
			&Node{value: 26, red: true},
		}

		result := make([]*Node, 0)
		flatten(tree.root, &result)

		for i := 0; i < len(expect); i++ {
			if expect[i].value != result[i].value || expect[i].red != result[i].red {
				t.Fatalf("expected: %v, got: %v", expect[i], result[i])
			}
		}
	})
}

func TestContains(t *testing.T) {
	t.Run("1 to 7", func(t *testing.T) {
		tree := NewTree(1)
		tree.Insert(2)
		tree.Insert(3)
		tree.Insert(4)
		tree.Insert(5)
		tree.Insert(6)
		tree.Insert(7)

		expect := []int{1, 2, 3, 4, 5, 6, 7}

		result := make([]*Node, 0)
		flatten(tree.root, &result)

		for i := 0; i < len(expect); i++ {
			if !tree.Contains(expect[i]) {
				t.Fatalf("failed to find: %d", expect[i])
			}
		}
	})
}
