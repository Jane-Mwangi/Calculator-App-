package prices

import (

	"calculator/conversions"
	"calculator/filemanager"
	"fmt"
	
)

type TaxIncludePriceJob struct {
	IOManager 	  filemanager.FileManager
	TaxRate          float64
	InputPrices      []float64
	TaxIncludePrices map[string]string
}

func (job *TaxIncludePriceJob) LoadData() {

	lines,err :=job.IOManager.ReadLines()

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
	job.IOManager.WriteResult(job)
	
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager,taxRate float64) *TaxIncludePriceJob {

	return &TaxIncludePriceJob{
		IOManager:  fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
