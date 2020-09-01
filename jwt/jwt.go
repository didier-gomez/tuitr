package jwt

import (
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/didier-gomez/tuitr/models"
)
func GenerateJWT(t models.User) (string, error) {
	payload := jwt.MapClaims {
		"email": t.Email,
		"name": t.Name,
		"surname": t.Surname,
		"bio": t.Bio,
		"location": t.Location,
		"_id": t.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	key := []byte("Sup3r53cr3T")
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}