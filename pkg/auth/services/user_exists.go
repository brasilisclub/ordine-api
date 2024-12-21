package services

import (
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func UserExists(userName string) error {
	_, err := getUserFromDb(userName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorf("Username não encontrado: %v", err.Error())
		} else {
			logrus.Errorf("Erro ao verificar existência do usuário: %v", err.Error())
		}

		return err
	}
	return nil
}
