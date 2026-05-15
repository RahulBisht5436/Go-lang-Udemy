package main

import (
	"fmt"

	"example.com/callBackFunction/utility"
)

func main() {
	DoubleFunction, err := createMultiple(2)
	if err != nil {
		fmt.Println(err.Error())
	}
	DoubleFunction(20)
	DoubleFunction(40)
	DoubleFunction(60)
}
func createMultiple(number int) func(int) (int, error) {
	valType := utility.CheckType(number)
	if valType != "int" {
		return 0, error.New("Invalid Data Type")
	}
	return func(number2 int) (int, error) {
		valTypeNumber := utility.CheckType(number)
		if valTypeNumber != "int" {
			return 0, error.New("Invalid Data Type")
		}
		return number * number2
	}
}
