package main

import (
	"fmt"
	"time"
)

/*Реализовать собственную функцию sleep.*/

func sleep(t time.Duration) {
	time := time.NewTimer(t) // создаем таймер и отправляем канал который не дает основной гоурутине завершиться
	<-time.C
}

func main() {
	sleep(4 * time.Second)
	fmt.Println("Готово")
}
