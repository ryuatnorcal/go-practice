package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"
const defaultValue = 1000

func writeBalanceToFile(balance float64) {
	balancedText := fmt.Sprintf("%.2f", balance)
	// filename, byte, permission
	os.WriteFile(accountBalanceFile, []byte(balancedText), 0644)
}
func readBalanceFromFile() (float64, error) {
	file, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return defaultValue, errors.New("Failed to find file") // default value
	}
	balanceText := string(file)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return defaultValue, errors.New("Failed to parse file") // default value
	}
	return balance, nil
}
func main() {
	var session bool = true
	accountBalance, err := readBalanceFromFile()
	if err != nil {
		fmt.Println("----- ERROR -----")
		fmt.Println(err)
		fmt.Println("-----------------")
		// developer mode failure message: comes with stack trace
		panic("Can't continue sorry")
	}

	for i := 0; session; i++ {
		fmt.Println("Welcome to Go BANK!")
		fmt.Println("What do you want to do? ")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")
		var choice int

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		// wantsCheckBalance := choice == 1
		switch choice {
		case 1:
			fmt.Println("Check Balance")
			fmt.Printf("Your Balance is: %.2f\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Println("Deposit")
			fmt.Println("Enter amount to deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than zero")
				continue
			}
			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Your New Balance is: %.2f\n", accountBalance)
		case 3:
			fmt.Println("Withdraw")
			var withdrawAmount float64
			fmt.Println("Enter amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("withdraw amount must be greater than zero")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds")
				continue
			}
			accountBalance -= withdrawAmount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Your New Balance is: %.2f\n", accountBalance)
		default:
			fmt.Println("Goodbye!")
			session = false
		}
		// if choice == 1 {
		// 	fmt.Println("Check Balance")
		// 	fmt.Printf("Your Balance is: %.2f\n", accountBalance)
		// } else if choice == 2 {
		// 	var depositAmount float64
		// 	fmt.Println("Deposit")
		// 	fmt.Println("Enter amount to deposit: ")
		// 	fmt.Scan(&depositAmount)
		// 	if depositAmount <= 0 {
		// 		fmt.Println("Deposit amount must be greater than zero")
		// 		continue
		// 	}
		// 	accountBalance += depositAmount
		// 	fmt.Printf("Your New Balance is: %.2f\n", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Println("Withdraw")
		// 	var withdrawAmount float64
		// 	fmt.Println("Enter amount to withdraw: ")
		// 	fmt.Scan(&withdrawAmount)
		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("withdraw amount must be greater than zero")
		// 		continue
		// 	}
		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("Insufficient funds")
		// 		continue
		// 	}
		// 	accountBalance -= withdrawAmount
		// 	fmt.Printf("Your New Balance is: %.2f\n", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	session = false

		// }
	}

}
