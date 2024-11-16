package main

import (
	"fmt"
	"sync"
	"time"
)

func adder1(ch chan int, mutex *sync.Mutex) {
	mutex.Lock()
	ch <- 1
	time.Sleep(1 * time.Second)
	mutex.Unlock()

}

func adder2(ch chan int, mutex *sync.Mutex) {
	mutex.Lock()
	ch <- 2
	time.Sleep(1 * time.Second)
	mutex.Unlock()

}

func adder3(ch chan int, mutex *sync.Mutex) {
	mutex.Lock()
	ch <- 3
	time.Sleep(1 * time.Second)
	mutex.Unlock()

}

func main() {

	var chInt chan int = make(chan int)
	var counter int = 0
	var mutex sync.Mutex

	for {
		go adder1(chInt, &mutex)
		go adder3(chInt, &mutex)
		go adder2(chInt, &mutex)

		a := <-chInt
		counter += a
		fmt.Printf("counter added %d now equals %d \n", a, counter)
	}

}
