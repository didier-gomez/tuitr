package routers

import (
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/didier-gomez/tuitr/bd"
)
func FindFollowedTuits(w http.ResponseWriter, r * http.Request) {
	page := r.URL.Query().Get("page")

	pgTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "'page' parameter required or invalid format", http.StatusBadRequest)
		return
	}
	result, success := bd.FindFollowedTuits(UserID, pgTemp)
	if !success {
		http.Error(w, "error retreiving tuits ", http.StatusBadRequest)
		return
	}
	
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}