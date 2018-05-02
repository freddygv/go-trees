package trees

import (
	"math/rand"
	"testing"

	"github.com/freddygv/go-trees/redblack"
	"github.com/freddygv/go-trees/skiplist"
	"github.com/freddygv/go-trees/splay"
	"github.com/freddygv/go-trees/treap"
)

var input []int

var insertTable = []struct {
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

// BenchmarkInsertSeqN tests sequential insert from 0 to N
func BenchmarkInsertSeq10(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 10; i++ {
		input = append(input, i)
	}

	for _, tc := range insertTable {
		b.StartTimer()
		b.Run(tc.desc, tc.bench)
	}
}

func BenchmarkInsertSeq100(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 100; i++ {
		input = append(input, i)
	}

	for _, tc := range insertTable {
		b.StartTimer()
		b.Run(tc.desc, tc.bench)
	}
}

func BenchmarkInsertSeq1000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 1000; i++ {
		input = append(input, i)
	}

	for _, tc := range insertTable {
		b.StartTimer()
		b.Run(tc.desc, tc.bench)
	}
}

func BenchmarkInsertSeq10000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		input = append(input, i)
	}

	for _, tc := range insertTable {
		b.StartTimer()
		b.Run(tc.desc, tc.bench)
	}
}

func BenchmarkInsertSeq100000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 100000; i++ {
		input = append(input, i)
	}

	for _, tc := range insertTable {
		b.StartTimer()
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkInsertRandN tests N random int inserts
func BenchmarkInsertRand10(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 10; i++ {
		r := rand.Int()
		input = append(input, r)
	}

	for _, tc := range insertTable {
		b.StartTimer()
		b.Run(tc.desc, tc.bench)
	}
}

func BenchmarkInsertRand100(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 100; i++ {
		r := rand.Int()
		input = append(input, r)
	}

	for _, tc := range insertTable {
		b.StartTimer()
		b.Run(tc.desc, tc.bench)
	}
}

func BenchmarkInsertRand1000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 1000; i++ {
		r := rand.Int()
		input = append(input, r)
	}

	for _, tc := range insertTable {
		b.StartTimer()
		b.Run(tc.desc, tc.bench)
	}
}

func BenchmarkInsertRand10000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		r := rand.Int()
		input = append(input, r)
	}

	for _, tc := range insertTable {
		b.StartTimer()
		b.Run(tc.desc, tc.bench)
	}
}

func BenchmarkInsertRand100000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 100000; i++ {
		r := rand.Int()
		input = append(input, r)
	}

	for _, tc := range insertTable {
		b.StartTimer()
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
