package services

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DeviationHandler(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		coin := r.URL.Query().Get("coin")
		if coin == "" {
			http.Error(w, "coin query param is required", http.StatusBadRequest)
			return
		}

		collection := client.Database("koinx").Collection("cryptostats")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Find last 100 records sorted by timestamp descending
		cursor, err := collection.Find(ctx,
			bson.M{"coin": coin},
			options.Find().SetSort(bson.D{{"timestamp", -1}}).SetLimit(100),
		)
		if err != nil {
			http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)

		var prices []float64
		for cursor.Next(ctx) {
			var record bson.M
			if err := cursor.Decode(&record); err != nil {
				http.Error(w, "Error decoding data", http.StatusInternalServerError)
				return
			}
			price, ok := record["price"].(float64)
			if !ok {
				continue // skip if price is missing or wrong type
			}
			prices = append(prices, price)
		}

		if len(prices) == 0 {
			http.Error(w, "No price data found", http.StatusNotFound)
			return
		}

		// Calculate standard deviation
		mean := mean(prices)
		variance := 0.0
		for _, p := range prices {
			variance += (p - mean) * (p - mean)
		}
		variance /= float64(len(prices))
		stdDev := math.Sqrt(variance)

		// Return JSON response
		resp := map[string]float64{
			"deviation": stdDev,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

// Helper function to calculate mean of slice
func mean(nums []float64) float64 {
	total := 0.0
	for _, n := range nums {
		total += n
	}
	return total / float64(len(nums))
}
