package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// функция подсчета дистанции по формуле
func distance(a *Point, b *Point) float64 {
	distance := math.Sqrt(math.Pow((b.X-a.X), 2) + math.Pow((b.Y-a.Y), 2))
	return distance
}

// конструктор
func newPoint(x float64, y float64) *Point {
	return &Point{x, y}
}

func main() {
	a := newPoint(2.34, 1.30)
	b := newPoint(5.32, 6.43)

	fmt.Println("Дистанция между точек равна: ", distance(a, b))
}
