package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/depermana12/go-notes/auth"
	"github.com/depermana12/go-notes/db"
	"github.com/depermana12/go-notes/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	hashedPwd, err := auth.HashedPassword(user.Password)
	if err != nil {
		http.Error(w, "error hashing password", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPwd
	if err := db.GetDB().Create(&user).Error; err != nil {
		http.Error(w, "failed to create user", http.StatusBadRequest)
		return
	}

	tokenString, _ := auth.CreateJWT(user.ID, user.Username)

	response := map[string]interface{}{
		"message": "user created",
		"data":    user,
		"token":   tokenString,
	}

	if err := JSONResponse(w, http.StatusCreated, response); err != nil {
		fmt.Printf("error sending json response %v", err)
	}
}
