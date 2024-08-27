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

<<<<<<<<<<<<<<  ✨ Codeium Command ⭐  >>>>>>>>>>>>>>>>
// muateAgeWithAdultYears subtract 18 from the value that the pointer point to
// the result is stored directly into the address location, so the value is mutated
// this is more efficient than returning the result and then having to assign it back
// to the original variable.
<<<<<<<  d7c5cfa7-caa1-459d-8548-6f0b1f9062b2  >>>>>>>
func muateAgeWithAdultYears(age *int) {
	// return *age - 18 // dereferencing the pointer to get the value
	*age = *age - 18 // mutating the value directly into the address location
}
