package services

import (
	"ordine-api/pkg/auth"
	"ordine-api/pkg/database"
)

func getUserFromDb(userName string) (auth.User, error) {
	db := database.GetConnector()
	var user auth.User
	err := db.Where("username = ?", userName).First(&user).Error
	return user, err

}
