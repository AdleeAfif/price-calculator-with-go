package prices

import (
	"fmt"
	"project/price-calculator/conversion"
	"project/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager        iomanager.IOManager `json:"-"` //exclude from json
	TaxRate          float64             `json:"tax_rate"`
	InputPrices      []float64           `json:"input_prices"`
	TaxIncludedPrice map[string]string   `json:"tax_included_price"`
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}

	pricesAfterTax := make(map[string]string)
	for _, price := range job.InputPrices {
		result := price * (1 + job.TaxRate)
		pricesAfterTax[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", result)
	}

	job.TaxIncludedPrice = pricesAfterTax

	return job.IOManager.WriteJSON(job)
}

func New(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job *TaxIncludedPriceJob) LoadData() error {

	fileLine, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	data, err := conversion.StringToFloat(fileLine)

	if err != nil {
		return err
	}

	job.InputPrices = data

	return nil
}
