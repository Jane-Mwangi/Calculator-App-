package main

import "fmt"

func main() {
	var prices []float64 = []float64{10, 20, 30}
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		var taxIncludePrices []float64 = make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludePrices[priceIndex] = price * (1 + taxRate)

		}
		result[taxRate] = taxIncludePrices
		fmt.Println(result)
	}
}
