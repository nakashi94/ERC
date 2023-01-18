package models

// import "time"

type Recipe struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	DishName    string `json:"dish_name"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	Media       string `json:"media"`
	Repeat      string `json:"repeat"`
	CookingTime int    `json:"cooking_time"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	// CreatedAt   time.Time `json:"created_at"`
	// UpdatedAt   time.Time `json:"updated_at"`
}
