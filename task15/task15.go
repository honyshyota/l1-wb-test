package main

import "fmt"

// создаем строку заданой в аргументе длины
func createHugeString(len int) string {
	var s string
	for i := 0; i < len; i++ {
		s += "й"
	}

	return s
}

func someFunc() {
	v := createHugeString(1 << 10) // вызываем функцию
	justString := v[:100]
	// так как строка по сути массив байт, делаем срез до 100 байта
	// основная проблема видимо здесь, ибо если искомая строка содержит
	// символы Unicode они могут занимать от 1 до 4 байт, если же ASCII
	// проблем не будет, избежать этого можно использую тип rune

	s := []rune(v)      // приводим строку к типу rune
	justRune := s[:100] // делаем срез до 100 байта

	fmt.Println(string(justRune)) // результат вывода среза по rune
	fmt.Println(justString)       // результат вывода среза по строке
}

func main() {
	someFunc()
}
