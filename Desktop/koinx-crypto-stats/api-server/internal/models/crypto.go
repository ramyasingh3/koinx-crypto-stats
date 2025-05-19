package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CryptoStats represents a cryptocurrency statistics document in MongoDB.
type CryptoStats struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Coin      string             `bson:"coin"`
	Price     float64            `bson:"price"`
	MarketCap float64            `bson:"marketCap"`
	Change24h float64            `bson:"change24h"`
	Timestamp time.Time          `bson:"timestamp"`
}