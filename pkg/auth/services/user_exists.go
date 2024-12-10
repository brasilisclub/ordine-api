package services

import (
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func UserExists(userName string) bool {
	user, err := getUserFromDb(userName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		logrus.Errorf("Erro ao verificar existência do usuário: %v", err.Error())
		return false
	}
	return user.ID != 0
}
