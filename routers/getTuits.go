package routers

import (
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/didier-gomez/tuitr/bd"
)
func GetTuits(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) <1 {
		http.Error(w, "ID parameter required", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("page")) <1 {
		http.Error(w, "'page' parameter required", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "'page' parameter invalid", http.StatusBadRequest)
		return
	}
	pag := int64(page)

	result, status := bd.GetTuits(ID, pag)

	if status == false {
		http.Error(w, "error searching tuits "+err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}