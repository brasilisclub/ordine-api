package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ordine-api/pkg/auth"
	"ordine-api/pkg/auth/services"
	"ordine-api/pkg/database"
	"ordine-api/pkg/migrations"
	"testing"

	"github.com/gin-gonic/gin"
)

func setUpPostLoginTest() {
	migrations.Migrate() // perguntar ao Gritzko
	c := database.GetConnector()
	hp, err := services.HashPassword("validpassword")
	if err != nil {
		panic("cant hashed the password")
	}
	user := auth.User{Username: "validuser", Password: hp}
	c.Create(&user)
}

func dropDownPostLoginTest() {
	c := database.GetConnector()
	c.Where("username = ?", "validuser").Delete(&auth.User{})
}

func TestPostLogin(t *testing.T) {

	tests := []struct {
		name           string
		requestBody    auth.AuthRequestBody
		expectedStatus int
	}{
		{
			name: "Success - Valid Credentials",
			requestBody: auth.AuthRequestBody{
				Username: "validuser",
				Password: "validpassword",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Failure - Invalid Credentials",
			requestBody: auth.AuthRequestBody{
				Username: "wronguser",
				Password: "wrongpassword",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Failure - Invalid Input",
			requestBody:    auth.AuthRequestBody{},
			expectedStatus: http.StatusBadRequest,
		},
	}
	setUpPostLoginTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			jsonBody, _ := json.Marshal(tt.requestBody)
			c.Request = httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(jsonBody))
			c.Request.Header.Set("Content-Type", "application/json")

			PostLogin(c)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
	dropDownPostLoginTest()
}
