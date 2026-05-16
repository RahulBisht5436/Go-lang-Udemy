package main

import (
	"fmt"
	"os"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	if cwd, err := os.Getwd(); err == nil {
		fmt.Println("Working directory:", cwd)
	}

	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChan := make([]chan bool, len(taxRates))

	for index, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		doneChan[index] = make(chan bool)
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChan[index])
	}

	for _, value := range doneChan {
		<-value
	}
}
