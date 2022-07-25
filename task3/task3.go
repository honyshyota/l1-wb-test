package main

import (
	"fmt"
	"sync"
)

/*Дана последовательность чисел: 2,4,6,8,10. Найти сумму
их квадратов(22+32+42….) с использованием конкурентных вычислений.*/

func main() {
	var arr = [5]int{2, 4, 6, 8, 10}
	fmt.Println(arrayToSquare(arr))
}

func arrayToSquare(arr [5]int) int {
	var wg sync.WaitGroup // инициализируем группу ожидания
	amount := len(arr)
	wg.Add(amount) // колличество рутин соответствует длине переданного массива

	ch := make(chan int, amount) // инициализируем канал длиной равной длине массива

	for _, val := range arr {
		go func(val int) { // явная передача аргумента
			defer wg.Done() // отложенный вызов завершения ожидания
			ch <- val * val
		}(val)
	}

	wg.Wait()
	close(ch)                                  // после завершения ожидания очереди закрываем канал
	result := <-ch + <-ch + <-ch + <-ch + <-ch // считываем данные из канала и суммируем их
	return result
}
