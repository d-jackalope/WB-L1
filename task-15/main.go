package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

func createHugeString(size int) string {
	chars := []rune("йцукенгшщзфывапролдячсмить") // все символы состоят из 2-х байт

	var b strings.Builder
	for i := 0; i < size; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	// 1. Некорректное взятие символов
	// тк как взятие подстроки происходит побайтово, то подстрока будет состоять из первых 100 байт строки
	// в случае если строка v содержит многобайтовые символы - строку некорректно порежет
	justString = v[:100]
	fmt.Printf("%v символов из 100\n", len([]rune(justString))) // 50 символов из 100
	// поэтому корректней будет преобразовать строку в срез рун,
	// взять первые 100 символов и создать новую строку

	runes := []rune(v)
	justString = string(runes[:100])
	fmt.Printf("%v символов из 100\n", len([]rune(justString)))

	// 2. Хранение в памяти огромной строки, хотя нужна только часть
	// Срезы в го содержат в себе ссылку на базовый массив. Присваивая глобальной переменной justString значение с помощью v[:100]
	// создается подстрока (он же срез строк), таким образом удерживая в памяти огромную строку, но по факту
	// от основной строки используется только часть, сборщик мусора не может удалить основную строку
	// тк на нее есть ссылки, это излишний расход памяти. Лучше пересоздать с нужным размером новую строку
	// а сборщик мусора спокойно удалит огромную строку, тк она больше не используется.

}

func main() {
	someFunc()
}
