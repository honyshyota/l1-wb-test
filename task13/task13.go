package main

import "fmt"

func main() {
	x, y := 1, 10
	fmt.Printf("x = %v, y = %v\n", x, y)
	x, y = y, x
	fmt.Printf("x = %v, y = %v\n", x, y)
}
