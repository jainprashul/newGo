package auth

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"xpJain.co/bookserver/db"
	"xpJain.co/bookserver/models"
)

var UserDB = db.NewModel("users", &models.User{})



func Login(w http.ResponseWriter, r *http.Request) {
	
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username, ok := data["username"]
	if !ok {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	password, ok := data["password"]
	if !ok {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}

	// Find the user with the given username
	user, err := UserDB.GetByField("email", username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Hash the password
	hashedPassword := HashPassword(password)


	// Check if the password is correct
	if user.Password != hashedPassword {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Create a new token
	token, err := CreateToken(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the token in the response header
	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		token : token,
	})
}


func Register(w http.ResponseWriter, r *http.Request) {
	
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username, ok := data["username"]
	if !ok {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	password, ok := data["password"]
	if !ok {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}

	// Check if the user already exists
	_, err = UserDB.GetByField("email", username)
	if err == nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword := HashPassword(password)

	// Create a new user
	user := &models.User{
		Email:    username,
		Password: hashedPassword,
		Name:    data["name"],
		ID: `users-` + uuid.NewString(),
	}

	// Save the user to the database
	err = UserDB.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
		"username": username,
		"name": data["name"],
	})
}


func Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
