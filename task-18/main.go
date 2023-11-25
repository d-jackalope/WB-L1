package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

type count struct {
	mu sync.Mutex // мютекс для безопасного доступа к полям структуры
	v  int        // счетчик
}

func (c *count) increment() {
	c.mu.Lock() // блокировка доступа
	c.v++
	c.mu.Unlock()
}

func main() {

	var c count            // инициализация структуры
	wg := sync.WaitGroup{} // создание waitGroup для того, чтобы дождаться все горутины
	workers := 10000       // количество горутин, которые будут инкрементировать
	for i := 0; i < workers; i++ {
		wg.Add(1) // добавление каждой горутины в группу
		go func(wg *sync.WaitGroup, c *count) {
			c.increment() // прибавление единицы
			wg.Done()     // удаление горутины из группы по завершению работы
		}(&wg, &c) // передаем структуру и группу по указателю
	}

	wg.Wait() // ожидание всех горутин

	fmt.Printf("workers: %v\ncount struct: %v\n", workers, c.v) // сверка данных

	// вывод:
	// workers: 10000
	// count struct: 10000
}
