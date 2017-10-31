package linearsort_test

import (
	"sort"
	"testing"

	"github.com/mrap/linearsort"
)

func BenchmarkLibSort1000(b *testing.B)        { benchmarkLibSort(b, 1000) }
func BenchmarkLinearSort1000(b *testing.B)     { benchmarkLinearSort(b, 1000) }
func BenchmarkLibSort10000(b *testing.B)       { benchmarkLibSort(b, 10000) }
func BenchmarkLinearSort10000(b *testing.B)    { benchmarkLinearSort(b, 10000) }
func BenchmarkLibSort100000(b *testing.B)      { benchmarkLibSort(b, 100000) }
func BenchmarkLinearSort100000(b *testing.B)   { benchmarkLinearSort(b, 100000) }
func BenchmarkLibSort1000000(b *testing.B)     { benchmarkLibSort(b, 1000000) }
func BenchmarkLinearSort1000000(b *testing.B)  { benchmarkLinearSort(b, 1000000) }
func BenchmarkLibSort10000000(b *testing.B)    { benchmarkLibSort(b, 10000000) }
func BenchmarkLinearSort10000000(b *testing.B) { benchmarkLinearSort(b, 10000000) }

func benchmarkLibSort(b *testing.B, n int) {
	desc := intsDesc(n)
	a := make([]int, n)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(a, desc)
		b.StartTimer()
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	}
}

func benchmarkLinearSort(b *testing.B, n int) {
	desc := intsDesc(n)
	a := make([]int, n)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(a, desc)
		b.StartTimer()
		linearsort.Sort(&a)
	}
}

func intsDesc(n int) []int {
	a := make([]int, n)
	for i := 0; i < len(a); i++ {
		a[i] = len(a) - i
	}
	return a
}
