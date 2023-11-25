package main

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

// реализовывать ли float?

type number struct { // хранение основания и степени числа
	base int64
	exp  int64
}

// парсинг значений из слайса строк в структуру
func (n *number) parse(s []string) error {
	if len(s) < 2 {
		return errors.New("некорректный ввод")
	}
	base, err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}
	n.base = int64(base)

	exp, err := strconv.Atoi(s[1])
	if err != nil {
		return err
	}
	n.exp = int64(exp)
	return nil
}

func main() {
	var strA, strB string
	// считывание через основание и степень
	// сделано скорее для удобства
	// отрицательная степень даст - "1", тк работа с целочисленными данными
	fmt.Println("Введите основание и степень 2-х чисел, например, \"2^100 2^64\"...")
	fmt.Scan(&strA, &strB)
	partsA := strings.Split(strA, "^")
	partsB := strings.Split(strB, "^")

	a, b := number{}, number{}
	if err := a.parse(partsA); err != nil {
		fmt.Println(err, strA)
		return
	}
	if err := b.parse(partsB); err != nil {
		fmt.Println(err, strB)
		return
	}

	expBigInt(&a, &b)

	// вывод:
	// Введите основание и степень 2-х чисел, например, "2^100 2^64"...
	// 2^2 2^2
	// a=4
	// b=4
	// Произведение: 16
	// Частное: 1
	// Сложение: 8
	// Разность: 0

	// Введите основание и степень 2-х чисел, например, "2^100 2^64"...
	// 2^100 2^64
	// a=1267650600228229401496703205376
	// b=18446744073709551616
	// Произведение: 23384026197294446691258957323460528314494920687616
	// Частное: 68719476736
	// Сложение: 1267650600246676145570412756992
	// Разность: 1267650600209782657422993653760

	// Введите основание и степень 2-х чисел, например, "2^100 2^64"...
	// 2^0 0^1
	// a=1
	// b=0
	// Произведение: 0
	// Деление на ноль!
	// Сложение: 1
	// Разность: 1

}

// через ввод основания и степени числа
func expBigInt(a, b *number) {
	bigA := new(big.Int)
	bigB := new(big.Int)

	bigA.Exp(big.NewInt(a.base), big.NewInt(a.exp), nil)
	bigB.Exp(big.NewInt(b.base), big.NewInt(b.exp), nil)

	fmt.Printf("a=%v\nb=%v\n", bigA, bigB)

	mul := new(big.Int) // инициализация значения big.Int
	mul.Mul(bigA, bigB) // расчет произведения a и b
	fmt.Printf("Произведение: %v\n", mul)

	if bigB.IsInt64() && bigB.Int64() == 0 { // проверка на int64 и значения делителя на ноль
		fmt.Println("Деление на ноль!")
	} else {
		div := new(big.Int)
		div.Div(bigA, bigB)
		fmt.Printf("Частное: %v\n", div) // расчет частного a и b
	}

	sum := new(big.Int)
	sum.Add(bigA, bigB) // расчет сложения a и b
	fmt.Printf("Сложение: %v\n", sum)

	sub := new(big.Int)
	sub.Sub(bigA, bigB) // расчет разности a и b
	fmt.Printf("Разность: %v\n", sub)
}

// ввод больших чисел с помощью строки, из минусов - неудобно вводить
// первоначальный вариант, остановилась на варианте с основанием и степенью
func stringBigInt(a, b string) {

	bigA, ok := new(big.Int).SetString(a, 10) // преобразование строки в big.Int
	if !ok {                                  // проверка корректности ввода
		fmt.Println("Некорректное значение: ", a)
		return
	}
	bigB, ok := new(big.Int).SetString(b, 10)
	if !ok {
		fmt.Println("Некорректное значение: ", b)
		return
	}

	mul := new(big.Int) // инициализация значения big.Int
	mul.Mul(bigA, bigB) // расчет произведения a и b
	fmt.Printf("Произведение: %v\n", mul)

	if bigB.IsInt64() && bigB.Int64() == 0 { // проверка на int64 и значения делителя на ноль
		fmt.Println("Деление на ноль!")
	} else {
		div := new(big.Int)
		div.Div(bigA, bigB)
		fmt.Printf("Частное: %v\n", div) // расчет частного a и b
	}

	sum := new(big.Int)
	sum.Add(bigA, bigB) // расчет сложения a и b
	fmt.Printf("Сложение: %v\n", sum)

	sub := new(big.Int)
	sub.Sub(bigA, bigB) // расчет разности a и b
	fmt.Printf("Разность: %v\n", sub)
}
