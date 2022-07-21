package main

import (
	"fmt"
	"time"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10} // инициализируем массив с переменными

	writeCh := write(arr) // передаем массив в пишущую в канал функцию и на выход получаем канал

	readCh := read(writeCh) // передаем канал выше в слушающую этот канал функцию и пишущую произведение в другой канал

	for val := range readCh { 
		fmt.Println(val) // итерируемся по каналу и выводим значения в stdout
	}
}

func write(arr [5]int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, val := range arr {
			time.Sleep(1 * time.Second) // ждем для наглядности
			ch <- val
		}
		close(ch)
	}()

	return ch
}

func read(inCh <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		for val := range inCh {
			ch <- val * 2
		}
		close(ch)
	}()

	return ch
}
