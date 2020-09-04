package routers

import (
	"net/http"
	"encoding/json"
		"fmt"
	"github.com/didier-gomez/tuitr/models"
	"github.com/didier-gomez/tuitr/bd"
)

func CheckRelation(w http.ResponseWriter, r * http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relation
	t.UserID = UserID
	t.RelatedUserID = ID

	status, err := bd.CheckRelation(t)

	fmt.Println("status ", status, "err ", err)
	var res models.CheckRelationResponse
	if err != nil || status == false {
		res.Status = false
	} else {
		res.Status = true
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}