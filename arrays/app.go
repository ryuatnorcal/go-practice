package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

/**
* Static array
 */
// func main() {

// 	fmt.Println("----- Array -----")
// 	// if you creating strict length array use [length]
// 	prices := [5]float64{1.1, 2.2, 3.3, 4.4, 5.5}
// 	fmt.Println(prices)

// 	// Slice prices
// 	// python way of spliting
// 	// prices[0:3] => [startIndex: endIndex]
// 	// startIndex: inclusive
// 	// endIndex: exclusive
// 	fmt.Println("----- Slicing -----")
// 	featuredPrices := prices[1:3]
// 	fmt.Println(featuredPrices)
// 	// start from 0 to 2
// 	limitedLast := prices[:3]
// 	fmt.Println(limitedLast)

// 	// you cannot use -1 to reverse like python
// 	// select upper bound
// 	upperlimit := prices[2:]

// 	fmt.Println(upperlimit)
// 	fmt.Println("----- Length -----")
// 	// length and capacity of array
// 	// cap shows non selected length
// 	fmt.Println(len(prices), cap(prices))
// }

/**
* Dynamic array
 */
func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[1])

	// append new item and return new slice
	// since array in go is fixed length, append create new array with new element
	updatedPrices := append(prices, 7.99)
	fmt.Println(updatedPrices, prices)

	// this will overwirite original array
	prices = append(prices, 7.99)
	fmt.Println(prices)

	// remove item from slice
	prices = prices[1:]

	fmt.Println(prices)

	// add multiple items to slice
	// use ... to add multiple items to end of original array
	// just like deconstruction in js
	discountPrices := []float64{100.99, 800.99, 200.99}
	prices = append(prices, discountPrices...)

	fmt.Println(prices)
}
