package prices

import (
	"bufio"
	"fmt"
	"go/scanner"
	"os"
)

type taxIncludePriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	taxIncludePrices map[string]float64
}

func (job taxIncludePriceJob) LoadData() {
	file, err := os.Open("prices.txt")

	//error handling
	if err != nil {
		fmt.Println("Could not open file.")
		fmt.Println(err)
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
}

func (job taxIncludePriceJob) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%2f", price)] = price * (1 + job.TaxRate)
	}
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *taxIncludePriceJob {

	return &taxIncludePriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
