package validator

import (
	"net/mail"
)

func EmailValidator(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func PasswordValidator(password string) error {
	return nil
}
