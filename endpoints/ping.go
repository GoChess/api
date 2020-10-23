package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/GoChess/api/models"
)

// Ping returns with an alive status if all things are working
func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.Ping{Pong: true})
}
