package example

import "fmt"

type Human struct {
	act string
}

func (h *Human) Act() {
	fmt.Printf("Структура Human из пакета example %s\n", h.act)
}

func NewHuman(act string) *Human {
	return &Human{
		act: act,
	}
}

type Animal struct {
	act string
}

func (a *Animal) Act() {
	fmt.Printf("Структура Animal из пакета example %v\n", a.act)
}

func NewAnimal(act string) *Animal {
	return &Animal{
		act: act,
	}
}
