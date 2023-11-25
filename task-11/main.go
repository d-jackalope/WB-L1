package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	set1 := map[string]struct{}{
		"cat":   {},
		"dog":   {},
		"duck":  {},
		"human": {},
	}

	set2 := map[string]struct{}{
		"cat":  {},
		"dog":  {},
		"duck": {},
		"bird": {},
	}

	// множество для хранения пересечений
	intersect := make(map[string]struct{})
	// проходим по элементам одного из множеств и проверяем наличие во втором
	for k := range set1 {
		if _, ok := set2[k]; ok {
			intersect[k] = struct{}{} // если найдено, то записываем в intersect
		}
	}

	fmt.Printf("Ожидаемая длина пересечения: %v из 3\n", len(intersect))
	for k := range intersect {
		fmt.Println(k)
	}

	// вывод:
	// Ожидаемая длина пересечения: 3 из 3
	// dog
	// duck
	// cat

}
