package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) info() {
	fmt.Println("Имя: ", p.name, " Возраст: ", p.age)
}

func (p *person) birthday() {
	(*p).age += 1
}

func main() {

	var b = person{name: "Ben", age: 44}

	b.birthday()

	b.info()
}
