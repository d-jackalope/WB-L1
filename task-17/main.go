package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

// Алгоритм:
// 1. Сортируем массив данных.
// 2. Делим его пополам и находим середину.
// 3. Сравниваем срединный элемент с заданным искомым элементом.
// 4. Если искомое число больше среднего — продолжаем поиск в правой части массива (если он отсортирован по возрастанию): делим ее пополам, повторяя пункт 3.
//    Если же заданное число меньше — алгоритм продолжит поиск в левой части массива, снова возвращаясь к пункту 3.

// Реализация с итерационным методом, возвращает индекс искомого элемента и метку того, что значение было найдено
func binarySearch_iter(s []int, n int) (int, bool) {
	low := 0           // начальное значение нижней границы поиска
	high := len(s) - 1 //  начальное значение верхней границы

	// если искомый элемент меньше значения нижней границы
	// или больше, то сразу возвращаем результат отсутствия
	// так можно сделать, потому что слайс заранее отсортирован
	if n < s[low] || n > s[high] {
		return 0, false
	}

	// итерация до тех пор пока, нижняя граница меньше или равна верхней,
	// чтобы продолжать поиск, пока не будет найден элемент, в самом худшем случае - элемент будет последним
	var mid int
	for low <= high {
		mid = (low + high) / 2 // определение срединного элемента
		if s[mid] == n {
			break // остановка цикла, если элемент найден
		}
		// если текущее значение срединного элемента меньше искомого
		if s[mid] < n {
			low = mid + 1 // то сдвигаем нижнюю границу вправо
		} else {
			high = mid - 1 // иначе верхнюю влево для сужения области поиска
		}
	}
	return mid, true // возврат результата
}

// Реализация методом рекурсии, возвращает индекс искомого элемента и метку того, что значение было найдено
// в целом все тоже самое
func binarySearch_rec(s []int, n int, low, high int) (int, bool) {
	if n < s[low] || n > s[high] {
		return 0, false
	}

	if low <= high {
		mid := (low + high) / 2

		if s[mid] == n {
			return mid, true
		} else if s[mid] < n {
			return binarySearch_rec(s, n, mid+1, high)
		} else {
			return binarySearch_rec(s, n, low, mid-1)
		}
	}

	return 0, false
}

func main() {
	// Слайс обязательно должен быть отсортирован
	slice := []int{1, 2, 3, 4, 5}

	if result, ok := binarySearch_iter(slice, 1); !ok {
		fmt.Println("Искомого элемента нет")
	} else {
		fmt.Println(result) // 0
	}

	if result, ok := binarySearch_iter(slice, 3); !ok {
		fmt.Println("Искомого элемента нет")
	} else {
		fmt.Println(result) // 2
	}

	if result, ok := binarySearch_iter(slice, 0); !ok {
		fmt.Println("Искомого элемента нет") // Искомого элемента нет
	} else {
		fmt.Println(result)
	}

	if result, ok := binarySearch_rec(slice, 1, 0, len(slice)-1); !ok {
		fmt.Println("Искомого элемента нет")
	} else {
		fmt.Println(result) // 0
	}

	if result, ok := binarySearch_rec(slice, 3, 0, len(slice)-1); !ok {
		fmt.Println("Искомого элемента нет") // Искомого элемента нет
	} else {
		fmt.Println(result) // 2
	}

	if result, ok := binarySearch_rec(slice, 0, 0, len(slice)-1); !ok {
		fmt.Println("Искомого элемента нет") // Искомого элемента нет
	} else {
		fmt.Println(result)
	}
}

//

// В тестовом слайсе было 20 млн упорядоченных значений
// benchtime=100000x
// cpu: Intel(R) Core(TM) i5-4250U CPU @ 1.30GHz
// BenchmarkBinarySearch/binarySearch_iter_value=random-4            100000               479.8 ns/op             0 B/op          0 allocs/op
// BenchmarkBinarySearch/binarySearch_rec_value=random-4             100000               644.4 ns/op             0 B/op          0 allocs/op
// PASS
// ok      github.com/d-jackalope/WB-L1/task-17    0.248s
