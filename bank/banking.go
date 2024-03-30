package main

import (
	"fmt"
	"bank.io/u-bank/docops"
)

// declare and initialize the account balance file
const accountBalance = "balance.txt"

func main() {
	balance, err := docops.RetrieveFloatFromFile(accountBalance)
	// since the readBalance function is used here, you need to use the err package for error handling here too
	if err != nil {
		fmt.Println("Something isn't right.")
		fmt.Println(err)
		fmt.Println("<------------------------->")
	}

	var deposit, withdraw float64

	fmt.Println("Welcome to U Bank.")

	// run the app for as long as possible until the user selects option 4
	// to avoid using so many if and else ifs, a switch comes in handy
	for {
	presentOptions()

	// initialize and collect user option
	var option int
	fmt.Print("Option: ")
	fmt.Scan(&option)

	// print the chosen option
	fmt.Println("You chose: ", option)

	switch option {
	case 1:
		fmt.Printf("Your balance is $%v\n", balance)

	case 2:
		fmt.Print("Deposit Amount: ")
		fmt.Scan(&deposit)

		// check if deposit amount is greater than 0
		if deposit <= 0 {
			fmt.Println("Can't complete this operation. Deposit amount must be greater than 0.")
			continue
		} else {
			balance += deposit
			fmt.Printf("Your new balance is $%.2f.\n", balance)
			docops.PersistFloatToDoc(balance, accountBalance)
		}

	case 3:
		fmt.Print("Withdrawal Amount: ")
		fmt.Scan(&withdraw)

		// check if the withdrawal amount is less than the available balance or is less than 0
		if withdraw >= balance || withdraw <= 0 {
			fmt.Println("Your withdrawal amount must be equal to or less than your available balance.")
			continue
		} else {
			balance -= withdraw
			fmt.Printf("Your new balance is $%.2f\n", balance)
			docops.PersistFloatToDoc(balance, accountBalance)
		}

	default:
		fmt.Println("Thanks for banking with us.")
		fmt.Println("We hope to see you soon!")
		// break out of the app and execute the code outside the loop
		// break
		return
	}
	}
}