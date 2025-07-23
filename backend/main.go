package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MieAyammmm/recipe/backend/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CORS Middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	// 1. Load .env file (for local dev)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found (using Railway env vars)")
	}

	// 2. Get environment variables
	mongoURI := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	if mongoURI == "" {
		log.Fatal("❌ MONGODB_URI not set in environment variables")
	}
	if dbName == "" {
		log.Fatal("❌ DB_NAME not set in environment variables")
	}
	if port == "" {
		port = "8000" // fallback if PORT not set
	}

	log.Println("🔧 Environment variables loaded")

	// 3. Setup Gin
	router := gin.Default()
	router.Use(CORSMiddleware())

	// 4. Connect to MongoDB
	ctx := context.Background()
	log.Println("🔌 Connecting to MongoDB...")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("❌ MongoDB connection failed:", err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("⚠️ Error disconnecting MongoDB:", err)
		}
	}()

	// 5. Ping MongoDB
	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := client.Ping(ctxPing, nil); err != nil {
		log.Fatal("❌ MongoDB ping failed:", err)
	}
	log.Println("✅ Connected to MongoDB!")

	// 6. Init controller
	db := client.Database(dbName)
	recipeController := controllers.NewRecipeController(db)

	// 7. Register routes
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is running! 🚀")
	})
	router.POST("/recipes", recipeController.CreateRecipe)
	router.GET("/recipes", recipeController.GetRecipes)
	router.GET("/recipes/:id", recipeController.GetRecipeByID)
	router.PUT("/recipes/:id", recipeController.UpdateRecipe)
	router.DELETE("/recipes/:id", recipeController.DeleteRecipe)

	// 8. Start server
	log.Printf("🌐 Server running on http://0.0.0.0:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
