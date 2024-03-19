package helper

import (
	"errors"
	"regexp"
)

// ValidateLogin проверяет длину имени пользователя.
func ValidateLogin(username string) error {
	if len(username) < 3 || len(username) > 16 {
		return errors.New("имя пользователя должно быть от 3 до 16 символов")
	}
	// TODO: проверка на уникальность
	// TODO: Проверка на запрещенные символы или слова
	return nil
}

// ValidatePassword проверяет пароль на длину, наличие цифр, заглавных и строчных букв, специальных символов.
func ValidatePassword(password string) error {
	if len(password) < 6 || len(password) > 64 {
		return errors.New("пароль должен быть от 8 до 64 символов")
	}

	hasNumber := regexp.MustCompile(`[0-9]+`).MatchString
	hasUpper := regexp.MustCompile(`[A-Z]+`).MatchString
	hasLower := regexp.MustCompile(`[a-z]+`).MatchString
	hasSpecial := regexp.MustCompile(`[!@#$%^&*]+`).MatchString

	if !hasNumber(password) || !hasUpper(password) || !hasLower(password) || !hasSpecial(password) {
		return errors.New("пароль должен содержать хотя бы одну цифру, одну заглавную и строчную букву, а также один специальный символ")
	}

	// TODO: игнорировать общеизвестные пароли

	return nil
}
