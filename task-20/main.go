package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func main() {

	str := "snow dog sun"

	slice := strings.Split(str, " ") // разделение слов по пробелу
	reverse(slice)                   // переворачивание элементов слайса

	newstr := strings.Join(slice, " ")   // сборка новой строки
	fmt.Printf("%s — %s\n", str, newstr) // snow dog sun — sun dog snow

}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i] // меняет местами элементы с начала и конца слайса
	}
}
