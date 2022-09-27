package routes

import (
	"io"
	"net/http"
	"os"

	"github.com/makrozai/gobaserest/bd"
)

// GetAvaGetBannertar envia el banner al http
func GetBanner(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	user, err := bd.SearchUser(ID)
	if err != nil {
		http.Error(rw, "usuario no encontrado", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/banners/" + user.Banner)
	if err != nil {
		http.Error(rw, "imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(rw, openFile)
	if err != nil {
		http.Error(rw, "error al copiar la imagen", http.StatusBadRequest)
		return
	}
}
