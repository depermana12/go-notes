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

	// public routes
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", createUser)
	})

	r.Route("/api/v1/note", func(r chi.Router) {
		r.Get("/", listNotes)
		r.Get("/{id}", getNoteByID)
		r.Post("/", createNote)
		r.Put("/{id}", updateNote)
		r.Delete("/{id}", deleteNote)
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

	response := map[string]interface{}{
		"message": "user created",
		"data":    user,
	}

	err = JSONResponse(w, http.StatusCreated, response)

	if err != nil {
		log.Printf("error sending json response %v", err)
	}

}

func JSONResponse(w http.ResponseWriter, statusCode int, value any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(value)
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

func listNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listing all notes"))
}

func getNoteByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("Getting note with ID: " + id))
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("Updating note with ID: " + id))
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("Deleting note with ID: " + id))
}
