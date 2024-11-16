package main

import (
	"fmt"
	"project-pubsub/lib/logger"
	"project-pubsub/pkg/pubsub"
	"project-pubsub/pkg/stock"
	"time"
)

func main() {
	logger.Info("Starting Stock Price Notifier...")

	// Initialize PubSub
	ps := pubsub.NewPubSub()

	// Initialize StockService
	ss := stock.NewStockService(ps)

	// Start fetching stock prices
	go ss.FetchStockPrices()

	// Create subscribers
	sub1 := pubsub.NewSubscriber(1)
	sub2 := pubsub.NewSubscriber(2)
	ps.AddSubscriber(sub1)
	ps.AddSubscriber(sub2)

	// Start listening for updates
	go func() {
		// This loop will continue to run as long as there are messages in the channel
		for msg := range sub1.Chan {
			// For each message received, print it to the console with a prefix indicating it's for Subscriber 1
			fmt.Printf("Subscriber 1 received: %s\n", msg)
		}
	}()

	go func() {
		// This loop will continue to run as long as there are messages in the channel
		for msg := range sub2.Chan {
			// For each message received, print it to the console with a prefix indicating it's for Subscriber 2
			fmt.Printf("Subscriber 2 received: %s\n", msg)
		}
	}()

	// Simulate a running service
	time.Sleep(30 * time.Second)
	logger.Info("Shutting down...")
}
