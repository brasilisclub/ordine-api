package services

import (
	"errors"
	"fmt"
	"ordine-api/pkg/auth"
)

func Login(user *auth.User) (string, error) {

	err := UserExists(user.Username)

	if err != nil {
		return "", err
	}

	err = PasswordIsValid(user)

	if err != nil {
		return "", err
	}

	token, err := GenerateJWT(user)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Failed to generate token: %s", err.Error()))
	}

	return token, nil
}
