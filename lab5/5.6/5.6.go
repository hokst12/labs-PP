package main

import "fmt"

type stringer interface {
	showinfo()
}

type book struct {
	name   string
	author string
	year   int
}

func (b book) showinfo() {
	fmt.Printf("Автор книги: %s,  название книги: %s, год написания книги: %d \n", b.author, b.name, b.year)
}

func main() {
	b := book{name: "Flowers for elgernon", author: "Daniel Keyes", year: 1966}
	b.showinfo()
}
