package main

import (
	"fmt"

	"stringutils"
)

func main() {
	var a string
	fmt.Println("Введите строку: ")
	fmt.Scan(&a)
	fmt.Println(stringutils.Reverse(a))
}
