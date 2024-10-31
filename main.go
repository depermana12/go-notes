package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"github.com/depermana12/go-notes/config"
	"github.com/depermana12/go-notes/models"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = config.ConnectToDB()
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = DB.AutoMigrate(&models.User{}, models.Note{})
	if err != nil {
		log.Fatal("failed to migrate schema to database", err)
	}

}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/user", createUser)
		r.Post("/notes", createNote)
	})

	log.Println("server listening to port 3000")
	http.ListenAndServe(":3000", r)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err = DB.Create(&user).Error

	if err != nil {
		http.Error(w, "failed to create user", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func createNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note

	err := json.NewDecoder(r.Body).Decode(&note)

	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err = DB.Create(&note).Error

	if err != nil {
		http.Error(w, "failed to create note", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}
