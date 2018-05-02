package trees

import (
	"testing"

	"github.com/freddygv/go-trees/redblack"
	"github.com/freddygv/go-trees/skiplist"
	"github.com/freddygv/go-trees/splay"
	"github.com/freddygv/go-trees/treap"
)

var input []int

// BenchmarkInsertSeq10 tests sequential insert of 0 to 9
func BenchmarkInsertSeq10(b *testing.B) {
	for i := 0; i < 10; i++ {
		input = append(input, i)
	}

	tt := []struct {
		desc  string
		bench func(*testing.B)
	}{
		{
			desc:  "splay",
			bench: insertSplay,
		},
		{
			desc:  "redblack",
			bench: insertRedBlack,
		},
		{
			desc:  "treap",
			bench: insertTreap,
		},
		{
			desc:  "skiplist",
			bench: insertSkip,
		},
	}

	for _, tc := range tt {
		b.Run(tc.desc, tc.bench)
	}
}

func insertSplay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := splay.NewTree()

		for _, v := range input {
			tree.Insert(v)
		}
	}
}

func insertRedBlack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := redblack.NewTree()

		for _, v := range input {
			tree.Insert(v)
		}
	}
}

func insertTreap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := treap.NewTree()

		for _, v := range input {
			tree.Insert(v)
		}
	}
}

func insertSkip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := skiplist.New()

		for _, v := range input {
			list.Insert(v)
		}
	}
}
