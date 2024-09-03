package main

import (
	"fmt"

	"example.com/price_calculator/cmdmanager"
	"example.com/price_calculator/prices"
)

func main() {

	pricesa := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	doneChannels := make([]chan bool, len(taxRates))
	errChannels := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChannels[index] = make(chan bool)
		errChannels[index] = make(chan error)
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result%0.f.json", taxRate*100))
		cmdm := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(cmdm, pricesa, taxRate)
		// err := job.Process()
		go job.Process(doneChannels[index], errChannels[index])
		// if err != nil {
		// 	fmt.Println("Could not process job", err)
		// }

		// fmt.Println("Done!")
	}

	// error channel is most of case not coming back
	// so use select
	for index, _ := range taxRates {
		select {
		case err := <-errChannels[index]:
			fmt.Println(err)
		case done := <-doneChannels[index]:
			if done {
				fmt.Println("Done!")
			}
		}
	}

	// for _, doneChannel := range doneChannels {
	// 	<-doneChannel
	// }
}
