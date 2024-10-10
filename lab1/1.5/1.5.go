package main

import (
	"fmt"
)

func sum(a, b float32) float32 {
	return a + b
}

func razn(a, b float32) float32 {
	return a - b
}

func main() {
	var a, b float32 = 11.5, 3.7

	fmt.Println(sum(a, b), razn(a, b), razn(b, a))
}
