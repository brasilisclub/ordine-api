package services

import (
	"ordine-api/pkg/auth"

	"golang.org/x/crypto/bcrypt"
)

func PasswordIsValid(user *auth.User) bool {
	dbUser, err := getUserFromDb(user.Username)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	return err == nil
}