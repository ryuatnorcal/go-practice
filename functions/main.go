package main

import "fmt"

// anomous function
func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	// anomous function: pass function itself as arg
	// anomous function args take the same as you defined at original function definiiton
	// use in one off task without logic in the function
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)

	fmt.Println(doubled)
	fmt.Println(tripled)

	fact := factorial(5)

	fmt.Println(fact)

	// variadic function
	vnum := []int{1, 2, 3, 4, 5}
	sum := sumup(1, 2, 3, 4, 5)
	// val... means destruct list into individual args
	another := sumup(vnum[0], vnum[1:]...)

	fmt.Println(sum)

	fmt.Println(another)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// closure: your factor is higher scope so you can use in nested function
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// factorial of 5 => 5 * 4 * 3 * 2 * 1 = 120
func factorial(number int) int {
	if number == 1 {
		return 1
	}

	// recursion function call aka call itself
	return number * factorial(number-1)
}

// variadic function
// you can take as many args as you want as long as type matches
func sumup(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, number := range numbers {
		sum += number
	}
	return sum
}
