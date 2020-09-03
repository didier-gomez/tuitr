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
	/* auth endpoints */
	router.HandleFunc("/register",middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login",middlew.CheckDB(routers.Login)).Methods("POST")

	/* user endpoints */
	router.HandleFunc("/profile",	middlew.CheckDB(middlew.JwtAuth(routers.FindProfile))).Methods("GET")
	router.HandleFunc("/profile",	middlew.CheckDB(middlew.JwtAuth(routers.UpdateProfile))).Methods("PUT")
	
	router.HandleFunc("/profile/avatar",	middlew.CheckDB(middlew.JwtAuth(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/profile/avatar",	middlew.CheckDB(routers.FindAvatar)).Methods("GET")
	router.HandleFunc("/profile/banner",	middlew.CheckDB(middlew.JwtAuth(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/profile/banner",	middlew.CheckDB(routers.FindBanner)).Methods("GET")

	/* relations*/
	router.HandleFunc("/relation",	middlew.CheckDB(middlew.JwtAuth(routers.CreateRelation))).Methods("POST")
	router.HandleFunc("/relation",	middlew.CheckDB(middlew.JwtAuth(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/relation/check",	middlew.CheckDB(middlew.JwtAuth(routers.CheckRelation))).Methods("GET")

	/* tuit endpoints */
	router.HandleFunc("/tuit",	middlew.CheckDB(middlew.JwtAuth(routers.CreateTuit))).Methods("POST")
	router.HandleFunc("/tuit",	middlew.CheckDB(middlew.JwtAuth(routers.GetTuits))).Methods("GET")
	router.HandleFunc("/tuit",	middlew.CheckDB(middlew.JwtAuth(routers.DeleteTuit))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT ="8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}