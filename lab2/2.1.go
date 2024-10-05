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
	var a int
	fmt.Print("Введите число: ")
	fmt.Scan(&a)
	opred1(a)

}
