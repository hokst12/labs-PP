package main

import (
	"fmt"

	"mathutils"
)

func main() {
	var a int
	fmt.Println("Введите число: ")
	fmt.Scan(&a)
	fmt.Println(mathutils.Fact(a))
}