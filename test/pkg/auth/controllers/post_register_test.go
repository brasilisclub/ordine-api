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

func setUpRegisterTests() {
	database.TestConnect()
	db := database.GetConnector()

	db.AutoMigrate(&auth.User{})
}

func setUpUserAlreadyExistsRegisterTests() {
	db := database.GetConnector()

	passwd, _ := services.HashPassword("test")
	db.Save(&auth.User{Username: "test", Password: passwd})
}

func TestPostRegister(t *testing.T) {

	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "Success register",
			body:       `{"username": "test", "password": "test"}`,
			wantStatus: http.StatusCreated,
			wantBody:   `{"username":"test","password":"test"}`,
		},
		{
			name:       "Error: invalid body",
			body:       `{"username": "test"}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"message":"Invalid body`,
		},
	}
	setUpRegisterTests()
	defer tearDown()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)

			req, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			controllers.PostRegister(ctx)
			fmt.Println(recorder.Body.String())

			assert.Equal(t, tt.wantStatus, recorder.Code)
			assert.Contains(t, recorder.Body.String(), tt.wantBody)
		})
	}

}

func TestPostRegisterUserAlreadyExists(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "User already exists",
			body:       `{"username": "test", "password": "test"}`,
			wantStatus: http.StatusConflict,
			wantBody:   auth.ErrorUserAlreadyExists.Error(),
		},
	}
	setUpRegisterTests()
	setUpUserAlreadyExistsRegisterTests()
	defer tearDown()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)

			req, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			controllers.PostRegister(ctx)
			fmt.Println(recorder.Body.String())

			assert.Equal(t, tt.wantStatus, recorder.Code)
			assert.Contains(t, recorder.Body.String(), tt.wantBody)
		})
	}

}
