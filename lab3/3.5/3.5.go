package main

import (
	"fmt"
)

func main() {
	var a [8]int = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	srez := a[3:7]

	fmt.Println(srez)

	srez = append(srez, 11)
	fmt.Println(srez)

	srez = append(srez[:2], srez[3:]...)
	fmt.Println(srez)

}
