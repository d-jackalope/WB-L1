package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func main() {
	str := "abCd"
	fmt.Println(uniqueStringWithMap(str)) // true

	str = "abCdefAaf"
	fmt.Println(uniqueStringWithMap(str)) // false

	str = "aabcd"
	fmt.Println(uniqueStringWithMap(str)) // false

}

func uniqueStringWithMap(s string) bool {
	s = strings.ToLower(s) // приведение строки к единому регистру
	// хранение уникальных символов, значение опускается, поэтому struct{}
	charMap := make(map[rune]struct{})

	for _, char := range s { // итерация по символам строки
		if _, ok := charMap[char]; !ok { // если символа нет, то добавляется в мапу
			charMap[char] = struct{}{}
		} else { // если есть, то выход из функции, ибо символ повторился
			return false
		}

	}
	return true // после окончания цикла возвращает true, ибо все значения уникальные
}
