package stock

import (
	"fmt"
	"project-pubsub/pkg/pubsub"
	"sync"
	"time"
)

// StockService manages stock prices and notifications.
type StockService struct {
	PubSub      *pubsub.PubSub
	StockPrices map[string]float64
	sync.Mutex
}

// NewStockService creates a new StockService instance.
func NewStockService(ps *pubsub.PubSub) *StockService {
	return &StockService{
		PubSub:      ps,
		StockPrices: make(map[string]float64),
	}
}

// FetchStockPrices simulates fetching stock prices.
func (ss *StockService) FetchStockPrices() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ss.Lock()
			ss.StockPrices["AAPL"] = 150 + float64(time.Now().Second()%10)
			ss.StockPrices["GOOG"] = 2800 + float64(time.Now().Second()%50)
			updatedPrices := ss.StockPrices
			ss.Unlock()

			message := fmt.Sprintf("Updated Prices: %v", updatedPrices)
			ss.PubSub.Notify(message)
		}
	}
}
