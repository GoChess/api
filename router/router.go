package router

import (
	"net/http"

	"github.com/GoChess/api/endpoints"
	"github.com/go-chi/chi"
)

// PublicRouter all public endpoints are available here
func PublicRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", endpoints.Ping)
	r.Route("/blog", func(r chi.Router) {
		r.Get("/", endpoints.RetrievePreviews)
		r.Get("/{postID}", endpoints.RetrieveBlog)
	})
	return r
}

// PrivateRouter all private endpoints are available here, these often need context passed
func PrivateRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", endpoints.Ping)
	return r
}

// AdminRouter all administrative endpoints are available here, these are restricted to the website administrators
func AdminRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", endpoints.Ping)
	r.Route("/blog", func(r chi.Router) {
		r.Post("/", endpoints.PostBlog)
		r.Delete("/{postID}", endpoints.DeleteBlog)
		r.Patch("/{postID}", endpoints.OverwriteBlog)
	})
	return r
}
