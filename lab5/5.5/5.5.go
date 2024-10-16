package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() int
}

type rectangle struct {
	length int
	width  int
}

type circle struct {
	radius int
}

func (c circle) area() int {
	p := float64(c.radius) * float64(c.radius) * float64(math.Pi)
	return int(math.Round(p))
}

func (r rectangle) area() int {
	p := r.width * r.length
	return p
}

func main() {
	var b = []shape{rectangle{width: 11, length: 43}, circle{radius: 5}, rectangle{width: 9, length: 6}, circle{radius: 25}}

	var i int = 1

	for _, a := range b {
		fmt.Printf("Площадь фигуры номер %d равна %d\n", i, a.area())
		i += 1
	}
}
