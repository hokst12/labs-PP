package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fact(a int) {
	var b int = 1
	for i := 1; i <= a; i++ {
		b *= i
		fmt.Println("fact element ", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println(" факториал числа = ", b)
}

func random() {
	var a [5]int
	for i := range a {
		a[i] = rand.Intn(100-1) + 1
		fmt.Println("rand element ", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("случайные числа - ", a)
}

func ryad(a []int) {
	var sum int = 0
	for s, i := range a {
		sum += i
		fmt.Println("ryad element ", s)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("сумма ряда = ", sum)
}

func main() {
	go fact(7)
	go random()
	go ryad([]int{5, 6, 6, 8, 1, 5, 7})

	fmt.Scanln()

}
