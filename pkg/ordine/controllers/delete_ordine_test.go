package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"ordine-api/pkg/ordine"
	"ordine-api/tests"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDeleteOrdine(t *testing.T) {
	tests := []struct {
		name               string
		ordineId           string
		expectedStatusCode int
		setUpTest          func()
		dropDownTest       func()
	}{
		{
			name:               "Should return 404, ordine not founded",
			ordineId:           "1",
			expectedStatusCode: http.StatusBadRequest,
			setUpTest: func() {
				tests.MakeMigrationsForTests(&ordine.Ordine{})
			},
			dropDownTest: func() {
				tests.DropTablesForTests(&ordine.Ordine{})

			},
		},
		{
			name:               "Should return 500, internal error, table not found",
			ordineId:           "1",
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
			ordineId:           "1",
			expectedStatusCode: http.StatusOK,
			setUpTest: func() {
				tests.MakeMigrationsForTests(&ordine.Ordine{})

				tests.CreateInsertValueForTests(&ordine.Ordine{
					ID:         1,
					Table:      10,
					ClientName: "Test",
					Status:     false,
				})

			},
			dropDownTest: func() {
				tests.DropTablesForTests(&ordine.Ordine{})
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			tt.setUpTest()
			defer tt.dropDownTest()

			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			c.Request = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/ordine/%s", tt.ordineId), nil)
			c.Params = []gin.Param{{Key: "id", Value: tt.ordineId}}

			DeleteOrdine(c)

			if rr.Code != tt.expectedStatusCode {
				t.Errorf("status code error, expected %v got %v", tt.expectedStatusCode, rr.Code)
			}

		})

	}
}
