package models

import "time"

type CryptoStat struct {
    Coin      string    `bson:"coin"`
    Price     float64   `bson:"price"`
    MarketCap float64   `bson:"marketCap"`
    Change24h float64   `bson:"change24h"`
    Timestamp time.Time `bson:"timestamp"`
}
