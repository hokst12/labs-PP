package main

import (
	"fmt"
	"strings"
	"sync"
)

type Task struct {
	Original string
	Result   string
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func worker(id int, tasks <-chan Task, results chan<- Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		reversed := reverseString(task.Original)
		results <- Task{Original: task.Original, Result: reversed}
		fmt.Printf("Worker %d processed: %s -> %s\n", id, task.Original, reversed)
	}
}

func main() {
	var numWorkers int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&numWorkers)

	tasks := make(chan Task, 100)
	results := make(chan Task, 100)

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("Введите строки для реверсирования (введите 'exit' для выхода):")
	for {
		var input string
		fmt.Print("> ")
		fmt.Scanln(&input)

		if strings.ToLower(input) == "exit" {
			break
		}

		tasks <- Task{Original: input}
	}

	close(tasks)

	for result := range results {
		fmt.Printf("Результат: %s -> %s\n", result.Original, result.Result)
	}
}
