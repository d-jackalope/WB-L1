package main

import (
	"fmt"
	"os"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

func main() {
	fmt.Println("Введите количество секунд...")
	var N int
	fmt.Scan(&N)
	timer := time.NewTimer(time.Duration(N) * time.Second) // Создание таймера с заданным кол-вом секунд с ввода

	dataChan := make(chan int)      // создание канала передачи данных
	doneChan := make(chan struct{}) // создание канала для сигнала завершения работы горутины

	go worker(dataChan, doneChan) // создание горутины

	count, last := 0, 0
	for {
		select {
		case <-timer.C: // ожидание таймера
			fmt.Println("Время вышло!")
			close(dataChan) // закрытие канала данных
			<-doneChan      // ожидание сигнала завершения работы горутины
			fmt.Println("Горутина завершила работу")
			fmt.Println("Последнее значение из main: ", last)
			fmt.Println("Конец.")
			os.Exit(0) // завершение с кодом 0
		default:
			dataChan <- count //  отправка данных
			last = count      // последнее значение для сверки с горутиной
			count++           // меняем данные
		}
	}
}

func worker(dataChan <-chan int, doneChan chan struct{}) {
	last := 0
	for {
		data, ok := <-dataChan // ожидание данных из канала с проверкой на закрытие канала
		if !ok {               // если канал закрыт
			fmt.Println("Канал данных был закрыт по истечению времени!")
			fmt.Println("Последнее значение из горутины:", last)
			close(doneChan) // закрытие канала, тем самым сигнализируя о завершении работы горутины
			return          // выход из функции
		}
		last = data // последнее значение для сверки с main
	}
}

// Условный вывод:
// Введите количество секунд...
// 10
// Время вышло!
// Канал данных был закрыт по истечению времени!
// Последнее значение из горутины: 26790044
// Горутина завершила работу
// Последнее значение из main:  26790044
// Конец.
