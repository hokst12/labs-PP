package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var b []int
	var s string
	for {
		fmt.Println("Введите число, для остановки введите stop:")
		fmt.Scan(&s)
		if strings.ToLower(s) == "stop" {
			break
		} else {
			c, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			b = append(b, c)
		}
	}
	var sum int = 0
	for _, value := range b {
		sum += value
	}
	fmt.Println("Сумма всех введенных чисел: ", sum)

}
