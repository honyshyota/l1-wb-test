package main

import (
	"fmt"
	"strings"
)

/*Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».*/

func main() {
	s := "sun dog snow" // инициализируем готовую строку из примера

	sArr := strings.Split(s, " ") // разбиваем ее в массив
	var resArr []string           // инициализируем массив в который вернем значения после итерации

	for i := len(sArr) - 1; i >= 0; i-- { // итерируемся подобно примеру из предыдущей задачи
		resArr = append(resArr, sArr[i]) // и так же последни элемент подставляем первым в результат
	}

	var result string // инициализируем результирующую строку

	for i := 0; i < len(resArr); i++ { // итерируемся по массиву резльтирующему
		result = result + resArr[i] + " " // конкатенируем строки
	}

	result = strings.TrimRight(result, " ") // обрезаем в конце пробел
	fmt.Println(result)                     // ну и выводим результат на в stdout
}
