package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type CalculationRequest struct {
	Operation string
	Operand1  float64
	Operand2  float64
}

type CalculationResponse struct {
	Result float64
	Error  error
}

func calculator(requests <-chan CalculationRequest, results chan<- CalculationResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	for req := range requests {
		var result float64
		var err error

		switch req.Operation {
		case "+":
			result = req.Operand1 + req.Operand2
		case "-":
			result = req.Operand1 - req.Operand2
		case "*":
			result = req.Operand1 * req.Operand2
		case "/":
			if req.Operand2 == 0 {
				err = fmt.Errorf("division by zero")
			} else {
				result = req.Operand1 / req.Operand2
			}
		default:
			err = fmt.Errorf("unknown operation: %s", req.Operation)
		}

		results <- CalculationResponse{Result: result, Error: err}
	}
}

func main() {
	requests := make(chan CalculationRequest)
	results := make(chan CalculationResponse)

	var wg sync.WaitGroup
	wg.Add(1)

	go calculator(requests, results, &wg)

	go func() {
		for res := range results {
			if res.Error != nil {
				fmt.Println("Error:", res.Error)
			} else {
				fmt.Println("Result:", res.Result)
			}
		}
	}()

	fmt.Println("Введите команды в формате: операция число1 число2 (например: + 10 5). Введите 'exit' для выхода.")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Неверный формат. Пожалуйста, используйте: операция число1 число2")
			continue
		}

		operation := parts[0]
		operand1, err1 := strconv.ParseFloat(parts[1], 64)
		operand2, err2 := strconv.ParseFloat(parts[2], 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Неверные операнды. Пожалуйста, введите числа.")
			continue
		}

		requests <- CalculationRequest{Operation: operation, Operand1: operand1, Operand2: operand2}
	}

	close(requests)
	wg.Wait()
	close(results)
}
