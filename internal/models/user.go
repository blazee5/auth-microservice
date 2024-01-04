package models

type User struct {
	ID         string `json:"id" db:"id"`
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	Email      string `json:"email" db:"email"`
	Password   string `json:"password" db:"password"`
	Avatar     string `json:"avatar" db:"avatar"`
	IsVerified bool   `json:"is_verified" db:"is_verified"`
}
