package router

import (
	"net/http"

	"github.com/depermana12/go-notes/auth"
	"github.com/depermana12/go-notes/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

func Router() http.Handler {
	jwt := auth.GetTokenAuth()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// public routes
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", handler.CreateUser)
		r.Post("/login", handler.Login)
	})

	// protected route
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwt))
		r.Use(jwtauth.Authenticator(jwt))

		r.Route("/api/v1/note", func(r chi.Router) {
			r.Get("/", handler.ListNotes)
			r.Get("/{id}", handler.GetNoteByID)
			r.Post("/", handler.CreateNote)
			r.Put("/{id}", handler.UpdateNote)
			r.Delete("/{id}", handler.DeleteNote)
		})
	})

	return r
}
