package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (user User) outputUserData() {

	fmt.Println("User Data:")
	fmt.Println("First Name:", user.firstName)
	fmt.Println("Last Name:", user.lastName)
	fmt.Println("Birth Date:", user.birthDate)
	fmt.Println("Account Created At:", user.createdAt.Format("2006-01-02 15:04:05"))
}
func (user *User) clearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func newUser(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		fmt.Println("Error: All fields must be filled out.")
		return nil, errors.New("all fields must be filled out")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	fisrtName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthDate := getUserData("Enter your birth date(MM//DD/YYYY): ")

	var appUser *User
	appUser, err := newUser(fisrtName, lastName, birthDate)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	appUser.outputUserData()
	appUser.clearUserName()
}

func getUserData(promptText string) string {

	fmt.Println(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
