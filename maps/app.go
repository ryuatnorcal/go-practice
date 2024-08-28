package main

import "fmt"

// maps are key-value pairs just like vector in C++

func main() {
	// map[keyType]valueType
	airline_websites := map[string]string{
		"delta_airline":   "https://www.delta.com",
		"air_canada":      "https://www.aircanada.com",
		"united_airlines": "https://www.united.com",
		"west_jet":        "https://www.westjet.com",
	}
	fmt.Println(airline_websites)
	// when you retrive data from map use [key] not .key like js
	fmt.Print("Delta Airline: ", airline_websites["delta_airline"])
	// adding new item to map
	airline_websites["air_newzealand"] = "https://www.airnewzealand.com"

	fmt.Println(airline_websites)

	// delete item from map
	delete(airline_websites, "air_canada")
	fmt.Println(airline_websites)
}
