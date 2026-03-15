package models

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name" validate:"required, min:3, max:30"`
	Email string `json:"email" validate:"required, email"`
}
