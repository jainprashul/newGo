// User struct

package models

import (
	"time"
)

// User struct
type User struct {
	// ID
	ID int `json:"id"`
	// Name
	Name string `json:"name"`
	// Email
	Email string `json:"email"`
	// Password
	Password string `json:"password"`
	// CreatedAt
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt
	UpdatedAt time.Time `json:"updated_at"`
}

// UserResponse struct
type UserResponse struct {
	// ID
	ID int `json:"id"`
	// Name
	Name string `json:"name"`
	// Email
	Email string `json:"email"`
	// CreatedAt
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt
	UpdatedAt time.Time `json:"updated_at"`
}
