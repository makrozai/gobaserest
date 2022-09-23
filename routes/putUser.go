package routes

import (
	"encoding/json"
	"net/http"

	"github.com/makrozai/gobaserest/bd"
	"github.com/makrozai/gobaserest/models"
)

// PutUser modifica el usuario
func PutUser(rw http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(rw, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.UpdateUser(user, IDUser)
	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar modificar el usuario. Reintente nuevamente "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(rw, "no se ha logrado modificar el registro de usuario", 400)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
