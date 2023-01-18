package models

// import "time"

type User struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	RememberToken string `json:"remember_token"`
	RefreshToken  string `json:"refresh_token"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	// CreatedAt     time.Time `json:"created_at"`
	// UpdatedAt     time.Time `json:"updated_at"`
}
