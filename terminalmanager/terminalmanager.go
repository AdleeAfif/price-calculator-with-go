package terminalmanager

import "fmt"

type TermManager struct {
}

func (tm TermManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter the prices by hitting ENTER")

	var prices []string
	for {
		var price string
		fmt.Print("Enter price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (tm TermManager) WriteJSON(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() TermManager {
	return TermManager{}
}
