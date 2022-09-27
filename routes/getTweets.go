package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/makrozai/gobaserest/bd"
)

// GetTweets permite extraer los valores de los tweets
func GetTweets(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "Debe enviar un parametro id", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(rw, "Debe enviar un parametro page", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(rw, "Debe enviar un parametro page con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	pageInt := int64(page)

	response, state := bd.SearchTweets(ID, pageInt)
	if !state {
		http.Error(rw, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response)
}
