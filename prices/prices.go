package prices

import (
	"fmt"
	"project/price-calculator/conversion"
	"project/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]float64
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadDataFromFile()
	pricesAfterTax := make(map[string]string)
	for _, price := range job.InputPrices {
		result := price * (1 + job.TaxRate)
		pricesAfterTax[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", result)
	}

	fmt.Println(pricesAfterTax)
}

func New(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job *TaxIncludedPriceJob) LoadDataFromFile() {

	fileLine, err := filemanager.ReadLines("prices/pricesDB.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := conversion.StringToFloat(fileLine)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = data
}