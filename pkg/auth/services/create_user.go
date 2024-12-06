package services

import (
	"errors"
	"ordine-api/pkg/auth"
	"ordine-api/pkg/database"
)

func CreateUser(user *auth.User) error {
	if UserExists(user.Username) {
		return errors.New("User with user name %s alredy exists")
	}

	var err error

	db := database.GetConnector()
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		return err
	}
	db.Save(user)

	return nil

}
