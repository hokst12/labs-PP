package main

import "fmt"

func vozr(a map[string]int) int {
	i := 0
	sum := 0
	for _, c := range a {
		sum += c
		i += 1
	}
	return sum / i
}

func main() {
	var people = map[string]int{
		"Дмитрий": 53,
		"Оксана":  35,
		"Иван":    16,
		"Олег":    27}

	fmt.Println(vozr(people))
}
