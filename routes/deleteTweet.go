package routes

import (
	"net/http"

	"github.com/makrozai/gobaserest/bd"
)

// DeleteTweet permite borrar un tweet determinado
func DeleteTweet(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(ID, IDUser)
	if err != nil {
		http.Error(rw, "ocurrio un error al intentar borrar el tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
}
