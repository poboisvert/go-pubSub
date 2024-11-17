package main

import (
	"fmt"
	"log"
	"net/http"
	"project-pubsub/db"
	"project-pubsub/lib/logger"
	"project-pubsub/pkg/pubsub"
	"project-pubsub/pkg/stock"
	"time"
)

func main() {
	logger.Info("Starting Stock Price Notifier...")
	// Connect to database
	err := db.Connect()
	if err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}
	defer db.Close()

	// Initialize the StockHandler.
	stockHandler := stock.NewStockHandler()

	// Register routes.
	http.HandleFunc("/v1/stocks", stockHandler.GetAllStocks)

	// Start the HTTP server.
	logger.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

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
