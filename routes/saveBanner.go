package routes

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/makrozai/gobaserest/bd"
	"github.com/makrozai/gobaserest/models"
)

// SaveBanner sube el banner al servidor
func SaveBanner(rw http.ResponseWriter, r *http.Request) {
	file, handler, _ := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var fileName string = "uploads/banners/" + IDUser + "." + extension

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(rw, "error al subir imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(rw, "error al copiar la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Banner = IDUser + "." + extension
	status, err = bd.UpdateUser(user, IDUser)
	if err != nil || !status {
		http.Error(rw, "eerror al grabar el banner en la bd! "+err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}
