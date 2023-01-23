package persistance

import (
	"context"
	"web-service-api/infrastructures/database"
	"web-service-api/models"
)

type RecipeRepository struct {
	conn *database.Conn
}

func NewTaskRepository(conn *database.Conn) *RecipeRepository {
	return &RecipeRepository{
		conn: conn,
	}
}

func (r *RecipeRepository) GetAllRecipes() ([]models.Recipe, error) {
	var recipes []models.Recipe
	query := "select * from recipes"
	rows, err := r.conn.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var recipe models.Recipe
		err = rows.Scan(&recipe.ID, &recipe.UserID, &recipe.DishName, &recipe.URL, &recipe.Category, &recipe.Media, &recipe.Repeat, &recipe.CookingTime, &recipe.CreatedAt, &recipe.UpdatedAt)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}
