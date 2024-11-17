package stock

import (
	"encoding/json"
	"net/http"
	"project-pubsub/db"
	"project-pubsub/lib/logger"
)

// StockHandler handles stock-related API requests.
type StockHandler struct{}

// NewStockHandler creates a new StockHandler instance.
func NewStockHandler() *StockHandler {
	return &StockHandler{}
}

// GetAllStocks handles the GET /v1/stocks request to fetch all stock prices.
func (sh *StockHandler) GetAllStocks(w http.ResponseWriter, r *http.Request) {
	// Fetch stock prices from the database.
	stockPrices, err := db.GetAllStockPrices()
	if err != nil {
		logger.Error(err)
		http.Error(w, "Failed to fetch stock prices", http.StatusInternalServerError)
		return
	}

	// Set response headers and write the JSON response.
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(stockPrices); err != nil {
		logger.Error(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
