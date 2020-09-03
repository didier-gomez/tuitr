package routers

import (
	"net/http"
	"io"
	"os"
	"github.com/didier-gomez/tuitr/bd"
)
func FindBanner(w http.ResponseWriter, r * http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id parameter required", http.StatusBadRequest)
		return
	}
	profile, err := bd.FindProfile(ID)
	if err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}
	openFile, err := os.Open("uploads/banners/"+profile.Banner)
	if err!= nil {
		http.Error(w, "image not found", http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, openFile)
	if err!= nil {
		http.Error(w, "error copying img", http.StatusInternalServerError)
	}
}