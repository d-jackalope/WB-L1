package main

import (
	"math/rand"
	"testing"
)

func BenchmarkBinarySearch(b *testing.B) {

	data := make([]int, 20000000)
	for i := 0; i < 20000000; i++ {
		data[i] = i
	}

	b.ResetTimer()

	b.Run("binarySearch_iter value=random", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value := rand.Intn(1000000)
			_, _ = binarySearch_iter(data, value)
		}
	})

	b.Run("binarySearch_rec value=random", func(b *testing.B) {
		low, high := 0, len(data)-1
		for i := 0; i < b.N; i++ {
			value := rand.Intn(1000000)
			_, _ = binarySearch_rec(data, value, low, high)
		}
	})

}

// cpu: Intel(R) Core(TM) i5-4250U CPU @ 1.30GHz
// BenchmarkBinarySearch/binarySearch_iter_value=random-4            100000               479.8 ns/op             0 B/op          0 allocs/op
// BenchmarkBinarySearch/binarySearch_rec_value=random-4             100000               644.4 ns/op             0 B/op          0 allocs/op
// PASS
// ok      github.com/d-jackalope/WB-L1/task-17    0.248s
