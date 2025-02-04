package controllers

import (
	"bytes"
	"encoding/json"
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

func TestPostProductsToOrdine(t *testing.T) {
	tests := []struct {
		name               string
		ordineId           string
		requestBody        []order.OrderProductBody
		expectedStatusCode int
		setUpTest          func()
		dropDownTest       func()
	}{
		{
			name:               "Should return 400, invalid body",
			ordineId:           "1",
			requestBody:        []order.OrderProductBody{{}, {}},
			expectedStatusCode: http.StatusBadRequest,
			setUpTest: func() {
				// Do nothing
			},
			dropDownTest: func() {
				// Do nothing
			},
		},
		{
			name:     "Should return 400, ordine not founded",
			ordineId: "1",
			requestBody: []order.OrderProductBody{{
				ProductID: 1,
				Quantity:  2,
			}},
			expectedStatusCode: http.StatusBadRequest,
			setUpTest: func() {
				tests.MakeMigrationsForTests(&product.Product{}, &ordine.Ordine{}, order.OrderProducts{})
			},
			dropDownTest: func() {
				tests.DropTablesForTests(&product.Product{}, &ordine.Ordine{}, order.OrderProducts{})
			},
		},
		{
			name:     "Should return 400, product not founded",
			ordineId: "1",
			requestBody: []order.OrderProductBody{{
				ProductID: 1,
				Quantity:  2,
			}},
			expectedStatusCode: http.StatusBadRequest,
			setUpTest: func() {
				tests.MakeMigrationsForTests(&product.Product{}, &ordine.Ordine{}, order.OrderProducts{})
				tests.CreateInsertValueForTests(&ordine.Ordine{
					Table:      8,
					ClientName: "test",
					Status:     true,
				})
			},
			dropDownTest: func() {
				tests.DropTablesForTests(&product.Product{}, &ordine.Ordine{}, order.OrderProducts{})
			},
		},
		{
			name:     "Should return 500, tables not founded on database",
			ordineId: "1",
			requestBody: []order.OrderProductBody{{
				ProductID: 1,
				Quantity:  2,
			}},
			expectedStatusCode: http.StatusInternalServerError,
			setUpTest: func() {
				// Do nothing
			},
			dropDownTest: func() {
				// Do nothing
			},
		},
		{
			name:     "Should return 200, product not founded",
			ordineId: "1",
			requestBody: []order.OrderProductBody{{
				ProductID: 1,
				Quantity:  2,
			}},
			expectedStatusCode: http.StatusOK,
			setUpTest: func() {
				tests.MakeMigrationsForTests(&product.Product{}, &ordine.Ordine{}, order.OrderProducts{})
				tests.CreateInsertValueForTests(&ordine.Ordine{
					Table:      8,
					ClientName: "test",
					Status:     true,
				})
				tests.CreateInsertValueForTests(&product.Product{
					Name:        "test",
					Category:    "test",
					Price:       9.99,
					Description: "test",
					Stock:       1,
				})
			},
			dropDownTest: func() {
				tests.DropTablesForTests(&product.Product{}, &ordine.Ordine{}, order.OrderProducts{})
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

			c.Request = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/ordine/%s/products", tt.ordineId), bytes.NewBuffer(body))
			c.Params = gin.Params{{Key: "id", Value: tt.ordineId}}

			PostProductsToOrdine(c)

			if tt.expectedStatusCode != rr.Code {
				t.Errorf("Invalid status code, expected %v got %v", tt.expectedStatusCode, rr.Code)
			}

		})
	}
}
