package routes

import (
	"net/http"
	"web-service-api/models"

	"github.com/gin-gonic/gin"
)

var recipes = []models.Recipe{
	{ID: "1", UserID: "abc", DishName: "a", URL: "http", Category: "curry", Media: "YouTube", Repeat: "1", CookingTime: 20, CreatedAt: "2023-01-18 12:30:24.241566", UpdatedAt: "2023-01-18 12:30:24.241566"},
	{ID: "2", UserID: "abc", DishName: "a", URL: "http", Category: "curry", Media: "YouTube", Repeat: "1", CookingTime: 20, CreatedAt: "2023-01-18 12:30:24.241566", UpdatedAt: "2023-01-18 12:30:24.241566"},
	{ID: "3", UserID: "abc", DishName: "a", URL: "http", Category: "curry", Media: "YouTube", Repeat: "1", CookingTime: 20, CreatedAt: "2023-01-18 12:30:24.241566", UpdatedAt: "2023-01-18 12:30:24.241566"},
}

func (r *Router) NewTaskRouter() {
	g := r.Engine.Group("/v1/api")
	g.GET("/recipes", getAllRecipes)
	g.POST("/recipes", storeRecipe)
	g.GET("/recipes/:id", getRecipeByID)
}

func getAllRecipes(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, recipes)
}

func storeRecipe(ctx *gin.Context) {
	var recipe = models.Recipe{ID: "4", UserID: "abc", DishName: "pie", URL: "http", Category: "curry", Media: "YouTube", Repeat: "1", CookingTime: 20, CreatedAt: "2023-01-18 12:30:24.241566", UpdatedAt: "2023-01-18 12:30:24.241566"}
	recipes = append(recipes, recipe)
	ctx.IndentedJSON(http.StatusCreated, recipes)
}

func getRecipeByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var recipe models.Recipe
	for _, v := range recipes {
		if v.ID == id {
			recipe = v
			break
		}
	}
	ctx.IndentedJSON(http.StatusOK, recipe)
}
