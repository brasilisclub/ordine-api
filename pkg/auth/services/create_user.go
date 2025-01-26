package services

import (
	"ordine-api/pkg/auth"
	"ordine-api/pkg/database"
)

func CreateUser(user *auth.AuthRequestBody) error {
	var dbUser auth.User

	if UserExists(user.Username) {
		return auth.ErrorUserAlreadyExists
	}

	db := database.GetConnector()
	dbUser.Username = user.Username

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	dbUser.Password = hashedPassword
	err = db.Save(&dbUser).Error

	return err
}
