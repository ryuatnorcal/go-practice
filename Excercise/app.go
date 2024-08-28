package main

import (
	"fmt"

	"example.com/excercise/product"
)

// Time to practice what you learned!

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	var hobbies []string = []string{"MTB", "SKI", "SNB"}
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println("first hobby: ", hobbies[0])
	fmt.Println("rest of hobbies: ", hobbies[1:3])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	var top2 []string = hobbies[0:2]
	fmt.Println("first and second", top2)

	var top2slice []string = []string{hobbies[0], hobbies[1]}
	fmt.Println("first and second", top2slice)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	var top2slice2 []string = top2slice[1:]
	top2slice2 = append(top2slice2, hobbies[2])
	fmt.Println("first and thrid", top2slice2)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	var dynamicGoals []string = []string{"finish course", "learn go"}
	fmt.Println("dynamic goals: ", dynamicGoals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	dynamicGoals[1] = "new goal: become pro at go"
	dynamicGoals = append(dynamicGoals, "additional goal: become fullstack dev with react go and solidity")
	fmt.Println("dynamic goals: ", dynamicGoals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	var products []product.Product
	product1, err := product.NewProduct("Line Ski: Vision 108", 599.99)
	if err != nil {
		fmt.Println(err)
		return
	}
	product2, err := product.NewProduct("Fox 36 Suspension", 1100.99)
	if err != nil {
		fmt.Println(err)
		return
	}
	product3, err := product.NewProduct("Union Binding", 199.99)
	if err != nil {
		fmt.Println(err)
		return
	}
	products = append(products, product1)
	products = append(products, product2)
	products = append(products, product3)
	fmt.Println(products)
}
