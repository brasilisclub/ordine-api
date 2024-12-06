package services

import (
	"errors"
	"fmt"
	"ordine-api/pkg/auth"
)

func Login(user *auth.User) (string, error) {

	if !UserExists(user.Username) {
		return "", errors.New(fmt.Sprintf("User %s not exists", user.Username))
	}

	if !PasswordIsValid(user) {
		return "", errors.New("Invalid password")
	}

	token, err := GenerateJWT(user)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Failed to generate token: %s", err.Error()))
	}

	return token, nil
}
