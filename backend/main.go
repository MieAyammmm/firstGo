package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MieAyammmm/recipe/backend/controllers"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found (using default settings)")
	}

	// 2. Setup Gin
	router := gin.Default()

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
	router.Run(":" + port)
}