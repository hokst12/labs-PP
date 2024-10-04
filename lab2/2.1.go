package main

import "fmt"

func opred1(a int) {
	if a == 0 {
		fmt.Println("Число рано нулю")
	} else if a%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}
}

func main() {
	opred1(0)
	opred1(4)
	opred1(9)
	opred1(112)
}
