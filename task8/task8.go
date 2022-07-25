package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

/*Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.*/

var number int
var positionBit int
var variable int

func init() {
	flag.IntVar(&number, "n", 400, "number")
	flag.IntVar(&positionBit, "p", 1, "postion bit")
	flag.IntVar(&variable, "v", 1, "variable value")
}

func main() {
	flag.Parse() // парсим аргументы командной строки

	intToBin := strconv.FormatInt(int64(number), 2) // перевод числа в двоичную систему просто для проверки на несоответствие длины

	if positionBit > len(intToBin) {
		log.Fatalln("позиция бита выходит за рамки представления искомого числа в двоичном исчеслении")
	} else if positionBit < 0 {
		log.Fatalln("значение позиции бита должно быть положительным")
	}

	if variable > 1 || variable < 0 { // проведрка значения бита 1 или 0
		log.Fatalln("значение бита должно быть либо 1, либо 0")
	}

	fmt.Printf("Число %d в двоичной системе измерений до изменения %b\n", number, number)

	switch variable {
	case 1: // используем подразрядную дизюнкцию
		number |= 1 << positionBit
	case 0: // используем сброс бита
		number &^= 1 << positionBit
	}

	fmt.Printf("Число %d в двоичной системе измерений после изменения %b\n", number, number)
}
