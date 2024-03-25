package main

import (
	"fmt"
)

func main() {
	balance := 1000.00
	var deposit, withdraw float64

	fmt.Println("Welcome to U Bank.")

	// run the app for as long as possible until the user selects option 4
	for {
	fmt.Println("Choose a service: ")
	fmt.Println("1. Check my balance")
	fmt.Println("2. Make a deposit")
	fmt.Println("3. Withdraw funds")
	fmt.Println("4. Exit")

	// initialize and collect user option
	var option int
	fmt.Print("Option: ")
	fmt.Scan(&option)

	// print the chosen option
	fmt.Println("You chose: ", option)

	if option == 1 {
		fmt.Printf("Your balance is $%v\n", balance)
	} else if option == 2 {
		fmt.Print("Deposit Amount: ")
		fmt.Scan(&deposit)

		// check if deposit amount is greater than 0
		if deposit <= 0 {
			fmt.Println("Can't complete this operation. Deposit amount must be greater than 0.")
		} else {
			balance += deposit
			fmt.Printf("Your new balance is $%.2f.\n", balance)
		}
	} else if option == 3 {
		fmt.Print("Withdrawal Amount: ")
		fmt.Scan(&withdraw)

		// check if the withdrawal amount is less than the available balance or is less than 0
		if withdraw >= balance || withdraw <= 0 {
			fmt.Println("Your withdrawal amount must be equal to or less than your available balance.")
		} else {
			balance -= withdraw
			fmt.Printf("Your new balance is $%.2f\n", balance)
		}
	} else {
		fmt.Println("Thanks for banking with us.")
		// break out of the app and execute the code outside the loop
		break
	}
	}
	
	fmt.Println("We hope to see you soon!")
}