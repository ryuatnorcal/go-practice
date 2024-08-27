package main

import "fmt"

func main() {
	// using pointer reduce amount of copy of data
	// performance increased!
	age := 32
	agePointer := &age
	fmt.Println("Age:", age)
	fmt.Println("Age Address:", agePointer)
	fmt.Println("Get Value From address", *agePointer) // dereference pointer
	// fmt.Println("Adult years:", getAdultYears(agePointer)) // passing adress with &
	muateAgeWithAdultYears(agePointer)
	fmt.Println("Adult years:", age)
}

func muateAgeWithAdultYears(age *int) {
	// return *age - 18 // dereferencing the pointer to get the value
	*age = *age - 18 // mutating the value directly into the address location
}
