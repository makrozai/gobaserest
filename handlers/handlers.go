package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/makrozai/gobaserest/middlew"
	"github.com/makrozai/gobaserest/routes"
	"github.com/rs/cors"
)

// Manipulators setea mi puerto, el handler y pone a escuchar al servidor
func Manipulators() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckedBD(routes.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckedBD(routes.Login)).Methods("POST")
	router.HandleFunc("/user", middlew.CheckedBD(middlew.ValidateJWT(routes.GetUser))).Methods("GET")
	router.HandleFunc("/user", middlew.CheckedBD(middlew.ValidateJWT(routes.PutUser))).Methods("PUT")
	router.HandleFunc("/tweets", middlew.CheckedBD(middlew.ValidateJWT(routes.SaveTweet))).Methods("POST")
	router.HandleFunc("/tweets", middlew.CheckedBD(middlew.ValidateJWT(routes.GetTweets))).Methods("GET")
	router.HandleFunc("/tweets", middlew.CheckedBD(middlew.ValidateJWT(routes.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/avatar", middlew.CheckedBD(middlew.ValidateJWT(routes.SaveAvatar))).Methods("POST")
	router.HandleFunc("/avatar", middlew.CheckedBD(middlew.ValidateJWT(routes.GetAvatar))).Methods("GET")
	router.HandleFunc("/banner", middlew.CheckedBD(middlew.ValidateJWT(routes.SaveBanner))).Methods("POST")
	router.HandleFunc("/banner", middlew.CheckedBD(middlew.ValidateJWT(routes.GetBanner))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
