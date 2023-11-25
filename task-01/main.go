package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	name string
	age  int
}

func (h *Human) GetAge() int {
	return h.age
}

func (h *Human) GetName() string {
	return h.name
}

type Action struct {
	Human //Тип Human встроен(embedded) в тип Action
	act   string
}

func main() {
	action := Action{
		Human: Human{name: "duck", age: 28},
		act:   "quack quack",
	}

	fmt.Println("Получение name и age из cтруктуры Action путем встраивания структуры Human и ее методов")
	age := action.GetAge()
	name := action.GetName()
	fmt.Printf("name: %s\nage: %v\n\n", name, age)
}
