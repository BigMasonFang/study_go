package _map

import "testing"

const mapSize = 10000

func BenchmarkMapInitWithoutCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make(map[int]int)
		for i := 0; i < mapSize; i++ {
			m[i] = i
		}
	}
}

func BenchmarkMapInitWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make(map[int]int, mapSize)
		for i := 0; i < mapSize; i++ {
			m[i] = i
		}
	}
}

// $ go test -benchmem -bench=. intermediate/map/map_benchamark_test.go
