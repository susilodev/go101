package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const balanceTextLocation = "balance.txt"

func main() {
	balance, err := fileops.GetFloatFromFile(balanceTextLocation)
	if err != nil {
		fmt.Print("Errors: ")
		fmt.Println(err)
		fmt.Println("----------")
	}

	fmt.Println("Welcome to the Salaf Bank")
	fmt.Println("What do you want ?")

	var choice int

	for {
		showOptions()
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
			fileops.WriteFloatToFile(balance, balanceTextLocation)

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
			fileops.WriteFloatToFile(balance, balanceTextLocation)
			continue
		} else {
			fmt.Println("Thank You, Bye!")
			break
		}
	}
}
