package middlew

import (
	"net/http"

	"github.com/makrozai/gobaserest/bd"
)

// interceptor de endpoints, donde recibe el handlerFunc, procesa un validacion y si esta pasa correctamente devuelve el handlerFunc
func CheckedBD(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(rw, "conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
