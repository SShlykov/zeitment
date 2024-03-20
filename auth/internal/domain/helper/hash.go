package helper

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	// PasswordHashCost TODO: перенести в конфиг
	PasswordHashCost = 14
)

// HashPassword создает хеш пароля
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordHashCost)
	return string(bytes), err
}

func СomparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}
