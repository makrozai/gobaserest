package bd

import (
	"github.com/makrozai/gobaserest/models"
	"golang.org/x/crypto/bcrypt"
)

// TryLogin realiza el chequeo de login a la BD
func TryLogin(email string, password string) (models.User, bool) {
	user, duplicate, _ := CheckedDuplicateUser(email)
	if !duplicate {
		return user, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return user, false
	}
	return user, true
}
