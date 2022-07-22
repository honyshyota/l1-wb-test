package main

import (
	"fmt"
	"sync"
)

type count struct { // инициализируем структуру
	count          int // значние которое будем инкременировать
	sync.WaitGroup     // группа ожидания для гоурутин
	sync.Mutex         // мютекс для блокировки переменной от записи
}

func newCount() *count { // функция для инициализации переменной типа *count
	return &count{}
}

func (c *count) incrementCount() { // инкреминирующая функция
	c.Lock()   // блокируем переменную
	c.count++  // пишем +1 в переменную
	c.Unlock() // разблокируем переменную для следующих гоурутин
}

func (c *count) get() int { // функция для извлечения итогового значения
	return c.count
}

func main() {
	c := newCount() // инициализируем переменную

	for i := 0; i < 10000; i++ { // заупскаем цикл до нужного нам значения
		c.WaitGroup.Add(1) // добавляем в группу ожидания 1 элемент для гоурутины

		go func(c *count) {
			c.incrementCount() // в гоурутине инкременируем значение
			c.WaitGroup.Done() // и закрываем позицию в очереди
		}(c)
	}

	c.WaitGroup.Wait() // ждем завершения очередей
	fmt.Println(c.get())

}
