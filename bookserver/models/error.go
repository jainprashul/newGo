package models

type Error struct {
	// Error message
	Message string `json:"message"`
	// Error code
	Code int `json:"code"`
}