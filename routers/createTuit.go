package routers

import (
	"net/http"
	"encoding/json"
	"time"

	"github.com/didier-gomez/tuitr/models"
	"github.com/didier-gomez/tuitr/bd"
)

func CreateTuit(w http.ResponseWriter, r * http.Request) {
	var msg models.Tuit
	err := json.NewDecoder(r.Body).Decode(&msg)
	register := models.TuitInsert {
		UserID: UserID,
		Message: msg.Message,
		Date: time.Now(),
	}
	_, status, err := bd.CreateTuit(register)
	if err != nil {
		http.Error(w, "Insert tuit failed, try again "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Insert tuit failed ", 400)
		return
	}
}