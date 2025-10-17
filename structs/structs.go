package main

import (
	"fmt"

	"example.com/structs/user"
)

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {
	add(1, 2)

	fisrtName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthDate := getUserData("Enter your birth date(MM//DD/YYYY): ")

	admin, err := user.NewAdmin(fisrtName, lastName, birthDate, "ILOLA", "12234")

	if err != nil {
		fmt.Println("Error creating user or admin:", err)
		return
	}
	admin.OutputUserData()
	admin.OutputAdminData()
	admin.User.ClearUserName()
}

func getUserData(promptText string) string {

	fmt.Println(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
