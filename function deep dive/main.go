package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	var multiplicationFactor int
	fmt.Println("What is the factor of multiplication : ")
	fmt.Scan(&multiplicationFactor)
	transformed := transformNumbers(&numbers, createTranformer(multiplicationFactor))

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTranformer(multiplicationFactor int) func(int) int {
	return func(val int) int {
		return val * multiplicationFactor
	}
}
