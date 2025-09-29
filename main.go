package main 

import (
	"fmt"
)

const USD_TO_EUR = 0.83
const USD_TO_RUB = 83
const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR

func main() {
	readUserInput()

}

func readUserInput() string {
	var userInput string
	fmt.Print("Enter number: ")
	fmt.Scanln(&userInput)
	
	return userInput
}

func convertMoney(number float32, currencyFrom, currencyTo string) {

}