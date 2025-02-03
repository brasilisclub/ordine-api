package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ordine-api/pkg/auth"
	"ordine-api/pkg/auth/services"
	"ordine-api/tests"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPostRegister(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    auth.AuthRequestBody
		expectedStatus int
		setUpTest      func()
		dropDownTest   func()
	}{
		{
			name:           "Bad Request - invalid Body",
			requestBody:    auth.AuthRequestBody{},
			expectedStatus: http.StatusBadRequest,
			setUpTest: func() {
				// Do nothing
			},
			dropDownTest: func() {
				// Do nothing
			},
		},
		{
			name: "Conflict user already exists",
			requestBody: auth.AuthRequestBody{
				Username: "test",
				Password: "test",
			},
			expectedStatus: http.StatusConflict,
			setUpTest: func() {
				tests.MakeMigrationsForTests(&auth.User{})
				hp, _ := services.HashPassword("test")

				tests.CreateInsertValueForTests(&auth.User{Username: "test", Password: hp})

			},
			dropDownTest: func() {
				tests.DropTablesForTests(&auth.User{})
			},
		},
		{
			name:           "Internal Server error Table not found",
			requestBody:    auth.AuthRequestBody{Username: "test", Password: "test"},
			expectedStatus: http.StatusInternalServerError,
			setUpTest: func() {
				// Do nothing
			},
			dropDownTest: func() {
				// Do nothing
			},
		},
		{
			name:           "Created - Should Create a new user",
			requestBody:    auth.AuthRequestBody{Username: "test", Password: "test"},
			expectedStatus: http.StatusCreated,
			setUpTest: func() {
				tests.MakeMigrationsForTests(&auth.User{})
			},
			dropDownTest: func() {
				tests.DropTablesForTests(&auth.User{})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUpTest()
			defer tt.dropDownTest()

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			jsonBody, _ := json.Marshal(tt.requestBody)
			c.Request = httptest.NewRequest(
				http.MethodPost,
				"/auth/register",
				bytes.NewBuffer(jsonBody),
			)
			c.Request.Header.Set("Content-Type", "application/json")

			PostRegister(c)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}
