package handlers

import (
	"net/http"
	"web-service-api/infrastructures/persistance"
	"web-service-api/models"

	"github.com/gin-gonic/gin"
)

var recipes = []models.Recipe{
	{ID: "1", UserID: "abc", DishName: "a", URL: "http", Category: "curry", Media: "YouTube", Repeat: "1", CookingTime: 20, CreatedAt: "2023-01-18 12:30:24.241566", UpdatedAt: "2023-01-18 12:30:24.241566"},
	{ID: "2", UserID: "abc", DishName: "a", URL: "http", Category: "curry", Media: "YouTube", Repeat: "1", CookingTime: 20, CreatedAt: "2023-01-18 12:30:24.241566", UpdatedAt: "2023-01-18 12:30:24.241566"},
	{ID: "3", UserID: "abc", DishName: "a", URL: "http", Category: "curry", Media: "YouTube", Repeat: "1", CookingTime: 20, CreatedAt: "2023-01-18 12:30:24.241566", UpdatedAt: "2023-01-18 12:30:24.241566"},
}

type RecipeHandler struct {
	persistance.RecipeRepository
}

func NewRecipeHandler() *RecipeHandler {
	return &RecipeHandler{}
}

func getAllRecipes(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, recipes)
}
