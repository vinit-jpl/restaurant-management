package utils

import "golang.org/x/crypto/bcrypt"

func Hashpassword(password string) (string, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashed), err
}
