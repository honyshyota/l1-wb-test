package main

import (
	"flag"
	"fmt"
)

/*Удалить i-ый элемент из слайса.*/

func main() {
	var i int
	flag.IntVar(&i, "i", 3, "delete slice element")
	flag.Parse()

	arr := []int{0, 1, 2, 3, 4, 5}

	arr[i] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]

	fmt.Println("Удаление с нарушением сортировки: ", arr)

	arr = []int{0, 1, 2, 3, 4, 5}
	var res []int

	for index := range arr {
		if i != index {
			res = append(res, arr[index])
		}
	}

	fmt.Println("Удаление без нарушения сортировки, но через еще один слайс: ", res)
}
