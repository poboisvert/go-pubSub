package utils

import "fmt"

// FormatPrices formats stock prices for display.
func FormatPrices(prices map[string]float64) string {
	return fmt.Sprintf("Stock Prices: %v", prices)
}
