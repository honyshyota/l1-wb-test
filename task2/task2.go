package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

func main() {
	var arr = []int{2, 4, 6, 8, 10}
	fmt.Println(arrayToSquare(arr))
}

func arrayToSquare(arr []int) []int {
	var wg sync.WaitGroup // инициализируем группу ожидания
	routinesAmount := len(arr)
	wg.Add(routinesAmount) // колличество рутин соответствует длине переданного массива
	var result []int       // инициализируем возвращаемый срез

	for _, val := range arr {
		go func(val int) { // явная передача аргумента
			defer wg.Done()                                         // отложенный вызов завершения ожидания
			result = append(result, int(math.Pow(float64(val), 2))) // возведение значения в квадрат
		}(val)
	}

	wg.Wait()

	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] }) // сортировка возращаемого результата с помощью встроенной библиотеки sort
	return result
}
