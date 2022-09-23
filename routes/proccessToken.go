package routes

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/makrozai/gobaserest/bd"
	"github.com/makrozai/gobaserest/models"
)

// Email valor de email usado en todos los endpoints
var Email string

// IDUser es el id devuelte del modelo, que se usara en todos los enpoints
var IDUser string

// ProccessToken proceso token para extraer sus valores
func ProccessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("Makrozai_restApi")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, checkedUser, _ := bd.CheckedDuplicateUser(claims.Email)
		if checkedUser {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, checkedUser, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
