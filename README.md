# Price Calculator with Hardcoded Tax

This Go program calculates the total price of items with a hardcoded tax rate. It reads input prices from a text file and calculates the prices including taxes (in JSON file).

## Usage

1. Clone the repository:
   ```
   git clone https://github.com/AdleeAfif/price-calculator-with-go
   ```

2. Navigate to the project directory:
   ```
   cd price-calculator-with-go
   ```

3. Ensure you have Go installed. If not, you can download it from the official website: https://golang.org/

4. Create a text file named `pricesDB.txt` in the project directory and enter the prices of items, each on a new line.

5. Run the program:
   ```
   go run main.go
   ```

6. The program will read the prices from the `pricesDB.txt` file, calculate the total price including tax, and display the result.

## Customization

If you want to change the tax rate, you can modify the `taxRates` constant in the `main.go` file.

## Example

Suppose `pricesDB.txt` contains the following prices:

```
9.99
19.22
11.50
29.10
```

If the tax rate is set to 10% the output will be:
```
{
    "tax_rate": 0.1,
    "input_prices": [
        9.99,
        19.22,
        11.5,
        29.1
    ],
    "tax_included_price": {
        "11.50": "12.65",
        "19.22": "21.14",
        "29.10": "32.01",
        "9.99": "10.99"
    }
}
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
