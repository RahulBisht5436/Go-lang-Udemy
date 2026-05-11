package main

import "fmt"

func main() {
	// This is the Syntax for decalring array
	prices := []float32{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	//prints the 2 element of the array
	fmt.Printf("This is the Second Element of the Array : %v \n", prices[1])

	// Append new value
	prices = append(prices, 120)

	//prints the prices array
	fmt.Println(prices)

	// prints the len of the array
	fmt.Printf("This is the length of Array %v \n", len(prices))

	//Slice prints a particular segment of the array based
	// on the provied initial and end position
	fmt.Println(prices[:])

	// Capacity shows the number of element in original array
	fmt.Print(cap(prices))

	// to append slice inside a array 
	discountedArray := []float32{12,23,34,45,56}
	prices = append(prices, discountedArray... )
}

