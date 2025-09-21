package main

import "fmt"

type Human struct {
	Name    string
	Age     int
	Address string
}

// String() специальный метод в Go, который позволяет кастомизировать строковое представление структур
func (h Human) String() string {
	return fmt.Sprintf("Имя: %s\nВозраст: %d\nАдрес: %s", h.Name, h.Age, h.Address)
}

type Action struct {
	Human
}

func main() {
	action := Action{Human{
		Name:    "Platon",
		Age:     22,
		Address: "Moscow",
	}}

	fmt.Println(action) // Происходит неявный вызов String()

	//	Вывод:
	//	——————————————
	//	Имя: Platon
	//	Возраст: 22
	//	Адрес: Moscow

}
