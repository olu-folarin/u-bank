package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

// declare and initialize the account balance file
const accountBalance = "balance.txt"

// read from the account balance file
// added error as a type so i could use the errors package to generate custom error messages
func readBalance() (float64, error) {
	data, err := os.ReadFile(accountBalance)

	// error handling
	if err != nil {
		return 700, errors.New("accountBalance file could not be found.")
	}

	statementTxt := string(data) // convert the data that's converted to byte using writetofile to string format
	accBalance, err := strconv.ParseFloat(statementTxt, 64) // convert the string data back to float64
	// error handling for parsefloat as it could also generate an error
	if err != nil {
		return 700, errors.New("Could not parse stored balance value to float.")
	}

	// add nil to the return to show there's no error
	return accBalance, nil
}

// persist the account balance to a file
func persistAccountBalance(balance float64) {
	statementTxt := fmt. Sprint(balance)
	os.WriteFile(accountBalance, []byte(statementTxt), 644)
}

func main() {
	balance, err := readBalance()
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
			persistAccountBalance(balance)
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
			persistAccountBalance(balance)
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