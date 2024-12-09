package services

import (
	"errors"
	"ordine-api/pkg/auth"
	"ordine-api/pkg/database"
)

func CreateUser(user *auth.LoginRequestBody) error {
	var dbUser auth.User

	if UserExists(user.Username) {
		return errors.New("User with user name %s alredy exists")
	}

	var err error

	db := database.GetConnector()
	dbUser.Username = user.Username
	dbUser.Password, err = HashPassword(user.Password)
	if err != nil {
		return err
	}
	db.Save(dbUser)

	return nil

}
