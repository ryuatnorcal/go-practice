package main

import "fmt"

type transform func(int) int

// write generic function first then write specific function pass as args
// keep Dry code
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// without () means pass function without execution
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformsFunction(&numbers)
	transformedNumber := transformNumbers(&doubled, transformerFn1)

	fmt.Println(transformedNumber)

}

// you can path functions as arg; type definition is func(arg_type) return_type
func transformNumbers(numbers *[]int, transform transform) []int {
	var doubledNumbers []int
	for _, number := range *numbers {
		// you can pass functions as arg
		doubledNumbers = append(doubledNumbers, transform(number))
	}
	return doubledNumbers
}

// func doubleNumbers(numbers *[]int) []int {
// 	var doubledNumbers []int
// 	for _, number := range *numbers {
// 		// you can pass functions as arg
// 		doubledNumbers = append(doubledNumbers, double(number))
// 	}
// 	return doubledNumbers
// }

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// you can return function as a value
func getTransformsFunction(numbers *[]int) transform {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}
