package main

import (
	"net/http"
	"time"

	"github.com/GoChess/api/router"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httprate"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	r.Mount("/public", router.PublicRouter())
	r.Mount("/private", router.PrivateRouter())
	r.Mount("/admin", router.AdminRouter())
	http.ListenAndServe(":3000", r)
}
