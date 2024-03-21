package main

import (
	"fmt"
	"project/price-calculator/filemanager"
	"project/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.1, 0.2, 0.3}

	for _, tax := range taxRates {
		fm := filemanager.New("prices/pricesDB.txt", "result_"+fmt.Sprintf("%.0f", tax*100)+".json")
		priceJob := prices.New(fm, tax)
		priceJob.Process()
	}
}
