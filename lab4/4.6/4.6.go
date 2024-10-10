package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Reverse(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

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
	fmt.Println(Reverse(b))

}
