package main

import (
	"fmt"

	"example.com/structs/users"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	customerInfo, error := users.NewCustomerData(firstName, lastName, birthdate)
	if error != nil {
		println(error.Error())
		return
	}
	// ... do something awesome with that gathered data!
	customerInfo.PrintUserDara()
	customerInfo.ChangeName("rahul", "bisht")
	customerInfo.PrintUserDara()
	// fmt.Println(firstName, lastName, birthdate)
	// fmt.Println(customerInfo)

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
