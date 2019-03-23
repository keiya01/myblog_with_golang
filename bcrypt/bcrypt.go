package bcrypt

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func EncryptPassword(password string) (string, error) {
	plainPassword := []byte(password)
	encryptedPassword, err := bcrypt.GenerateFromPassword(plainPassword, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "EncryptPassword")
	}

	return string(encryptedPassword), nil
}

func ComparePassword(inputPassword string, encryptedPassword string) bool {
	plainPassword := []byte(inputPassword)
	hashedPassword := []byte(encryptedPassword)

	err := bcrypt.CompareHashAndPassword(hashedPassword, plainPassword)
	if err != nil {
		log.Printf("Occured error in ComparePassword: %v", err)
		return false
	}

	return true
}
