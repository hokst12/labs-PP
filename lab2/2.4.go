package main

import "fmt"

func length(a string) int {
	return len(a)
}

func main() {
	a := "Kobachok"
	fmt.Println(length(a))
}
