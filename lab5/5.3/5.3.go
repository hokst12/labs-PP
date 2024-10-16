package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius int
}

func (c circle) plosh() float32 {
	p := float32(c.radius) * float32(c.radius) * math.Pi
	return p
}

func main() {
	b := circle{radius: 12}

	fmt.Println(b.plosh())
}
