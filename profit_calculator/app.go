package main

/**
I want you to build a profit calculator, a command line tool built with Go, of course,
that asks users for 	, expenses, and tax rate input, and that then calculates earnings before tax
and earnings after tax, as well as the ratio between these two values
and that then outputs these three values.

So that we got such a profit calculator that we can use
to calculate any profit based on these three input values.
**/
import (
	"errors"
	"fmt"
	"os"
)

const earningsBeforeTaxFile = "earningsBeforeTax.txt"
const earningsAfterTaxFile = "earningsAfterTax.txt"
const profitBalanceFile = "profit.txt"

func writeProfitToFile(result float64, t string) {
	profitText := fmt.Sprintf("%.2f", result)
	switch t {
	case "before":
		os.WriteFile(earningsBeforeTaxFile, []byte(profitText), 0644)
	case "after":
		os.WriteFile(earningsAfterTaxFile, []byte(profitText), 0644)
	case "profit":
		os.WriteFile(profitBalanceFile, []byte(profitText), 0644)
	}
}
func main() {
	var revenue, expenses, taxRate float64
	revenue, revErr := getUserInput("Enter revenue: ")
	if revErr != nil {
		fmt.Println("Revenue Input Error: ", revErr)
		return
	}
	expenses, expErr := getUserInput("Enter expenses: ")
	if expErr != nil {
		fmt.Println("Expenses Input Error: ", expErr)
		return
	}
	taxRate, taxErr := getUserInput("Enter tax rate: ")
	if taxErr != nil {
		fmt.Println("Tax Rate Input Error: ", taxErr)
		return
	}

	earningsBeforeTax, earningsAfterTax, profit := calculateErning(revenue, expenses, taxRate)
	writeProfitToFile(earningsBeforeTax, "before")
	writeProfitToFile(earningsAfterTax, "after")
	writeProfitToFile(profit, "profit")
	// earningsAfterTax := calculateAfterTax(earningsBeforeTax, taxRate)
	// profit := earningsAfterTax - expenses
	fmt.Printf("Earnings before tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Earnings after tax: %.2f\n", earningsAfterTax)
	// fmt.Printf("Profit: %.2f\n", profit)

	// Sprintf is used to format a string and store it in a variable
	formattedFV := fmt.Sprintf("Profit: %.2f\n", profit)
	formattedRFV := fmt.Sprintf("Earnings after tax: %.2f\n", earningsAfterTax)

	// Print doesn't allow formamting
	fmt.Print(formattedFV, formattedRFV)
}

func getUserInput(prompt string) (float64, error) {
	var input float64
	fmt.Print(prompt)
	fmt.Scan(&input)
	if input <= 0 {
		fmt.Println("Invalid input: input should be greater than 0")
		return input, errors.New("Invalid input: input should be greater than 0")
	}
	return input, nil
}

// func outputText(text string) {
// 	fmt.Print(text)
// }

func calculateErning(revenue, expenses, taxRate float64) (earningsBeforeTax float64, earningsAfterTax float64, profit float64) {
	earningsBeforeTax = revenue - expenses
	earningsAfterTax = earningsBeforeTax * (1 - taxRate/100)
	profit = earningsAfterTax - expenses
	return earningsBeforeTax, earningsAfterTax, profit
}

// func calculateEarningsBeforeTax(revenue, expenses float64) float64 {
// 	return revenue - expenses
// }

// func calculateAfterTax(earningsBeforeTax, taxRate float64) float64 {
// 	return earningsBeforeTax * (1 - taxRate/100)
// }
