package bd

import (
	"github.com/JosseMontano/InstagramReactGolang/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentLogin(gmail string, password string) (models.User, bool) {
	use, find, _ := CheckExitsUser(gmail)
	if !find {
		return use, false
	}
	passBytes := []byte(password)
	passBD := []byte(use.Password)
	err := bcrypt.CompareHashAndPassword(passBD, passBytes)
	if err != nil {
		return use, false
	}
	return use, true
}
