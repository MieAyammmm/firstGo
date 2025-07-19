package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/MieAyammmm/recipe/backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeController struct {
	Collection *mongo.Collection
}

func NewRecipeController(db *mongo.Database) *RecipeController {
	return &RecipeController{
		Collection: db.Collection("recipes"),
	}
}

// CreateRecipe - POST /recipes
func (rc *RecipeController) CreateRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi input
	if recipe.Difficulty == "" || len(recipe.Ingredients) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	// Insert ke MongoDB
	result, err := rc.Collection.InsertOne(context.Background(), recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create recipe"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":      result.InsertedID,
		"message": "Recipe created successfully",
	})
}

// GetRecipes - GET /recipes
func (rc *RecipeController) GetRecipes(c *gin.Context) {
	var recipes []models.Recipe

	cursor, err := rc.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipes"})
		return
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &recipes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode recipes"})
		return
	}

	c.JSON(http.StatusOK, recipes)
}

// GetRecipeByID - GET /recipes/:id
func (rc *RecipeController) GetRecipeByID(c *gin.Context) {
    id := c.Param("id")
    fmt.Println("Received ID:", id)
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
        return
    }

    var recipe models.Recipe
    if err := rc.Collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&recipe); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
        return
    }

    c.JSON(http.StatusOK, recipe)
}
// UpdateRecipe - PUT /recipes/:id
func (rc *RecipeController) UpdateRecipe(c *gin.Context) {
    id := c.Param("id")
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
        return
    }

    var updatedRecipe models.Recipe
    if err := c.ShouldBindJSON(&updatedRecipe); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validasi input (opsional, jika tidak pakai validator)
    if updatedRecipe.Difficulty == "" || len(updatedRecipe.Ingredients) == 0 || len(updatedRecipe.Steps) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
        return
    }

    // Update data
    update := bson.M{
        "$set": bson.M{
            "title":       updatedRecipe.Title,
            "ingredients": updatedRecipe.Ingredients,
            "steps":       updatedRecipe.Steps,
            "duration":    updatedRecipe.Duration,
            "difficulty":  updatedRecipe.Difficulty,
            "updatedAt":   time.Now(),
        },
    }

    result, err := rc.Collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, update)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update recipe"})
        return
    }

    if result.ModifiedCount == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found or no changes made"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Recipe updated successfully"})
}
// DeleteRecipe - DELETE /recipes/:id
func (rc *RecipeController) DeleteRecipe(c *gin.Context) {
    id := c.Param("id")
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
        return
    }

    result, err := rc.Collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete recipe"})
        return
    }

    if result.DeletedCount == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted successfully"})
}