package prices

import (

	"calculator/conversions"
	"calculator/filemanager"
	"fmt"
	
)

type TaxIncludePriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludePrices map[string]string
}

func (job *TaxIncludePriceJob) LoadData() {

	lines,err :=filemanager.ReadLines("prices/prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}


	prices, err := conversions.StringsToFloat(lines)

		if err != nil {
			fmt.Println(err)
			return
		}

		job.InputPrices = prices
		
	}


func (job *TaxIncludePriceJob) Process() {

	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludePrice:=price*(1+job.TaxRate)
		result[fmt.Sprintf("%2f", price)] = fmt.Sprintf("%2f",taxIncludePrice)
	}
	job.TaxIncludePrices =result
	filemanager.WriteJSON(fmt.Sprintf("result_%0f.json",job.TaxRate*100),job)
	
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludePriceJob {

	return &TaxIncludePriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
