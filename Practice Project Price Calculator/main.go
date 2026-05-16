package main

import (
	"fmt"

	"example.com/price_calculator/prices"
	"example.com/price_calculator/utility"
)

func main() {
	pricesInfo, priceInfoErr := utility.ReadFiles("prices.txt")
	if priceInfoErr != nil {
		fmt.Println(priceInfoErr.Error())
		return
	}

	//  pricesInfo:= []float64{10, 20, 30}
	taxes := []float64{0, 0.07, 0.1, 0.15}
	fmt.Println("line no 13 printed ")
	results, err := prices.CalcPrices(pricesInfo, taxes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, taxRate := range taxes {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

	fmt.Println(results)
}
