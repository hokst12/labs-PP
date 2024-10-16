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

	a := rectangle{length: 11, width: 32}
	b := circle{radius: 14}

	fmt.Printf("Площадь квадрата: %d , площадь круга: %d", a.area(), b.area())

}
