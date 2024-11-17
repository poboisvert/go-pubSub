package db

import (
	"database/sql"
	"fmt"
	"time"

	"project-pubsub/config"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var dbInstance *sql.DB

// Connect initializes the PostgreSQL connection and ensures the required table exists.
func Connect() error {
	var err error
	dbConfig := config.GetConfig()

	// Build the connection string from the config.
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Database)

	// Open the connection.
	dbInstance, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Test the connection.
	if err = dbInstance.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Ensure the required table exists.
	if err := ensureTableExists(); err != nil {
		return fmt.Errorf("failed to ensure table exists: %w", err)
	}

	fmt.Println("Connected to PostgreSQL successfully.")
	return nil
}

// Close closes the database connection gracefully.
func Close() {
	if dbInstance != nil {
		_ = dbInstance.Close()
		fmt.Println("Database connection closed.")
	}
}

// GetDB provides the active database connection.
func GetDB() *sql.DB {
	return dbInstance
}

// ensureTableExists checks if the `stock_prices` table exists and creates it if not.
func ensureTableExists() error {
	query := `
	CREATE TABLE IF NOT EXISTS stock_prices (
		id SERIAL PRIMARY KEY,
		stock_name VARCHAR(255) NOT NULL,
		price NUMERIC(10, 2) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := dbInstance.Exec(query)
	if err != nil {
		return fmt.Errorf("error ensuring table exists: %w", err)
	}
	return nil
}

// SaveStockPrice inserts a stock price into the database.
func SaveStockPrice(stockName string, price float64) error {
	query := "INSERT INTO stock_prices (stock_name, price) VALUES ($1, $2)"
	_, err := dbInstance.Exec(query, stockName, price)
	if err != nil {
		return fmt.Errorf("failed to save stock price: %w", err)
	}
	return nil
}

// StockPrice represents a stock price record from the database.
type StockPrice struct {
	ID        int       `json:"id"`
	StockName string    `json:"stock_name"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"created_at"`
}

// GetAllStockPrices retrieves all stock prices from the database.
func GetAllStockPrices() ([]StockPrice, error) {
	query := "SELECT id, stock_name, price, created_at FROM stock_prices ORDER BY created_at DESC"
	rows, err := dbInstance.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query stock prices: %w", err)
	}
	defer rows.Close()

	var stockPrices []StockPrice
	for rows.Next() {
		var sp StockPrice
		if err := rows.Scan(&sp.ID, &sp.StockName, &sp.Price, &sp.Timestamp); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		stockPrices = append(stockPrices, sp)
	}

	return stockPrices, nil
}
