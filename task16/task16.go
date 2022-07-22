package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{5, 1, 8, 12, 34, 3, 63, 25, 645, 23, 865, 45, 324, 563, 546, 234}
	quicksort(arr)
	fmt.Println(arr)
}

func quicksort(arr []int) []int {
	if len(arr) < 2 { // если длина массива меньше двух то возращаем массив
		return arr
	}

	left := 0             // инициализируем крайнюю левую переменную
	right := len(arr) - 1 // а здесь крайнюю правую от длины массива

	pivot := rand.Int() % len(arr) // инициализируем случайную точку вокруг которой будем сортировать

	arr[pivot], arr[right] = arr[right], arr[pivot] // меняем эту переменной в этой точке массива с крайним правым элементом

	for i := range arr {
		if arr[i] < arr[right] { // проверяем меньше ли i-й элемент меньше нашего рандомного элемента
			arr[left], arr[i] = arr[i], arr[left] // и если меньше то меняем его с левым элементом
			left++                                // увеличиваем счетчик левого элемента на 1
		}
	}

	arr[left], arr[right] = arr[right], arr[left]
	// меняем местами наш рандомный элемент и левый элемент счетчик которого вырос во время итерация
	// что позволит нам дальше сортировать часть слева и справа
	quicksort(arr[:left])
	quicksort(arr[left+1:])

	return arr
}