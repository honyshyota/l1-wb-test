package main

import "fmt"

/*Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).*/

type human struct { // родительская структура
	age    int
	growth int
	name   string
}

type action struct { // дочерняя
	*human
}

func (h *human) setParameters(age int, growth int, name string) { // метод родительской структуры
	h.age = age
	h.growth = growth
	h.name = name
}

func (h *human) getParameters() *human { // 2й метод
	return h
}

func main() {
	action := &action{ // инициализируем переменную action
		&human{}, // в анонимном поле human инициализируем пустую переменную типа human
	}

	action.setParameters(26, 186, "Petya") // и так как поле в дочерней структуре анонимное
	fmt.Println(action.getParameters())    // мы можем пользоваться методами родительской структуры
}
