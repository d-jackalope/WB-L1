package main

import (
	"math/rand"
	"strings"
	"testing"
)

func BenchmarkReverseString(b *testing.B) {
	chars := []rune("йцукенгшщзфывапролдячсмить") // все символы состоят из 2-х байт
	var build strings.Builder
	for i := 0; i < 10000; i++ {
		build.WriteRune(chars[rand.Intn(len(chars))])
	}

	str := build.String()
	runeStr1 := []rune(str)
	runeStr2 := []rune(str)

	b.ResetTimer()

	b.Run("reverseString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reverseString(str)
		}
	})

	b.Run("reverseRune", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reverseRune(runeStr1)
		}
	})

	b.Run("sortRunes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sortRunes(runeStr2)
		}
	})

}

// cpu: Intel(R) Core(TM) i5-4250U CPU @ 1.30GHz
// BenchmarkReverseString/reverseString-4               100          36934874 ns/op        135127134 B/op     10006 allocs/op
// BenchmarkReverseString/reverseRune-4                 100              9466 ns/op               0 B/op          0 allocs/op
// BenchmarkReverseString/sortRunes-4                   100           1250394 ns/op              56 B/op          2 allocs/op
// PASS
// ok      github.com/d-jackalope/WB-L1/task-19    3.865s
