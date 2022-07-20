package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) // создаем канал
	defer close(ch)         // закрываем канал при выходе из функции

	go func() {
		for {
			select {
			case a := <-ch: // гоурутина читает из канала
				fmt.Println(a)
				// когда данные в канале закончились гоурутина передает пустую структуру в канал и закрывается
			case ch <- "гоурутина передает данные в канал и завершает работу":
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second) // для наглядности ждем 1 секунду перед каждой итерацией
		ch <- "гоурутина работает"  // передаем данные в канал
	}

	fmt.Println(<-ch) // ожидаем данных из канала в основной гоурутине
	fmt.Println("функция завршает работу")
}
