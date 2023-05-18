package slice

import "testing"

const slicesize = 10000

func BenchmarkSliceInitWithoutCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s1 := make([]int, 0)
		for i := 0; i < slicesize; i++ {
			s1 = append(s1, i)
		}
	}
}

func BenchmarkSliceInitWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s1 := make([]int, 0, slicesize)
		for i := 0; i < slicesize; i++ {
			s1 = append(s1, i)
		}
	}
}

// enter folder
// go test -benchmem -bench=. slice_benchmark_test.go
// or
//  go test -benchmem -bench=. ./intermediate/slice/slice_benchmark_test.go
