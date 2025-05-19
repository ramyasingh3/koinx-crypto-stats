package services

import (
    "context"
    "encoding/json"
    "net/http"

    "github.com/ramyasingh3/koinx-assignment/api-server/internal/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// StatsHandler returns an HTTP handler function
// 1. Read coin query param
// 2. Get the Mongo collection
// 3. Query: Find the latest record for this coin
// 4. Prepare response JSON (match sample response keys)
// 5. Set JSON headers and send response
func StatsHandler(client *mongo.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 1. Read coin query param
        coin := r.URL.Query().Get("coin")
        if coin == "" {
            http.Error(w, "missing coin query param", http.StatusBadRequest)
            return
        }

        // 2. Get the Mongo collection
        collection := client.Database("koinx").Collection("cryptostats")

        // 3. Query: Find the latest record for this coin
        ctx := context.Background()
        opts := options.FindOne().SetSort(bson.D{{"timestamp", -1}})  // newest first

        var stat models.CryptoStats

        err := collection.FindOne(ctx, bson.M{"coin": coin}, opts).Decode(&stat)
        if err != nil {
            if err == mongo.ErrNoDocuments {
                http.Error(w, "no data found for coin", http.StatusNotFound)
                return
            }
            http.Error(w, "db error: "+err.Error(), http.StatusInternalServerError)
            return
        }

        // 4. Prepare response JSON (match sample response keys)
        res := map[string]interface{}{
            "price":     stat.Price,
            "marketCap": stat.MarketCap,
            "24hChange": stat.Change24h,
        }

        // 5. Set JSON headers and send response
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(res)
    }
}

// DeviationHandler returns an HTTP handler function for price deviation
// (Removed duplicate implementation; see deviation.go)
