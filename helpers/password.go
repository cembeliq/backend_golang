package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword is generate hash password
func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	return string(hash), err
}

// CompareHashAndPassword is compare password hash and plain password
func CompareHashAndPassword(pwd, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))

	if err != nil {
		return false, err
	}

	return true, nil
}
