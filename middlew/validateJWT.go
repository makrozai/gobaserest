package middlew

import (
	"net/http"

	"github.com/makrozai/gobaserest/routes"
)

// ValidateJWT permite validar el JWT que nos viene en la peticion
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, _, _, err := routes.ProccessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(rw, "error en el token !"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
