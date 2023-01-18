package main

import (
	"net/http"
	"web-service-api/models"
	"web-service-api/routes"

	"github.com/gin-gonic/gin"
)

var recipes = []models.Recipe{
	{ID: "fd", UserID: "abc", DishName: "a", URL: "http", Category: "curry", Media: "YouTube", Repeat: "1", CookingTime: 20, CreatedAt: "2023-01-18 12:30:24.241566", UpdatedAt: "2023-01-18 12:30:24.241566"},
	{ID: "fd", UserID: "abc", DishName: "a", URL: "http", Category: "curry", Media: "YouTube", Repeat: "1", CookingTime: 20, CreatedAt: "2023-01-18 12:30:24.241566", UpdatedAt: "2023-01-18 12:30:24.241566"},
	{ID: "fd", UserID: "abc", DishName: "a", URL: "http", Category: "curry", Media: "YouTube", Repeat: "1", CookingTime: 20, CreatedAt: "2023-01-18 12:30:24.241566", UpdatedAt: "2023-01-18 12:30:24.241566"},
}

func main() {
	router := routes.NewRouter()

	router.SetMiddleware()
	router.SetProxy()
	router.SetHealthChecker()
	// router.Engine.Group("/v1/api")
	// router.Group("/v1/api")
	// router.GET("/recipes", getAllRecipes)
	router.Engine.GET("/v1/api/recipes", getAllRecipes)

	router.Serve()
}

// 全てのレシピデータを取得する
func getAllRecipes(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, recipes)
}
