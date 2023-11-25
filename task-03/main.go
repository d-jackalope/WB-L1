package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

func main() {
	arr := []int{2, 4, 6, 8, 10}
	numChan := make(chan int, len(arr))

	var wg sync.WaitGroup
	for _, v := range arr {
		wg.Add(1)
		go func(wg *sync.WaitGroup, v int, numChan chan int) {
			numChan <- v * v
			wg.Done()
		}(&wg, v, numChan)
	}

	wg.Wait()
	close(numChan)
	sum := 0
	for square := range numChan {
		sum += square
	}
	fmt.Println(sum)
}
