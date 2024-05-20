package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x float64
	y float64
}

// вычисление дистанции между точками по формуле
func Distance(p1, p2 *Point) float64 {
	dx, dy := p2.x-p1.x, p2.y-p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

// конструктор, возвращает указатель на структуру
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func main() {
	var x, y float64

	fmt.Println("Введите x, y первой точки...")
	fmt.Scan(&x, &y)
	p1 := NewPoint(x, y)

	fmt.Println("Введите x, y второй точки...")
	fmt.Scan(&x, &y)
	p2 := NewPoint(x, y)

	dis := Distance(p1, p2)
	fmt.Printf("Дистанция между 2-мя точками: %.2f\n", dis)

}
