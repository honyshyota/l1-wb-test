package main

import (
	"context"
	"flag"
	"fmt"
	"time"
)

/*Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.*/

var duration time.Duration

func init() {
	flag.DurationVar(&duration, "t", 20*time.Second, "time out")
}

func main() {
	flag.Parse()                                                       // парсим аргумент из коммандной строки -t
	ctx, cancel := context.WithTimeout(context.Background(), duration) // контекст с таймаутом
	defer cancel()

	ch := make(chan int)
	defer close(ch)

	go func() {
		for {
			select {
			case a := <-ch:
				fmt.Println(a)
			case <-ctx.Done(): // ждем данные из канала контекста и завершаем работу функции
				fmt.Println("Время вышло")
				return
			}
		}
	}()

LOOP1:
	for i := 0; ; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done(): // ждем данные из канала контекст и завершаем основную петлю
			break LOOP1
		case ch <- i:
		}
	}

}
