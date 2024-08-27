package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct { // its called struct literal
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}
type Admin struct {
	email    string
	password string
	User     // you can embed User struct in Admin struct
}

/*******************************************************
* CONSTRUCTORS
*******************************************************/
func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, Last name and Birthdate are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(firstName, lastName, birthdate, email, password string) (*Admin, error) {
	user, err := NewUser(firstName, lastName, birthdate)
	if err != nil {
		return nil, err
	}
	return &Admin{
		email:    email,
		password: password,
		User:     *user,
	}, nil
}

/*******************************************************
* METHODS
*******************************************************/
// it overkill to use pointer here but practice purpose
// func outputUserDetails(user *User) {
// Reciver argument: (User) functionName  attach to the struct if this case you don't need parameters
func (user User) OutputUserDetails() {
	// referencing struct is the same as js object
	fmt.Println(user.firstName, user.lastName, user.birthdate)
}

// to mutation data in struct use pointer
func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}
