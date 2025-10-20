package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var numbersSlice []float64
	var err error
	var result float64
	var operation string

	EnterOperation:
	for {
		fmt.Print("Enter operation (avg/sum/med): ")
		operation, err = readUserInput()
		if err != nil {
			continue
		}
		operation = strings.ToLower(operation)
		if !isItCorrectOperation(operation) {
			fmt.Println("Wrong operation. Try again")
			continue
		}
		break EnterOperation
	}

	EnterNumbers:
	for {
		fmt.Print("Enter number row separated by comma: ")
		numberRow, err := readUserInput()
		if err != nil {
			continue
		}
		numbersSlice, err = splitInput(numberRow)
		if err != nil {
			continue
		}
		break EnterNumbers
	}

	switch {
	case operation == "avg":
		result = calcAverage(numbersSlice)
	case operation == "sum":
		result = calcSum(numbersSlice)
	case operation == "med":
		result = calcMed(numbersSlice)
	}
	fmt.Printf("Operation %s = %.2f\n", operation, result)

}

func isItCorrectOperation (operation string) bool {
	return (operation == "avg" || operation == "sum" || operation == "med")
}

func readUserInput() (string, error) {
	var err error = nil
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while user input")
	}
	return strings.TrimSpace(userInput), err
}

func splitInput(input string) ([]float64, error) {
	result := make([]float64, 0)
	for _, numString := range strings.Split(input, ",") {
		number, err := strconv.ParseFloat(strings.TrimSpace(numString), 64)
		if err != nil {
			return nil, errors.New("you entered numbers by wrong format. Enter a numbers in correct way, please")
		}
		result = append(result, number)
	}
	return result, nil
}

func calcAverage(numbers []float64) float64 {
	return calcSum(numbers) / float64(len(numbers))
}

func calcSum(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	} 
	return sum
}

func calcMed(numbers []float64) float64 {
	sort.Float64s(numbers)
	if len(numbers) % 2 != 0 {
		return numbers[len(numbers) / 2]
	} else {
		return (numbers[len(numbers) / 2] + numbers[(len(numbers) / 2) + 1]) / 2
	}

}