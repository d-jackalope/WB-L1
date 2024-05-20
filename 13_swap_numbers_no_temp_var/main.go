package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	a := 100
	b := 200

	fmt.Println("Начальные значения")
	fmt.Printf("a: %v, b: %v\n", a, b)

	fmt.Println("Обмен c помощью синтаксиса go")
	a, b = b, a
	fmt.Printf("a: %v, b: %v\n", a, b)

	fmt.Println("Обмен c помощью xor")
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Printf("a: %v, b: %v\n", a, b)

}
