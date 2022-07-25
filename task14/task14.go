package main

import "fmt"

/*Разработать программу, которая в рантайме способна определить
тип переменной: int, string, bool, channel из переменной типа interface{}.*/

func main() {
	// объявляем переменный с дефолтными значениями
	var i int
	var s string
	var b bool
	var ch chan int

	//вызываем нашу функция для опеределения типа переменной и передаем ей заранее заготовленные переменные
	do(i)
	do(s)
	do(b)
	do(ch)

}

func do(i interface{}) {
	// используем switch type для определения типа переменной в интерфейсе
	switch val := i.(type) {
	case int:
		fmt.Println("value is int ", val)
	case string:
		fmt.Println("value is string ", val)
	case bool:
		fmt.Println("value is boolean ", val)
	case chan int:
		fmt.Println("value is channel ", val)
	}
}
