package session3

import (
	"math/rand"
	"testing"
)

const benchN, benchQ = 100_000_000, 1_000

func benchData() ([]int, [][2]int) {
	r := rand.New(rand.NewSource(1))
	nums := make([]int, benchN)
	for i := range nums {
		nums[i] = r.Intn(1000)
	}
	queries := make([][2]int, benchQ)
	for i := range queries {
		l, rr := r.Intn(benchN), r.Intn(benchN)
		queries[i] = [2]int{min(l, rr), max(l, rr)}
	}
	return nums, queries
}

func BenchmarkSumRangeNaive(b *testing.B) {
	nums, queries := benchData()
	b.ResetTimer()
	for range b.N {
		for _, q := range queries {
			SumRangeNaive(nums, q[0], q[1])
		}
	}
}

func BenchmarkSumRange(b *testing.B) {
	nums, queries := benchData()
	b.ResetTimer()
	for range b.N {
		numArr := Constructor(nums)
		for _, q := range queries {
			numArr.SumRange(q[0], q[1])
		}
	}
}
