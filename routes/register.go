package routes

import (
	"encoding/json"
	"net/http"

	"github.com/makrozai/gobaserest/bd"
	"github.com/makrozai/gobaserest/models"
)

// Register es la funcion para crear en la database el registro de usuario
func Register(rw http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(rw, "error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(rw, "El email de usuario es requerido", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(rw, "El password de usuario es requerido", 400)
		return
	}

	_, duplicateUser, _ := bd.CheckedDuplicateUser(user.Email)
	if duplicateUser {
		http.Error(rw, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertRegister(user)
	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(rw, "No se logro insertar el registro del usuario "+err.Error(), 400)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}
