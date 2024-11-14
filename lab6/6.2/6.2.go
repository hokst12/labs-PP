package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sender(ch chan int) {
	for {
		ch <- rand.Intn(100-1) + 1
		time.Sleep(2 * time.Second)
	}
}

func reciever(ch chan int) {
	var a int
	for {
		if val, opened := <-ch; opened {
			a = val
			fmt.Println(a)
		} else {
			fmt.Println("Канал закрыт")
			break
		}
	}
}

func main() {

	var ch chan int = make(chan int)

	go sender(ch)
	go reciever(ch)
	time.Sleep(9 * time.Second)
	close(ch)
	fmt.Scanln()
}
