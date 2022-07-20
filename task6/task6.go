package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // создаем группу ожидания

	ch := make(chan string) // создаем канал
	defer close(ch)         // при выходе из фукнции закрываем его
	wg.Add(1)               // добавляем очередь в wg

	go func() {
		for {
			select {
			case a := <-ch: // ждем данные из канала, при получении закрываем его
				fmt.Println(a)
				wg.Done()

				return
			default: // пока не получены данные из канала гоурутина работает
				fmt.Println("гоурутина работает")
			}
		}
	}()

	time.Sleep(2 * time.Second)                // главная гоурутина ждет 2 секунды
	ch <- "останавливаем выполнение гоурутины" // после чего передаем данные в канал и гоурутина закрывается
	wg.Wait()                                  // ожидание очереди
	fmt.Println("завершение функции")
}
