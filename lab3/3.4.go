package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a [5]int
	for i, _ := range a {
		a[i] = rand.Intn(100-1) + 1
	}

	fmt.Println(a)
}
