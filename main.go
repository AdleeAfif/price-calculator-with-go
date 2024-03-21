package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.1, 0.2, 0.3}

	pricesAfterTax := make(map[float64][]float64)

	for _, tax := range taxRates {
		for _, price := range prices {
			pricesAfterTax[tax] = append(pricesAfterTax[tax], calculatePriceAfterTax(price, tax))
		}
	}

	fmt.Println(pricesAfterTax)
}

func calculatePriceAfterTax(price float64, taxRate float64) float64 {
	return price * (1 + taxRate)
}
