package main

import (
	"fmt"
	"sort"
)

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode.

func main() {
	text := "главрыба Ê"
	expected := "Ê абырвалг" // строка для проверки

	// Метод итерирования по строке
	revText := reverseString(text)
	if expected == revText {
		fmt.Println(true, revText) // true Ê абырвалг
	}

	// Преобразование строки в слайс рун, где напрямую меняется порядок элементов
	runesRev := []rune(text)
	reverseRune(runesRev)
	revText = string(runesRev)
	if expected == revText {
		fmt.Println(true, revText) // true Ê абырвалг
	}

	// Преобразование строки в слайс рун, где меняется порядок элементов с помощью sort
	runesSort := []rune(text)
	sortRunes(runesSort)
	revText = string(runesRev)
	if expected == revText {
		fmt.Println(true, revText) // true Ê абырвалг
	}

}

// не очень хороший способ, ибо на каждой итерации создается новая строка
// к тому же тут нужно проитерироваться по всей длине
func reverseString(r string) string {
	var rev string           // создание пустой строки
	for _, char := range r { // итерация по каждому символу (руне)
		rev = string(char) + rev // добаление текущей строки в начала созданной
	}
	return rev // возвращение новой строки
}

// эта реализация лучше, как минимум она не выделяет память
// и количество итераций равно len(runes)/2, тк цикл доходит до середины и прекращает работу
func reverseRune(runes []rune) {
	// инициализация 2-х индексов, где i - первый элемент, j - последний
	// переворачивание строки идет до тех пор, пока i меньше j,
	// каждая итерация будет перемещать индексы к середине
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // меняет местами элементы с начала и конца слайса
	}
}

// Эта реализация для сравнения
func sortRunes(runes []rune) {
	sort.Slice(runes, func(i, j int) bool {
		return i > j // условие определяющее обратный порядок
	})
}

// В тесте была строка длиной 10.000 рандомных символов
// benchtime=100x
// cpu: Intel(R) Core(TM) i5-4250U CPU @ 1.30GHz
// BenchmarkReverseString/reverseString-4               100          30433561 ns/op        105389424 B/op     10004 allocs/op
// BenchmarkReverseString/reverseRune-4                 100              8859 ns/op               0 B/op          0 allocs/op
// BenchmarkReverseString/sortRunes-4                   100           1264668 ns/op              56 B/op          2 allocs/op
// PASS
// ok      github.com/d-jackalope/WB-L1/task-19    3.218s
// Как ожидалось, reverseString - выделяет много памяти и долго работает
// reverseRune - показывает наилучший результат по времени и не выделяет память
// встроенный метод сортировки с условием работает хуже, чем если напрямую менять элементы местами
