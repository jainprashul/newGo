// User struct

package models

import (
	"time"

	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	ID string `json:"id"`
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


func (u User) GetID() string {
	return string(u.ID)
}

func (u *User) SetID(id string) {
	u.ID = id
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
