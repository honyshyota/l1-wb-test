package main

import "fmt"

type human struct {
	age    int
	growth int
	name   string
}

type action struct {
	human *human
}

func (a *action) setParameters(age int, growth int, name string) {
	a.human = &human{
		age:    age,
		growth: growth,
		name:   name,
	}
}

func (a *action) getParameters() *human {
	return a.human
}

func main() {
	var human action                      // Инициализация переменной типа action, nil значение в памяти
	human.setParameters(24, 186, "Petya") // Передача значений в переменную human по указателю, присвоение адресса в памяти
	fmt.Println(human.getParameters())
}
