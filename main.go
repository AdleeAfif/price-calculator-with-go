package main

import (
	"project/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.1, 0.2, 0.3}

	for _, tax := range taxRates {
		priceJob := prices.New(tax)
		priceJob.Process()
	}
}
