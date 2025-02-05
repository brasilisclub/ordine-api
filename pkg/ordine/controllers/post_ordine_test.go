package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ordine-api/pkg/ordine"
	"ordine-api/pkg/ordine/order"
	"ordine-api/pkg/product"
	"ordine-api/tests"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPostOrdine(t *testing.T) {

	tests := []struct {
		name               string
		expectedStatusCode int
		requestBody        ordine.OrdineRequestBody
		setUpTest          func()
		dropDownTest       func()
	}{
		{
			name:               "Should return 400, invalid body",
			expectedStatusCode: http.StatusBadRequest,
			requestBody:        ordine.OrdineRequestBody{},
			setUpTest: func() {
				// Do nothing
			},
			dropDownTest: func() {
				// Do nothing
			},
		},
		{
			name:               "Should return 500, no such table in database",
			expectedStatusCode: http.StatusInternalServerError,
			requestBody: ordine.OrdineRequestBody{
				Table:      10,
				ClientName: "test",
				Status:     true,
			},
			setUpTest: func() {
				// Do nothing
			},
			dropDownTest: func() {
				// Do nothing
			},
		},
		{
			name:               "Should return 201, ordine created successfully",
			expectedStatusCode: http.StatusCreated,
			requestBody: ordine.OrdineRequestBody{
				Table:      10,
				ClientName: "test",
				Status:     true,
			},
			setUpTest: func() {
				tests.MakeMigrationsForTests(&product.Product{}, &ordine.Ordine{}, &order.OrderProducts{})
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

			body, _ := json.Marshal(tt.requestBody)

			c.Request = httptest.NewRequest(http.MethodPost, "/ordine", bytes.NewBuffer(body))
			c.Request.Header.Set("Content-Type", "application/json")

			PostOrdine(c)

			if tt.expectedStatusCode != rr.Code {
				t.Errorf("Invalid status code, expected %v got %v", tt.expectedStatusCode, rr.Code)
			}

		})
	}
}
