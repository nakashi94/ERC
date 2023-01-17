package main

import (
	"net/http"
	"web-service-api/routes"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var recipes = []Recipe{
	{ID: "1", Name: "a"},
	{ID: "2", Name: "b"},
}

func main() {
	router := routes.NewRouter()

	router.SetMiddleware()
	router.SetProxy()
	// router.Engine.Group("/v1/api")
	// router.Group("/v1/api")
	// router.GET("/recipes", getAllRecipes)
	router.Engine.GET("/v1/api/recipes", getAllRecipes)

	router.Engine.Run("localhost:8080")
}

func getAllRecipes(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, recipes)
}
