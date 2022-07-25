package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
)

/*Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные
данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.*/

var workersAmount int

func init() {
	// колличество гоурутин принимается на вход из аргумента коммандной строки, дефолтное значение 10
	flag.IntVar(&workersAmount, "wa", 10, "workers amount")
}

func main() {
	flag.Parse() // парсим аргумент

	var wg sync.WaitGroup // создаем группу ожидания

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT) // контекст канал по сигналу ctrl+c

	ch := make(chan int) // основной канал в который пишет гоурутина

	go func() {
	LOOP1: // главная петля
		for i := 0; ; i++ {
			select {
			case <-ctx.Done(): // ожидание сигнала от канала контекст
				cancel()
				break LOOP1
			default: // запись данных в канал
				ch <- i
			}
		}
	}()

	for i := 0; workersAmount > i; i++ { // реализация колличества воркеров
		wg.Add(1)
		go func(i int) {
		LOOP2:
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					break LOOP2
				case a := <-ch: // чтение данных из канала
					fmt.Printf("Гоурутина %v, данные: %v\n", i, a)
				}
			}
			fmt.Printf("Гоурутина %v завершила работу\n", i)
		}(i)
	}

	wg.Wait() // ожидание гоурутин
	fmt.Println("конец")
}
