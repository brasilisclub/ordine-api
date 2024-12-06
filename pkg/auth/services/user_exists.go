package services

import (
	"ordine-api/pkg/auth"
	"ordine-api/pkg/database"
)

func UserExists(userName string) bool {
	db := database.GetConnector()
	var user auth.User
	if err := db.Where("username = ?", userName).First(&user).Error; err != nil {
		return false
	}
	return true
}
