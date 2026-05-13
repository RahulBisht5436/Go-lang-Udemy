package customer

import (
	"errors"
	"fmt"
	"os"
)

type setDefaults interface {
	ChangeInsuranceStatus(data bool) error
	ChangeEmployedStatus(data bool) error
}
type additionalInformation struct {
	address_secondary string
	insurance         bool
	employed          bool
}

type customer struct {
	firstName             string
	lastName              string
	phoneNumber           string
	age                   int
	address               string
	additionalInformation additionalInformation
}

func (c customer) PrintCustomerInfo() {
	fmt.Println(c)
}

func (c *customer) ChangeInsuranceStatus(status bool) error {
	c.additionalInformation.insurance = status
	fmt.Printf("Insurance status changed to : %v\n", status)
	c.PrintCustomerInfo()
	return nil
}

func (c *customer) ChangeEmployedStatus(status bool) error {
	c.additionalInformation.employed = status
	fmt.Printf("Employed status changed to : %v\n", status)
	c.PrintCustomerInfo()
	return nil
}

func (c customer) SaveData() {
	os.WriteFile("customerData.txt", []byte(fmt.Sprintf("%+v", c)), 0644)
}

func New(firstName string, lastName string, phoneNumber string, age int, address string, address_secondary string) (customer, error) {
	if firstName == "" || lastName == "" || phoneNumber == "" || age == 0 || address == "" {
		return customer{}, errors.New("All mandatory fields are not entered")
	}
	return customer{
		firstName:   firstName,
		lastName:    lastName,
		phoneNumber: phoneNumber,
		address:     address,
		age:         age,
		additionalInformation: additionalInformation{
			address_secondary: address_secondary,
			employed:          false,
			insurance:         false,
		},
	}, nil
}

func SetDefaultsFunction(customer setDefaults) error {
	errEmployment := customer.ChangeEmployedStatus(true)
	if errEmployment != nil {
		return errors.New("failed to set the employment")
	}

	errInsurance := customer.ChangeInsuranceStatus(true)
	if errInsurance != nil {
		return errors.New("failed to set the isurance")
	}

	return nil

}
