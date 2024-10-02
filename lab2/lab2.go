package main

import "fmt"

func opred1(a int) {
	if a == 0 {
		fmt.Println("Число рано нулю")
	} else if a%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}
}

func opred2(a int) string {
	if a == 0 {
		return ("Zero")
	} else if a > 0 {
		return ("Positive")
	} else {
		return ("Negative")
	}
}

func length(a string) int {
	return len(a)
}

type rectangle struct {
	length int
	width  int
}

func (p rectangle) plosh() int {
	return p.length * p.width
}

func sred(a, b int) int {
	return (a + b) / 2
}

func main() {
	opred1(0)
	opred1(4)
	opred1(9)
	opred1(112)

	fmt.Println(opred2(11), opred2(-2), opred2(0))

	for i := 1; i < 11; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	a := "Kobachok"
	fmt.Println(length(a))

	var b = rectangle{5, 10}
	fmt.Println(b.plosh())

	c := 11
	d := 24

	fmt.Println(sred(c, d))
}
