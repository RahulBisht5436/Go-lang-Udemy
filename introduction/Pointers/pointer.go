package main

import "fmt"

func main() {
	age := 32
	ageAddress := &age
	fmt.Println(ageAddress)
	getAdultAge(ageAddress)
	println(age)
}

func getAdultAge(age *int) {
	*age -= 18
}
