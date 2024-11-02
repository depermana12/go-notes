package main

import (
	"fmt"
	"net/http"

	"github.com/depermana12/go-notes/db"
	"github.com/depermana12/go-notes/router"
)

func init() {
	db.ConnectToDB()
}

func main() {
	addr := ":3000"
	r := router.Router()

	fmt.Printf("server listening on port %s", addr)
	http.ListenAndServe(addr, r)
}
