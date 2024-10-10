package main

import "fmt"

func del(s string, a map[string]int) {
	delete(a, s)
}

func main() {
	var people = map[string]int{
		"Дмитрий": 53,
		"Оксана":  35,
		"Иван":    16,
		"Олег":    27}

	fmt.Println(people)
	var s string
	fmt.Println("Введите имя для удаления:")
	fmt.Scan(&s)

	del(s, people)
	fmt.Println(people)

}
