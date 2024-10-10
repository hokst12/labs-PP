package main

import "fmt"

func sred(a, b int) int {
	return (a + b) / 2
}

func main() {
	c := 11
	d := 24

	fmt.Println(sred(c, d))
}
