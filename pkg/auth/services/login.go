package services

import (
	"ordine-api/pkg/auth"
)

func Login(user *auth.User) (string, error) {

	if !UserExists(user.Username) {
		return "", auth.ErrorUserNotFound
	}

	err := PasswordIsValid(user)

	if err != nil {
		return "", err
	}

	token := GenerateJWT(user)

	return token, nil
}
