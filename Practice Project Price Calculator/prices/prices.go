package prices

import (
	"errors"
	"fmt"
)

type TaxIncludedPriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]float64
}

func NewTaxIncludedPriceJob(rate float64) *TaxIncludedPriceJob {
	InputPrices := []float64{10, 20, 30}
	return &TaxIncludedPriceJob{
		InputPrices: InputPrices,
		TaxRate:     rate,
	}
}

func (t TaxIncludedPriceJob) Process() {
	results := make(map[float64][]float64)

	tempArray := []float64{}
	for _, valuePrices := range t.InputPrices {
		tempArray = append(tempArray, valuePrices-valuePrices*t.TaxRate)
	}

	results[t.TaxRate] = tempArray

	fmt.Println(results)

}

func CalcPrices(prices []float64, taxes []float64) (map[float64][]float64, error) {

	if len(prices) <= 0 || len(taxes) <= 0 {
		return map[float64][]float64{}, errors.New("invalid arguments passed")
	}

	results := make(map[float64][]float64)

	for _, valueTaxed := range taxes {
		tempArray := []float64{}
		for _, valuePrices := range prices {
			tempArray = append(tempArray, valuePrices-valuePrices*valueTaxed)
		}
		results[valueTaxed] = tempArray
	}

	return results, nil
}
