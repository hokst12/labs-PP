package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func (p rectangle) plosh() int {
	return p.length * p.width
}

func main() {
	var b = rectangle{5, 10}
	fmt.Println(b.plosh())
}
