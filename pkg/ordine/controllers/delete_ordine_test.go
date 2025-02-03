package controllers

import (
	"net/http"
	"net/http/httptest"
	"ordine-api/pkg/database"
	"ordine-api/pkg/ordine"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestDeleteOrdine(t *testing.T) {
	tests := []struct {
		name               string
		ordineId           int
		expectedStatusCode int
		setUpTest          func()
		dropDownTest       func()
	}{
		{
			name:               "Should return 404, ordine not founded",
			ordineId:           1,
			expectedStatusCode: http.StatusBadRequest,
			setUpTest: func() {
				c := database.GetConnector()
				err := c.AutoMigrate(&ordine.Ordine{})
				if err != nil {
					logrus.Error(err.Error())
				}
			},
			dropDownTest: func() {
				c := database.GetConnector()
				err := c.Migrator().DropTable(&ordine.Ordine{})
				if err != nil {
					logrus.Error(err.Error())
				}
			},
		},
		{
			name:               "Should return 500, internal error, table not found",
			ordineId:           1,
			expectedStatusCode: http.StatusInternalServerError,
			setUpTest: func() {
				// Do nothing
			},
			dropDownTest: func() {
				// Do nothing
			},
		},
		{
			name:               "Should delete the ordine successfully",
			ordineId:           1,
			expectedStatusCode: http.StatusOK,
			setUpTest: func() {
				c := database.GetConnector()
				err := c.AutoMigrate(&ordine.Ordine{})
				if err != nil {
					logrus.Error(err.Error())
				}
			},
			dropDownTest: func() {
				c := database.GetConnector()
				err := c.Migrator().DropTable(&ordine.Ordine{})
				if err != nil {
					logrus.Error(err.Error())
				}
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			tt.setUpTest()
			defer tt.dropDownTest()

			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			c.Request = httptest.NewRequest(http.MethodDelete, "/ordine/1", nil)
			c.Request.Header.Set("Content-Type", "application/json")

			DeleteOrdine(c)

			if rr.Code != tt.expectedStatusCode {
				t.Errorf("status code error, expected %v got %v", tt.expectedStatusCode, rr.Code)
			}

		})

	}
}
