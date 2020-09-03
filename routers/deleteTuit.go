package routers

import (
	"net/http"
	"github.com/didier-gomez/tuitr/bd"
)
func DeleteTuit(w http.ResponseWriter, r*http.Request) {
	ID:= r.URL.Query().Get("id")
	if len(ID) <1 {
		http.Error(w, "ID parameter required", http.StatusBadRequest)
		return
	}
	err := bd.DeleteTuit(ID, UserID)
	if err != nil {
		http.Error(w, "error deleting tuit "+ err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}