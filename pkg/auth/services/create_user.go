package services

import (
	"fmt"
	"ordine-api/pkg/auth"
	"ordine-api/pkg/database"
)

func CreateUser(user *auth.AuthRequestBody) error {
	var dbUser auth.User

	if UserExists(user.Username) {
		return fmt.Errorf("user with username %s already exists", user.Username)
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
