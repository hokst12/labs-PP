package main

import (
	"fmt"

	"lab3/stringutils"
)

func main() {
	var a string
	fmt.Println("Введите строку: ")
	fmt.Scan(&a)
	fmt.Println(stringutils.Reverse(a))
}
