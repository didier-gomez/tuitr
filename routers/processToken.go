package routers

import (
	"errors"
	"strings"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/didier-gomez/tuitr/bd"
	"github.com/didier-gomez/tuitr/models"
)
var Email string
var UserID string

func ProcessToken(token string) (*models.Claim, bool, string, error) {
	key := []byte("Sup3r53cr3T")
	claims := &models.Claim {}
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}
	token = strings.TrimSpace(splitToken[1])
	parsed, err := jwt.ParseWithClaims(token, claims, func(toke * jwt.Token)(interface{}, error) {
		return key, nil
	})
	// token succesfull
	if err == nil {
		usr, found, _ := bd.CheckUserExists(claims.Email)
		if found == true {
			Email = claims.Email
			UserID = usr.ID.Hex()
		}
		return claims, found, UserID, nil
	}
	if !parsed.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}