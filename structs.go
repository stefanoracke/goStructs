package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	promptFirstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(promptFirstName, lastName, birthdate)
	if err != nil {
		fmt.Print("Error:", err)
		return
	}

	appAdmin := user.NewAdmin("test@admin.com", "123456")
	appAdmin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUser()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
