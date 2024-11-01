package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/depermana12/go-notes/db"
	"github.com/depermana12/go-notes/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func JSONResponse(w http.ResponseWriter, statusCode int, value any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(value)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note

	_, claims, _ := jwtauth.FromContext(r.Context())
	userId, ok := claims["user_id"].(float64)
	if !ok {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	note.AuthorId = uint(userId)

	err = db.GetDB().Create(&note).Error
	if err != nil {
		http.Error(w, "failed to create note", http.StatusBadRequest)
	}

	response := map[string]interface{}{
		"message": "note created",
		"data":    note,
	}

	if err = JSONResponse(w, http.StatusCreated, response); err != nil {
		fmt.Printf("error sending json response %v", err)
	}
}

func ListNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listing all notes"))
}

func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("Getting note with ID: " + id))
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("Updating note with ID: " + id))
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("Deleting note with ID: " + id))
}
