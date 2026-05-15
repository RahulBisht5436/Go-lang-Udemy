package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int
	fmt.Println("Enter the Value")
	fmt.Scan(&number)

	numberDouble, errDouble := apply(number, double)
	if errDouble != nil {
		fmt.Println(errDouble.Error())
		return
	}
	fmt.Println(numberDouble)

	numberSquare, errSquare := apply(number, square)
	if errSquare != nil {
		fmt.Println(errSquare.Error())
		return
	}
	fmt.Println(numberSquare)

}

func apply(number int, function func(int) (int, error)) (int, error) {
	if checkType(number) != "int" {
		return 0, errors.New("Invalid param passed")
	}
	result, error := function(number)
	if error != nil {
		return 0, errors.New(error.Error())
	}
	return result, nil
}

func checkType(value any) string {
	switch value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case float64:
		return "float64"
	default:
		return "error"
	}
}

func double(number int) (int, error) {
	value := checkType(number)
	if value != "int" {
		return 0, errors.New("Enter a Valid Digit")
	}
	return number * 2, nil
}

func square(number int) (int, error) {
	value := checkType(number)
	if value != "int" {
		return 0, errors.New("Enter a Valid Digit")
	}
	return number * number, nil
}
