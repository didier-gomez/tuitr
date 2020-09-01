package middlew

import (
	"net/http"
	"github.com/didier-gomez/tuitr/routers"
)

func JwtAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r* http.Request) {
		_,_,_, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Token error "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w,r)
	}
}