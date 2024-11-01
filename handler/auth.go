package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/depermana12/go-notes/auth"
	"github.com/depermana12/go-notes/db"
	"github.com/depermana12/go-notes/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var inputLoginForm struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&inputLoginForm); err != nil {
		http.Error(w, "inavlid input login", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.GetDB().Where("email = ?", inputLoginForm.Email).First(&user).Error; err != nil {
		http.Error(w, "invalid user not registered", http.StatusUnauthorized)
		return
	}
	if err := auth.ComparePassword(user.Password, inputLoginForm.Password); err != nil {
		fmt.Printf("mismatch %s", err)
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	tokenString, _ := auth.CreateJWT(user.ID, user.Username)

	response := map[string]interface{}{
		"message": "user logged in",
		"data":    user,
		"token":   tokenString,
	}

	if err := JSONResponse(w, http.StatusOK, response); err != nil {
		fmt.Printf("error sending JSON response: %v", err)
	}
}
