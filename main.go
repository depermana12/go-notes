package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/depermana12/go-notes/config"
	"github.com/depermana12/go-notes/models"
)

var (
	DB        *gorm.DB
	tokenAuth *jwtauth.JWTAuth
)

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

	tokenAuth = jwtauth.New("HS256", []byte("mie-ayam"), nil)

}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// public routes
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", createUser)
		r.Post("/login", login)
	})

	// protected route
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Route("/api/v1/note", func(r chi.Router) {
			r.Get("/", listNotes)
			r.Get("/{id}", getNoteByID)
			r.Post("/", createNote)
			r.Put("/{id}", updateNote)
			r.Delete("/{id}", deleteNote)
		})
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		http.Error(w, "error hashing password", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)

	err = DB.Create(&user).Error

	if err != nil {
		http.Error(w, "failed to create user", http.StatusBadRequest)
		return
	}

	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{
		"user_id":  user.ID,
		"username": user.Username,
	})

	response := map[string]interface{}{
		"message": "user created",
		"data":    user,
		"token":   tokenString,
	}

	if err = JSONResponse(w, http.StatusCreated, response); err != nil {
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

	err = DB.Create(&note).Error
	if err != nil {
		http.Error(w, "failed to create note", http.StatusBadRequest)
	}

	response := map[string]interface{}{
		"message": "note created",
		"data":    note,
	}

	if err = JSONResponse(w, http.StatusCreated, response); err != nil {
		log.Printf("error sending json response %v", err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user logged in"))
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
