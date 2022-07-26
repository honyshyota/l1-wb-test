package main

import (
	"flag"
	"fmt"
)

/*Разработать программу, которая переворачивает подаваемую на ход
строку (например: «главрыба — абырвалг»). Символы могут быть unicode.*/

func main() {
	var s string
	flag.StringVar(&s, "s", "golang", "income string")
	flag.Parse() // получаем строку с помощью аргумента командной строки

	// переводим строку в массив рун, ибо если символы в Unicode то могу занимать
	// от 1 до 4 байт, и итерация по такой строке не даст нужного результата
	sRune := []rune(s)
	var result []rune // результат тоже в rune

	// итерируемся по массиву где i будет равно максимальной длине массива
	// условие работает пока i >= 0
	// ну и при итерации отнимаем по 1 из i
	// это нужна нам чтобы проще было брать последние элементы из исходного массива
	// и аппендить их к результирующему массиву но уже с нулевой позиции
	for i := len(sRune) - 1; i >= 0; i-- { // итерируемся по массиву
		result = append(result, sRune[i]) // и добавляем
	}

	fmt.Println(string(result)) // ну и сам результат приводим к типу string
}
