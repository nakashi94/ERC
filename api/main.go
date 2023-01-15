package main

import (
	"fmt"
	"net/http"

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

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	fmt.Println("Hello GO!")

	router := gin.Default()
	router.GET("/v1/api/recipes", getAllRecipes)
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

func getAllRecipes(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, recipes)
}

func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}
