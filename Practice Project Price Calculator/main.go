package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxes := []float64{0, 0.07, 0.1, 0.15}
	results := make(map[float64][]float64)

	for _, valueTaxed := range taxes {
		tempArray := []float64{}
		for _, valuePrices := range prices {
			tempArray = append(tempArray, valuePrices-valuePrices*valueTaxed)
		}
		results[valueTaxed] = tempArray
	}

	fmt.Println(results)
}
