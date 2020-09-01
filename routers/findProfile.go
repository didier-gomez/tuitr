package routers

import (
	"net/http"
	"encoding/json"
	"github.com/didier-gomez/tuitr/bd"
)
func FindProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) <1 {
		http.Error(w, "ID parameter required", http.StatusBadRequest)
		return
	}
	profile, err := bd.FindProfile(ID)
	if err != nil {
		http.Error(w, "error searching user "+err.Error(), 400)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}