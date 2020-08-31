package routers
import (
	"encoding/json"
	"net/http"
	"github.com/didier-gomez/tuitr/bd"
	"github.com/didier-gomez/tuitr/models"
)

/* Register creates a user register in db*/
func Register(w http.ResponseWriter, r* http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	
	if(err != nil) {
		http.Error(w, "Invalid data: " +err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email required", 400)
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password  min length is 6", 400)
	}
	_, found, _ := bd.CheckUserExists(t.Email)

	if found == true {
		http.Error(w, "User already exists", 400)
	}
	_, status, err := bd.CreateRegister(t)
	if err != nil {
		http.Error(w, "Error creating user "+ err.Error(), 400)
	}
	if status == false {
		http.Error(w, "User creation failed ", 400)
	}
	w.WriteHeader(http.StatusCreated)

}
