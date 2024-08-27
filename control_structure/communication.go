package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func presentOptions() {
	fmt.Println("Welcome to Go BANK!")
	fmt.Println("Reach us:", randomdata.PhoneNumber())
	fmt.Println("What do you want to do? ")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")
}
