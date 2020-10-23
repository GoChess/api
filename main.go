package main

import (
	"net/http"

	"github.com/GoChess/api/endpoints"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", endpoints.Ping)
	http.ListenAndServe(":3000", r)
}
