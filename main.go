package main

import (
	"log"

	"github.com/makrozai/gobaserest/bd"
	"github.com/makrozai/gobaserest/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("sin conexion a la BD")
		return
	}
	handlers.Manipulators()
}
