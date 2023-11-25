package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	ctx, cansel := context.WithCancel(context.Background())
	doneChan := make(chan struct{})

	go stopWithContext(ctx)
	go stopWithChannel(doneChan)

	cansel()        // отмена контекста
	close(doneChan) // закрытие канала завершения

	time.Sleep(time.Second * 2) // ожидание завершающихся горутин

	fmt.Println("Программа завершила работу")

}

func stopWithContext(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("Завершение с контекстом")
}

func stopWithChannel(doneChan chan struct{}) {
	<-doneChan
	fmt.Println("Завершение с каналом")
}
