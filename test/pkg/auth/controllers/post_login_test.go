package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"ordine-api/pkg/auth"
	"ordine-api/pkg/auth/controllers"
	"ordine-api/pkg/auth/services"
	"ordine-api/pkg/database"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setUp() {
	database.TestConnect()
	db := database.GetConnector()

	db.AutoMigrate(&auth.User{})
	passwd, _ := services.HashPassword("test")
	db.Save(&auth.User{Username: "test", Password: passwd})
}

func TestPostLogin(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "Valid login request",
			body:       `{"username": "test", "password": "test"}`,
			wantStatus: http.StatusOK,
			wantBody:   `"token":`, // Exemplo: valida apenas se contém a chave "token".
		},
		{
			name:       "Invalid input",
			body:       `{"username": "testuser"}`, // Falta "password".
			wantStatus: http.StatusBadRequest,
			wantBody:   `"message":"Invalid input:`,
		},
		{
			name:       "Invalid username",
			body:       `{"username": "error", "password": "test"}`,
			wantStatus: http.StatusUnauthorized,
			wantBody:   `"message":"Invalid credentials:`, // Exemplo: valida apenas se contém a chave "token".
		},
		{
			name:       "Invalid password",
			body:       `{"username": "test", "password": "error"}`,
			wantStatus: http.StatusUnauthorized,
			wantBody:   `"message":"Invalid credentials:`, // Exemplo: valida apenas se contém a chave "token".
		},
	}

	setUp()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Configura o mock do contexto e da requisição.
			gin.SetMode(gin.TestMode)
			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)

			// Configura o corpo da requisição.
			req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			// Executa a função.
			controllers.PostLogin(ctx)
			fmt.Println(recorder.Body.String())
			// Verifica o status e o corpo da resposta.
			assert.Equal(t, tt.wantStatus, recorder.Code)
			assert.Contains(t, recorder.Body.String(), tt.wantBody)
		})
	}
}
