package routes

import (
	"github.com/MieAyammmm/recipe/backend/controllers"

	"github.com/gin-gonic/gin"
)

func RecipeRoutes(router *gin.Engine, rc *controllers.RecipeController) {
	router.POST("/recipes", rc.CreateRecipe)
	router.GET("/recipes", rc.GetRecipes)
	router.GET("/recipes/:id", rc.GetRecipeByID)
	router.PUT("/recipes/:id", rc.UpdateRecipe)
	router.DELETE("/recipes/:id", rc.DeleteRecipe)
}