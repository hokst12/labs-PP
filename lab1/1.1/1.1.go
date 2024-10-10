package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Год:", now.Year(), " Месяц:", now.Month(), " День:", now.Day(), " Час:", now.Hour(), " Минута:", now.Minute())
}
