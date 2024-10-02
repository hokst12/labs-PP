package main

import (
	"fmt"
	"time"
)

func sum(a, b int) int {
	return a + b
}

func razn(a, b int) int {
	return a - b
}

func sred(a, b, c int) int {
	return (a + b + c) / 3
}

func main() {
	now := time.Now()
	fmt.Println("Год:", now.Year(), " Месяц:", now.Month(), " День:", now.Day(), " Час:", now.Hour(), " Минута:", now.Minute())
	var (
		a int     = 5
		b float64 = 7.25
		c string  = "gnome"
		d bool    = false
	)
	fmt.Println("Переменные по порядку: ", a, b, c, d)

	f := 11
	fmt.Println("Кратко созданная переменная: ", f)

	fmt.Println("Сумма a и f = ", a+f)
	fmt.Println(sum(a, f), razn(a, f), razn(f, a))
	fmt.Println(sred(a, f, 4))

}
