package routers

import (
	"encoding/json"
	"net/http"
	"github.com/didier-gomez/tuitr/bd"
	"github.com/didier-gomez/tuitr/jwt"
	"github.com/didier-gomez/tuitr/models"
)

func Login(w http.ResponseWriter, r * http.Request) {
	w.Header().Add("Content-type", "application/json")
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "invalid user/password " +err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email required", 400)
		return
	}
	document, exists := bd.LoginTry(t.Email, t.Password)
	if exists == false {
		http.Error(w, "wrong user or password ", 400)
		return
	}
	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Token generation failed: " + err.Error(), 400)
		return
	}
	result := models.ResponseLogin {
		Token: "Bearer"+jwtKey,
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

	/*
	expirationTime := time.Now().add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie {
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime
	}
	*/

}