package user

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
type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password, firstName, lastName, birthDate string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password must be provided")
	}
	user, err := New(firstName, lastName, birthDate)
	if err != nil {
		return nil, err
	}
	return &Admin{
		email:    email,
		password: password,
		User:     *user,
	}, nil
}

func (admin Admin) OutputAdminData() {
	fmt.Println("Admin Data:")
	fmt.Println("Email:", admin.email)
	fmt.Println("Password:", admin.password)
	fmt.Println("First Name:", admin.firstName)
	fmt.Println("Last Name:", admin.lastName)
	fmt.Println("Birth Date:", admin.birthDate)
	fmt.Println("Account Created At:", admin.createdAt.Format("2006-01-02 15:04:05"))
}
func (user User) OutputUserData() {

	fmt.Println("User Data:")
	fmt.Println("First Name:", user.firstName)
	fmt.Println("Last Name:", user.lastName)
	fmt.Println("Birth Date:", user.birthDate)
	fmt.Println("Account Created At:", user.createdAt.Format("2006-01-02 15:04:05"))
}
func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func New(firstName, lastName, birthDate string) (*User, error) {

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
