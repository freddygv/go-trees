package trees

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/freddygv/go-trees/redblack"
	"github.com/freddygv/go-trees/skiplist"
	"github.com/freddygv/go-trees/splay"
	"github.com/freddygv/go-trees/treap"
)

var (
	inputs []int
	reads  []int
)

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

var readTable = []struct {
	desc  string
	bench func(*testing.B)
}{
	{
		desc:  "splay",
		bench: readSplay,
	},
	{
		desc:  "redblack",
		bench: readRedBlack,
	},
	{
		desc:  "treap",
		bench: readTreap,
	},
	{
		desc:  "skiplist",
		bench: readSkip,
	},
}

// BenchmarkInsertSeqN tests sequential insert from 0 to N
func BenchmarkInsertSeq10(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 10; i++ {
		inputs = append(inputs, i)
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
		inputs = append(inputs, r)
	}

	for _, tc := range insertTable {
		b.StartTimer()
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadSeqN tests sequential reads from 0 to N
func BenchmarkReadSeq10(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 10; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
	}

	copy(reads, inputs)
	sort.Ints(reads)

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRandN tests N random reads
func BenchmarkReadRand10(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 10; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
	}

	copy(reads, inputs)
	rand.Shuffle(len(reads), func(i, j int) {
		reads[i], reads[j] = reads[j], reads[i]
	})

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRepeatedN tests N repeated reads
func BenchmarkReadRepeat10(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 10; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, inputs[0])
	}

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

func BenchmarkInsertSeq100(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 100; i++ {
		inputs = append(inputs, i)
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
		inputs = append(inputs, r)
	}

	for _, tc := range insertTable {
		b.StartTimer()
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadSeqN tests sequential reads from 0 to N
func BenchmarkReadSeq100(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 100; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
	}

	copy(reads, inputs)
	sort.Ints(reads)

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRandN tests N random reads
func BenchmarkReadRand100(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 100; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
	}

	copy(reads, inputs)
	rand.Shuffle(len(reads), func(i, j int) {
		reads[i], reads[j] = reads[j], reads[i]
	})

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRepeatedN tests N repeated reads
func BenchmarkReadRepeat100(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 100; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, inputs[0])
	}

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

func BenchmarkInsertSeq1000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 1000; i++ {
		inputs = append(inputs, i)
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
		inputs = append(inputs, r)
	}

	for _, tc := range insertTable {
		b.StartTimer()
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadSeqN tests sequential reads from 0 to N
func BenchmarkReadSeq1000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 1000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
	}

	copy(reads, inputs)
	sort.Ints(reads)

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRandN tests N random reads
func BenchmarkReadRand1000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 1000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
	}

	copy(reads, inputs)
	rand.Shuffle(len(reads), func(i, j int) {
		reads[i], reads[j] = reads[j], reads[i]
	})

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRepeatedN tests N repeated reads
func BenchmarkReadRepeat1000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 1000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, inputs[0])
	}

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

func BenchmarkInsertSeq10000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		inputs = append(inputs, i)
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
		inputs = append(inputs, r)
	}

	for _, tc := range insertTable {
		b.StartTimer()
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadSeqN tests sequential reads from 0 to N
func BenchmarkReadSeq10000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
	}

	copy(reads, inputs)
	sort.Ints(reads)

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRandN tests N random reads
func BenchmarkReadRand10000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
	}

	copy(reads, inputs)
	rand.Shuffle(len(reads), func(i, j int) {
		reads[i], reads[j] = reads[j], reads[i]
	})

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRepeatedN tests N repeated reads
func BenchmarkReadRepeat10000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, inputs[0])
	}

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

func BenchmarkInsertSeq100000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 100000; i++ {
		inputs = append(inputs, i)
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
		inputs = append(inputs, r)
	}

	for _, tc := range insertTable {
		b.StartTimer()
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadSeqN tests sequential reads from 0 to N
func BenchmarkReadSeq100000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 100000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
	}

	copy(reads, inputs)
	sort.Ints(reads)

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRandN tests N random reads
func BenchmarkReadRand100000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 100000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
	}

	copy(reads, inputs)
	rand.Shuffle(len(reads), func(i, j int) {
		reads[i], reads[j] = reads[j], reads[i]
	})

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRepeatedN tests N repeated reads
func BenchmarkReadRepeat100000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 100000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, inputs[0])
	}

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

func insertSplay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := splay.NewTree()

		for _, v := range inputs {
			tree.Insert(v)
		}
	}
}

func readSplay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := splay.NewTree()

		for _, v := range inputs {
			tree.Insert(v)
		}

		b.StartTimer()
		for _, v := range reads {
			tree.Get(v)
		}
	}
}

func insertRedBlack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := redblack.NewTree()

		for _, v := range inputs {
			tree.Insert(v)
		}
	}
}

func readRedBlack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := redblack.NewTree()

		for _, v := range inputs {
			tree.Insert(v)
		}

		b.StartTimer()
		for _, v := range reads {
			tree.Get(v)
		}
	}
}

func insertTreap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := treap.NewTree()

		for _, v := range inputs {
			tree.Insert(v)
		}
	}
}

func readTreap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := treap.NewTree()

		for _, v := range inputs {
			tree.Insert(v)
		}

		b.StartTimer()
		for _, v := range reads {
			tree.Get(v)
		}
	}
}

func insertSkip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := skiplist.New()

		for _, v := range inputs {
			list.Insert(v)
		}
	}
}

func readSkip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := skiplist.New()

		for _, v := range inputs {
			list.Insert(v)
		}

		b.StartTimer()
		for _, v := range reads {
			list.Get(v)
		}
	}
}
