package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	fmt.Println(randomdata.PhoneNumber())

	var accountBalance, error = fileops.GetFloatFromFile(accountBalanceFile)
	if error != nil {
		fmt.Println("ERROR")
		fmt.Println(error)
		fmt.Println("--------------------------")
		return
	}
	fmt.Println("--------------------------------------------------")
	fmt.Println("Welcome to the Bank Application!")
	for {
		presentMenu()
		var choice int
		fmt.Print("Enter your choice (1-4): ")
		fmt.Scan(&choice)
		fmt.Println("You selected option:", choice)

		if choice == 1 {
			fmt.Printf("Your current balance is: $%.2f\n", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be positive!")
				continue
			}
			accountBalance += depositAmount
			fmt.Printf("Successfully deposited $%.2f. New balance is: $%.2f\n", depositAmount, accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		} else if choice == 3 {
			var withdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Withdrawal amount must be positive!")
				continue
			} else if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds!")
				continue
			} else {
				accountBalance -= withdrawAmount
				fmt.Printf("Successfully withdrew $%.2f. New balance is: $%.2f\n", withdrawAmount, accountBalance)
			}
		} else if choice == 4 {
			fmt.Println("Thank you for using the Bank Application. Goodbye!")
			return
		} else {
			fmt.Println("Invalid choice. Please select a valid option.")
			return
		}
	}

}
