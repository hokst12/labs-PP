package main

import "fmt"

func main() {
	var a [6]string = [6]string{"kom", "ghfks", "satlkrwjk", "gs", "gsghh", "wqtgx"}
	var b, c int

	fmt.Println("Введите начало и конец среза")
	fmt.Scan(&b, &c)

	srez := a[b:c]
	fmt.Println(srez)
	var max int = 0
	var el string

	for _, value := range srez {
		if len(value) > max {
			max = len(value)
			el = value
		}

	}
	fmt.Println(el)
}
