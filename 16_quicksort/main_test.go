package main

import (
	"math/rand"
	"sort"
	"testing"
)

func BenchmarkQuickSort(b *testing.B) {

	data := make([]int, 1000000)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(10000)
	}

	b.ResetTimer()

	b.Run("quickSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slice := make([]int, len(data))
			copy(slice, data)
			quickSort(slice, 0, len(slice)-1)
		}
	})

	b.Run("quickSort_2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slice := make([]int, len(data))
			copy(slice, data)
			quickSort_2(slice)
		}
	})

	b.Run("sort.Ints", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slice := make([]int, len(data))
			copy(slice, data)
			sort.Ints(slice)
		}
	})

}

// benchtime=10x
// cpu: Intel(R) Core(TM) i5-4250U CPU @ 1.30GHz
// BenchmarkQuickSort/quickSort-4                10         122076663 ns/op         8003593 B/op          1 allocs/op
// BenchmarkQuickSort/quickSort_2-4              10        1380974570 ns/op        1919205616 B/op  6844813 allocs/op
// BenchmarkQuickSort/sort.Ints-4                10         157580699 ns/op         8003608 B/op          2 allocs/op
// PASS
// ok      github.com/d-jackalope/WB-L1/task-16    18.174s
