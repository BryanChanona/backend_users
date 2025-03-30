package helpers

import (


	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	pass := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)

	if err != nil {
        panic(err)
    }

	return string(hash), nil
}