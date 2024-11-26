package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceText = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceText)
	if err != nil {
		return 0.0, errors.New("Errors ReadFile")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0.0, errors.New("failed parse text")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) (float64 error) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceText, []byte(balanceText), 0644)
	return
}

func main() {
	balance, err := getBalanceFromFile()
	if err != nil {
		fmt.Print("Errors: ")
		fmt.Println(err)
		fmt.Println("----------")
	}

	fmt.Println("Welcome to the Salaf Bank")
	fmt.Println("What do you want ?")

	var choice int

	for {
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. WithDraw Money")
		fmt.Println("4. Exit")

		fmt.Print("Enter Your choice: ")
		fmt.Scan(&choice)
		fmt.Println("Your Choice is: ", choice)

		if choice == 1 {
			// show balance
			fmt.Println("Your balance is ", balance)
			continue
		}

		if choice == 2 {
			// doing Deposit
			var depositAmount float64
			fmt.Print("Your Deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount < 0 {
				fmt.Println("invalid input")
				continue
			}

			balance += depositAmount
			fmt.Println("Your Total Balance is: ", balance)
			writeBalanceToFile(balance)

			continue
		}

		if choice == 3 {
			// doing withDraw
			var withDraw float64
			fmt.Print("Type your WithDraw: ")
			fmt.Scan(&withDraw)

			if withDraw > balance {
				fmt.Println("invalid withDraw")
				continue
			}

			balance -= withDraw
			fmt.Println("Your Total Balance is: ", balance)
			writeBalanceToFile(balance)
			continue
		} else {
			fmt.Println("Thank You, Bye!")
			break
		}
	}
}
