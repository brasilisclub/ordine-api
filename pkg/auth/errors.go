package auth

import "errors"

var (
	ErrorUserNotFound      = errors.New("A user with this username was not found")
	ErrorUserAlreadyExists = errors.New("There is already a user with this username")
)
