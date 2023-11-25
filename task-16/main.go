package main

import "fmt"

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

// Алгоритм:
// 1. Выбрать элемент из массива. Назовём его опорным.
// 2. Разбиение: перераспределение элементов в массиве таким образом, что элементы,
//    меньшие опорного, помещаются перед ним, а большие или равные - после.
// 3. Рекурсивно применить первые два шага к двум подмассивам слева и справа от
// 	  опорного элемента. Рекурсия не применяется к массиву, в котором только один
// 	  элемент или отсутствуют элементы.

// quickSort - это рекурсивная функция, которая имеет 3 параметра,
// s - срез, который требуется отсортировать,
// start - индекс начала области среза,
// end - индекс конца области среза
func quickSort(s []int, start, end int) {
	if start >= end { // базовый случай для выхода из функции
		return // так как подсрез состоит из одного элемента или такого среза вообще не может быть
	}

	// определение элемента "опорной точки", чтобы сравнивать его с другими элементами
	pivot := s[start]

	// установка индекса для следующего элемента после "опорного",
	// это позиция нужна, чтобы указывать место для элементов, которые меньше "опорного"
	index := start + 1

	// проход по всем элементам от начала подсреза до его конца относительно опорного элемента
	for j := start; j <= end; j++ {
		if pivot > s[j] { // если опорный элемент больше текущего элемента в подсрезе,
			s[index], s[j] = s[j], s[index] // то элемент в позиции index меняем местами с текущим
			// и инкрементируем index для того, чтобы указать новое место для следующего элемента,
			// который меньше опорного
			index++
		}
	}

	// После того, как цикл отработал, нужно переместить опорный элемент
	// для того, чтобы все элементы слева были меньше или равны ему, а справа были больше или равны
	// index-1 тк опорный элемент нужно поменять с последним наименьшим элементом при сравнении
	s[start], s[index-1] = s[index-1], s[start]

	// далее рекурсивный вызов для сортировки левой и правой части относительно опоры
	// не включая саму опору, так как она находится на своем месте
	quickSort(s, start, index-2) // для левой части
	quickSort(s, index, end)     // для правой части
}

func main() {
	slice := []int{4, 654, 5, 3, 6, -3, -1000, 4000, 80000, 64000}

	newSlise := quickSort_2(slice)
	fmt.Println(newSlise) // [-1000 -3 3 4 5 6 654 4000 64000 80000]

	quickSort(slice, 0, len(slice)-1)
	fmt.Println(slice) // [-1000 -3 3 4 5 6 654 4000 64000 80000]

}

// Этот вариант был найден в интернете, я решила его разобрать и сделать бенчмарк, чтобы сравнить с моим вариантом
// Разница в подходе. В моей версии опорный элемент выбирается как самый левый, тут же срединный,
// от этого зависит эффективность алгоритма. Так же в моей реализации нет прямого деления и создания подсрезов,
// идет оперирование значений индексов, в рекурсивный вызов каждый раз подается цельный слайс (ссылка по факту) и новые значения индексов.
// Тут же создаются новые подсрезы и они в цикле наполняются значениями, а потом все склеивается в единый слайс.
// Скорее всего при больших данных подобная реализация будет жрать больше памяти за счет выделения каждый раз новых подслайсов
// В моем варианте функция работает с оригинальным срезом и не возвращает копию. Так как срез является ссылочным типом данных
func quickSort_2(s []int) []int {
	if len(s) <= 1 { // базовый случай
		return s
	}

	// выбор "опорного" элемента путем взятия среднего элемента
	index := len(s) / 2
	pivot := s[index]
	// создание нового среза без опорного элемента
	s = append(s[:index], s[index+1:]...)

	// создание пустых подсрезов для наименьших и наибольших элементов относительно опорного
	var less, greater []int
	for _, v := range s {
		if v <= pivot { // если элемент в срезе меньше или равен опорного
			less = append(less, v) // то он добавляется в подсрез наименьших
		} else {
			greater = append(greater, v) // иначе в наибольшие
		}
	}

	less = quickSort_2(less)       // рекурсивный вызов для подсреза с наименьшими элементами относительно опорного
	greater = quickSort_2(greater) // для наибольших

	// слияние отсортированных подсрезов
	sorted := append(less, pivot)       // наименьший + опорный
	sorted = append(sorted, greater...) // остальное + наибольшие элементы

	return sorted // возврат отсортированного массива
}

// В тестовом слайсе было 1 млн рандомных значений
// benchtime=10x
// cpu: Intel(R) Core(TM) i5-4250U CPU @ 1.30GHz
// BenchmarkQuickSort/quickSort-4                10         122076663 ns/op         8003593 B/op          1 allocs/op
// BenchmarkQuickSort/quickSort_2-4              10        1380974570 ns/op        1919205616 B/op  6844813 allocs/op какая жесть
// BenchmarkQuickSort/sort.Ints-4                10         157580699 ns/op         8003608 B/op          2 allocs/op
// PASS
// ok      github.com/d-jackalope/WB-L1/task-16    18.174s

// quickSort_2 выделил и использовал кучу памяти, а так же выполнился один раз с большим временем выполнения
// quickSort и sort.Ints работают примерно одинаково хорошо
