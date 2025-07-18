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
	// 1. Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found (using default settings)")
	}

	// 2. Setup Gin with CORS
	router := gin.Default()
	router.Use(CORSMiddleware()) 

	// 3. Connect to MongoDB
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(os.Getenv("MONGODB_URI")),
	)
	if err != nil {
		log.Fatal("❌ MongoDB connection failed:", err)
	}
	defer client.Disconnect(context.Background())

	// 4. Check MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("❌ MongoDB ping failed:", err)
	}
	log.Println("✅ Connected to MongoDB!")

	// 5. Initialize Controller
	db := client.Database(os.Getenv("DB_NAME"))
	recipeController := controllers.NewRecipeController(db)

	// 6. Register Routes
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is running! 🚀")
	})
	router.POST("/recipes", recipeController.CreateRecipe)
	router.GET("/recipes", recipeController.GetRecipes)
	router.GET("/recipes/:id", recipeController.GetRecipeByID)
	router.PUT("/recipes/:id", recipeController.UpdateRecipe)
	router.DELETE("/recipes/:id", recipeController.DeleteRecipe)

	// 7. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("🌐 Server running on http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}