package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/practice/customer"
)

func main() {
	createCustomer()
}

func createCustomer() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What is your first name?")
	customerFirstName, _ := reader.ReadString('\n')
	customerFirstName = strings.TrimSpace(customerFirstName)

	fmt.Println("What is your last name?")
	customerLastName, _ := reader.ReadString('\n')
	customerLastName = strings.TrimSpace(customerLastName)

	fmt.Println("What is your phone number?")
	customerPhoneNumber, _ := reader.ReadString('\n')
	customerPhoneNumber = strings.TrimSpace(customerPhoneNumber)

	fmt.Println("What is your age?")
	var customerAge int
	fmt.Scan(&customerAge)

	// Clear leftover newline after fmt.Scan
	reader.ReadString('\n')

	fmt.Println("What is your Primary Address?")
	customerAddress, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Not able to read Customer Address")
		return
	}
	customerAddress = strings.TrimSpace(customerAddress)

	fmt.Println("What is your Secondary Address?")
	customerAddressSecondary, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Unable to read Customer Secondary Address")
		return
	}
	customerAddressSecondary = strings.TrimSpace(customerAddressSecondary)

	customerStruct, err := customer.New(
		customerFirstName,
		customerLastName,
		customerPhoneNumber,
		customerAge,
		customerAddress,
		customerAddressSecondary,
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	errCustomer := customer.SetDefaultsFunction(&customerStruct)
	if errCustomer != nil {
		fmt.Println(errCustomer.Error())
		return
	}

	fmt.Println(customerStruct)
}
