package prices

type taxIncludePriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	taxIncludePrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *taxIncludePriceJob {

	return &taxIncludePriceJob{
		InputPrices: 	[]float64{10, 20, 30},
		TaxRate: 		taxRate,
	}
}
