package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) // создаем канал

	go func() {
		for val := range ch { // получаем данные из канала
			fmt.Println(val)
		}

		close(ch) // когда данные закончились закрываем канал встроенной функцией close
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second) // для наглядности ждем 1 секунду перед каждой итерацией
		ch <- "гоурутина работает"  // передаем данные в канал
	}

	fmt.Println("функция завршает работу")
}
