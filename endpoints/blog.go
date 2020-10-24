package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/GoChess/api/models"
)

// RetrievePreviews writes a json object with an array of previews
func RetrievePreviews(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.Test{Test: "Retrieve Blog Previews"})
}

// RetrieveBlog writes a json object that returns the blog post in full
func RetrieveBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.Test{Test: "Retrieve Blog"})
}

// PostBlog posts blog to the db
func PostBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.Test{Test: "Post Blog"})
}

// OverwriteBlog allows you to edit and overwrite the entry
func OverwriteBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.Test{Test: "Overwrite Blog"})
}

// DeleteBlog deletes a given blog post
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.Test{Test: "Delete Blog"})
}
