package routes

import (
	"encoding/json"
	"net/http"

	"github.com/makrozai/gobaserest/bd"
)

// GetProfile permite extraer los valores del perfil
func GetUser(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	user, err := bd.SearchUser(ID)
	if err != nil {
		http.Error(rw, "ocurrio un error al intentar buscar el registro"+err.Error(), 400)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(user)
}
