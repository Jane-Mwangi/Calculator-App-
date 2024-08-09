package prices

type taxIncludePrices struct {
	TaxRate float64
	InputPrices []float64
	taxIncludePrices map[string]float64
}