package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/didier-gomez/tuitr/bd"
	"github.com/didier-gomez/tuitr/models"
)

func UploadAvatar(w http.ResponseWriter, r * http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var filePath string = "uploads/avatars/" + UserID + "." + extension

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error uploading image! " + err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f,file)
	if err !=nil {
		http.Error(w, "Error uploading image! " + err.Error(), http.StatusBadRequest)
		return
	}
	var usr models.User
	var status bool
	usr.Avatar = UserID + "." + extension
	status, err = bd.UpdateRegister(usr, UserID)
	if err != nil || status == false {
		http.Error(w, "Error uploading file to db " + err.Error(), http.StatusBadRequest)
		return
	}
	
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}