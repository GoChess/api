package endpoints

import (
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
	w.WriteHeader(200)
}
