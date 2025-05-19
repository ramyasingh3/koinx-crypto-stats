package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/ramyasingh3/koinx-assignment/api-server/internal/config"
	"github.com/ramyasingh3/koinx-assignment/api-server/internal/services"
	"github.com/redis/go-redis/v9"
)

type Event struct {
	Trigger string `json:"trigger"`
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to MongoDB
	client, err := config.ConnectMongo()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(nil)

	// Initialize Redis client
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost:6379"
	}
	
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "", // No password for default Redis setup
		DB:       0,  // Default DB
	})

	// Create a context for Redis operations
	ctx := context.Background()

	// Test Redis connection
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")

	// Subscribe to Redis channel
	pubsub := rdb.Subscribe(ctx, "crypto_updates")
	defer pubsub.Close()

	// Start a goroutine to handle incoming messages
	go func() {
		ch := pubsub.Channel()
		for msg := range ch {
			var event Event
			if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
				log.Printf("Error unmarshaling event: %v", err)
				continue
			}

			if event.Trigger == "update" {
				log.Println("Received update event, triggering storeCryptoStats...")
				if err := services.StoreCryptoStats(client); err != nil {
					log.Printf("Error storing crypto stats: %v", err)
				} else {
					log.Println("Successfully stored crypto stats")
				}
			}
		}
	}()

	// Register the endpoints
	http.HandleFunc("/stats", services.StatsHandler(client))
	http.HandleFunc("/deviation", services.DeviationHandler(client))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}