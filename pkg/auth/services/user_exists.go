package services

import (
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func UserExists(userName string) bool {
	_, err := getUserFromDb(userName)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorf("Error trying to get user from DB: %s", err.Error())
		}

		return false
	}
	return true
}
