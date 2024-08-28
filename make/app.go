package main

import "fmt"

// type alias
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}
func main() {
	// you can preallocate array with make
	// make([]type, initial_length, capacity)
	// if you know capacity, make can reduce memory usage by preallocating the array
	userName := make([]string, 2, 5)

	userName[0] = "John wick"

	userName = append(userName, "John")
	userName = append(userName, "Jane")
	fmt.Println(userName)

	// you can preallocate map
	// make(map[keyType]valueType, capacity)
	// coursesRating := make(map[string]float64, 3)

	// use type alias to shorten the type
	coursesRating := make(floatMap, 3)
	coursesRating["math"] = 4.5
	coursesRating["english"] = 4.0
	coursesRating["history"] = 3.5
	fmt.Println(coursesRating)
	// reallocate memory now to add it
	coursesRating["science"] = 4.5

	// fmt.Println(coursesRating)
	coursesRating.output()

	// for loop to iterate over array
	for i := 0; i < len(userName); i++ {
		fmt.Println(userName[i])
	}

	// you can get index value combination too
	for index, value := range userName {
		fmt.Println(index, value)
	}

	// for loop to iterate over map
	fmt.Println("Key | Value")
	for key, value := range coursesRating {
		fmt.Println(key, value)
	}
}
