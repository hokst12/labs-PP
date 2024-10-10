package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Введите строку:")
	var s string
	fmt.Scan(&s)
	fmt.Println(strings.ToUpper(s))
}
