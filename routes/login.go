package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/makrozai/gobaserest/bd"
	"github.com/makrozai/gobaserest/jwt"
	"github.com/makrozai/gobaserest/models"
)

// Login realiza el login
func Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(rw, "usuario y/o contrasena invalidos"+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(rw, "el email del usuario es requerido", 400)
		return
	}
	document, exist := bd.TryLogin(user.Email, user.Password)
	if !exist {
		http.Error(rw, "usuario y/o contrasena no coinciden", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(rw, "ocurrio un erro al generar el token correspondiente"+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
