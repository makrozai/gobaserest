package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/makrozai/gobaserest/models"
)

func GenerateJWT(user models.User) (string, error) {
	myKey := []byte("Makrozai_restApi")

	payload := jwt.MapClaims{
		"email":       user.Email,
		"name":        user.Name,
		"last_name":   user.LastName,
		"date_birth":  user.DateBirth,
		"description": user.Description,
		"location":    user.Location,
		"web_site":    user.WebSite,
		"_id":         user.ID.Hex(),
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
