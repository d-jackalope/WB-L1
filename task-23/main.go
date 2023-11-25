package main

import (
	"fmt"
)

// Удалить i-ый элемент из слайса.

func main() {
	slice := []int{1, 2, 3, 4, 5}
	i := 4

	slice = removeWithAppend(slice, i)
	fmt.Println(slice)

	slice2 := []int{1, 2, 3, 4, 5}
	i = 2

	slice2 = removeWithCopy(slice2, i)
	fmt.Println(slice2)

	slice3 := []int{1, 2, 3, 4, 5}
	i = 4

	slice2 = removeWithLastElement(slice3, i)
	fmt.Println(slice2)

}

func removeWithAppend(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice
	}
	return append(slice[:i], slice[i+1:]...) // вращает новый слайс без элемента, который нужно было удалить
}

func removeWithCopy(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice
	}

	copy(slice[i:], slice[i+1:]) // сдвиг влево на один элемент, тем самым перезаписывая ненужный элемент
	return slice[:len(slice)-1]  // возвращает урезанный срез c конца, тк там дублируется последний элемент
}

// если порядок не важен, самый эффективный вариант
func removeWithLastElement(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice
	}

	slice[i] = slice[len(slice)-1] // меняет местами последний элемент с тем, который надо удалить
	return slice[:len(slice)-1]    // // возвращает урезанный срез c конца, тк там элемент, который надо удалить
}
