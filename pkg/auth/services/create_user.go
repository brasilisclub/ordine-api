package services

import (
	"ordine-api/pkg/auth"
	"ordine-api/pkg/database"
)

func CreateUser(user *auth.AuthRequestBody) error {
	var dbUser auth.User

	err := UserExists(user.Username)

	if err != nil {
		return err
	}

	db := database.GetConnector()
	dbUser.Username = user.Username

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	dbUser.Password = hashedPassword

	if err := db.Save(&dbUser).Error; err != nil {
		return err
	}

	return nil
}
