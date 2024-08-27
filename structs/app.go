package main

import (
	"fmt"

	user "example.com/struct/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	// var appUser = User{ // you cannot put space between type and {}
	// 	FirstName: firstName,
	// 	LastName:  lastName,
	// 	Birthdate: birthdate,
	// 	createdAt: time.Now(),
	// }
	// you could refactor ^ as below
	// order matter!
	// appUser := User{
	// 	firstName,
	// 	lastName,
	// 	birthdate,
	// 	time.Now(),
	// }
	// you can define empty struct as below
	// var appUser = User{}

	// you could insert partial data
	// Birthday will be nil
	// appUser := User{
	// 	FirstName: firstName,
	// 	LastName:  lastName,
	// 	createdAt: time.Now(),
	// }

	// use constructor
	appUser, err := user.NewUser(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	admin, err := user.NewAdmin(firstName, lastName, birthdate, "admin@admin.com", "admin")
	admin.OutputUserDetails()
	// use reference to call the method
	// outputUserDetails(&appUser)
	// use methods as part of struct as below
	appUser.OutputUserDetails()
	// fmt.Println(firstName, lastName, birthdate)
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// Scan doesn't allow empty break
	fmt.Scanln(&value)
	return value
}
