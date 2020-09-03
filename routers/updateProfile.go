package routers

import (
	"net/http"
	"encoding/json"
	"github.com/didier-gomez/tuitr/bd"
	"github.com/didier-gomez/tuitr/models"
)
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "invalid data "+err.Error(), http.StatusBadRequest)
		return
	}
	var status bool
	status, err = bd.UpdateRegister(t, UserID)

	if (err != nil || status == false) {
		http.Error(w, "error updating user "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}