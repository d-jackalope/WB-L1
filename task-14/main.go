package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func main() {
	checkType(555)
	checkType("строка")
	checkType(true)

	c := make(chan int)
	checkType(c)

	checkType(45.33)

	// Вывод:
	// Тип int 555
	// Тип string строка
	// Тип bool true
	// Тип chan int
	// Неизвестный тип

}

func checkType(i interface{}) {
	switch reflect.TypeOf(i).Kind() { //проверка на соответствие с различными типами, используя рефлексию, чтобы можно было проверить типы в рантайме
	case reflect.Int:
		fmt.Println("Тип int", i)
	case reflect.String:
		fmt.Println("Тип string", i)
	case reflect.Bool:
		fmt.Println("Тип bool", i)
	case reflect.Chan:
		fmt.Println("Тип chan int")
	default: // в случае, если нет ожидаемых типов
		fmt.Println("Неизвестный тип")
	}
}
