package main

import (
	"fmt"

	"github.com/d-jackalope/WB-L1/task-21/example"
)

// Реализовать паттерн «адаптер» на любом примере.

// Суть: обеспечить возможность использования объектов разных типов (структур) через общий интерфейс

// создание единого интерфейса
type Act interface {
	Action() // c определением метода для структур
}

// оболочка для импортируемой структуры с ее методами
// он же адаптер структуры Human из пакета example
type HumanAdapter struct {
	*example.Human
}

// реализация метода интерфейса, которая реализует логику адаптируемой структуры
func (a *HumanAdapter) Action() {
	a.Human.Act()
}

// конструктор адаптера
func NewHumanAdapter(human *example.Human) Act {
	return &HumanAdapter{Human: human}
}

type AnimalAdapter struct {
	*example.Animal
}

func (a *AnimalAdapter) Action() {
	a.Animal.Act()
}

func NewAnimalAdapter(animal *example.Animal) Act {
	return &AnimalAdapter{Animal: animal}
}

// структура в пакете main, которая так же реализует интерфейс Act
type Food struct {
	act string
}

// реализация метода интерфейса
func (f *Food) Action() {
	fmt.Printf("Структура Food %v в пакете main\n", f.act)
}

func NewFood(act string) Act {
	return &Food{
		act: act,
	}
}

func main() {

	animal := NewAnimalAdapter(example.NewAnimal("лезет в холодильник"))
	human := NewHumanAdapter(example.NewHuman("в данный момент кушает"))
	food := NewFood("была съедена")

	actions := []Act{animal, human, food}
	// Независимо от конкретного типа теперь можно применить метод action()
	// Каждый тип данных вызовет свой метод action()
	for _, act := range actions {
		act.Action()
	}

	// вывод:
	// Структура Animal из пакета example лезет в холодильник
	// Структура Human из пакета example в данный момент кушает
	// Структура Food была съедена в пакете main
}
