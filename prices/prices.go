package prices

import (
	"bufio"
	"calculator/conversions"
	"fmt"
	"os"
	
)

type TaxIncludePriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	taxIncludePrices map[string]float64
}

func (job *TaxIncludePriceJob) LoadData() {
	file, err := os.Open("prices/prices.txt")

	//error handling
	if err != nil {
		fmt.Println("Could not open file.")
		fmt.Println(err)

		return
	}
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Could not open file.")
		fmt.Println(err)
		file.Close()

		return
	}

	prices, err := conversions.StringsToFloat(lines)


		if err != nil {
		
			fmt.Println(err)
			file.Close()
			return
		}


		job.InputPrices = prices
		file.Close()
	}


func (job *TaxIncludePriceJob) Process() {

	job.LoadData()

	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%2f", price)] = price * (1 + job.TaxRate)
	}
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludePriceJob {

	return &TaxIncludePriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
