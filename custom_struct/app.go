package main

import "fmt"

type customString string // type alias

func (text customString) log() {
	fmt.Println(text)
}
func main() {
	var name customString = "John"

	name.log()
}
