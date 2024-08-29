package main

import (
	"fmt"

	"example.com/price_calculator/cmdmanager"
	"example.com/price_calculator/prices"
)

func main() {

	pricesa := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result%0.f.json", taxRate*100))
		cmdm := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(cmdm, pricesa, taxRate)
		err := job.Process()
		if err != nil {
			fmt.Println("Could not process job", err)
		}

		fmt.Println("Done!")
	}

}
