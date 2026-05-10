package users

import (
	"errors"
	"fmt"
	"time"
)

type CustomerData struct {
	firstName string
	lastName  string
	birthdate string
	createAt  time.Time
}

type Admin struct {
	email    string
	password string
	CustomerData
}

func (customerInfo CustomerData) PrintUserDara() {
	fmt.Println(customerInfo.firstName)
	fmt.Println(customerInfo.lastName)
	fmt.Println(customerInfo.birthdate)
	fmt.Println(customerInfo.createAt)
}
func (customerInfoAddress *CustomerData) ChangeName(firstName string, lastName string) {
	if len(firstName) != 0 {
		println("inside the condition")
		customerInfoAddress.firstName = firstName
	}
	if len(lastName) != 0 {
		customerInfoAddress.lastName = lastName
	}
}
func NewCustomerData(firstName string, lastName string, birthdate string) (CustomerData, error) {
	if firstName == "" || lastName == "" || birthdate == "" {

		return CustomerData{}, errors.New("Required Fields are empty , kindly check the entered Data")
	}
	return CustomerData{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createAt:  time.Now(),
	}, nil
}

func NewAdmin(email string, password string, firstName string, lastName string, birthdate string) (Admin, nil) {
	if email == "" || password == "" {
		return Admin{}, errors.New("Email or password is not set")
	}
	if firstName == "" || lastName == "" || birthdate == "" {

		return Admin{}, errors.New("Required Fields are empty , kindly check the entered Data")
	}

	return Admin{
		email:    email,
		password: password,
		CustomerData: CustomerData{
			firstName: firstName,
			lastName:  lastName,
			birthdate: birthdate,
			createAt:  time.Now(),
		},
	}, nil
}
