package main

import (
	"fmt"
)

func sred(a, b, c int) int {
	return (a + b + c) / 3
}

func main() {
	var a, b int = 11, 3

	fmt.Println(sred(a, b, 4))
}
