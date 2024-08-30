package user

import (
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) ([]byte, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return pwd, nil
}
