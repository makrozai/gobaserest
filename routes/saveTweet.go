package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/makrozai/gobaserest/bd"
	"github.com/makrozai/gobaserest/models"
)

// SaveTweet permite grabar el tweet en la database
func SaveTweet(rw http.ResponseWriter, r *http.Request) {
	var message models.RequestTweet
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(rw, "ocurrio un error al obtener el mensaje "+err.Error(), 400)
		return
	}

	messageRegister := models.Tweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.InsertTweet(messageRegister)
	if err != nil {
		http.Error(rw, "ocurrio un error al intentar insertar el registro, reintente nuevamente "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado insertar el Tweet", 400)
	}

	rw.WriteHeader(http.StatusOK)
}
