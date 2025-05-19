package services

import (
	"os"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/ramyasingh3/koinx-assignment/api-server/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// CoinGeckoResponse represents the CoinGecko API response structure.
// StoreCryptoStats fetches crypto stats from CoinGecko and stores them in MongoDB.
// Build CoinGecko API request
// Prepare stats for MongoDB
// Insert into MongoDB

func StoreCryptoStats(client *mongo.Client) error {
	coins := []string{"bitcoin", "ethereum", "matic-network"}

	// Build CoinGecko API request
	params := url.Values{}
	params.Add("ids", "bitcoin,ethereum,matic-network")
	params.Add("vs_currencies", "usd")
	params.Add("include_market_cap", "true")
	params.Add("include_24hr_change", "true")

	reqURL := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?%s", params.Encode())

	resp, err := http.Get(reqURL)
	if err != nil {
		return fmt.Errorf("failed to fetch from CoinGecko: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("CoinGecko API returned status: %s", resp.Status)
	}

	var data CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return fmt.Errorf("failed to decode CoinGecko response: %v", err)
	}

	// Prepare stats for MongoDB
	var stats []interface{}
	for _, coin := range coins {
		coinData, ok := data[coin]
		if !ok {
			return fmt.Errorf("no data for coin: %s", coin)
		}

		stats = append(stats, models.CryptoStats{
			Coin:      coin,
			Price:     coinData.USD,
			MarketCap: coinData.USDMarketCap,
			Change24h: coinData.USD24hChange,
			Timestamp: time.Now(),
		})
	}

	// Insert into MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbName := os.Getenv("MONGO_DB_NAME")
    collection := client.Database(dbName).Collection("cryptostats")

	_, err = collection.InsertMany(ctx, stats)
	if err != nil {
		return fmt.Errorf("failed to insert stats: %v", err)
	}

	log.Println("Crypto stats stored successfully")
	return nil
}