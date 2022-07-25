package main

import (
	"fmt"
	"sync"
)

/*Реализовать конкурентную запись данных в map.*/

func main() {
	a := make(map[int]int) // инициализируем мапу
	mu := &sync.Mutex{}    // инициализируем мютекс
	var wg sync.WaitGroup  // инициализируем вэйт группу

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(a map[int]int, i int, mu *sync.Mutex) {
			for j := 0; j < 10; j++ {
				mu.Lock()     // мютекс лочится при работе с мапой и по завершению анлок для других гоурутин
				a[i*10+j] = j // при записи в мапу нужен уникальный ключ ибо данные при одинаковом ключе перезаписываются
				mu.Unlock()   // тот самый анлок
			}
			wg.Done()
		}(a, i, mu)
	}
	wg.Wait()
	fmt.Println(a)
}
