package services

import (
	"ordine-api/config"
	"ordine-api/pkg/auth"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user *auth.User) string {

	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expira em 24 horas
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(config.Envs.JWT_SECRET_TOKEN)) // Use uma chave secreta

	return signedToken
}
