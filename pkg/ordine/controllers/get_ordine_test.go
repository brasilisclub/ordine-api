package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"ordine-api/pkg/ordine"
	"ordine-api/pkg/ordine/order"
	"ordine-api/pkg/product"
	"ordine-api/tests"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetOrdine(t *testing.T) {
	tests := []struct {
		name               string
		ordineId           string
		expectedStatusCode int
		setupTest          func()
		dropDownTest       func()
	}{
		{
			name:               "Should return 400, ordine not founded",
			ordineId:           "1",
			expectedStatusCode: http.StatusBadRequest,
			setupTest: func() {
				tests.MakeMigrationsForTests(&ordine.Ordine{})
			},
			dropDownTest: func() {
				tests.DropTablesForTests(&ordine.Ordine{})
			},
		},
		{
			name:               "Should return 500, table not founded",
			ordineId:           "1",
			expectedStatusCode: http.StatusInternalServerError,
			setupTest: func() {
				// Do nothing
			},
			dropDownTest: func() {
				//Do nothing
			},
		},
		{
			name:               "Should return 200, get ordine successfully",
			ordineId:           "1",
			expectedStatusCode: http.StatusOK,
			setupTest: func() {
				tests.MakeMigrationsForTests(&ordine.Ordine{}, product.Product{}, order.OrderProducts{})
				tests.CreateInsertValueForTests(&ordine.Ordine{
					ID:         1,
					Table:      10,
					ClientName: "test",
					Status:     true,
				})
			},
			dropDownTest: func() {
				tests.DropTablesForTests(&ordine.Ordine{}, product.Product{}, order.OrderProducts{})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.setupTest()
			defer tt.dropDownTest()

			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			c.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/ordine/%s", tt.ordineId), nil)
			c.Params = gin.Params{{Key: "id", Value: "1"}}

			GetOrdine(c)

			if rr.Code != tt.expectedStatusCode {
				t.Errorf("Unexpected status code, expected %v got %v", tt.expectedStatusCode, rr.Code)
			}

		})
	}
}
