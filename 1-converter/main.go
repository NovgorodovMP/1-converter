package main

import (
	"fmt"
	"strconv"
)

const USD_TO_EUR float64 = 0.83
const USD_TO_RUB float64 = 83
const EUR_TO_RUB float64 = USD_TO_RUB / USD_TO_EUR

func main() {
	for {
		
		var currencyFrom, currencyTo string
		var number, resultNumber float64
		fmt.Print("Enter first currency (usd/eur/rub): ")
		currencyFrom = inputCurrency()
		number = inputNumber()
		fmt.Print("Enter target currency: ")
		for {
			currencyTo = inputCurrency()
			if currencyTo != currencyFrom {
				break
			} else {
				fmt.Print("I don't know this currency or it is first currency. Enter again")
			}
		}
		resultNumber = convertMoney(number, currencyFrom, currencyTo)
		fmt.Println(resultNumber)
	}

}

func readUserInput() string {
	var userInput string
	fmt.Scanln(&userInput)
	return userInput
}

func inputNumber() float64 {
	fmt.Print("Enter value: ")
	for {
		number, err := strconv.ParseFloat(readUserInput(), 64)
		if err != nil {
			fmt.Println("You entered a string. Enter a number, please")
		} else {
			return number
		}
	}
}

func inputCurrency() string{
	for {
		currencyFrom := readUserInput()
		if isItCurrency(currencyFrom) {
			return currencyFrom
		} else {
			fmt.Print("I don't know this currency. Enter usd or eur or rub")
		}
	}
}
func isItCurrency (curr string) bool {
	return (curr == "eur" || curr == "usd" || curr == "rub")
}

func convertMoney(number float64, currencyFrom, currencyTo string) float64 {
	var resultNumber float64
	switch {
	case currencyFrom == "usd" && currencyTo == "eur":
		resultNumber = number * USD_TO_EUR
	case currencyFrom == "usd" && currencyTo == "rub":
		resultNumber = number * USD_TO_RUB
	case currencyFrom == "eur" && currencyTo == "usd":
		resultNumber = number / USD_TO_EUR
	case currencyFrom == "eur" && currencyTo == "rub":
		resultNumber = number * EUR_TO_RUB
	case currencyFrom == "rub" && currencyTo == "usd":
		resultNumber = number / USD_TO_RUB
	case currencyFrom == "rub" && currencyTo == "eur":
		resultNumber = number / EUR_TO_RUB
	}
	return resultNumber
}