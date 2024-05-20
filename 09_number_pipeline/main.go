package main

import (
	"fmt"
)

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {
	arr := []int{}
	for num := 1; num <= 10000; num++ {
		arr = append(arr, num)
	}

	xChan := make(chan int)
	x2Chan := make(chan int)

	// горутина, которая отправляет значения из массива в канал
	go func(arr []int, xChan chan int) {
		for _, x := range arr {
			xChan <- x
		}
		close(xChan) // закрытие канала после отправки всех чисел, чтобы сигнализировать о том, что чисел больше не будет
	}(arr, xChan)

	// горутина принимает значения из канала с данными и пересылает удвоенное значение в другой канал
	go func(xChan chan int, x2Chan chan int) {
		for x := range xChan {
			x2Chan <- x * 2
		}
		close(x2Chan) // закрытие канала после отправки всех обработанных данных
	}(xChan, x2Chan)

	// прием данных из канала завершится, когда будет послан сигнал о его закрытии
	for v := range x2Chan {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	fmt.Println("Конец.")

}
