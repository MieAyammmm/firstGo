package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// 1. Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found (using default settings)")
	}

	// 2. Setup HTTP Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server is running! 🚀")
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	// 3. Start HTTP Server (di goroutine terpisah)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: nil, // Gunakan http.DefaultServeMux
	}

	go func() {
		log.Printf("Server running on http://localhost:%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// 4. Connect to MongoDB (opsional, tidak menghentikan server jika gagal)
	connectToMongoDB()

	// Block main thread
	select {}
}

func connectToMongoDB() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Println("MONGODB_URI not set, skipping database connection")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Printf("⚠️ MongoDB connection failed: %v", err)
		return
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("⚠️ MongoDB ping failed: %v", err)
		return
	}

	log.Println("✅ Connected to MongoDB!")
}