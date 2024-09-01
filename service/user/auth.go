package user

import (
	"time"

	"github.com/Brian-Hsieh/ecomm/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) ([]byte, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return pwd, nil
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString(config.Env.Secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
