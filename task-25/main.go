package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func main() {
	var s int
	fmt.Println("Введите количество секунд...")
	fmt.Scan(&s)

	d := time.Second * time.Duration(s)
	customSleepWithTimer(d)
	customSleepWithAfter(d)

	fmt.Println("Программа проснулась")
}

func customSleepWithTimer(d time.Duration) {
	timer := time.NewTimer(d) // создание нового таймера
	<-timer.C                 // ожидание сигнала завершения таймера, возвращает текущее время
	fmt.Printf("timer: Прошло %v\n", d)
}

func customSleepWithAfter(d time.Duration) {
	<-time.After(d) // time.After возвращает канал, который по истечению времени возвращает текущее время
	fmt.Printf("after: Прошло %v\n", d)
}
