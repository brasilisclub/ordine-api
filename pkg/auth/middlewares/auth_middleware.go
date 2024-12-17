package middleware

import (
	"net/http"
	"ordine-api/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Chave secreta usada para assinar o token
var jwtSecret = []byte(config.Envs.API_SECRET_KEY)

// Middleware para validar JWT
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extrair o token do cabeçalho Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		// Verifica se o token é fornecido no formato correto
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token deve usar o esquema Bearer"})
			c.Abort()
			return
		}

		// Validar o token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verificar o método de assinatura do token
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		// Verificar erros de parsing ou se o token é inválido
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Passa o token decodificado para o contexto, caso precise acessar claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token malformado"})
			c.Abort()
			return
		}

		// Armazenar informações dos claims no contexto
		c.Set("userID", claims["user_id"])
		c.Next()
	}
}
