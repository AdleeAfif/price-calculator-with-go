package main

import (
	"fmt"
	"project/price-calculator/filemanager"
	"project/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.1, 0.2, 0.3}
	channels := make([]chan bool, len(taxRates))
	errChan := make([]chan error, len(taxRates))

	for i, tax := range taxRates {

		fm := filemanager.New("prices/pricesDB.txt", "result_"+fmt.Sprintf("%.0f", tax*100)+".json")
		// tm := terminalmanager.New()

		priceFileJob := prices.New(fm, tax)
		// priceTerminalJob := prices.New(tm, tax)

		channels[i] = make(chan bool)
		errChan[i] = make(chan error)

		go priceFileJob.Process(channels[i], errChan[i])

		// go priceTerminalJob.Process(channels[i], errChan[i])
	}

	for index := range taxRates {
		select {
		case err := <-errChan[index]:
			if err != nil {
				fmt.Println("Could not process job!")
				fmt.Println(err)
			}
		case <-channels[index]:
			fmt.Println("Job done!")
		}
	}

	for _, channel := range channels {
		<-channel
	}
}
