package main

import (
	"math/rand"
	"testing"
)

func BenchmarkRemoveWithAppend(b *testing.B) {
	slice := make([]int, 20000000)
	for i := 0; i < 20000000; i++ {
		slice[i] = i
	}

	var index int

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index = rand.Intn(20000000)
		removeWithAppend(slice, index)
	}

}

func BenchmarkRemoveWithCopy(b *testing.B) {
	slice := make([]int, 20000000)
	for i := 0; i < 20000000; i++ {
		slice[i] = i
	}

	var index int

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index = rand.Intn(20000000)
		removeWithCopy(slice, index)
	}
}

func BenchmarkRemoveWithLastElement(b *testing.B) {
	slice := make([]int, 20000000)
	for i := 0; i < 20000000; i++ {
		slice[i] = i
	}

	var index int

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index = rand.Intn(20000000)
		removeWithLastElement(slice, index)
	}
}

// benchtime=1000x
// cpu: Intel(R) Core(TM) i5-4250U CPU @ 1.30GHz
// BenchmarkRemoveWithAppend-4                 1000          23149561 ns/op               0 B/op          0 allocs/op
// BenchmarkRemoveWithCopy-4                   1000          10638923 ns/op               0 B/op          0 allocs/op
// BenchmarkRemoveWithLastElement-4            1000               106.1 ns/op             0 B/op          0 allocs/op
