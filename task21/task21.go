package main

import (
	"fmt"
)

type target interface { // интерфейс с которым мы хотим работать
	print() string
}

type adapter struct { // адаптер для структуры которую невозможно переписать
	*adaptee
}

type adaptee struct{} // сама искомая структура

func newAdapter(adaptee *adaptee) target { // функция встраивания типа который нам нужен в наш адаптер
	return &adapter{adaptee}
}

func (a *adaptee) specPrint() string { // метод структуры adaptee к которому мы получим доступ через адаптер
	return "hello world"
}

func (a *adapter) print() string { // и метод нашего адаптера в котором вызываем метод для нашей искомой структуры
	return a.specPrint()
}

func main() {
	adapter := newAdapter(&adaptee{}) // инициализируем адаптер и встаиваем пустую переменную нашей структуры adaptee

	print := adapter.print() // вызывает метод Print нашего адаптера который в свою очередь обращается к методу структуры adaptee

	fmt.Println(print) // для наглядности выведем в Stdout
}
