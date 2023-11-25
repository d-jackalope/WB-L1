package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {

	values := []string{"cat", "cat", "dog", "cat", "tree"}

	// создание пустого множества с помощью map
	set := make(map[string]struct{})

	// Добавление строк в множество, сохраняя только уникальные элементы
	for _, v := range values {
		set[v] = struct{}{}
	}

	for k := range set {
		fmt.Println(k)
	}

	// вывод:
	// cat
	// dog
	// tree
}
