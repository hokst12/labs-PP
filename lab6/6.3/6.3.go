package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sender1(ch chan int) {
	for {
		ch <- rand.Intn(100-1) + 1
		time.Sleep(2 * time.Second)
	}
}

func sender2(ch chan int, ch1 chan string) {
	for {
		a := <-ch
		fmt.Println("Пришло число ", a)
		opred1(a, ch1)
		time.Sleep(3 * time.Second)
	}
}

func opred1(a int, ch chan string) {
	if a == 0 {
		ch <- ("Число равно нулю")
	} else if a%2 == 0 {
		ch <- ("Число чётное")
	} else {
		ch <- ("Число нечётное")
	}
}

func main() {

	var ch1 chan int = make(chan int)
	var ch2 chan string = make(chan string)
	a := 0

	go sender1(ch1)
	go sender2(ch1, ch2)
	for {
		select {
		case res := <-ch1:
			a = res
			fmt.Println("Пришло число ", a)
		case res := <-ch2:
			fmt.Println("Пришло сообщение: ", res)

		}
	}
}
