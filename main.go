package main

import (
	"calculator/prices"
)

func main() {
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}
