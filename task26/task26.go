package main

import (
	"flag"
	"fmt"
	"strings"
)

/*Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func main() {
	var s string

	flag.StringVar(&s, "s", "aVkrEDqw", "some string")
	flag.Parse() // парсим строку из командной строки

	s = strings.ToLower(s) // приводим строку к нижнему регистру

	fmt.Println(checkUnique(s))
}

func checkUnique(str string) bool {
	strArr := []rune(str) // приводим строку к срезу рун

	for i := range strArr { // итерируемся пока не найдем одинаковые значения
		for j := range strArr {
			if i != j && strArr[i] == strArr[j] {
				return false
			}
		}
	}

	return true // если не нашли возращаем true
}
