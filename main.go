package main

import (
	"fmt"
	"project/price-calculator/filemanager"
	"project/price-calculator/prices"
	"project/price-calculator/terminalmanager"
)

func main() {
	taxRates := []float64{0, 0.1, 0.2, 0.3}

	for _, tax := range taxRates {

		fm := filemanager.New("prices/pricesDB.txt", "result_"+fmt.Sprintf("%.0f", tax*100)+".json")
		tm := terminalmanager.New()

		priceFileJob := prices.New(fm, tax)
		priceTerminalJob := prices.New(tm, tax)

		err := priceFileJob.Process()
		if err != nil {
			fmt.Println("Could not process file job!")
			fmt.Println(err)
		}

		err = priceTerminalJob.Process()
		if err != nil {
			fmt.Println("Could not process terminal job!")
			fmt.Println(err)
		}
	}
}
