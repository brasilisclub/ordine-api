package controllers

import (
	"net/http"
	"net/http/httptest"
	"ordine-api/pkg/ordine"
	"ordine-api/pkg/ordine/order"
	"ordine-api/pkg/product"
	"ordine-api/tests"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetOrdines(t *testing.T) {

	tests := []struct {
		name               string
		expectedStatusCode int
		setUpTest          func()
		dropDownTest       func()
	}{
		{
			name:               "Should return 500, no table",
			expectedStatusCode: http.StatusInternalServerError,
			setUpTest: func() {
				// Do nothing
			},
			dropDownTest: func() {
				// Do nothing
			},
		},
		{
			name:               "Should return 200, get ordines successfully",
			expectedStatusCode: http.StatusOK,
			setUpTest: func() {
				tests.MakeMigrationsForTests(&product.Product{}, &ordine.Ordine{}, &order.OrderProducts{})
				tests.CreateInsertValueForTests(&ordine.Ordine{
					ID:         1,
					Table:      10,
					ClientName: "Test 1",
					Status:     true,
				})
				tests.CreateInsertValueForTests(&ordine.Ordine{
					ID:         2,
					Table:      8,
					ClientName: "Test 2",
					Status:     true,
				})
			},
			dropDownTest: func() {
				tests.DropTablesForTests(&product.Product{}, &ordine.Ordine{}, &order.OrderProducts{})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.setUpTest()
			defer tt.dropDownTest()

			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			c.Request = httptest.NewRequest(http.MethodGet, "/ordines", nil)

			GetOrdines(c)

			if tt.expectedStatusCode != rr.Code {
				t.Errorf("Invalid status code, got %v expected %v", tt.expectedStatusCode, rr.Code)
			}
		})
	}
}
