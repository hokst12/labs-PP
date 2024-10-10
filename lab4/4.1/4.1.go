package main

import "fmt"

func main() {
	var people = map[string]int{
		"Дмитрий": 53,
		"Оксана":  35,
		"Иван":    16,
		"Олег":    27}

	people["Светлана"] = 54

	for a, v := range people {
		fmt.Printf("Имя %s возраст %d\n", a, v)
	}
}
