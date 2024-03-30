// every function must be part of a package so it can be used across the given project
package main

import (
	"fmt"
)

func presentOptions() {
	fmt.Println("Choose a service: ")
	fmt.Println("1. Check my balance")
	fmt.Println("2. Make a deposit")
	fmt.Println("3. Withdraw funds")
	fmt.Println("4. Exit")
}