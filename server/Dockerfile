FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o github.com/poboisvert/go-pubSub

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/github.com/poboisvert/go-pubSub .
COPY .env .

# Optional: install necessary packages if required (like PostgreSQL client) 
# RUN apk add postgresql-client 

CMD ["/app/voting-app"]