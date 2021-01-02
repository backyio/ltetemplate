package models

import "time"

type (
	// User is the base model for site user
	User struct {
		Id        string    `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Avatar    string    `json:"avatar"`
		Roles     []string  `json:"roles"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
