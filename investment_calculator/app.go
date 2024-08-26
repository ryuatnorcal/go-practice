package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	// you can assign multiple variables at once with a type
	var investmentAmount, years float64 = 1000, 10

	// := use default type and assign value
	expectedReturnRate := 5.5
	// var years float64 = 10
	fmt.Print("Enter amount to invest (defalut: 1000): ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter number of years (default: 10): ")
	fmt.Scan(&years)

	fmt.Print("Enter expected return rate (default: 5.5): ")
	fmt.Scan(&expectedReturnRate)

	var futureValue = investmentAmount * math.Pow(1+(expectedReturnRate/100), years)
	futureRealVlaue := futureValue / math.Pow(1+(inflationRate/100), years)

	fmt.Println(futureValue)
	fmt.Println(futureRealVlaue)
}
