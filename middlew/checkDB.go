package middlew

import(
	"net/http"
	"github.com/didier-gomez/tuitr/bd"
)

/* CHeckDB middleware que valida la conexión a la bd antes de empezar a procesar*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r * http.Request){
		if(bd.CheckConnection() == 0) {
			http.Error(w, "Conexión perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w,r);
	}
}