package treap

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
			input:  []int{5, 5, 5, 5, 5, 5},
			expect: []int{5, 5, 5, 5, 5, 5},
		},
		{
			desc:   "mix",
			input:  []int{7, 6, 5, 4, 3, 2, 1, 1, 2, 3, 4, 5, 6, 7},
			expect: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7},
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
					t.Fatalf("failed to get: %d", v)
				}
			}
		})
	}
}
