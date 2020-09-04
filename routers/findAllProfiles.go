package routers

import (
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/didier-gomez/tuitr/bd"
)

func FindAllProfiles(w http.ResponseWriter, r * http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pgTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "'page' parameter required or invalid format", http.StatusBadRequest)
		return
	}
	pag := int64(pgTemp)
	result, status := bd.FindAllProfiles(UserID, pag, search, typeUser)
	if !status {
		http.Error(w, "error retreiving users ", http.StatusBadRequest)
		return
	}
	
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}