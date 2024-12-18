package services

import (
	"ordine-api/config"
	"ordine-api/pkg/auth"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user *auth.User) (string, error) {

	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expira em 24 horas
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.Envs.API_SECRET_KEY)) // Use uma chave secreta
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
