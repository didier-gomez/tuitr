package handlers
import (
	"log"
	"net/http"
	"os"
	"github.com/didier-gomez/tuitr/middlew"
	"github.com/didier-gomez/tuitr/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Handlers genera un router y le agrega el middleware de CORS*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register",middlew.CheckDB(routers.Register)).Methods("POST")



	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT ="8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}