package bd

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/didier-gomez/tuitr/models"
)

func LoginTry(email string, password string) (models.User, bool) {
	foundUser, found,_ := CheckUserExists(email)
	if found == false {
		return foundUser, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(foundUser.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return foundUser, false
	}
	return foundUser, true
}