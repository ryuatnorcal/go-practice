package main

import (
	"fmt"

	"example.com/bank_app/fileops" // when you import from custom package, import starts from go.mod's package name + folder name of your package
)

const accountBalanceFile = "balance.txt"
const defaultValue = 1000

func main() {
	var session bool = true
	// when you use custom package, use package name of your package as object
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile, defaultValue)
	if err != nil {
		fmt.Println("----- ERROR -----")
		fmt.Println(err)
		fmt.Println("-----------------")
		// developer mode failure message: comes with stack trace
		panic("Can't continue sorry")
	}

	for i := 0; session; i++ {
		presentOptions()
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
