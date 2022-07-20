package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	var arr = [5]int{2, 4, 6, 8, 10}
	fmt.Println(arrayToSquare(arr))
	fmt.Println(time.Since(t))
}

// создаем структуру с ээлементом блокировки
type result struct {
	result int
	sync.Mutex
}

func arrayToSquare(arr [5]int) int {
	var result result     // инициализируем переменную типа result
	var wg sync.WaitGroup // инициализируем группу ожидания
	amount := len(arr)
	wg.Add(amount) // колличество рутин соответствует длине переданного массива

	for _, val := range arr {
		go func(val int) { // явная передача аргумента
			result.Lock()   // блокируем запись в перменную
			defer wg.Done() // отложенный вызов завершения ожидания
			result.result += val * val
			result.Unlock() // разблокируем запись в переменную
		}(val)
	}

	wg.Wait()

	return result.result
}
