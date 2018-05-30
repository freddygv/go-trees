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

// BenchmarkInsertSequentialN tests sequential insert from 0 to N
func BenchmarkInsertSequential10(b *testing.B) {
	for i := 0; i < 10; i++ {
		inputs = append(inputs, i)
	}

	for _, tc := range insertTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkInsertRandomN tests N random int inserts
func BenchmarkInsertRandom10(b *testing.B) {
	for i := 0; i < 10; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
	}

	for _, tc := range insertTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkInsertRepeatededN tests N repeated inserts
func BenchmarkInsertRepeated10(b *testing.B) {
	r := rand.Int()

	for i := 0; i < 10; i++ {
		inputs = append(inputs, r)
	}

	for _, tc := range insertTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadSequentialN tests sequential reads from 0 to N
func BenchmarkReadSequential10(b *testing.B) {
	for i := 0; i < 10; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, r)
	}

	sort.Ints(reads)

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRandomN tests N random reads
func BenchmarkReadRandom10(b *testing.B) {
	for i := 0; i < 10; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, r)
	}

	rand.Shuffle(len(reads), func(i, j int) {
		reads[i], reads[j] = reads[j], reads[i]
	})

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRepeatededN tests N repeated reads
func BenchmarkReadRepeated10(b *testing.B) {
	for i := 0; i < 10; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, inputs[0])
	}

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkInsertSequentialN tests sequential insert from 0 to N
func BenchmarkInsertSequential100(b *testing.B) {
	for i := 0; i < 100; i++ {
		inputs = append(inputs, i)
	}

	for _, tc := range insertTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkInsertRandomN tests N random int inserts
func BenchmarkInsertRandom100(b *testing.B) {
	for i := 0; i < 100; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
	}

	for _, tc := range insertTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkInsertRepeatededN tests N repeated inserts
func BenchmarkInsertRepeated100(b *testing.B) {
	r := rand.Int()

	for i := 0; i < 100; i++ {
		inputs = append(inputs, r)
	}

	for _, tc := range insertTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadSequentialN tests sequential reads from 0 to N
func BenchmarkReadSequential100(b *testing.B) {
	for i := 0; i < 100; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, r)
	}

	sort.Ints(reads)

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRandomN tests N random reads
func BenchmarkReadRandom100(b *testing.B) {
	for i := 0; i < 100; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, r)
	}

	rand.Shuffle(len(reads), func(i, j int) {
		reads[i], reads[j] = reads[j], reads[i]
	})

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRepeatededN tests N repeated reads
func BenchmarkReadRepeated100(b *testing.B) {
	for i := 0; i < 100; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, inputs[0])
	}

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkInsertSequentialN tests sequential insert from 0 to N
func BenchmarkInsertSequential1000(b *testing.B) {
	for i := 0; i < 1000; i++ {
		inputs = append(inputs, i)
	}

	for _, tc := range insertTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkInsertRandomN tests N random int inserts
func BenchmarkInsertRandom1000(b *testing.B) {
	for i := 0; i < 1000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
	}

	for _, tc := range insertTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkInsertRepeatededN tests N repeated inserts
func BenchmarkInsertRepeated1000(b *testing.B) {
	r := rand.Int()

	for i := 0; i < 1000; i++ {
		inputs = append(inputs, r)
	}

	for _, tc := range insertTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadSequentialN tests sequential reads from 0 to N
func BenchmarkReadSequential1000(b *testing.B) {
	for i := 0; i < 1000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, r)
	}

	sort.Ints(reads)

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRandomN tests N random reads
func BenchmarkReadRandom1000(b *testing.B) {
	for i := 0; i < 1000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, r)
	}

	rand.Shuffle(len(reads), func(i, j int) {
		reads[i], reads[j] = reads[j], reads[i]
	})

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRepeatededN tests N repeated reads
func BenchmarkReadRepeated1000(b *testing.B) {
	for i := 0; i < 1000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, inputs[0])
	}

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkInsertSequentialN tests sequential insert from 0 to N
func BenchmarkInsertSequential10000(b *testing.B) {
	for i := 0; i < 10000; i++ {
		inputs = append(inputs, i)
	}

	for _, tc := range insertTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkInsertRandomN tests N random int inserts
func BenchmarkInsertRandom10000(b *testing.B) {
	for i := 0; i < 10000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
	}

	for _, tc := range insertTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkInsertRepeatededN tests N repeated inserts
func BenchmarkInsertRepeated10000(b *testing.B) {
	r := rand.Int()

	for i := 0; i < 10000; i++ {
		inputs = append(inputs, r)
	}

	for _, tc := range insertTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadSequentialN tests sequential reads from 0 to N
func BenchmarkReadSequential10000(b *testing.B) {
	for i := 0; i < 10000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, r)
	}

	sort.Ints(reads)

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRandomN tests N random reads
func BenchmarkReadRandom10000(b *testing.B) {
	for i := 0; i < 10000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, r)
	}

	rand.Shuffle(len(reads), func(i, j int) {
		reads[i], reads[j] = reads[j], reads[i]
	})

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRepeatededN tests N repeated reads
func BenchmarkReadRepeated10000(b *testing.B) {
	for i := 0; i < 10000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, inputs[0])
	}

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkInsertSequentialN tests sequential insert from 0 to N
func BenchmarkInsertSequential100000(b *testing.B) {
	for i := 0; i < 100000; i++ {
		inputs = append(inputs, i)
	}

	for _, tc := range insertTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkInsertRandomN tests N random int inserts
func BenchmarkInsertRandom100000(b *testing.B) {
	for i := 0; i < 100000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
	}

	for _, tc := range insertTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkInsertRepeatededN tests N repeated inserts
func BenchmarkInsertRepeated100000(b *testing.B) {
	r := rand.Int()

	for i := 0; i < 100000; i++ {
		inputs = append(inputs, r)
	}

	for _, tc := range insertTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadSequentialN tests sequential reads from 0 to N
func BenchmarkReadSequential100000(b *testing.B) {
	for i := 0; i < 100000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, r)
	}

	sort.Ints(reads)

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRandomN tests N random reads
func BenchmarkReadRandom100000(b *testing.B) {
	for i := 0; i < 100000; i++ {
		r := rand.Int()
		inputs = append(inputs, r)
		reads = append(reads, r)
	}

	rand.Shuffle(len(reads), func(i, j int) {
		reads[i], reads[j] = reads[j], reads[i]
	})

	for _, tc := range readTable {
		b.Run(tc.desc, tc.bench)
	}
}

// BenchmarkReadRepeatededN tests N repeated reads
func BenchmarkReadRepeated100000(b *testing.B) {
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
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree := splay.NewTree()

		for _, v := range inputs {
			tree.Insert(v)
		}
	}
}

func readSplay(b *testing.B) {
	tree := splay.NewTree()

	for _, v := range inputs {
		tree.Insert(v)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range reads {
			tree.Get(v)
		}
	}
}

func insertRedBlack(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree := redblack.NewTree()

		for _, v := range inputs {
			tree.Insert(v)
		}
	}
}

func readRedBlack(b *testing.B) {
	tree := redblack.NewTree()

	for _, v := range inputs {
		tree.Insert(v)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range reads {
			tree.Get(v)
		}
	}
}

func insertTreap(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree := treap.NewTree()

		for _, v := range inputs {
			tree.Insert(v)
		}
	}
}

func readTreap(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()

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
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list := skiplist.New()

		for _, v := range inputs {
			list.Insert(v)
		}
	}
}

func readSkip(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()

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
