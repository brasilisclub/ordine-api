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

func TestPutOrdine(t *testing.T) {
	tests := []struct {
		name               string
		ordineId           string
		expectedStatusCode int
		requestBody        ordine.OrdineRequestBody
		setUpTest          func()
		dropDownTest       func()
	}{
		{
			name:               "Should return 400, invalid Body",
			ordineId:           "1",
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
			name:               "Should return 400, ordine not founded",
			ordineId:           "1",
			expectedStatusCode: http.StatusBadRequest,
			requestBody: ordine.OrdineRequestBody{
				Table:      8,
				ClientName: "new name",
				Status:     true,
			},
			setUpTest: func() {
				tests.MakeMigrationsForTests(&ordine.Ordine{}, &product.Product{}, &order.OrderProducts{})
			},
			dropDownTest: func() {
				tests.DropTablesForTests(&ordine.Ordine{}, &product.Product{}, &order.OrderProducts{})
			},
		},
		{
			name:               "Should return 500, table not founded on database",
			ordineId:           "1",
			expectedStatusCode: http.StatusInternalServerError,
			requestBody: ordine.OrdineRequestBody{
				Table:      8,
				ClientName: "new name",
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
			name:               "Should return 200, ordine not founded",
			ordineId:           "1",
			expectedStatusCode: http.StatusOK,
			requestBody: ordine.OrdineRequestBody{
				Table:      8,
				ClientName: "new name",
				Status:     true,
			},
			setUpTest: func() {
				tests.MakeMigrationsForTests(&ordine.Ordine{}, &product.Product{}, &order.OrderProducts{})
				tests.CreateInsertValueForTests(&ordine.Ordine{
					Table:      8,
					ClientName: "test",
					Status:     true,
				})
			},
			dropDownTest: func() {
				tests.DropTablesForTests(&ordine.Ordine{}, &product.Product{}, &order.OrderProducts{})
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
			c.Request = httptest.NewRequest(http.MethodPut, fmt.Sprintf("/ordine/%s", tt.ordineId), bytes.NewBuffer(body))
			c.Params = gin.Params{{Key: "id", Value: tt.ordineId}}

			PutOrdine(c)

			if tt.expectedStatusCode != rr.Code {
				t.Errorf("Invalid status code, expected %v got %v", tt.expectedStatusCode, rr.Code)
			}
		})
	}
}
