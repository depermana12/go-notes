package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func handleFunc(r http.ResponseWriter, w *http.Request) {
	r.Write([]byte("hello world"))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handleFunc)

	http.ListenAndServe(":3000", r)
}
