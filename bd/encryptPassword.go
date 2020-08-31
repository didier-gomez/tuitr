package bd
import (
	"golang.org/x/crypto/bcrypt"
)
func EncryptPassword(pwd string) (string, error) {
	cost := 7
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	return string(bytes), err
}