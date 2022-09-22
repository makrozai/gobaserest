package bd

import "golang.org/x/crypto/bcrypt"

func EncriptPassword(pass string) (string, error) {
	amount := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), amount)
	return string(bytes), err
}
