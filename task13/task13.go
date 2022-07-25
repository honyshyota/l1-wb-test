package main

import "fmt"

/*Поменять местами два числа без создания временной переменной.*/

func main() {
	x, y := 1, 10
	fmt.Printf("x = %v, y = %v\n", x, y)
	x, y = y, x
	fmt.Printf("x = %v, y = %v\n", x, y)
}
