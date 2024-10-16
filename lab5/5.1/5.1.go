package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) info() {
	fmt.Println("Имя: ", p.name, " Возраст: ", p.age)
}

func main() {

	var b = person{name: "Ben", age: 44}

	b.info()
}
