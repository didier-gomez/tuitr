package routers

import (
	"net/http"
	"github.com/didier-gomez/tuitr/models"
	"github.com/didier-gomez/tuitr/bd"
)

func DeleteRelation(w http.ResponseWriter, r * http.Request) {
	
	ID:= r.URL.Query().Get("id")
	if len(ID) <1 {
		http.Error(w, "ID parameter required", http.StatusBadRequest)
		return
	}
	var t models.Relation
	t.UserID = UserID
	t.RelatedUserID = ID
	err := bd.DeleteRelation(t)
	if err != nil {
		http.Error(w, "delete relation failed, try again "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}