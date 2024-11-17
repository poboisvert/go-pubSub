# Stock Price Channel/Goroutines

A Go-based service for fetching stock prices from an external source and notifying subscribers in real-time. The service demonstrates core Go concepts like **goroutines**, **channels**, and **mutexes** while persisting data in PostgreSQL.

- Pub/Sub Architecture: Efficiently decouples the producer (price fetcher) from the consumers (subscribers).
- Database Persistence: Saves each stock price update for historical reference.
- Goroutines & Channels: Demonstrates concurrency for fetching and notifying updates.

Feature to implement:

- WebSocket support for real-time notifications.

## Prerequisites

Go 1.20 or higher.
PostgreSQL installed and running.
Set up a .env file with the following content:

```
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=voting_app
```

## Features

- Fetch stock prices at regular intervals using goroutines.
- Notify subscribers about stock price updates using a pub/sub model.
- Persist stock price data in PostgreSQL.
- Expose an API to retrieve stock prices.

### # 1. Goroutines

Goroutines are lightweight threads managed by the Go runtime. This project uses goroutines for:

- **Fetching Stock Prices**: A goroutine periodically fetches stock prices from an API and pushes them to a channel.
- **Notifying Subscribers**: Another goroutine listens to the stock price update channel and notifies all subscribers.

**Example**:

```go
go func() {
    ticker := time.NewTicker(1 * time.Minute)
    for range ticker.C {
        fetcher.FetchStockPrices()
    }
}()

```

### 2. Channel

priceChannel := make(chan StockPrice)
go fetcher.FetchStockPrices(priceChannel)
go pubsub.NotifySubscribers(priceChannel)

- Stock Price Updates: A channel passes updated stock prices from the fetcher to the pub/sub system.

```
priceChannel := make(chan StockPrice)
go fetcher.FetchStockPrices(priceChannel)
go pubsub.NotifySubscribers(priceChannel)
```

### 3. Mutex

A mutex (mutual exclusion lock) ensures thread-safe access to shared resources. Here, it protects the subscriber list during concurrent read/write operations in the pub/sub model.

```
var mu sync.Mutex

func (ps *PubSub) AddSubscriber(subscriber chan StockPrice) {
    mu.Lock()
    defer mu.Unlock()
    ps.subscribers = append(ps.subscribers, subscriber)
}
```

### 4. Endpoints

- GET /v1/stocks

```
[
  {
    "id": 1,
    "stock_name": "AAPL",
    "price": 152.34,
    "created_at": "2024-11-16T15:04:05Z"
  }
]

```
