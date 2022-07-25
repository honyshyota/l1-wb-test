package main

import (
	"flag"
	"fmt"
	"math"
)

/*Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.*/

func main() {
	var a, b uint64
	var operation string

	// парсим данные из командной строкиб дефолтные значения для числовых переменных равны 2^20
	// дефолтное значение строковой переменной равно +
	flag.Uint64Var(&a, "a", 1048576, "a number")
	flag.Uint64Var(&b, "b", 1048576, "b number")
	flag.StringVar(&operation, "op", "+", "math operation")
	flag.Parse()

	// делаем проверку по 20 степени числа 2
	if a < uint64(math.Pow(2, 20)) || b < uint64(math.Pow(2, 20)) {
		fmt.Println("ввели значение меньше чем 2^20")
		return
	}

	// ну и операционная часть
	switch operation {
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	default:
		fmt.Println("неправильный ввод")
	}
}
