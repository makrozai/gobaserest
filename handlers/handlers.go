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

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
