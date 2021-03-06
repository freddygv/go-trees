package splay

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tt := []struct {
		desc   string
		input  []int
		expect []int
	}{
		{
			desc:   "1 to 7",
			input:  []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			desc:   "7 to 1",
			input:  []int{7, 6, 5, 4, 3, 2, 1},
			expect: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			desc:   "duplicated",
			input:  []int{5, 5, 5, 5, 5, 5, 5, 5},
			expect: []int{5, 5, 5, 5, 5, 5, 5, 5},
		},
		{
			desc:   "random",
			input:  []int{1298498081, 2019727887, 1427131847, 939984059, 911902081, 1474941318, 140954425},
			expect: []int{140954425, 911902081, 939984059, 1298498081, 1427131847, 1474941318, 2019727887},
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			tree := NewTree()
			for _, v := range tc.input {
				tree.Insert(v)
			}

			result := tree.toSlice()
			for i, v := range result {
				if v.Value != tc.expect[i] {
					t.Fatalf("expected: '%v', got '%v'", tc.input[i], v.Value)
				}
			}
		})
	}
}

func TestGet(t *testing.T) {
	tt := []struct {
		desc  string
		input []int
	}{
		{
			desc:  "1 to 7",
			input: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			desc:  "7 to 1",
			input: []int{7, 6, 5, 4, 3, 2, 1},
		},
		{
			desc:  "random",
			input: []int{1298498081, 2019727887, 1427131847, 939984059, 911902081, 1474941318, 140954425},
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			tree := NewTree()
			for _, v := range tc.input {
				tree.Insert(v)
			}

			for _, v := range tc.input {
				if _, ok := tree.Get(v); !ok {
					t.Fatalf("failed to get: %v", v)
				}
			}
		})

	}

	t.Run("1 to 7 and two gets", func(t *testing.T) {
		t.Parallel()

		input := []int{1, 2, 3, 4, 5, 6, 7}

		tree := NewTree()
		for _, v := range input {
			tree.Insert(v)
		}

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
