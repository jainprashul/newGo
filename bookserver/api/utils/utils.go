package utils

import "github.com/google/uuid"

// Function to Create a UUID
func CreateUUID() string {
	// Create a UUID
	return string(uuid.New().String())
}