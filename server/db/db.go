package db

import (
	"database/sql"
	"fmt"

	"project-pubsub/config"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var dbInstance *sql.DB

func Connect() error {
	var err error
	dbConfig := config.GetConfig()

	dbInstance, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Database))
	if err != nil {
		return err
	}

	if err = dbInstance.Ping(); err != nil {
		return err
	}

	fmt.Println("Connected to PostgreSQL")
	return nil
}

func Close() {
	if dbInstance != nil {
		_ = dbInstance.Close()
	}
}

func GetDB() *sql.DB {
	return dbInstance
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
