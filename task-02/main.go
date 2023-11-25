package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	arr := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup // waitGroup для того, чтобы дождаться все горутины
	for _, v := range arr {
		wg.Add(1) // добавление в группу
		go func(wg *sync.WaitGroup, v int) {
			fmt.Println(v * v)
			wg.Done() // удаление из группы
		}(&wg, v)

	}

	wg.Wait() // ожидание всех горутин
}
