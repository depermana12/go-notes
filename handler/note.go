package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/depermana12/go-notes/auth"
	"github.com/depermana12/go-notes/db"
	"github.com/depermana12/go-notes/models"
	"github.com/depermana12/go-notes/service"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm/clause"
)

func JSONResponse(w http.ResponseWriter, statusCode int, value any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(value)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note

	userId, err := auth.GetIdFromAuthCtx(r)
	if err != nil {
		http.Error(w, "invalid user id authctx", http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	newNote, err := service.AddNote(userId, &note)
	if err != nil {
		http.Error(w, "failed to create note", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "note created",
		"data":    newNote,
	}

	if err = JSONResponse(w, http.StatusCreated, response); err != nil {
		fmt.Printf("error sending json response %v", err)
	}
}

func ListNotes(w http.ResponseWriter, r *http.Request) {
	var notes []models.Note

	userId, err := auth.GetIdFromAuthCtx(r)
	if err != nil {
		http.Error(w, "invalid user id authctx", http.StatusUnauthorized)
		return
	}

	if err := db.GetDB().Where("author_id = ?", userId).Find(&notes).Error; err != nil {
		http.Error(w, "failed to get notes", http.StatusBadRequest)
	}

	response := map[string]interface{}{
		"message": "success",
		"data":    notes,
	}

	if err := JSONResponse(w, http.StatusOK, response); err != nil {
		fmt.Printf("error sending json response %v", err)
	}
}

func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	var note models.Note

	userId, err := auth.GetIdFromAuthCtx(r)
	if err != nil {
		http.Error(w, "invalid user id authctx", http.StatusUnauthorized)
		return
	}

	id := chi.URLParam(r, "id")

	if err := db.GetDB().Where("author_id = ? AND id = ?", userId, id).First(&note).Error; err != nil {
		http.Error(w, "failed to get notes", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "success",
		"data":    note,
	}

	if err := JSONResponse(w, http.StatusOK, response); err != nil {
		fmt.Printf("error sending json response %v", err)
	}
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	var UpdateNote struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	err := json.NewDecoder(r.Body).Decode(&UpdateNote)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	userId, err := auth.GetIdFromAuthCtx(r)
	if err != nil {
		http.Error(w, "invalid user id authctx", http.StatusUnauthorized)
		return
	}

	id := chi.URLParam(r, "id")

	err = db.GetDB().Where("author_id = ? AND id = ?", userId, id).First(&note).Error
	if err != nil {
		http.Error(w, "note not found", http.StatusNotFound)
		return
	}

	note.Title = UpdateNote.Title
	note.Content = UpdateNote.Content

	if err = db.GetDB().Save(&note).Error; err != nil {
		http.Error(w, "failed to update note", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "note updated",
		"data":    note,
	}

	if err := JSONResponse(w, http.StatusOK, response); err != nil {
		fmt.Printf("error sending json response %v", err)
	}
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note

	userId, err := auth.GetIdFromAuthCtx(r)
	if err != nil {
		http.Error(w, "invalid user id authctx", http.StatusUnauthorized)
		return
	}

	id := chi.URLParam(r, "id")

	if err := db.GetDB().Clauses(clause.Returning{}).Where("author_id = ? AND id = ?", uint(userId), id).Delete(&note).Error; err != nil {
		http.Error(w, "failed to delete note", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "success deleted note",
		"data":    note,
	}

	if err := JSONResponse(w, http.StatusOK, response); err != nil {
		fmt.Printf("error sending json response %v", err)
	}
}
