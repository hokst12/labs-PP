package main

import "fmt"

func opred2(a int) string {
	if a == 0 {
		return ("Zero")
	} else if a > 0 {
		return ("Positive")
	} else {
		return ("Negative")
	}
}

func main() {
	fmt.Println(opred2(11), opred2(-2), opred2(0))
}
