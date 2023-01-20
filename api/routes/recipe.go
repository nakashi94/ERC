package routes

import (
	"net/http"
	"web-service-api/models"

	"github.com/gin-gonic/gin"
)

func (r *Router) NewTaskRouter() {
	g := r.Engine.Group("/v1/api")
	g.GET("recipes", getAllRecipes)
}

func getAllRecipes(ctx *gin.Context) {
	var recipes = []models.Recipe{
		{ID: "fd", UserID: "abc", DishName: "a", URL: "http", Category: "curry", Media: "YouTube", Repeat: "1", CookingTime: 20, CreatedAt: "2023-01-18 12:30:24.241566", UpdatedAt: "2023-01-18 12:30:24.241566"},
		{ID: "fd", UserID: "abc", DishName: "a", URL: "http", Category: "curry", Media: "YouTube", Repeat: "1", CookingTime: 20, CreatedAt: "2023-01-18 12:30:24.241566", UpdatedAt: "2023-01-18 12:30:24.241566"},
		{ID: "fd", UserID: "abc", DishName: "a", URL: "http", Category: "curry", Media: "YouTube", Repeat: "1", CookingTime: 20, CreatedAt: "2023-01-18 12:30:24.241566", UpdatedAt: "2023-01-18 12:30:24.241566"},
	}
	ctx.IndentedJSON(http.StatusOK, recipes)
}
