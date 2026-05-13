package functionsarevalues

import (
	"errors"
	"fmt"
)

type factoreAlias func(int, int) int

func main() {
	numbers := []int{
		1, 2, 3, 4, 5,
	}
	// here we are passing the function as a param(factor)
	doubleNumbers, err := mulArray(numbers, 5, factor)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(doubleNumbers)

}

// here we are reciving a function as a Params
// function is declared like func(type_ofParam)return type
func mulArray(numbers []int, factor int, factorFunction factoreAlias) ([]int, error) {
	if len(numbers) == 0 || factor <= 0 {
		return []int{}, errors.New("Invalid Params passed kindly check the entered value")
	}
	doubleNumbers := []int{}
	for _, value := range numbers {
		// fmt.Println(value)
		doubleNumbers = append(doubleNumbers, factorFunction(value, factor))
	}

	return doubleNumbers, nil
}

// this is the function used as a param above
func factor(number int, multiple int) int {
	return multiple * number
}

// here we can also return the functoion
func getTransformedFunction() factoreAlias {
	return factor
}
